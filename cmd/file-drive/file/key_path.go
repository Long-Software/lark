package file

import (
	"fmt"
	"strings"
)

type KeyPath struct {
	Folder   string
	Filename string
}

func (p *KeyPath) Path() string {
	return fmt.Sprintf("%s/%s", p.Folder, p.Filename)
}

func (p *KeyPath) Root() string {
	paths := strings.Split(p.Path(), "/")
	if len(paths) == 0 {
		return ""
	}
	return paths[0]
}
