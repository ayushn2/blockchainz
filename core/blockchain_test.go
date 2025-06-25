package core

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesis(t *testing.T) *Blockchain{
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)

	return bc
}

func TestBlockchain(t *testing.T){
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))

	fmt.Println(bc.Height())
}

func TestHasBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.True(t, bc.HasBlock(0))
}

func TestAddBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	lenBlock := 100
	for i:= 0; i<lenBlock;i++{
		block := randomBlock(uint32(i+1))
		err := bc.AddBlock(block)
		assert.Nil(t, err)
	}

	assert.Equal(t, bc.Height(), uint32(lenBlock))
	assert.Equal(t, len(bc.headers), lenBlock +1)
}