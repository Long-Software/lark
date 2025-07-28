package main

import (
	"fmt"
	"os"

	arch "github.com/Long-Software/lark/cmd/analyser/internal/archaeologist"
	"github.com/Long-Software/lark/cmd/analyser/internal/utils"
	"github.com/Long-Software/lark/pkg/log"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run analyzer.go <directory>")
		fmt.Println("Example: go run analyzer.go ./src")
		os.Exit(1)
	}

	rootDir := os.Args[1]

	// Check if directory exists
	if _, err := os.Stat(rootDir); os.IsNotExist(err) {
		utils.Log.NewLog(log.FATAL, err.Error())
	}

	analyzer := arch.NewCodeArchaeologist(rootDir)
	if err := analyzer.Excavate(); err != nil {
		utils.Log.NewLog(log.FATAL, err.Error())
	}

}
