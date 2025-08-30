/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hexCmd represents the hex command
var hexCmd = &cobra.Command{
	Use:   "hex",
	Short: "Generate Hex numbers",
	Long: `Provided length, it generate hex number.
	For example:
	securerandom hex -l 10`,
	Run: func(cmd *cobra.Command, args []string) {
		length, _:= cmd.Flags().GetInt("length")
		fmt.Println(length )
	},
}

func init() {
	rootCmd.AddCommand(hexCmd)

	// Here you will define your flags and configuration settings.
	hexCmd.Flags().IntP("length", "l", 4, "length of hex")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hexCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hexCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
