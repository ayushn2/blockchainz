# 📡 Communication Layer for a Decentralized System

This document explains the architecture and theoretical foundation of a basic peer-to-peer communication system built in Go. It covers the use of transports, channels, goroutines, and message passing as a base layer for decentralized systems like blockchains or distributed storage systems.

---

## 🧱 1. Overview

A **decentralized system** requires nodes to communicate directly with each other without relying on a central authority. This communication layer serves as the foundation by handling:

- Peer-to-peer messaging  
- Abstraction over transport protocols  
- Concurrent message handling  
- Graceful startup and shutdown

The architecture is modular and extensible, enabling the addition of features like encryption, discovery, and consensus on top of the core message system.

---

## 💡 2. Server Abstraction

 The `Server` struct is the **central coordinator** of all network events. It manages message routing via channels and coordinates different transports.

```go
type Server struct {
	ServerOpts
	rpcCh   chan RPC         // Channel for internal RPC message handling
	quitch  chan struct{}    // Signal channel to shut down the server
}

- 	rpcCh: used for internal routing of messages between peers.
-   quitch: allows the server to be shut down cleanly.
-   ServerOpts: contains the list of transports to use.

### ✅ Creating a New Server

```go
func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts: opts,
		rpcCh:      make(chan RPC),
		quitch:     make(chan struct{}, 1),
	}
}```

This creates a new Server with initialized channels and transport options, making it ready to handle peer communications.

---

## 🚚 3. Transport Layer Abstraction

A `Transport` is an interface that abstracts how messages are sent and received — it could be TCP, UDP, or a mock in-memory transport like `LocalTransport`.

### 🔌 Interface Definition

```go
type Transport interface {
	Addr() string
	SendMessage(to string, payload []byte) error
	Consume() <-chan RPC
}```


Each Transport is responsible for:
- 	Defining a node address (Addr)
-	Sending messages to a target address (SendMessage)
-	Receiving incoming messages via a channel (Consume)

This abstraction allows flexibility in choosing how nodes communicate without modifying the server logic.

---

## ⚙️ 4. `initTransports()` – Concurrent Message Listening

The `initTransports()` method initializes all transports by starting a separate goroutine for each transport to listen for incoming messages concurrently.

```go
func (s *Server) initTransports() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				// Handle incoming RPC messages from this transport
				fmt.Printf("Received message from %s: %s\n", rpc.From, string(rpc.Payload))
				// Forward the message to the server's internal RPC channel for processing
				s.rpcCh <- rpc
			}
		}(tr)
	}
}```

-   Each transport’s Consume() returns a channel from which messages can be read.
-   A dedicated goroutine per transport allows the server to listen to multiple transports simultaneously without blocking.
-   Messages received on transport channels are forwarded to the server’s central rpcCh channel for unified handling.

This design enables scalable and asynchronous message handling, crucial for decentralized network communication.

---

##  ▶️ 5. `Start()` Method – Server’s Main Event Loop

The `Start()` method runs the server’s main loop, which listens for incoming messages, shutdown signals, and periodic status updates.

```go
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
}```

### The server listens concurrently on:
-   rpcCh for incoming messages,
-   quitch to detect shutdown signals,
-   ticker.C for periodic “heartbeat” logs.
### break free exits the infinite loop on shutdown signal.
### The final print statement confirms graceful shutdown.

This event-driven loop is essential for maintaining responsive, concurrent communication in decentralized systems.