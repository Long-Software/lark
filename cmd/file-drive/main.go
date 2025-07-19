package main

import (
	"fmt"

	"github.com/Long-Software/lark/cmd/file-drive/p2p"
	"github.com/Long-Software/lark/cmd/file-drive/utils"
	"github.com/Long-Software/lark/pkg/log"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		Addr:    ":3000",
		Decoder: &p2p.DefaultDecoder{},
	}
	ttr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-ttr.Consume()
			utils.Log.NewLog(log.DEBUG, fmt.Sprintf("%+v", msg))
		}
	}()
	err := ttr.Listen()
	if err != nil {
		utils.Log.NewLog(log.FATAL, err.Error())
	}
	select {}
}
