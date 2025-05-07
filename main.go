package main

import (
	"time"

	"github.com/ayushn2/blockchainz/network"
)

// Server
// Transport => tcp, udp
// Block
// TX

func main() {
	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")
	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)
	// Send a message from LOCAL to REMOTE

	go func(){
		for {
			// Simulate sending messages
			// local will send message to remote every second in a seperate go routine
			trRemote.SendMessage(trLocal.Addr(), []byte("Hello from REMOTE"))
			time.Sleep(1 * time.Second)
		}
	}()

	opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}