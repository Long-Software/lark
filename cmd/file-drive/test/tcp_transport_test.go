package test

import (
	"testing"

	"github.com/Long-Software/lark/cmd/file-drive/p2p"
	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	listenAddr := ":4000"
	ttr := p2p.NewTCPTransport(listenAddr)
	assert.Nil(t, ttr.Listen())

}
