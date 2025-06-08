package types

import "encoding/hex"

type Address [20]uint8

func (a Address) ToSlice() []byte{
	b := make([]byte, 20)

	for i := 0; i<20; i++ {
		b[i] = a[i]
	}

	return b
}

func (a Address) String() string {
	return hex.EncodeToString(a.ToSlice())
}

func AddressFromBytes(b []byte) Address {
	if len(b) != 20 {
		panic("Address must be 20 bytes")
	}

	var a Address
	for i := 0; i < 20; i++ {
		a[i] = b[i]
	}

	return a
}