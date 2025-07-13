package core

import (
	"crypto/elliptic"
	"encoding/gob"
	"io"
)

type Encoder[T any] interface {
	Encode(T) error

}

type Decoder[T any] interface {
	Decode(T) error
}

// GobTxEncoder is an encoder for transactions using the gob encoding format.

type GobTxEncoder struct{
	w io.Writer
}

func NewGobTxEncoder(w io.Writer) *GobTxEncoder {
	// Register elliptic.P256 to ensure it can be encoded properly
	// when encoding transactions that contain public keys.
	gob.Register(elliptic.P256())
	return &GobTxEncoder{w: w}
}

func (e *GobTxEncoder) Encode(tx *Transaction) error {
	enc := gob.NewEncoder(e.w)
	return enc.Encode(tx)
}

// GobTxDecoder is a decoder for transactions using the gob encoding format.
type GobTxDecoder struct{
	r io.Reader
}

func NewGobTxDecoder(r io.Reader) *GobTxDecoder {
	// Register elliptic.P256 to ensure it can be decoded properly
	// when decoding transactions that contain public keys.
	gob.Register(elliptic.P256())
	return &GobTxDecoder{r: r}
}	
func (d *GobTxDecoder) Decode(tx *Transaction) error {
	dec := gob.NewDecoder(d.r)
	return dec.Decode(tx)
}	