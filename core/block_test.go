package core

import (
	"fmt"
	"testing"
	"time"

	"github.com/ayushn2/blockchainz/crypto"
	"github.com/ayushn2/blockchainz/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32, prevBlockHash types.Hash) *Block {
	h := &Header{
		Version:   1,
		PrevHash:  prevBlockHash,
		Timestamp: uint64(time.Now().UnixNano()),
		Height:    height,
		Nonce:     0,
	}

	tx := []Transaction{
		{Data: []byte("test transaction")},
	}

	return NewBlock(h, tx)
}

func randomBlockWithSignature(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(height, prevBlockHash)
	err := b.Sign(privKey)
	assert.Nil(t, err)
 
	return b
}

func TestHashBlock(t *testing.T){
	
	b := randomBlock(0, types.Hash{})
	fmt.Println(b.Hash(BlockHasher{}))
}

func TestSignBlock(t *testing.T){
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.Hash{})
	err := b.Sign(privKey)
	assert.Nil(t, err, "Block should be signed successfully")
	assert.NotNil(t, b.Signature, "Block signature should not be nil after signing")
}

func TestBlockVerify(t *testing.T){
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.Hash{})
	
	assert.Nil(t, b.Sign(privKey), "Block should be signed successfully")
	assert.Nil(t, b.Verify(), "Block signature should not be nil after signing")

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()

	assert.NotNil(t, b.Verify())

	b.Height = 1 // Change height to simulate a different block
	assert.NotNil(t, b.Verify(), "Block verification should fail with different validator public key")
}