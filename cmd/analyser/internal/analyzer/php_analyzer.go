package analyzer

import (
	"strings"
)

type PHPAnalyzer struct{}

func (p *PHPAnalyzer) AnalyzeFile(content string, filePath string) *FileAnalysis {
	a := NewAnalyzer(AnalyzerOpts{
		Complexity: []string{
			`\bif\b`, `\belse\b`, `\bfor\b`, `\bwhile\b`, `\bswitch\b`,
			`\bcase\b`, `\bcatch\b`, `\btry\b`, `&&`, `\|\|`, `\?`,
			`\bforeach\b`,
		},
		Functions: `function\s+\w+\s*\(`,
		Classes:   `class\s+\w+`,
		Imports:   `(?:include|require)(?:_once)?\s*\(?['"]([^'"]+)['"]|use\s+([^;]+);`,
		Exports:   `(?:public\s+)?(?:function|class)\s+(\w+)`,
	})

	res := a.Run(content)
	analysis := &FileAnalysis{
		Path:       filePath,
		Language:   "PHP",
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

	p.analyzeLines(content, analysis)
	p.detectIssues(content, analysis)

	return analysis
}

func (p *PHPAnalyzer) analyzeLines(content string, analysis *FileAnalysis) {
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

		if inBlockComment || strings.HasPrefix(trimmed, "//") || strings.HasPrefix(trimmed, "#") {
			analysis.CommentLines++
		} else {
			analysis.CodeLines++
		}
	}
}

func (p *PHPAnalyzer) detectIssues(content string, analysis *FileAnalysis) {
	lines := strings.Split(content, "\n")

	for i, line := range lines {
		// SQL injection risks
		if strings.Contains(line, "$_GET") || strings.Contains(line, "$_POST") {
			analysis.Issues = append(analysis.Issues, Issue{
				Type:        "potential_injection",
				Severity:    "warning",
				Line:        i + 1,
				Description: "Direct use of user input detected",
			})
		}
	}
}
