package types

import "crypto/rand"

type Hash [32]uint8

func HashFromBytes(b []byte) Hash {
	if len(b) != 32 {
		panic("HashFromBytes: input must be exactly 32 bytes")
	}

	var value [32]uint8
	for i :=0 ; i<32; i++ {
		value[i] = b[i]
	}

	return Hash(value)
}

func RandomBytes(size int) []byte{
	token := make([]byte, size)
	rand.Read(token)
	return token
}

func RandomHash() Hash {
	return HashFromBytes(RandomBytes(32))
}