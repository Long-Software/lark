package p2p

// Interface for the remote node
type Peer interface {
	Close() error
}

// Handles the communication between the nodes in the network
type Transport interface {
	Listen() error
	Consume() <-chan RPC
}
