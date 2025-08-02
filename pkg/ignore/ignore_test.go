package ignore

import (
	"io/fs"
	"path/filepath"
	"testing"

	"github.com/Long-Software/lark/pkg/log"
)

var lg = log.Logger{
	isproduction
}

func TestIgnore(t *testing.T) {
	files := []string{
		".gitignore",
	}
	ig, err := NewIgnorer(files)
	if err != nil {

	}
	err = filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if ig.ShouldIgnoreFolder(d.Name()) {
				return filepath.SkipDir
			}
			return nil
		} else {
			if ig.ShouldIgnoreFile(path) {
				return filepath.SkipDir
			}
			return nil
		}

	})
}
