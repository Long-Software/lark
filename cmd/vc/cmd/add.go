package cmd

import (
	"os/exec"

	"github.com/Long-Software/Bex/cmd/vc/internal/logger"
	"github.com/Long-Software/Bex/packages/log"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "git add",
	Long:  "a similar command to git add",
	Run: func(c *cobra.Command, args []string) {
		// TODO: add "add" flag to the command
		cmd := exec.Command("git", args...)
		err := cmd.Run()
		if err != nil {
			logger.New(log.ERROR, err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
