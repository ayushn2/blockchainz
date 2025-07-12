package core

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"github.com/ayushn2/blockchainz/crypto"
	"github.com/ayushn2/blockchainz/types"
)

// Hash, prev hash, timestamp, nonce, transactions
type Header struct {
	Version	uint32
	PrevHash types.Hash
	Timestamp uint64
	Height uint32
	Nonce uint64
}

func (h *Header) Bytes() []byte{
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	enc.Encode(h)

	return buf.Bytes()
}

type Block struct{
	*Header
	Transactions []Transaction
	Validator crypto.PublicKey // public key of the validator who created the block
	Signature *crypto.Signature // signature of the block header by the validator
	// Height uint32 // height of the block in the blockchain, can be used to verify the order of blocks
	// cached version of the header hash
	hash types.Hash // hash of the block, can be calculated from header and transactions
}

func NewBlock(h *Header, tx []Transaction) *Block {
	return &Block{
		Header: h,
		Transactions: tx,
	}
}

func (b *Block) AddTransaction(tx *Transaction){
	b.Transactions = append(b.Transactions, *tx)
}

func (b *Block) Sign(privKey crypto.PrivateKey) error{
	sig, err := privKey.Sign(b.Header.Bytes())
	if err != nil {
		return err
	}

	b.Validator = privKey.PublicKey()
	b.Signature = sig

	return nil
}

func (b *Block) Verify() error{
	if b.Signature == nil {
		return fmt.Errorf("block has no signature")
	}

	if !b.Signature.Verify(b.Validator, b.Header.Bytes()) {
		return fmt.Errorf("block has invalid signature")
	}

	for _, tx := range b.Transactions{
		if err := tx.Verify(); err != nil {
			return err
		}
	}

	return nil
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error{
	return dec.Decode(r, b)
}

func (b *Block) Encode(r io.Writer, enc Encoder[*Block]) error{
	return enc.Encode(r, b)
}

// Hash computes the hash of the block using the provided hasher.
// Hasher[*Block] means the hasher works specifically with *Block.
func (b *Block) Hash(hasher Hasher[*Header]) types.Hash{
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b.Header)
	}

	return b.hash
}
