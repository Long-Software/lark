package main

import (
	"fmt"
	"os"
	"github.com/Long-Software/Sonality/internal/consts"
	"github.com/Long-Software/Sonality/cmd/rename/remove"
	"time"
)

type Options struct {
	Add    string
	Remove string
}

var opt = Options{
	Add:    "add",
	Remove: "remove",
}


func main() {
	start := time.Now()
	var c consts.CommandInterface
	switch os.Args[1] {
	case opt.Remove:
		c = &remove.RemoveCommand{}
	}
	c.Run()
	elapsed := time.Since(start)
	fmt.Printf("Finished in %v ms\n", elapsed.Milliseconds())
}
