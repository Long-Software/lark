package archaeologist

import "github.com/Long-Software/lark/cmd/analyser/internal/hotspot"

type ProjectStatistics struct {
	TotalFiles         int                 `json:"total_files"`
	TotalLines         int                 `json:"total_lines"`
	TotalCodeLines     int                 `json:"total_code_lines"`
	TotalFunctions     int                 `json:"total_functions"`
	TotalClasses       int                 `json:"total_classes"`
	AverageComplexity  float64             `json:"average_complexity"`
	LanguageBreakdown  map[string]int      `json:"language_breakdown"`
	Hotspots           []hotspot.Entry     `json:"hotspots"`
	RefactoringTargets []RefactoringTarget `json:"refactoring_targets"`
	DependencyGraph    map[string][]string `json:"dependency_graph"`
}

type RefactoringTarget struct {
	Path       string   `json:"path"`
	Reasons    []string `json:"reasons"`
	Priority   string   `json:"priority"`
	Complexity int      `json:"complexity"`
	LineCount  int      `json:"line_count"`
}
