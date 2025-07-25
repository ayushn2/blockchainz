package core

import (
	"crypto/sha256"
	"github.com/ayushn2/blockchainz/types"
)

// Generic hasher interface for any type T.
// Requires a Hash method that takes T and returns types.Hash.
type Hasher[T any] interface {
	Hash(T) types.Hash // Hash computes the hash of the given type T
}

type BlockHasher struct{}

func (BlockHasher) Hash(head *Header) types.Hash {
	 h := sha256.Sum256(head.Bytes())

	 return types.Hash(h)
}

type TxHasher struct{}

func (TxHasher) Hash(tx *Transaction) types.Hash {
	h := sha256.Sum256(tx.Data)

	return types.Hash(h)
}