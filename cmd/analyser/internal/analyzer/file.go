package analyzer

import "time"

type FileAnalyzerConfig struct {
	MaxFileSize         int64
	SkipDirectories     []string
	IncludeExtensions   []string
	ComplexityThreshold int
	FunctionLengthLimit int
	MaxNestingDepth     int
}

func NewDefaultFileAnalyzerConfig() FileAnalyzerConfig {
	return FileAnalyzerConfig{
		MaxFileSize:         1024 * 1024, // 1MB
		SkipDirectories:     []string{".git", "node_modules", "vendor", ".dart_tool", "build", ".next", "dist"},
		ComplexityThreshold: 10,
		FunctionLengthLimit: 50,
		MaxNestingDepth:     4,
	}
}

type FileAnalysis struct {
	Path         string         `json:"path"`
	Language     string         `json:"language"`
	Size         int64          `json:"size"`
	Lines        int            `json:"lines"`
	CodeLines    int            `json:"code_lines"`
	CommentLines int            `json:"comment_lines"`
	BlankLines   int            `json:"blank_lines"`
	Functions    int            `json:"functions"`
	Classes      int            `json:"classes"`
	Complexity   int            `json:"complexity"`
	LastModified time.Time      `json:"last_modified"`
	Imports      []string       `json:"imports"`
	Exports      []string       `json:"exports"`
	TODOs        []TODO         `json:"todos"`
	Dependencies []string       `json:"dependencies"`
	HotspotScore float64        `json:"hotspot_score"`
	Issues       []Issue        `json:"issues"`
	Metrics      map[string]int `json:"metrics"`
}

type Issue struct {
	Type        string `json:"type"`
	Severity    string `json:"severity"`
	Line        int    `json:"line"`
	Description string `json:"description"`
}
