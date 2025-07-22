package network

import (
	"bytes"
	"fmt"
	"os"

	"time"

	"github.com/ayushn2/blockchainz/core"
	"github.com/ayushn2/blockchainz/crypto"
	"github.com/go-kit/log"
	"github.com/sirupsen/logrus"
)

var defaultBlockTime = 5 * time.Second // Default block time if not provided

type ServerOpts struct{
	ID string // Unique identifier for the server
	Logger log.Logger
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

	if opts.Logger == nil {
		opts.Logger = log.NewLogfmtLogger(os.Stderr)
		opts.Logger = log.With(opts.Logger, "ID", opts.ID)
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

	if s.isValidator {
		go s.validatorLoop() // Start the validator loop if this server is a validator
	}
	return s
}

func (s *Server) Start() {
	s.initTransports()

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			msg, err := s.RPCDecodeFunc(rpc)
			if err != nil {
				s.Logger.Log(
					"error", err,
				)
			}

			if err := s.RPCProcessor.ProcessMessage(msg); err != nil {
				logrus.Error(err)
			}
		case <-s.quitch:
			break free
		
		}
	}
	
	s.Logger.Log("msg", "server is shutting down")
	
}

func (s *Server) validatorLoop(){
	ticker := time.NewTicker(s.BlockTime)
	
	s.Logger.Log("msg", "starting validator loop", "blockTime", s.BlockTime)

	for {	
		<- ticker.C 
		s.createNewBlock()
	}
}

func (s *Server) ProcessMessage(msg *DecodeMessage) error {
	switch t := msg.Data.(type) {
	case *core.Transaction:
		return s.processTransaction(t)
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

func (s *Server) processTransaction(tx *core.Transaction) error {

	hash := tx.Hash(core.TxHasher{})

	if s.memPool.Has(hash){
		return nil
	}

	if err := tx.Verify(); err != nil {
		return fmt.Errorf("failed to verify transaction: %w", err)
	}

	tx.SetFirstSeen(time.Now().UnixNano())

	logrus.WithFields(logrus.Fields{
		"hash": tx.Hash(core.TxHasher{}),
	}).Info("adding new transaction to mempool")

		s.Logger.Log(
			"msg", "adding new transaction to mempool",
			"hash",hash,
			"mempool_length",s.memPool.Len())

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