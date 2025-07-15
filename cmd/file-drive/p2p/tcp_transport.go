package p2p

import (
	"fmt"
	"net"
	"sync"

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
	utils.NewLog(log.INFO, fmt.Sprintf("New peer has been created connected to %v with outbound: %v", conn, outbound))
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	addr     string
	listener net.Listener

	decoder Decoder

	mu    sync.RWMutex // Read and write mutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(addr string) *TCPTransport {
	return &TCPTransport{
		addr: addr,
	}
}

func (t *TCPTransport) Listen() error {
	var err error
	t.listener, err = net.Listen("tcp", t.addr)
	utils.NewLog(log.INFO, "Start listening")
	if err != nil {
		return err
	}

	go t.startAccept()
	return nil
}

func (t *TCPTransport) startAccept() {
	utils.NewLog(log.INFO, "Start Accepting connection")
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			utils.NewLog(log.ERROR, err.Error())
		}
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	utils.NewLog(log.INFO, "Start handling connection")
	peer := NewTCPPeer(conn, true)
	utils.NewLog(log.INFO, fmt.Sprintf("New incoming connection: %v", peer))
	if err := t.handshake(conn); err != nil {
		utils.NewLog(log.ERROR, err.Error())
	}

	for {
		err := t.decoder.Decode(conn, "")
		if err != nil {
			utils.NewLog(log.ERROR, err.Error())
		}
	}
}

func (t *TCPTransport) handshake(conn net.Conn) error {
	return nil
}
