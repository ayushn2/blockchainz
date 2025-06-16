package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/ayushn2/blockchainz/types"
)

func randomBlock(height uint32) *Block {
	h := &Header{
		Version:   1,
		PrevHash:  types.Hash{},
		Timestamp: uint64(time.Now().UnixNano()),
		Height:    height,
		Nonce:     0,
	}

	tx := []Transaction{
		{Data: []byte("test transaction")},
	}

	return NewBlock(h, tx)
}

func TestHashBlock(t *testing.T){
	
	b := randomBlock(0)
	fmt.Println(b.Hash(BlockHasher{}))
}