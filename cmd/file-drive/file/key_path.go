package file

import "fmt"

type KeyPath struct {
	Folder   string
	Filename string
}

func (p *KeyPath) Path() string {
	return fmt.Sprintf("%s/%s", p.Folder, p.Filename)
}
