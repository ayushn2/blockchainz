# 🧱 Header Encoding & Hashing

This module defines a `Header` struct used in a blockchain and implements binary encoding/decoding logic using Go's `encoding/binary` package. It also includes a 32-byte hash type with secure random generation.

## 🔹 Header

The `Header` struct includes:

- `Version` (uint32)
- `PrevHash` (types.Hash)
- `Timestamp` (uint64)
- `Height` (uint32)
- `Nonce` (uint64)

### Methods

- `EncodeBinary(io.Writer) error`: writes the header to a binary stream.
- `DecodeBinary(io.Reader) error`: reads a header from a binary stream.

## 🔸 Hash

The `types.Hash` is a fixed-size 32-byte array.

### Functions

- `RandomHash()`: generates a secure random 32-byte hash.
- `HashFromBytes([]byte)`: converts 32-byte input to a `Hash`.

## ✅ Test

Test function `TestHeader_Encode_Decode` checks:

- Whether a header can be encoded and then decoded without data loss.
- Uses `bytes.Buffer` and `testify/assert` for round-trip verification.

## 📦 Example Usage

```go
buf := &bytes.Buffer{}
h.EncodeBinary(buf)
hDecode := &Header{}
hDecode.DecodeBinary(buf)
```