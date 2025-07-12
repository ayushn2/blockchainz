package network

import (
	"strconv"
	"testing"
	"time"

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

func TestSortTransactions (t *testing.T){
	p := NewTxPool()
	txLen := 1000

	now := time.Now().UnixMilli()

	for i := 0; i < txLen; i++ {
		tx := core.NewTransaction([]byte(strconv.FormatInt(int64(i), 10)))
		tx.SetFirstSeen(now + int64(i))
		err := p.Add(tx)
		assert.NoError(t, err, "Adding a transaction should not return an error")
	}

	assert.Equal(t, p.Len(), txLen, "Transaction pool should have all transactions after adding")

	txx := p.Transactions()
	for i :=0 ;i< txLen-1; i++{
		assert.True(t, txx[i].FirstSeen() < txx[i+1].FirstSeen(), "Transactions should be sorted by first seen time")
	}
}