package core

import (
	"fmt"

	"github.com/ayushn2/blockchainz/crypto"
)

type Transaction struct{
	Data []byte//any type of data can be stored in a transaction, as this is a generic blockchain

	PubKey crypto.PublicKey // public key of the sender
	Signature *crypto.Signature // signature of the transaction by the sender
}

func (tx *Transaction) Sign(privKey crypto.PrivateKey) ( error){
	sig, err := privKey.Sign(tx.Data)
	if err != nil {
		return err
	}

	tx.PubKey = privKey.PublicKey()
	tx.Signature = sig

	return nil
}

func (tx *Transaction) Verify() error{
	if tx.Signature == nil {
		return fmt.Errorf("transaction has no signature")
	}

	if !tx.Signature.Verify(tx.PubKey, tx.Data) {
		return fmt.Errorf("invalid transaction signature")
	}

	return nil
	
}