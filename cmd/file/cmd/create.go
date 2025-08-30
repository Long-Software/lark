/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Long-Software/Bex/apps/cmd/file/template"
	"github.com/Long-Software/Bex/packages/errors"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a empty file",
	Long:  `Create a new empty file from the given file path.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		tem, _ := cmd.Flags().GetString("template")
		dirPath := filepath.Dir(name)
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println(errors.FileError{Error: err})
			return
		}
		file, err := os.Create(name)
		if err != nil {
			fmt.Println(errors.FileError{Error: err})
			return
		}
		defer file.Close()

		if tem != "" {
			_, err = file.WriteString(template.NewTemplate(filepath.Ext(name), tem))
			if err != nil {
				fmt.Println(errors.FileError{Error: err})
				return
			}
		}
		fmt.Println("File created successfully at:", name)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("name", "n", "example.txt", "Path to a file. (default=\"example.txt\")")
	createCmd.Flags().StringP("template", "t", "", "init|model|repo (default=\"init\")")
}
