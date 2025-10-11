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
	// when decoding transactions that contain public keys.(done in init())
	return &GobTxDecoder{r: r}
}	
func (d *GobTxDecoder) Decode(tx *Transaction) error {
	dec := gob.NewDecoder(d.r)
	return dec.Decode(tx)
}	

type GobBlockEncoder struct{
	w io.Writer
}

func NewGobBlockEncoder(w io.Writer) *GobBlockEncoder{
	return &GobBlockEncoder{w: w}
}

func (enc *GobBlockEncoder) Encode(b *Block) error{
	return gob.NewEncoder(enc.w).Encode(b)
}

type GobBlockDecoder struct{
	r io.Reader
}

func NewGobBlockDecoder(r io.Reader) *GobBlockDecoder{
	return &GobBlockDecoder{r: r}
}

func (dec *GobBlockDecoder) Decode(b *Block) error{
	return gob.NewDecoder(dec.r).Decode(b)
}

// Ensure elliptic.P256 is registered with gob on package initialization.
// init() is called automatically when the package is imported.
func init() {
	gob.Register(elliptic.P256())
}