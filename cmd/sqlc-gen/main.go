package main

import (
	"fmt"
	"os"
	"github.com/Long-Software/Sonality/internal/consts"
	"github.com/Long-Software/Sonality/cmd/sqlc-gen/crud"
	"time"
)

type Options struct {
	CRUD string
}

var opt = Options{
	CRUD: "crud",
}

func main() {
	start := time.Now()
	var c consts.CommandInterface
	switch os.Args[1] {
	case opt.CRUD:
		c = &crud.CRUDCommand{}
	default:
		c = &crud.CRUDCommand{}
		c.Help()
		return
	}
	c.Run()
	elapsed := time.Since(start)
	fmt.Printf("Finished in %v ms\n", elapsed.Milliseconds())
}
