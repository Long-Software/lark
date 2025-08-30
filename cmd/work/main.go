package main

import (
	"fmt"
	"os"
	"github.com/Long-Software/Sonality/internal/consts"
	"github.com/Long-Software/Sonality/cmd/work/sync"
	"github.com/Long-Software/Sonality/cmd/work/use"
	"time"
)

type Options struct {
	Sync string
	Use  string
}

var opt = Options{
	Sync: "sync",
	Use:  "use",
}

func main() {
	start := time.Now()
	var c consts.CommandInterface
	switch os.Args[1] {
	case opt.Sync:
		c = &sync.SyncCommand{}
	case opt.Use:
		c = &use.UseCommand{}
	default:
		c = &sync.SyncCommand{}
		c.Help()
		c = &use.UseCommand{}
		c.Help()
		return
	}
	c.Run()
	elapsed := time.Since(start)
	fmt.Printf("Finished in %v ms\n", elapsed.Milliseconds())
}
