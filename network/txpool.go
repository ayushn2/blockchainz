package network

import (
	"sort"

	"github.com/ayushn2/blockchainz/core"
	"github.com/ayushn2/blockchainz/types"
)

type TxMapSorter struct{
	transactions []*core.Transaction
}

func NewTxMapSorter(txs map[types.Hash]*core.Transaction) *TxMapSorter {
	txx := make([]*core.Transaction, 0, len(txs))

	i := 0
	for _, tx := range txs {
		txx = append(txx, tx)
		i++
	}

	s := &TxMapSorter{
		transactions: txx,
	}

	sort.Sort(s)
	return s
}

func (s *TxMapSorter) Len() int {
	return len(s.transactions)
}

func (s *TxMapSorter) Swap(i, j int) {
	s.transactions[i], s.transactions[j] = s.transactions[j], s.transactions[i]
}

func (s *TxMapSorter) Less(i, j int) bool {
	return s.transactions[i].FirstSeen() < s.transactions[j].FirstSeen()
}

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

func (p *TxPool) Transactions() []*core.Transaction {
	s := NewTxMapSorter(p.transactions)
	return s.transactions
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