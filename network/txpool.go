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

// Adds a transaction to the pool, the caller is responsible for 
// ensuring the transaction is already present
func (p *TxPool) Add(tx *core.Transaction) error{
	hash := tx.Hash(core.TxHasher{})
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