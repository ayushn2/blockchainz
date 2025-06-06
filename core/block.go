package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"io"

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

func (h *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, &h.Version); err != nil{
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.PrevHash); err != nil{
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Timestamp); err != nil{
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Height); err != nil{
		return err
	}
	
	return binary.Write(w,binary.LittleEndian, &h.Nonce)
}


func (h *Header) DecodeBinary(r io.Reader) error{
	if err := binary.Read(r, binary.LittleEndian, &h.Version); err != nil{
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.PrevHash); err != nil{
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Timestamp); err != nil{
		return err
	}
	if err := binary.Read(r, binary.LittleEndian, &h.Height); err != nil{
		return err
	}
	
	return binary.Read(r,binary.LittleEndian, &h.Nonce)
	
}

type Block struct{
	Header Header
	Transactions []Transaction

	// cachd version of the header hash
	hash types.Hash // hash of the block, can be calculated from header and transactions
}

func (b *Block) Hash() types.Hash {
	buf := &bytes.Buffer{}
	b.Header.EncodeBinary(buf)

	if b.hash.IsZero() {
		b.hash = types.Hash(sha256.Sum256(buf.Bytes()))
	}

	
	return b.hash
}

func (b *Block) EncodeBinary(w io.Writer) error{
	if err := b.Header.EncodeBinary(w); err != nil {
		return err
	}

	for _, tx := range b.Transactions {
		if err := tx.EncodeBinary(w); err != nil {
			return err
		}
	}
	return nil

}

func (b *Block) DecodeBinary(r io.Reader) error{
	if err := b.Header.DecodeBinary(r); err != nil {
		return err
	}

	for _, tx := range b.Transactions {
		if err := tx.DecodeBinary(r); err != nil {
			return err
		}
	}
	return nil
}