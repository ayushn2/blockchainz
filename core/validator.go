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

	if b.Height != v.bc.Height()+1 {
		return fmt.Errorf("block height (%d) is not equal to the current chain height (%d)", b.Height, v.bc.Height()+1)
	}

	prevHeader , err := v.bc.GetHeader(b.Height - 1)

	if err != nil {
		return err
	}

	hash := BlockHasher{}.Hash(prevHeader)
	
	if hash != b.PrevHash {
		return fmt.Errorf("block (%d) has invalid previous hash, expected (%s), got (%s)", b.Height, b.PrevHash, hash)
	}

	if err := b.Verify(); err != nil{
		return err		
	}
	return nil
}

// TODO: learn interface, rpc, make chan, struct, and other golang features like 