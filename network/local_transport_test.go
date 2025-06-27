package network

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")
	tra.Connect(trb)
	trb.Connect(tra)
	// assert.Equal(t, tra.peers[trb.addr], trb)
	// assert.Equal(t, trb.peers[tra.addr], tra)
}

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("Hello, World!")
	// err := tra.SendMessage(trb.addr, msg)
	// assert.NoError(t, err)

	select {
	case rpc := <-trb.Consume():
		// assert.Equal(t, rpc.From, tra.addr)
		assert.Equal(t, rpc.Payload, msg)
	default:
		t.Error("Expected message not received")
	}
}