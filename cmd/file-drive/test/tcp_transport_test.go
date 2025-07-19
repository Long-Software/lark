package test

import (
	"testing"

	"github.com/Long-Software/lark/cmd/file-drive/p2p"
	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	opts := p2p.TCPTransportOpts{
		Addr: ":4000",
	}
	ttr := p2p.NewTCPTransport(opts)
	assert.Nil(t, ttr.Listen())
}
