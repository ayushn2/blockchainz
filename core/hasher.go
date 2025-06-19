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

func (BlockHasher) Hash(b *Block) types.Hash {
	 h := sha256.Sum256(b.HeaderData())

	 return types.Hash(h)
}