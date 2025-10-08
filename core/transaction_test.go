package core

import (
	"bytes"
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
	tx.From = otherPrivKey.PublicKey()

	assert.NotNil(t, tx.Verify(), "Transaction should not verify with other public key")

}

func TestTxEncodeDecode(t *testing.T){
	tx := randomTxWithSignature(t)
	buf := &bytes.Buffer{}
	assert.Nil(t, tx.Encode(NewGobTxEncoder(buf)), "Transaction should encode without error")

	txDecoded := new(Transaction)
	assert.Nil(t, txDecoded.Decode(NewGobTxDecoder(buf)), "Transaction should decode without error")
	assert.Equal(t, &tx, txDecoded)
}

func randomTxWithSignature(t *testing.T) Transaction {
	privKey := crypto.GeneratePrivateKey()
	tx := Transaction{
		Data: []byte("test transaction"),
	}
	err := tx.Sign(privKey)
	assert.Nil(t, err, "Transaction should be signed successfully")
	return tx
}