package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "image",
	Short: "A CLI tool for working with images",
	Long:  "image is a CLI application that helps working with image file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to image CLI!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
