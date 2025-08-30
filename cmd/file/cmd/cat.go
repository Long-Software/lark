/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Long-Software/Bex/packages/errors"
	"github.com/spf13/cobra"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Output the content in a file",
	Long:  `The command accept a file path as input and output the content inside the file.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("file")
		file, err := os.Open(path)
		if err != nil {
			fmt.Println(errors.FileError{Error: err})
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(errors.FileError{Error: err})
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(catCmd)
	catCmd.Flags().StringP("file", "f", "", "Path to file")
}
