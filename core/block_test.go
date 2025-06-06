package core

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/ayushn2/blockchainz/types"
	"github.com/stretchr/testify/assert"
)
func TestHeader_Encode_Decode(t *testing.T){
	h := &Header{
		Version:   1,
		PrevHash: types.RandomHash(),
		Timestamp: uint64(time.Now().UnixNano()),
		Height: 10,
		Nonce: 989394,
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, h.EncodeBinary(buf))

	hDecode:= &Header{}
	assert.Nil(t, hDecode.DecodeBinary(buf))
	assert.Equal(t, h, hDecode)
}

func TestBlock_Decode(t *testing.T){
	b := &Block{
		Header: Header{
			Version:   1,
			PrevHash: types.RandomHash(),
			Timestamp: uint64(time.Now().UnixNano()),
			Height: 10,
			Nonce: 989394,
		},
		Transactions: []Transaction(nil),// use nil slice; using an empty slice []Transaction{} will cause test to fail
	}

	buf := &bytes.Buffer{}
	assert.Nil(t, b.EncodeBinary(buf))

	bDecode := &Block{}
	assert.Nil(t, bDecode.DecodeBinary(buf))
	assert.Equal(t, b, bDecode)
}

func TestBlockHash( t *testing.T){
	b := &Block{
		Header: Header{
			Version:   1,
			PrevHash: types.RandomHash(),
			Timestamp: uint64(time.Now().UnixNano()),
			Height: 10,
			Nonce: 989394,
		},
		Transactions: []Transaction(nil),// use nil slice; using an empty slice []Transaction{} will cause test to fail
	}

	h := b.Hash()
	fmt.Println(h)
	assert.False(t, h.IsZero(), "Block hash should not be zero")
}