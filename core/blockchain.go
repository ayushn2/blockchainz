package core

// The Blockchain is a state machine, that transitions from one state to another
// The genesis block is the initial state of the blockchain

type Blockchain struct{
	store Storage
	headers []*Header
	validator Validator
}

func NewBlockchain(genesis *Block) (*Blockchain, error) {
	bc := &Blockchain{
		headers: []*Header{},
		store: NewMemoryStorage(),
	}
	bc.validator = NewBlockValidator(bc)

	err := bc.addBlockWithoutValidation(genesis)

	// return &Blockchain{
	// 	store: store,
	// 	headers: []*Header{},
	// 	validator: NewBlockValidator(bc),
	// }

	return  bc, err
}

func (bc *Blockchain) SetValidator(v Validator){
	bc.validator = v
}

func (bc *Blockchain) AddBlock(b *Block) error{
	// validate
	if err := bc.validator.ValidateBlock(b); err != nil {
		return err
	}

	return bc.addBlockWithoutValidation(b)
}

func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

func (bc *Blockchain) Height() uint32{
	return uint32(len(bc.headers) -1)
}

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error{
	bc.headers = append(bc.headers, b.Header)
	return bc.store.Put(b)
}