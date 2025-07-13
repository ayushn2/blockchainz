package core

import (
	"fmt"

	"github.com/ayushn2/blockchainz/crypto"
	"github.com/ayushn2/blockchainz/types"
)

type Transaction struct{
	Data []byte//any type of data can be stored in a transaction, as this is a generic blockchain

	From crypto.PublicKey // public key of the sender
	Signature *crypto.Signature // signature of the transaction by the sender

	hash types.Hash // hash of the transaction, computed from Data
	firstSeen int64 // timestamp when the transaction was first seen locally
}

func NewTransaction(data []byte) *Transaction {
	return &Transaction{
		Data: data,
	}
}

func (tx *Transaction) Hash(hasher Hasher[*Transaction]) types.Hash{
	if tx.hash.IsZero() {
		tx.hash = hasher.Hash(tx)
	}
	return hasher.Hash(tx)
}

func (tx *Transaction) Sign(privKey crypto.PrivateKey) error{
	sig, err := privKey.Sign(tx.Data)
	if err != nil {
		return err
	}

	tx.From = privKey.PublicKey()
	tx.Signature = sig

	return nil
}

func (tx *Transaction) Verify() error{
	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	if !tx.Signature.Verify(tx.From, tx.Data) {
		return fmt.Errorf("invalid transaction signature")
	}

	return nil
	
}

func (tx *Transaction) Decode(dec Decoder[*Transaction]) error {
	return dec.Decode(tx)
}

func (tx *Transaction) Encode(enc Encoder[*Transaction]) error {
	return enc.Encode(tx)
}

func (tx *Transaction) SetFirstSeen(timestamp int64) {
	tx.firstSeen = timestamp
}

func (tx *Transaction) FirstSeen() int64 {
	return tx.firstSeen
}