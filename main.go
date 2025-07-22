package main

import (
	"bytes"
	"math/rand"
	"strconv"
	"time"

	"github.com/ayushn2/blockchainz/core"
	"github.com/ayushn2/blockchainz/crypto"
	"github.com/ayushn2/blockchainz/network"
	"github.com/sirupsen/logrus"
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
			if err:= sendTransaction(trRemote, trLocal.Addr()); err != nil {
				logrus.Error(err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	privKey := crypto.GeneratePrivateKey()

	opts := network.ServerOpts{
		PrivateKey: &privKey,
		ID: "LOCAL",
		Transports: []network.Transport{trLocal},
	}

	s := network.NewServer(opts)
	s.Start()
}

func sendTransaction(tr network.Transport, to network.NetAddr) error {
	// Send a transaction to the transport
	privateKey := crypto.GeneratePrivateKey()
	data := []byte(strconv.FormatInt(int64(rand.Intn(1000)), 10))
	tx := core.NewTransaction(data)
	tx.Sign(privateKey)
	buf := &bytes.Buffer{}
	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		return err
	}
	msg := network.NewMessage(network.MessageTypeTxn, buf.Bytes())

	return tr.SendMessage(to, msg.Bytes())
}