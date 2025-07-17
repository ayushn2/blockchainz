package network

import (
	"bytes"
	"fmt"
	"time"

	"github.com/ayushn2/blockchainz/core"
	"github.com/ayushn2/blockchainz/crypto"
	"github.com/sirupsen/logrus"
)

var defaultBlockTime = 5 * time.Second // Default block time if not provided

type ServerOpts struct{
	RPCDecodeFunc RPCDecodeFunc // Function to decode RPC messages
	RPCProcessor RPCProcessor
	Transports []Transport
	BlockTime time.Duration
	PrivateKey *crypto.PrivateKey
}

type Server struct {
	ServerOpts
	memPool *TxPool // Memory pool for transactions
	isValidator bool // Indicates if the server/node is a validator
	rpcCh chan RPC
	quitch chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	
	if opts.BlockTime == time.Duration(0){
		opts.BlockTime = defaultBlockTime
	}

	if opts.RPCDecodeFunc == nil{
		opts.RPCDecodeFunc = DefaultRPCDecodeFunc
	}

	s := &Server{
		ServerOpts: opts,
		memPool: NewTxPool(), // Initialize a new transaction pool
		isValidator: opts.PrivateKey != nil, // If a private key is provided, this server/node is a validator
		rpcCh: make(chan RPC), 
		quitch: make(chan struct{}, 1),
	}

	// if we don't have a RPCProcessor from the server options, we will use the server itself as default
	if s.RPCProcessor == nil {
		s.RPCProcessor = s
	}

	return s
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(s.BlockTime)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			msg, err := s.RPCDecodeFunc(rpc)
			if err != nil {
				logrus.Error(err)
			}

			if err := s.RPCProcessor.ProcessMessage(msg); err != nil {
				logrus.Error(err)
			}
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

func (s *Server) ProcessMessage(msg *DecodeMessage) error {
	

	switch t := msg.Data.(type) {
	case *core.Transaction:
		return s.ProcessTransaction(t)
	}
	return nil
}

func (s *Server) broadcast(msg []byte) error{
	for _, peer := range s.Transports {
		if err := peer.Broadcast(msg); err != nil {
			return fmt.Errorf("failed to broadcast message: %w", err)
		}
	}
	return nil
}

func (s *Server) ProcessTransaction(tx *core.Transaction) error {

	hash := tx.Hash(core.TxHasher{})

	if s.memPool.Has(hash){
		logrus.WithFields(logrus.Fields{
			"hash": hash,
			"mempool length": s.memPool.Len(),
		}).Info("mempool already contains this transaction")

	// TODO(@ayushn2): broadcast this tx to peers

		return nil
	}

	if err := tx.Verify(); err != nil {
		return fmt.Errorf("failed to verify transaction: %w", err)
	}

	tx.SetFirstSeen(time.Now().UnixNano())

	logrus.WithFields(logrus.Fields{
		"hash": tx.Hash(core.TxHasher{}),
	}).Info("adding new transaction to mempool")

	go s.broadcastTransaction(tx)

	return  s.memPool.Add(tx)
}

func (s *Server) broadcastTransaction(tx *core.Transaction) error {
	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		return fmt.Errorf("failed to encode transaction: %w", err)
	}

	msg := NewMessage(MessageTypeTxn, buf.Bytes())

	return s.broadcast(msg.Bytes())
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
				
				s.rpcCh <- rpc // Forward the RPC to the server's main channel
			}
		}(tr)
	}
}