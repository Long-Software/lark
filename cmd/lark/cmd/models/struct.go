package models

import (
	"fmt"

	"github.com/Long-Software/Bex/packages/word"
)

type ModelCmdState interface {
	Run()
}

type ModelCmdMake struct {
	Name string
	Ext  string
}

type ModelCmdInit struct {
	Ext string
}

type ModelInterface interface {
	FilePath() string
	Content() string
}
type Model struct {
	Base string
	Name string
	Ext  string
	Body string
}

func (m *Model) FilePath() string {
	// TODO: add switch case to return the file path based on the diff extension
	return fmt.Sprintf("%s/%s.%s", m.Base, word.Lowercase(m.Name), m.Ext)
}
func (m *Model) Content() string {
	switch m.Ext {
	case "go":
		return fmt.Sprintf(`package models

type %sModel struct {

}`, word.Capitalize(m.Name))
	case "dart":
		return ""
	default:
		return ""
	}
}

func NewModelFromExt(ext string, name string) Model {
	switch ext {
	case "go":
		return NewGoModel(name)
	case "dart":
		return NewDartModel(name)
	default:
		return NewGoModel(name)
	}
}

func NewGoModel(name string) Model {
	return Model{
		Base: "models",
		Name: name,
		Ext:  "go",
		Body: "",
	}
}

func NewDartModel(name string) Model {
	return Model{
		Base: "lib/models",
		Name: name,
		Ext:  "dart",
		Body: "",
	}
}
