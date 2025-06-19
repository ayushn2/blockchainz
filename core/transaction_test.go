package core

import (
	"testing"
	"github.com/ayushn2/blockchainz/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSignTransaction(t *testing.T){
	privKey := crypto.GeneratePrivateKey()
	data :=[]byte("test data")
	tx := &Transaction{
		Data: data,
	}

	assert.Nil(t, tx.Sign(privKey), "Successfully signed transaction")
	assert.NotNil(t, tx.Signature, "Signature should not be nil after signing")
}

func TestVerifyTransaction(t *testing.T){
	privKey := crypto.GeneratePrivateKey()
	data :=[]byte("test data")
	tx := &Transaction{
		Data: data,
	}

	assert.Nil(t, tx.Sign(privKey), "Successfully signed transaction")
	assert.NotNil(t, tx.Signature, "Signature should not be nil after signing")

	assert.Nil(t, tx.Verify(), "Transaction verification should succeed")

	otherPrivKey := crypto.GeneratePrivateKey()
	tx.PubKey = otherPrivKey.PublicKey()

	assert.NotNil(t, tx.Verify(), "Transaction should not verify with other public key")

}