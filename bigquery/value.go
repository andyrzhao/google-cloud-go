// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bigquery

import (
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"cloud.google.com/go/civil"

	bq "google.golang.org/api/bigquery/v2"
)

// Value stores the contents of a single cell from a BigQuery result.
type Value interface{}

// ValueLoader stores a slice of Values representing a result row from a Read operation.
// See RowIterator.Next for more information.
type ValueLoader interface {
	Load(v []Value, s Schema) error
}

// valueList converts a []Value to implement ValueLoader.
type valueList []Value

// Load stores a sequence of values in a valueList.
// It resets the slice length to zero, then appends each value to it.
func (vs *valueList) Load(v []Value, _ Schema) error {
	*vs = append((*vs)[:0], v...)
	return nil
}

// valueMap converts a map[string]Value to implement ValueLoader.
type valueMap map[string]Value

// Load stores a sequence of values in a valueMap.
func (vm *valueMap) Load(v []Value, s Schema) error {
	if *vm == nil {
		*vm = map[string]Value{}
	}
	// TODO(jba): treat nested records as nested maps.
	for i, f := range s {
		(*vm)[f.Name] = v[i]
	}
	return nil
}

type structLoader struct {
	typ reflect.Type // type of struct
	err error
	ops []structLoaderOp

	vstructp reflect.Value // pointer to current struct value; changed by set
}

// A setFunc is a function that sets a struct field to a value.
type setFunc func(field reflect.Value, val interface{}) error

// A structLoaderOp instructs the loader to set a struct field to a row value.
type structLoaderOp struct {
	fieldIndex []int
	valueIndex int
	setFunc    setFunc
	repeated   bool
	nested     []structLoaderOp // for nested schemas
}

func setAny(v reflect.Value, x interface{}) error {
	v.Set(reflect.ValueOf(x))
	return nil
}

func setInt(v reflect.Value, x interface{}) error {
	xx := x.(int64)
	if v.OverflowInt(xx) {
		return fmt.Errorf("bigquery: value %v overflows struct field of type %v", xx, v.Type())
	}
	v.SetInt(xx)
	return nil
}

func setFloat(v reflect.Value, x interface{}) error {
	xx := x.(float64)
	if v.OverflowFloat(xx) {
		return fmt.Errorf("bigquery: value %v overflows struct field of type %v", xx, v.Type())
	}
	v.SetFloat(xx)
	return nil
}

func setBool(v reflect.Value, x interface{}) error {
	v.SetBool(x.(bool))
	return nil
}

func setString(v reflect.Value, x interface{}) error {
	v.SetString(x.(string))
	return nil
}

func setBytes(v reflect.Value, x interface{}) error {
	v.SetBytes(x.([]byte))
	return nil
}

// set remembers a value for the next call to Load. The value must be
// a pointer to a struct. (This is checked in RowIterator.Next.)
func (sl *structLoader) set(structp interface{}, schema Schema) error {
	if sl.err != nil {
		return sl.err
	}
	sl.vstructp = reflect.ValueOf(structp)
	typ := sl.vstructp.Type().Elem()
	if sl.typ == nil {
		// First call: remember the type and compile the schema.
		sl.typ = typ
		ops, err := compileToOps(typ, schema)
		if err != nil {
			sl.err = err
			return err
		}
		sl.ops = ops
	} else if sl.typ != typ {
		return fmt.Errorf("bigquery: struct type changed from %s to %s", sl.typ, typ)
	}
	return nil
}

// compileToOps produces a sequence of operations that will set the fields of a
// value of structType to the contents of a row with schema.
func compileToOps(structType reflect.Type, schema Schema) ([]structLoaderOp, error) {
	var ops []structLoaderOp
	fields := fieldCache.Fields(structType)
	for i, schemaField := range schema {
		// Look for an exported struct field with the same name as the schema
		// field, ignoring case (BigQuery column names are case-insensitive,
		// and we want to act like encoding/json anyway).
		structField := fields.Match(schemaField.Name)
		if structField == nil {
			// Ignore schema fields with no corresponding struct field.
			continue
		}
		op := structLoaderOp{
			fieldIndex: structField.Index,
			valueIndex: i,
		}
		if schemaField.Type == RecordFieldType {
			// Field can be a struct or a pointer to a struct.
			t := structField.Type
			if t.Kind() == reflect.Ptr {
				t = t.Elem()
			}
			if t.Kind() != reflect.Struct {
				return nil, fmt.Errorf("bigquery: field %s has type %s, expected struct or *struct",
					structField.Name, structField.Type)
			}
			nested, err := compileToOps(t, schemaField.Schema)
			if err != nil {
				return nil, err
			}
			op.nested = nested
		} else {
			t := structField.Type
			if schemaField.Repeated {
				// TODO(jba): handle pointers to slices and arrays
				if t.Kind() != reflect.Slice && t.Kind() != reflect.Array {
					return nil, fmt.Errorf("bigquery: repeated schema field %s requires slice or array, but struct field %s has type %s",
						schemaField.Name, structField.Name, t)
				}
				t = t.Elem()
				op.repeated = true
			}
			op.setFunc = determineSetFunc(t, schemaField.Type)
			if op.setFunc == nil {
				return nil, fmt.Errorf("bigquery: schema field %s of type %s is not assignable to struct field %s of type %s",
					schemaField.Name, schemaField.Type, structField.Name, t)
			}
		}
		ops = append(ops, op)
	}
	return ops, nil
}

