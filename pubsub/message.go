// Copyright 2016 Google LLC
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

package pubsub

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	pb "google.golang.org/genproto/googleapis/pubsub/v1"
)

// Message represents a Pub/Sub message.
type Message struct {
	// ID identifies this message. This ID is assigned by the server and is
	// populated for Messages obtained from a subscription.
	//
	// This field is read-only.
	ID string

	// Data is the actual data in the message.
	Data []byte

	// Attributes represents the key-value pairs the current message is
	// labelled with.
	Attributes map[string]string

	// PublishTime is the time at which the message was published. This is
	// populated by the server for Messages obtained from a subscription.
	//
	// This field is read-only.
	PublishTime time.Time

	// DeliveryAttempt is the number of times a message has been delivered.
	// This is part of the dead lettering feature that forwards messages that
	// fail to be processed (from nack/ack deadline timeout) to a dead letter topic.
	// If dead lettering is enabled, this will be set on all attempts, starting
	// with value 1. Otherwise, the value will be nil.
	// This field is read-only.
	DeliveryAttempt *int

	// size is the approximate size of the message's data and attributes.
	size int

	// OrderingKey identifies related messages for which publish order should
	// be respected. If empty string is used, message will be sent unordered.
	OrderingKey string

	// ackh handles Ack() or Nack().
	ackh ackHandler
}

// NewMessage creates a message with a custom ack/nack handler, which should not
// be nil.
func NewMessage(ackh ackHandler) *Message {
	return &Message{ackh: ackh}
}

func toMessage(resp *pb.ReceivedMessage) (*Message, error) {
	if resp.Message == nil {
		return &Message{ackh: &psAckHandler{ackID: resp.AckId}}, nil
	}

	pubTime, err := ptypes.Timestamp(resp.Message.PublishTime)
	if err != nil {
		return nil, err
	}

	var deliveryAttempt *int
	if resp.DeliveryAttempt > 0 {
		da := int(resp.DeliveryAttempt)
		deliveryAttempt = &da
	}

	return &Message{
		Data:            resp.Message.Data,
		Attributes:      resp.Message.Attributes,
		ID:              resp.Message.MessageId,
		PublishTime:     pubTime,
		DeliveryAttempt: deliveryAttempt,
		OrderingKey:     resp.Message.OrderingKey,
		ackh:            &psAckHandler{ackID: resp.AckId},
	}, nil
}

// Ack indicates successful processing of a Message passed to the Subscriber.Receive callback.
// It should not be called on any other Message value.
// If message acknowledgement fails, the Message will be redelivered.
// Client code must call Ack or Nack when finished for each received Message.
// Calls to Ack or Nack have no effect after the first call.
func (m *Message) Ack() {
	if m.ackh != nil {
		m.ackh.OnAck()
	}
}

// Nack indicates that the client will not or cannot process a Message passed to the Subscriber.Receive callback.
// It should not be called on any other Message value.
// Nack will result in the Message being redelivered more quickly than if it were allowed to expire.
// Client code must call Ack or Nack when finished for each received Message.
// Calls to Ack or Nack have no effect after the first call.
func (m *Message) Nack() {
	if m.ackh != nil {
		m.ackh.OnNack()
	}
}

// ackHandler performs a safe cast of the message's ack handler to psAckHandler.
func (m *Message) ackHandler() (*psAckHandler, bool) {
	ackh, ok := m.ackh.(*psAckHandler)
	return ackh, ok
}

func (m *Message) ackID() string {
	if ackh, ok := m.ackh.(*psAckHandler); ok {
		return ackh.ackID
	}
	return ""
}

// ackHandler implements ack/nack handling.
type ackHandler interface {
	OnAck()
	OnNack()
}

// psAckHandler handles ack/nack for the pubsub package.
type psAckHandler struct {
	// ackID is the identifier to acknowledge this message.
	ackID string

	// receiveTime is the time the message was received by the client.
	receiveTime time.Time

	calledDone bool

	// The done method of the iterator that created this Message.
	doneFunc func(string, bool, time.Time)
}

func (ah *psAckHandler) OnAck() {
	ah.done(true)
}

func (ah *psAckHandler) OnNack() {
	ah.done(false)
}

func (ah *psAckHandler) done(ack bool) {
	if ah.calledDone {
		return
	}
	ah.calledDone = true
	if ah.doneFunc != nil {
		ah.doneFunc(ah.ackID, ack, ah.receiveTime)
	}
}
