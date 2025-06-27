package core

import "fmt"

type Validator interface{
	// ValidateBlock checks if the block is valid according to the blockchain rules.
	ValidateBlock(*Block) error
}

type BlockValidator struct{
	bc *Blockchain
}

func NewBlockValidator(bc *Blockchain) *BlockValidator {
	return &BlockValidator{
		bc: bc,
	}
}

func (v *BlockValidator) ValidateBlock(b *Block) error{
	if v.bc.HasBlock(b.Height){
		return fmt.Errorf("chain already contains block (%d) with hash (%s)",b.Height,b.Hash(BlockHasher{}))
	}

	if err := b.Verify(); err != nil{
		return err		
	}
	return nil
}

// TODO: learn interface, rpc, make chan, struct, and other golang features like 