package network

import (
	"github.com/ayushn2/blockchainz/core"
	"github.com/ayushn2/blockchainz/types"
)

type TxPool struct {
	transactions map[types.Hash]*core.Transaction // Map of transaction ID to Transaction
}

func NewTxPool() *TxPool {
	return &TxPool{
		transactions: make(map[types.Hash]*core.Transaction),
	}
}

func (p *TxPool) Len() int{
	return len(p.transactions)
}

func (p *TxPool) Add(tx *core.Transaction) error{
	hash := tx.Hash(core.TxHasher{})
	if p.Has(hash) {
		return nil // Transaction already exists in the pool
		// no need to return an error becoz we are going to frequently face repeated transactions
	}

	p.transactions[hash] = tx

	return nil
}

func (p *TxPool) Has(hash types.Hash) bool {
	_, exists := p.transactions[hash]
	return exists
}

func (p *TxPool) Flush() {
	p.transactions = make(map[types.Hash]*core.Transaction)
}