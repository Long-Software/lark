package analyzer

import (
	"strings"
)

type DartAnalyzer struct{}

func (d *DartAnalyzer) AnalyzeFile(content string, filePath string) *FileAnalysis {
	a := NewAnalyzer(AnalyzerOpts{
		Complexity: []string{
			`\bif\b`, `\belse\b`, `\bfor\b`, `\bwhile\b`, `\bswitch\b`,
			`\bcase\b`, `\bcatch\b`, `\btry\b`, `&&`, `\|\|`, `\?`,
		},
		Functions: `(?:void|int|String|bool|double|var|\w+)\s+\w+\s*\(`,
		Classes:   `class\s+\w+`,
		Imports:   `import\s+['"]([^'"]+)['"]`,
		Exports:   `(?:class|(?:void|int|String|bool|double)\s+)([A-Z]\w*)`,
	})

	res := a.Run(content)

	analysis := &FileAnalysis{
		Path:       filePath,
		Language:   "Dart",
		Size:       int64(len(content)),
		Complexity: res.Complexity,
		Functions:  res.Functions,
		Classes:    res.Classes,
		Imports:    res.Imports,
		Exports:    res.Exports,
		TODOs:      res.TODOs,
		Issues:     []Issue{},
		Metrics:    make(map[string]int),
	}

	d.analyzeLines(content, analysis)
	return analysis
}

func (d *DartAnalyzer) analyzeLines(content string, analysis *FileAnalysis) {
	lines := strings.Split(content, "\n")
	analysis.Lines = len(lines)

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			analysis.BlankLines++
		} else if strings.HasPrefix(trimmed, "//") || strings.HasPrefix(trimmed, "/*") {
			analysis.CommentLines++
		} else {
			analysis.CodeLines++
		}
	}
}
