package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "securerandom",
	Short: "A CLI tool for generating secure random values",
	Long:  "Securerandom is a CLI application that helps generate secure random numbers and strings.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to SecureRandom CLI!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
