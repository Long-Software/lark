package main

import (
	"fmt"
	"os"
	"github.com/Long-Software/Sonality/internal/consts"
	"github.com/Long-Software/Sonality/cmd/pdf/create"
	"time"
)

type Options struct {
	Make string
	Remove string
}

var opt = Options{
	Make: "make",
	Remove: "remove",
}

func main() {
	start := time.Now()
	var c consts.CommandInterface
	switch os.Args[1] {
	case opt.Make:
		c = &create.CreateCommand{}
	case opt.Remove:
		
	default:
		c = &DefaultPdfCommand{}
		c.Help()
		return
	}
	c.Run()
	elapsed := time.Since(start)
	fmt.Printf("Finished in %v ms\n", elapsed.Milliseconds())
}

type DefaultPdfCommand struct{}

func (d *DefaultPdfCommand) Run() {}

func (d *DefaultPdfCommand) Help() {
	c := &create.CreateCommand{}
	c.Help()
}
