package cmd

import (
	"fmt"
	"os"

	"github.com/Long-Software/Bex/cmd/vc/internal/logger"
	"github.com/Long-Software/Bex/packages/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vc",
	Short: "A CLI tool for custom git commit",
	Long:  "vc CLI help you write better and consistent Git commit messages.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to vc CLI!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.New(log.ERROR, err.Error())
		os.Exit(1)
	}
}
