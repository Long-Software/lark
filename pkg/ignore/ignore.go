package ignore

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type Ignorer struct {
	Imports []string
	Rules   []Rule
}

func NewIgnorer(importFiles []string) (*Ignorer, error) {
	ignorer := &Ignorer{
		Imports: importFiles,
	}

	for _, file := range importFiles {
		err := ignorer.loadIgnoreFile(file)
		if err != nil {
			return nil, err
		}
	}

	return ignorer, nil
}

func (i *Ignorer) loadIgnoreFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		rule := NewRule(line)
		i.Rules = append(i.Rules, rule)
	}

	return scanner.Err()
}

func (i *Ignorer) ShouldIgnoreFolder(folderPath string) bool {
	return i.matchRules(folderPath, true)
}

func (i *Ignorer) ShouldIgnoreFile(filePath string) bool {
	return i.matchRules(filePath, false)
}

func (i *Ignorer) matchRules(path string, isDir bool) bool {
	path = filepath.ToSlash(path)
	ignored := false
	for _, rule := range i.Rules {
		if rule.IsDir && !isDir {
			continue
		}
		if rule.Regex.MatchString(path) {
			if rule.IsNegated {
				ignored = false
			} else {
				ignored = true
			}
		}
	}
	return ignored
}
