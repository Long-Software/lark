package archaeologist

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"time"

	anl "github.com/Long-Software/lark/cmd/analyser/internal/analyzer"
	"github.com/Long-Software/lark/cmd/analyser/internal/hotspot"
	"github.com/Long-Software/lark/cmd/analyser/internal/utils"
	"github.com/Long-Software/lark/pkg/log"
)

type CodeArchaeologist struct {
	RootDir    string
	FileMap    map[string]*anl.FileAnalysis
	Config     anl.FileAnalyzerConfig
	Statistics ProjectStatistics
}

func NewCodeArchaeologist(rootDir string) *CodeArchaeologist {
	ca := &CodeArchaeologist{
		RootDir: rootDir,
		FileMap: make(map[string]*anl.FileAnalysis),
		Config:  anl.NewDefaultFileAnalyzerConfig(),
		Statistics: ProjectStatistics{
			LanguageBreakdown: make(map[string]int),
			DependencyGraph:   make(map[string][]string),
		},
	}
	return ca
}

func (ca *CodeArchaeologist) Excavate() error {
	err := ca.scanDirectory(ca.RootDir)
	if err != nil {
		return fmt.Errorf("scanning failed: %w", err)
	}

	for filePath, analysis := range ca.FileMap {
		ca.Statistics.DependencyGraph[filePath] = analysis.Imports
	}

	ca.identifyHotspots()
	ca.calculateStatistics()
	ca.generateReport()
	ca.saveJSONReport()
	return nil
}

func (ca *CodeArchaeologist) scanDirectory(dir string) error {
	return filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if slices.Contains(ca.Config.SkipDirectories, d.Name()) {
				return filepath.SkipDir
			}
			return nil
		}

		return ca.analyzeFile(path)
	})
}

func (ca *CodeArchaeologist) analyzeFile(filePath string) error {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return err
	}

	if fileInfo.Size() > ca.Config.MaxFileSize {
		utils.Log.NewLog(log.WARNING, fmt.Sprintf("⚠️  Skipping large file: %s (%d bytes)", filePath, fileInfo.Size()))
		return nil
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	analyzer := ca.getAnalyzerForFile(filePath)
	if analyzer == nil {
		return nil
	}

	analysis := analyzer.AnalyzeFile(string(content), filePath)
	analysis.LastModified = fileInfo.ModTime()
	analysis.HotspotScore = hotspot.Calculate(analysis.Complexity, analysis.Lines)

	ca.FileMap[filePath] = analysis

	return nil
}

func (ca *CodeArchaeologist) getAnalyzerForFile(filePath string) anl.LanguageAnalyzer {
	ext := filepath.Ext(filePath)
	switch ext {
	case ".go":
		return &anl.GoAnalyzer{}
	case ".js", ".ts", ".jsx", ".tsx":
		return &anl.JavaScriptAnalyzer{}
	case ".php":
		return &anl.PHPAnalyzer{}
	case ".dart":
		return &anl.DartAnalyzer{}
	default:
		return nil
	}
}

func (ca *CodeArchaeologist) identifyHotspots() {
	var hotspots []hotspot.Entry
	for path, analysis := range ca.FileMap {
		if analysis.HotspotScore > 50 {
			hotspots = append(hotspots, hotspot.Entry{Path: path, Score: analysis.HotspotScore})
		}
	}

	sort.Slice(hotspots, func(i, j int) bool {
		return hotspots[i].Score > hotspots[j].Score
	})

	// Take top 10
	for i, h := range hotspots {
		if i >= 10 {
			break
		}
		ca.Statistics.Hotspots = append(ca.Statistics.Hotspots, hotspot.Entry{
			Path:  h.Path,
			Score: h.Score,
		})
	}
}

