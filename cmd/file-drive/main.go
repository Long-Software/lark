package main

import (
	"github.com/Long-Software/lark/cmd/file-drive/p2p"
	"github.com/Long-Software/lark/cmd/file-drive/utils"
	"github.com/Long-Software/lark/pkg/log"
)

func main() {
	ttr := p2p.NewTCPTransport(":8080")

	err := ttr.Listen()
	if err != nil {
		utils.NewLog(log.FATAL, err.Error())
	}
	select {}
}
