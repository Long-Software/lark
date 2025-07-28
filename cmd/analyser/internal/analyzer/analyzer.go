package analyzer

import (
	"regexp"
	"strings"
)

type LanguageAnalyzer interface {
	AnalyzeFile(string, string) *FileAnalysis
}
type AnalyzerOpts struct {
	Complexity []string
	Functions  string
	Classes    string
	Imports    string
	Exports    string
}

type Analyzer struct {
	AnalyzerOpts
	// complexity []string
	// functions  string
	// classes    string
	// imports    string
}

func NewAnalyzer(opt AnalyzerOpts) *Analyzer {
	return &Analyzer{
		opt,
	}
}

func (a *Analyzer) calculateComplexity(content string) int {
	complexity := 1
	for _, pattern := range a.Complexity {
		re := regexp.MustCompile(pattern)
		matches := re.FindAllString(content, -1)
		complexity += len(matches)
	}

	return complexity
}

func (a *Analyzer) countFunctions(content string) int {
	re := regexp.MustCompile(a.Functions)
	matches := re.FindAllString(content, -1)
	return len(matches)
}

func (a *Analyzer) countClasses(content string) int {
	re := regexp.MustCompile(a.Classes)
	matches := re.FindAllString(content, -1)
	return len(matches)
}

func (a *Analyzer) findImports(content string) []string {
	var imports []string
	re := regexp.MustCompile(a.Imports)
	matches := re.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		imports = append(imports, match[1])
	}
	return imports
}

func (a *Analyzer) findExports(content string) []string {
	var exports []string
	re := regexp.MustCompile(a.Exports)
	matches := re.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		exports = append(exports, match[1])
	}

	return exports
}

type TODO struct {
	Line    int    `json:"line"`
	Content string `json:"content"`
	Type    string `json:"type"` // TODO, FIXME, HACK, BUG, etc.
}

func (a *Analyzer) FindTODOs(content string) []TODO {
	var todos []TODO

	patterns := map[string]string{
		"TODO":  `(?i)todo:?\s*(.*)`,
		"FIXME": `(?i)fixme:?\s*(.*)`,
		"HACK":  `(?i)hack:?\s*(.*)`,
		"BUG":   `(?i)bug:?\s*(.*)`,
		"NOTE":  `(?i)note:?\s*(.*)`,
	}

	lines := strings.Split(content, "\n")
	for i, line := range lines {
		for todoType, pattern := range patterns {
			re := regexp.MustCompile(pattern)
			if matches := re.FindStringSubmatch(line); matches != nil {
				todos = append(todos, TODO{
					Line:    i + 1,
					Content: strings.TrimSpace(matches[1]),
					Type:    todoType,
				})
			}
		}
	}

	return todos
}

type AnalyzerResult struct {
	Complexity int
	Functions  int
	Classes    int
	Imports    []string
	Exports    []string
	TODOs      []TODO
}

func (a *Analyzer) Run(content string) *AnalyzerResult {
	var res = &AnalyzerResult{
		Complexity: a.calculateComplexity(content),
		Functions:  a.countFunctions(content),
		Classes:    a.countClasses(content),
		Imports:    a.findImports(content),
		Exports:    a.findExports(content),
		TODOs:      a.FindTODOs(content),
	}
	return res
}
