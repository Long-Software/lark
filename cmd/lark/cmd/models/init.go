package models

import (
	"fmt"

	"github.com/Long-Software/Bex/packages/errors"
	"github.com/Long-Software/Bex/packages/file"
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the model directory or environment",
	Run: func(cmd *cobra.Command, args []string) {
		ext, _ := cmd.Flags().GetString("ext")
		model := &ModelCmdInit{Ext: ext}
		model.Run()
	},
}

func (m *ModelCmdInit) Run() {
	model := NewModelFromExt(m.Ext, "model")
	err := file.Create(model.FilePath())
	if err != nil {
		fmt.Println(errors.FileError{Error: err})
		return
	}
	err = file.Write(model.FilePath(), model.Content())
	if err != nil {
		fmt.Println(errors.FileError{Error: err})
		return
	}
}
func init() {
	InitCmd.Flags().StringP("ext", "e", "go", "File extension")
}
