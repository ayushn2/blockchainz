package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"

	"github.com/ayushn2/blockchainz/types"
)

// Generic hasher interface for any type T.
// Requires a Hash method that takes T and returns types.Hash.
type Hasher[T any] interface {
	Hash(T) types.Hash // Hash computes the hash of the given type T
}

type BlockHasher struct{}

func (BlockHasher) Hash(b *Block) types.Hash {
	 buf := &bytes.Buffer{}
	 enc := gob.NewEncoder(buf)
	 if err := enc.Encode(b.Header); err != nil {
		 panic(err)
	 }

	 h := sha256.Sum256(buf.Bytes())

	 return types.Hash(h)
}