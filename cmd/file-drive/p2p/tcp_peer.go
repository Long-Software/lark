package p2p

import (
	"fmt"
	"net"

	"github.com/Long-Software/lark/cmd/file-drive/utils"
	"github.com/Long-Software/lark/pkg/log"
)

// the remote node over the TCP connection
type TCPPeer struct {
	// the connection between the peer
	conn net.Conn

	// if we send a connection => outbound true else false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	utils.Log.NewLog(log.INFO, fmt.Sprintf("New peer has been created connected to %v with outbound: %v", conn, outbound))
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// close the peer connection
func (p *TCPPeer) Close() error {
	return p.conn.Close()
}