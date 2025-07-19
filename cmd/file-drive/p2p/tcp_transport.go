package p2p

import (
	"errors"
	"fmt"
	"net"
	"sync"

	"github.com/Long-Software/lark/cmd/file-drive/utils"
	"github.com/Long-Software/lark/pkg/log"
)

type TCPTransport struct {
	TCPTransportOpts
	listener net.Listener
	rpc      chan RPC

	mu sync.RWMutex // Read and write mutex
	// peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		TCPTransportOpts: opts,
		rpc:              make(chan RPC),
	}
}

type TCPTransportOpts struct {
	Addr    string
	Decoder Decoder
	OnPeer  func(Peer) error
}

func (t *TCPTransport) Listen() error {
	var err error
	t.listener, err = net.Listen("tcp", t.Addr)
	utils.Log.NewLog(log.INFO, "Start listening")
	if err != nil {
		return err
	}

	go t.startAccept()
	return nil
}

// return a read-only channel from the remote peer
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpc
}

func (t *TCPTransport) startAccept() {
	utils.Log.NewLog(log.INFO, "Start Accepting connection")
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			utils.Log.NewLog(log.ERROR, err.Error())
		}
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	var err error
	utils.Log.NewLog(log.INFO, "Start handling connection")
	peer := NewTCPPeer(conn, true)

	defer func() {
		utils.Log.NewLog(log.INFO, fmt.Sprintf("dropping peer connection: %s", err))
		conn.Close()
	}()

	if t.OnPeer != nil {
		if err = t.OnPeer(peer); err != nil {
			return
		}
	}

	utils.Log.NewLog(log.DEBUG, fmt.Sprintf("New incoming connection: %+v", peer))
	if err := t.handshake(conn); err != nil {
		conn.Close()
		utils.Log.NewLog(log.ERROR, err.Error())
		return
	}

	// read loop 
	rpc := RPC{}
	for {
		err := t.Decoder.Decode(conn, &rpc)
		if errors.Is(err, &net.OpError{}) || err == net.ErrClosed {
			return
		}
		if err != nil {
			utils.Log.NewLog(log.ERROR, err.Error())
			continue
		}
		rpc.Src = conn.RemoteAddr()
		t.rpc <- rpc
		utils.Log.NewLog(log.DEBUG, fmt.Sprintf("message: %+v", t.rpc))
	}
}

// returned when the connection failed between the local and the remote node
var ErrInvalidTCPHandshake = errors.New("Could not establish the handshake")

func (t *TCPTransport) handshake(_ net.Conn) error {
	return nil
}
