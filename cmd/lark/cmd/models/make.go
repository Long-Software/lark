package models

import (
	"fmt"

	"github.com/Long-Software/Bex/packages/errors"
	"github.com/Long-Software/Bex/packages/file"
	"github.com/spf13/cobra"
)

var MakeCmd = &cobra.Command{
	Use:   "make",
	Short: "Generate a new model",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		ext, _ := cmd.Flags().GetString("ext")
		model := &ModelCmdMake{Name: name, Ext: ext}
		model.Run()
	},
}

func (m *ModelCmdMake) Run() {
	model := NewModelFromExt(m.Ext, m.Name)
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
	MakeCmd.Flags().StringP("name", "n", "example", "Model name")
	MakeCmd.Flags().StringP("ext", "e", "go", "File extension")
}
