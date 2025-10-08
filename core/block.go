package core

import (
	"bytes"
	"encoding/gob"
	"crypto/sha256" 
	"fmt"
	"io"
	"github.com/ayushn2/blockchainz/crypto"
	"github.com/ayushn2/blockchainz/types"
)

// Hash, prev hash, timestamp, nonce, transactions
type Header struct {
	Version	uint32
	DataHash types.Hash
	PrevHash types.Hash
	Timestamp uint64
	Height uint32
	
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

	for i := 0; i < len(b.Transactions); i++ {
    tx := b.Transactions[i]
    if err := tx.Verify(); err != nil {
        return err
    }
}
	dataHash, err := CalculateDataHash(b.Transactions)

	if err != nil{
		return fmt.Errorf("failed to calculate data hash: %w", err)
	}

	if dataHash != b.DataHash {
		return fmt.Errorf("block (%s) has invalid data hash", b.Hash(BlockHasher{}))
	}

	return nil
}

func (b *Block) Decode(r io.Reader, dec Decoder[*Block]) error{
	return dec.Decode(b)
}

func (b *Block) Encode(r io.Writer, enc Encoder[*Block]) error{
	return enc.Encode(b)
}

// Hash computes the hash of the block using the provided hasher.
// Hasher[*Block] means the hasher works specifically with *Block.
func (b *Block) Hash(hasher Hasher[*Header]) types.Hash{
	if b.hash.IsZero() {
		b.hash = hasher.Hash(b.Header)
	}

	return b.hash
}

func CalculateDataHash(txx []Transaction)(hash types.Hash,err error){
	
		buf := &bytes.Buffer{}
		
	
	

	for i := 0; i< len(txx); i++ {
		tx := txx[i]
		if err = tx.Encode(NewGobTxEncoder(buf)); err != nil {
			return 
		}
	}
	hash = sha256.Sum256(buf.Bytes())
	return 
}