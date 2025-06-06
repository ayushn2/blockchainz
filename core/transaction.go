package core

import "io"

type Transaction struct{
	Data []byte//any type of data can be stored in a transaction, as this is a generic blockchain
}

func (tx Transaction) DecodeBinary(r io.Reader) error {
	// Implement decoding logic for Transaction
	return nil
}


func (tx Transaction) EncodeBinary(w io.Writer) error {
	// Implement decoding logic for Transaction
	return nil
}