package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyPair_Sign_Verify_Success(t *testing.T){
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()

	msg := []byte("Hello, Blockchainz!")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	assert.True(t, sig.Verify(pubKey, msg), "Signature verification failed")
}

func TestKeyPair_Sign_Verify_Fail(t *testing.T){
	privKey := GeneratePrivateKey()
	

	msg := []byte("Hello, Blockchainz!")
	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	attackPrivKey := GeneratePrivateKey()
	attackPubKey := attackPrivKey.PublicKey()

	assert.False(t, sig.Verify(attackPubKey, msg), "Attack successfully verified a signature that should not match")
	assert.False(t, sig.Verify(privKey.PublicKey(), []byte("Tampered message")), "Signature verification should fail for tampered message")
}