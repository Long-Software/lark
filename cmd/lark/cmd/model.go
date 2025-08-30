/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Long-Software/Bex/apps/cmd/lark/cmd/models"
	"github.com/spf13/cobra"
)

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to File CLI!")
	},
}


var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show existing models",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing all models...")
		// Your logic here
	},
}

func init() {
	rootCmd.AddCommand(modelCmd)
	// modelCmd.Flags().StringP("name", "n", "example", "Model name")
	// modelCmd.Flags().StringP("ext", "e", "go", "File extension")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	modelCmd.AddCommand(models.MakeCmd)
	modelCmd.AddCommand(models.InitCmd)
	// modelCmd.AddCommand(initCmd)
	// modelCmd.AddCommand(showCmd)
}
