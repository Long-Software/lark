package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "lark",
	Short: "A CLI tool for working with files",
	Long:  "file CLI helps with handling file and creating file ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to File CLI!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
