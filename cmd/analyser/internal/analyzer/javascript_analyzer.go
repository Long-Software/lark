package analyzer

import (
	"path/filepath"
	"strings"
)

type JavaScriptAnalyzer struct{}

func (js *JavaScriptAnalyzer) AnalyzeFile(content string, filePath string) *FileAnalysis {
	ext := filepath.Ext(filePath)
	language := "JavaScript"

	switch ext {
	case ".ts":
		language = "TypeScript"
	case ".jsx":
		language = "JSX"
	case ".tsx":
		language = "TSX"
	}

	a := NewAnalyzer(AnalyzerOpts{
		Complexity: []string{
			`\bif\b`, `\belse\b`, `\bfor\b`, `\bwhile\b`, `\bswitch\b`,
			`\bcase\b`, `\bcatch\b`, `\btry\b`, `&&`, `\|\|`, `\?`,
		},
		Functions: `(?:function\s+\w+\s*\(|\w+\s*:\s*function\s*\(|\w+\s*=\s*function\s*\(|\w+\s*=>\s*|\w+\s*=\s*\([^)]*\)\s*=>)`,
		Classes:   `class\s+\w+`,
		Imports:   `(?:require\s*\(\s*['"]([^'"]+)['"]\s*\)|import\s+.*?\s+from\s+['"]([^'"]+)['"])`,
		Exports:   `export\s+(?:const|let|var|function|class)\s+(\w+)|export\s*{\s*([^,}\s]+)`,
	})

	res := a.Run(content)

	analysis := &FileAnalysis{
		Path:       filePath,
		Language:   language,
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

	js.analyzeLines(content, analysis)
	js.detectIssues(content, analysis)

	return analysis
}

func (js *JavaScriptAnalyzer) analyzeLines(content string, analysis *FileAnalysis) {
	lines := strings.Split(content, "\n")
	analysis.Lines = len(lines)

	inBlockComment := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if trimmed == "" {
			analysis.BlankLines++
			continue
		}

		// Handle block comments
		if strings.Contains(trimmed, "/*") {
			inBlockComment = true
		}
		if strings.Contains(trimmed, "*/") {
			inBlockComment = false
			analysis.CommentLines++
			continue
		}

		if inBlockComment || strings.HasPrefix(trimmed, "//") {
			analysis.CommentLines++
		} else {
			analysis.CodeLines++
		}
	}
}

func (js *JavaScriptAnalyzer) detectIssues(content string, analysis *FileAnalysis) {
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		// Console.log detection
		if strings.Contains(line, "console.log") {
			analysis.Issues = append(analysis.Issues, Issue{
				Type:        "console_log",
				Severity:    "info",
				Line:        i + 1,
				Description: "Console.log statement found",
			})
		}

		// Eval usage
		if strings.Contains(line, "eval(") {
			analysis.Issues = append(analysis.Issues, Issue{
				Type:        "eval_usage",
				Severity:    "error",
				Line:        i + 1,
				Description: "Dangerous eval() usage detected",
			})
		}
	}
}