func (ca *CodeArchaeologist) calculateStatistics() {
	totalComplexity := 0

	for _, analysis := range ca.FileMap {
		ca.Statistics.TotalFiles++
		ca.Statistics.TotalLines += analysis.Lines
		ca.Statistics.TotalCodeLines += analysis.CodeLines
		ca.Statistics.TotalFunctions += analysis.Functions
		ca.Statistics.TotalClasses += analysis.Classes
		totalComplexity += analysis.Complexity

		ca.Statistics.LanguageBreakdown[analysis.Language]++

		// Identify refactoring targets
		if analysis.Complexity > ca.Config.ComplexityThreshold ||
			analysis.Lines > ca.Config.FunctionLengthLimit*2 {
			target := RefactoringTarget{
				Path:       analysis.Path,
				Complexity: analysis.Complexity,
				LineCount:  analysis.Lines,
				Priority:   "medium",
			}

			if analysis.Complexity > 20 {
				target.Priority = "high"
				target.Reasons = append(target.Reasons, "High complexity")
			}

			if analysis.Lines > 200 {
				target.Reasons = append(target.Reasons, "Large file")
			}

			if len(analysis.Issues) > 5 {
				target.Reasons = append(target.Reasons, "Multiple issues")
			}

			ca.Statistics.RefactoringTargets = append(ca.Statistics.RefactoringTargets, target)
		}
	}

	if ca.Statistics.TotalFiles > 0 {
		ca.Statistics.AverageComplexity = float64(totalComplexity) / float64(ca.Statistics.TotalFiles)
	}
}

func (ca *CodeArchaeologist) generateReport() {
	utils.Log.NewLog(log.INFO, fmt.Sprintf("📁 Total files analyzed: %d", ca.Statistics.TotalFiles))
	utils.Log.NewLog(log.INFO, fmt.Sprintf("📄 Total lines of code: %d", ca.Statistics.TotalLines))
	utils.Log.NewLog(log.INFO, fmt.Sprintf("⚙️ Total functions: %d", ca.Statistics.TotalFunctions))
	utils.Log.NewLog(log.INFO, fmt.Sprintf("🏗️ Total classes: %d", ca.Statistics.TotalClasses))
	utils.Log.NewLog(log.INFO, fmt.Sprintf("🔢 Average complexity: %.2f", ca.Statistics.AverageComplexity))

	fmt.Println("📊 Language Breakdown:")
	for lang, count := range ca.Statistics.LanguageBreakdown {
		percentage := float64(count) / float64(ca.Statistics.TotalFiles) * 100
		fmt.Printf("   %s: %d files (%.1f%%)\n", lang, count, percentage)
	}

	if len(ca.Statistics.Hotspots) > 0 {
		fmt.Println("🔥 Top Refactoring Hotspots:")
		for i, hotspot := range ca.Statistics.Hotspots {
			fmt.Printf("   %d. %s (Score: %.2f)\n", i+1, hotspot.Path, hotspot.Score)
		}
	}

	if len(ca.Statistics.RefactoringTargets) > 0 {
		fmt.Printf("🎯 Refactoring Targets (%d found):\n", len(ca.Statistics.RefactoringTargets))
		for i, target := range ca.Statistics.RefactoringTargets {
			if i >= 5 { // Show top 5
				break
			}
			fmt.Printf("   %d. %s (%s priority)\n", i+1, target.Path, target.Priority)
			fmt.Printf("      Complexity: %d, Lines: %d\n", target.Complexity, target.LineCount)
			if len(target.Reasons) > 0 {
				fmt.Printf("      Reasons: %s\n", strings.Join(target.Reasons, ", "))
			}
		}
	}
}

func (ca *CodeArchaeologist) saveJSONReport() {
	report := map[string]interface{}{
		"timestamp":  time.Now(),
		"root_dir":   ca.RootDir,
		"statistics": ca.Statistics,
		"files":      ca.FileMap,
	}

	jsonData, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		utils.Log.NewLog(log.ERROR, err.Error())
		return
	}

	filename := fmt.Sprintf("code_analysis_%s.json", time.Now().Format("20060102_150405"))
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		utils.Log.NewLog(log.ERROR, err.Error())
		return
	}
	utils.Log.NewLog(log.INFO, fmt.Sprintf("Report saved to: %s\n", filename))
}
