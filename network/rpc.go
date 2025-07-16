package network

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"

	"github.com/ayushn2/blockchainz/core"
)

type MessageType byte

const (
	MessageTypeTxn MessageType = 0x1
	MessageTypeBlock
)

type RPC struct{
	From NetAddr
	Payload io.Reader
}

type Message struct{
	Header MessageType
	Data []byte
}

func NewMessage(t MessageType, data []byte) *Message {
	return &Message{
		Header: t,
		Data:   data,
	}
}

func (msg *Message) Bytes() []byte{
	buf := &bytes.Buffer{}
	gob.NewEncoder(buf).Encode(msg)
	return buf.Bytes()
}

type DecodeMessage struct {
	From NetAddr
	Data any
}

type RPCDecodeFunc func(RPC) (*DecodeMessage, error)

func DefaultRPCDecodeFunc(rpc RPC) (*DecodeMessage, error) {
	msg := &Message{}
	if err := gob.NewDecoder(rpc.Payload).Decode(&msg); err != nil {
		return nil, fmt.Errorf("failed to decode message from %s: %s",rpc.From, err)
	}

	switch msg.Header {
	case MessageTypeTxn:
		tx := new(core.Transaction)
		if err := tx.Decode(core.NewGobTxDecoder(bytes.NewReader(msg.Data))); err != nil {
			return nil, err
		}
		
		return &DecodeMessage{
			From: rpc.From,
			Data: tx,
		}, nil
	
	default:
		return nil, fmt.Errorf("invalid message header: %v", msg.Header)
	}
}

type RPCProcessor interface {
	ProcessMessage(*DecodeMessage) error
}