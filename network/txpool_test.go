package network

import (
	"testing"

	"github.com/ayushn2/blockchainz/core"
	"github.com/stretchr/testify/assert"
)

func TestTxPool(t *testing.T){
	p := NewTxPool()
	assert.Equal(t, p.Len(), 0, "New transaction pool should be empty")
	
	
}

func TestTxPoolAddTx(t *testing.T){
	p := NewTxPool()
	tx := core.NewTransaction([]byte("test transaction"))
	err := p.Add(tx)
	assert.NoError(t, err, "Adding a transaction should not return an error")
	assert.Equal(t, p.Len(), 1, "Transaction pool should have one transaction after adding")

	txx := core.NewTransaction([]byte("test transaction")) // same transaction
	err = p.Add(txx) // adding the same transaction again
	assert.Nil(t, err, "Adding the same transaction again should not return an error")
	assert.Equal(t, p.Len(), 1, "Transaction pool should still have one transaction after adding the same transaction again")

	p.Flush()
	assert.Equal(t, p.Len(), 0, "Transaction pool should be empty after flushing")
}