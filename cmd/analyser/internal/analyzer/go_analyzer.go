package analyzer

import (
	"fmt"
	"strings"
)

type GoAnalyzer struct{}

func (g *GoAnalyzer) AnalyzeFile(content string, filePath string) *FileAnalysis {
	a := NewAnalyzer(AnalyzerOpts{
		Complexity: []string{
			`\bif\b`, `\belse\b`, `\bfor\b`, `\bswitch\b`, `\bselect\b`,
			`\bcase\b`, `\bgo\b`, `\bdefer\b`, `&&`, `\|\|`,
		},
		Functions: `func\s+(\w+|\([^)]*\)\s*\w+)\s*\(`,
		Classes:   `type\s+\w+\s+struct`,
		Imports:   `import\s+(?:\(\s*)?(?:"([^"]+)"|([^\s\)]+))`,
		Exports:   `(?:func|type)\s+([A-Z]\w*)\s*(?:\(|\s)`,
	})

	res := a.Run(content)
	analysis := &FileAnalysis{
		Path:       filePath,
		Language:   "Go",
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

	g.analyzeLines(content, analysis)
	g.detectIssues(content, analysis)

	return analysis
}

func (g *GoAnalyzer) analyzeLines(content string, analysis *FileAnalysis) {
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

func (g *GoAnalyzer) detectIssues(content string, analysis *FileAnalysis) {
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		// Long line detection
		if len(line) > 120 {
			analysis.Issues = append(analysis.Issues, Issue{
				Type:        "long_line",
				Severity:    "warning",
				Line:        i + 1,
				Description: fmt.Sprintf("Line is %d characters long (>120)", len(line)),
			})
		}

		// Potential issues
		if strings.Contains(line, "panic(") {
			analysis.Issues = append(analysis.Issues, Issue{
				Type:        "panic_usage",
				Severity:    "warning",
				Line:        i + 1,
				Description: "Usage of panic() detected",
			})
		}
	}
}
