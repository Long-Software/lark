package p2p

import "net"

type RPC struct {
	Src     net.Addr
	Payload []byte
}
