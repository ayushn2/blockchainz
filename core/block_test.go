package core

import (
	"bytes"
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