# 🧱 blockchainz

## 📚 Purpose

This project is a learning-oriented blockchain implementation in Go, built with the goal of understanding and documenting blockchain internals through real, working code. It combines hands-on development with detailed educational guides, making it valuable for:

- Developers new to blockchain who want to **learn by building**
- Students or researchers exploring **consensus, cryptography, and networking**
- Contributors seeking to understand **core blockchain mechanisms** without the complexity of production networks

The project is ongoing and aims to evolve into a complete educational resource with modular code and topic-specific guides.

## 📦 Components (So Far)

- `block.go`: Defines the `Block` structure, header separation, and logic for serialization and verification.
- `transaction.go`: Handles transaction structure, signing, and verification using public/private keys.
- `blockchain.go`: Maintains the chain of blocks, tracks height, and stores headers.
- `validator.go`: Contains logic to validate blocks (height, previous hash, signatures, etc.).
- `hasher.go`: Implements reusable hash logic with a `Hasher` interface and SHA-256.
- `encoding.go`: Binary encoder/decoder helpers for low-level data serialization.
- `crypto/`: Custom cryptographic utilities for keypair generation and ECDSA-based signatures.
- `test/`: Contains unit tests to verify correctness of blocks, transactions, and validation logic.

## 📘 Guides

## 📘 Guides

- [Guide 0: Communication System](https://github.com/ayushn2/blockchainz/blob/main/guide/0.%20communication_system.md)  
  Goroutines, TCP networking, and how peers exchange messages.

- [Guide 1: Block and Header](https://github.com/ayushn2/blockchainz/blob/main/guide/1.%20block_and_header.md)  
  Structure of a block and header, encoding with `encoding/binary`, and design decisions.

- [Guide 2: Keypairs and Digital Signatures](https://github.com/ayushn2/blockchainz/blob/main/guide/2.%20keypairs_and_digital_signatures.md)  
  Cryptographic key generation, signing data, and verifying signatures using ECDSA.

- [Guide 3: Block and Transaction Signing and Verification](https://github.com/ayushn2/blockchainz/blob/main/guide/3.%20block_and_transaction_signing_and_verification.md)  
  Signing headers and transactions, verifying authenticity, and abstracting with helper methods.

- [Guide 4: The Blockchain Data Structure](https://github.com/ayushn2/blockchainz/blob/main/guide/4.%20the_blockchain_data_structure.md)  
  Map-based blockchain implementation, adding blocks, storing headers, and retrieving data.

- [Guide 5: Block Validator](https://github.com/ayushn2/blockchainz/blob/main/guide/5.%20block_validator.md)  
  Block validation logic (height, hash, signatures), current testing logic, and theoretical responsibilities of validators.

This project is accompanied by a series of developer-friendly, educational guides that explain blockchain fundamentals through real Go code:

- [Guide 0: Intro to Networking & Goroutines](./docs/guide1_networking_goroutines.md)  
  Learn how Go’s concurrency model and goroutines support peer-to-peer networking.

- [Guide 1: Blocks, Headers & Encoding](./docs/guide2_blocks_headers_encoding.md)  
  Understand the structure of a block, separation of headers, and how binary encoding works.

- [Guide 3: Channels in Go](./docs/guide3_channels.md)  
  A beginner-friendly deep dive into Go channels and how they support concurrent message flow.

- [Guide 4: Signing & Verification](./docs/guide4_signing_verification.md)  
  Explore how transactions and blocks are signed and verified using public-key cryptography.

- [Guide 5: Block Validation](./docs/guide5_block_validation.md)  
  Covers block validation logic, validator role, and what’s implemented vs. what’s theoretical.

## 🚧 Work in Progress

This project is actively evolving and currently focuses on building a foundational blockchain core. Areas like validation, transaction signing, and networking are progressively being implemented. While not production-ready, the repository is structured to support gradual learning and modular expansion as features mature.

## 🤝 Contributions

This project is intended as an educational resource for learning blockchain internals through Go.

You're welcome to:
- Explore the code and use it for your own learning or experiments
- Read through the guides for conceptual clarity
- Open issues to ask questions, suggest improvements, or report bugs
- Fork the repo to build your own version

> 📌 Direct pushes to the main repository are disabled — please use forks and pull requests for any contributions.