// determineSetFunc chooses the best function for setting a field of type ftype
// to a value whose schema field type is sftype. It returns nil if stype
// is not assignable to ftype.
func determineSetFunc(ftype reflect.Type, stype FieldType) setFunc {
	switch stype {
	case StringFieldType:
		if ftype.Kind() == reflect.String {
			return setString
		}

	case BytesFieldType:
		if ftype == typeOfByteSlice {
			return setBytes
		}

	case IntegerFieldType:
		if isSupportedIntType(ftype) {
			return setInt
		}

	case FloatFieldType:
		switch ftype.Kind() {
		case reflect.Float32, reflect.Float64:
			return setFloat
		}

	case BooleanFieldType:
		if ftype.Kind() == reflect.Bool {
			return setBool
		}

	case TimestampFieldType:
		if ftype == typeOfGoTime {
			return setAny
		}

	case DateFieldType:
		if ftype == typeOfDate {
			return setAny
		}

	case TimeFieldType:
		if ftype == typeOfTime {
			return setAny
		}

	case DateTimeFieldType:
		if ftype == typeOfDateTime {
			return setAny
		}
	}
	return nil
}

func (sl *structLoader) Load(values []Value, _ Schema) error {
	if sl.err != nil {
		return sl.err
	}
	return runOps(sl.ops, sl.vstructp.Elem(), values)
}

// runOps executes a sequence of ops, setting the fields of vstruct to the
// supplied values.
func runOps(ops []structLoaderOp, vstruct reflect.Value, values []Value) error {
	for _, op := range ops {
		field := vstruct.FieldByIndex(op.fieldIndex)
		switch {
		case op.nested != nil:
			// Support both structs and pointers to structs.
			vsub := field
			if field.Kind() == reflect.Ptr {
				field.Set(reflect.New(field.Type().Elem()))
				vsub = vsub.Elem()
			}
			if err := runOps(op.nested, vsub, values[op.valueIndex].([]Value)); err != nil {
				return err
			}

		case op.repeated:
			if err := setRepeated(field, values[op.valueIndex].([]Value), op.setFunc); err != nil {
				return err
			}

		default:
			if err := op.setFunc(field, values[op.valueIndex]); err != nil {
				return err
			}
		}
	}
	return nil
}

func setRepeated(field reflect.Value, vslice []Value, setElem setFunc) error {
	vlen := len(vslice)
	var flen int
	switch field.Type().Kind() {
	case reflect.Slice:
		// Make a slice of the right size, avoiding allocation if possible.
		switch {
		case field.Len() < vlen:
			field.Set(reflect.MakeSlice(field.Type(), vlen, vlen))
		case field.Len() > vlen:
			field.SetLen(vlen)
		}
		flen = vlen

	case reflect.Array:
		flen = field.Len()
		if flen > vlen {
			// Set extra elements to their zero value.
			z := reflect.Zero(field.Type().Elem())
			for i := vlen; i < flen; i++ {
				field.Index(i).Set(z)
			}
		}
	}
	for i, val := range vslice {
		if i < flen { // avoid writing past the end of a short array
			if err := setElem(field.Index(i), val); err != nil {
				return err
			}
		}
	}
	return nil
}

// A ValueSaver returns a row of data to be inserted into a table.
type ValueSaver interface {
	// Save returns a row to be inserted into a BigQuery table, represented
	// as a map from field name to Value.
	// If insertID is non-empty, BigQuery will use it to de-duplicate
	// insertions of this row on a best-effort basis.
	Save() (row map[string]Value, insertID string, err error)
}

