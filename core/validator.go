package core

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
	return nil
}

// TODO: learn interface, rpc, make chan, struct, and other golang features like 