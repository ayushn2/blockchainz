package network

import (
	"fmt"
	"time"
)

type ServerOpts struct{
	Transports []Transport
}

type Server struct {
	ServerOpts
	rpcCh chan RPC
	quitch chan struct{}
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		rpcCh: make(chan RPC),
		quitch: make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(5 * time.Second)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			// Handle incoming RPC
			fmt.Printf("Received RPC from %s: %s\n", rpc.From, string(rpc.Payload))
		case <-s.quitch:
			break free
		case <-ticker.C:
			fmt.Println("Server is running...")
		}
	}

	fmt.Println("Server shutting down...")
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