// ValuesSaver implements ValueSaver for a slice of Values.
type ValuesSaver struct {
	Schema Schema

	// If non-empty, BigQuery will use InsertID to de-duplicate insertions
	// of this row on a best-effort basis.
	InsertID string

	Row []Value
}

// Save implements ValueSaver
func (vls *ValuesSaver) Save() (map[string]Value, string, error) {
	m, err := valuesToMap(vls.Row, vls.Schema)
	return m, vls.InsertID, err
}

func valuesToMap(vs []Value, schema Schema) (map[string]Value, error) {
	if len(vs) != len(schema) {
		return nil, errors.New("Schema does not match length of row to be inserted")
	}

	m := make(map[string]Value)
	for i, fieldSchema := range schema {
		if fieldSchema.Type == RecordFieldType {
			nested, ok := vs[i].([]Value)
			if !ok {
				return nil, errors.New("Nested record is not a []Value")
			}
			value, err := valuesToMap(nested, fieldSchema.Schema)
			if err != nil {
				return nil, err
			}
			m[fieldSchema.Name] = value
		} else {
			m[fieldSchema.Name] = vs[i]
		}
	}
	return m, nil
}

// convertRows converts a series of TableRows into a series of Value slices.
// schema is used to interpret the data from rows; its length must match the
// length of each row.
func convertRows(rows []*bq.TableRow, schema Schema) ([][]Value, error) {
	var rs [][]Value
	for _, r := range rows {
		row, err := convertRow(r, schema)
		if err != nil {
			return nil, err
		}
		rs = append(rs, row)
	}
	return rs, nil
}

func convertRow(r *bq.TableRow, schema Schema) ([]Value, error) {
	if len(schema) != len(r.F) {
		return nil, errors.New("schema length does not match row length")
	}
	var values []Value
	for i, cell := range r.F {
		fs := schema[i]
		v, err := convertValue(cell.V, fs.Type, fs.Schema)
		if err != nil {
			return nil, err
		}
		values = append(values, v)
	}
	return values, nil
}

func convertValue(val interface{}, typ FieldType, schema Schema) (Value, error) {
	switch val := val.(type) {
	case nil:
		return nil, nil
	case []interface{}:
		return convertRepeatedRecord(val, typ, schema)
	case map[string]interface{}:
		return convertNestedRecord(val, schema)
	case string:
		return convertBasicType(val, typ)
	default:
		return nil, fmt.Errorf("got value %v; expected a value of type %s", val, typ)
	}
}

func convertRepeatedRecord(vals []interface{}, typ FieldType, schema Schema) (Value, error) {
	var values []Value
	for _, cell := range vals {
		// each cell contains a single entry, keyed by "v"
		val := cell.(map[string]interface{})["v"]
		v, err := convertValue(val, typ, schema)
		if err != nil {
			return nil, err
		}
		values = append(values, v)
	}
	return values, nil
}

func convertNestedRecord(val map[string]interface{}, schema Schema) (Value, error) {
	// convertNestedRecord is similar to convertRow, as a record has the same structure as a row.

	// Nested records are wrapped in a map with a single key, "f".
	record := val["f"].([]interface{})
	if len(record) != len(schema) {
		return nil, errors.New("schema length does not match record length")
	}

	var values []Value
	for i, cell := range record {
		// each cell contains a single entry, keyed by "v"
		val := cell.(map[string]interface{})["v"]

		fs := schema[i]
		v, err := convertValue(val, fs.Type, fs.Schema)
		if err != nil {
			return nil, err
		}
		values = append(values, v)
	}
	return values, nil
}

// convertBasicType returns val as an interface with a concrete type specified by typ.
func convertBasicType(val string, typ FieldType) (Value, error) {
	switch typ {
	case StringFieldType:
		return val, nil
	case BytesFieldType:
		return base64.StdEncoding.DecodeString(val)
	case IntegerFieldType:
		return strconv.ParseInt(val, 10, 64)
	case FloatFieldType:
		return strconv.ParseFloat(val, 64)
	case BooleanFieldType:
		return strconv.ParseBool(val)
	case TimestampFieldType:
		f, err := strconv.ParseFloat(val, 64)
		return Value(time.Unix(0, int64(f*1e9))), err
	case DateFieldType:
		return civil.ParseDate(val)
	case TimeFieldType:
		return civil.ParseTime(val)
	case DateTimeFieldType:
		return civil.ParseDateTime(val)
	default:
		return nil, fmt.Errorf("unrecognized type: %s", typ)
	}
}
