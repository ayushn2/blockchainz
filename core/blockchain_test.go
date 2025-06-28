package core

import (
	"testing"

	"github.com/ayushn2/blockchainz/types"
	"github.com/stretchr/testify/assert"
)



func TestNewBlockchain(t *testing.T){
	bc, err := NewBlockchain(randomBlock(0, types.Hash{}))
	assert.Nil(t, err)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))
}

func TestHasBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.True(t, bc.HasBlock(0))
}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	lenBlock := 100
	for i := range(lenBlock){
		block := randomBlockWithSignature(t, uint32(i+1), getPrevBlockHash(t, bc, uint32(i+1)))
		err := bc.AddBlock(block)
		assert.Nil(t, err)
	}
	
	assert.Equal(t, bc.Height(), uint32(lenBlock))
	assert.Equal(t, len(bc.headers), lenBlock +1)

	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, 98,types.Hash{}))) //should not have added the new block
}

func TestGetHeader(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	// Add a block with height 1
	for i := range(10){
		block := randomBlockWithSignature(t, uint32(i+1), getPrevBlockHash(t, bc, uint32(i+1)))
		err := bc.AddBlock(block)
		assert.Nil(t, err)
		header, err := bc.GetHeader(block.Height)
		assert.Nil(t, err)
		assert.Equal(t, block.Header ,header)
	}
}

func TestAddBlockToHigh(t *testing.T){
	bc := newBlockchainWithGenesis(t)

	// Add a block with height 10
	block := randomBlockWithSignature(t, 10, types.Hash{})
	err := bc.AddBlock(block)
	assert.NotNil(t, err)
}

func newBlockchainWithGenesis(t *testing.T) *Blockchain{
	bc, err := NewBlockchain(randomBlock(0, types.Hash{}))
	assert.Nil(t, err)

	return bc
}

func getPrevBlockHash(t *testing.T, bc *Blockchain, height uint32) types.Hash {
	prevHeader, err := bc.GetHeader(height -1)
	assert.Nil(t, err)
	return BlockHasher{}.Hash(prevHeader)
}