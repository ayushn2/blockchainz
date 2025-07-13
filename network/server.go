package network

import (
	"fmt"
	"time"

	"github.com/ayushn2/blockchainz/core"
	"github.com/ayushn2/blockchainz/crypto"
	"github.com/sirupsen/logrus"
)

var defaultBlockTime = 5 * time.Second // Default block time if not provided

type ServerOpts struct{
	Transports []Transport
	BlockTime time.Duration
	PrivateKey *crypto.PrivateKey
}

type Server struct {
	ServerOpts
	blockTime time.Duration // the time after which server needs to consume the mempool and put it in the block and sign it
	memPool *TxPool // Memory pool for transactions
	isValidator bool // Indicates if the server/node is a validator
	rpcCh chan RPC
	quitch chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	if opts.BlockTime == time.Duration(0){
		opts.BlockTime = defaultBlockTime
	}
	return &Server{
		ServerOpts: opts,
		blockTime: opts.BlockTime,
		memPool: NewTxPool(), // Initialize a new transaction pool
		isValidator: opts.PrivateKey != nil, // If a private key is provided, this server/node is a validator
		rpcCh: make(chan RPC), 
		quitch: make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(s.blockTime)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			// Handle incoming RPC
			fmt.Printf("Received RPC from %s: %s\n", rpc.From, string(rpc.Payload))
		case <-s.quitch:
			break free
		case <-ticker.C:
			// Check if the server is a validator and if there are transactions in the mempool
			//this is just consensus placeholder for now
			if s.isValidator && s.memPool.Len() > 0 {
			s.createNewBlock()
			}
		}
	}

	fmt.Println("Server shutting down...")
}

func (s *Server) handleTransaction(tx *core.Transaction) error {
	if err := tx.Verify(); err != nil {
		return fmt.Errorf("failed to verify transaction: %w", err)
	}

	hash := tx.Hash(core.TxHasher{})

	if s.memPool.Has(hash){
		logrus.WithFields(logrus.Fields{
			"hash": hash,
		}).Info("mempool already contains this transaction")

		return nil
	}
	
	logrus.WithFields(logrus.Fields{
		"hash": tx.Hash(core.TxHasher{}),
	}).Info("adding new transaction to mempool")

	if err := s.memPool.Add(tx); err != nil {
		return fmt.Errorf("failed to add transaction to mempool: %w", err)
	}

	fmt.Printf("Transaction added to mempool: %s\n", tx.Hash(core.TxHasher{}))
	return nil
}

func (s *Server) createNewBlock() error{
	fmt.Println("creating a new block...")
	return nil
}

func (s *Server) initTransports() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {//keep consuming the message channels
				// Handle incoming RPC
				// For example, you might want to process the message or forward it
				// to another transport
				fmt.Printf("Received message from %s: %s\n", rpc.From, string(rpc.Payload))
				s.rpcCh <- rpc // Forward the RPC to the server's main channel
			}
		}(tr)
	}
}