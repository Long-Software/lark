package note

import (
	"context"
	"path/filepath"

	"github.com/Long-Software/Bex/apps/cmd/everia/internal/logger"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/Long-Software/Bex/packages/file"
	"github.com/Long-Software/Bex/packages/log"
)

type Controller struct {
	baseDir string `json:"-"`
}

func NewController() *Controller {
	rootDir, err := file.GetExecDir()
	if err != nil {
		logger.NewLog(log.FATAL, err.Error())
	}
	dir := filepath.Join(rootDir, NoteDirName)
	err = file.MkdirAll(dir)
	if err != nil {
		logger.NewLog(log.FATAL, err.Error())
	}

	return &Controller{
		baseDir: dir,
	}
}

func (n *Controller) Init() []Info {
	return []Info{
		{Title: "Welcome", LastEditTime: "12/10/23, 1:42 PM"},
		{Title: "Note 1", LastEditTime: "12/10/23, 1:42 PM"},
		{Title: "Note 2", LastEditTime: "12/10/23, 1:42 PM"},
		{Title: "Note 3", LastEditTime: "12/10/23, 1:42 PM"},
		{Title: "Note 4", LastEditTime: "12/10/23, 1:42 PM"},
	}
}
func (n *Controller) Notes() []Info {
	var notes []Info
	noteFiles, err := file.ListFilesWithExtension(n.baseDir, NoteFileExt)
	if err != nil {
		logger.NewLog(log.FATAL, err.Error())
		return []Info{}
	}
	for _, file := range noteFiles {
		note, err := info(n.FilePath(file))
		if err != nil {
			logger.NewLog(log.FATAL, err.Error())
			continue
		}
		notes = append(notes, note)
	}
	return notes
}

func (n *Controller) Read(filename string) string {
	content, err := file.ReadFile(filepath.Join(n.baseDir, filename))
	if err != nil {
		logger.NewLog(log.FATAL, err.Error())
	}
	return content
}

func (n *Controller) Write(filename, content string) {
	err := file.Write(filepath.Join(n.baseDir, filename), content)
	if err != nil {
		logger.NewLog(log.ERROR, err.Error())
	}
}
func (n *Controller) Create(filename, content string) {
	err := file.Write(filename, content)
	if err != nil {
		logger.NewLog(log.ERROR, err.Error())
	}
}
func (n *Controller) ShowSaveDialog(ctx context.Context) {
	options := runtime.SaveDialogOptions{
		Title:                "Save New Note",
		DefaultFilename:      "Untitle.md",
		DefaultDirectory:     n.baseDir,
		CanCreateDirectories: false,
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Markdown Files",
				Pattern:     "*.md",
			},
		},
	}

	filename, err := runtime.SaveFileDialog(ctx, options)
	if err != nil {
		logger.NewLog(log.ERROR, err.Error())
		return
	}
	err = file.Write(filename, "")
	if err != nil {
		logger.NewLog(log.ERROR, err.Error())
	}
}

func (n *Controller) Delete(ctx context.Context, filename string) {
	options := runtime.MessageDialogOptions{
		Type:    runtime.QuestionDialog,
		Message: "Are you sure you want to delete " + filename,
		Title:   "Delete Note",
	}
	response, err := runtime.MessageDialog(ctx, options)
	if err != nil {
		logger.NewLog(log.ERROR, err.Error())
	}
	if response == "Yes" {
		err := file.Delete(filepath.Join(n.baseDir, filename))
		if err != nil {
			logger.NewLog(log.ERROR, err.Error())
		}
	}
}

func info(filename string) (Info, error) {
	stat, err := file.GetFileInfo(filename)
	if err != nil {
		return Info{}, err
	}
	return Info{
		Title:        stat.Name(),
		LastEditTime: stat.ModTime().Format("2006-01-02 15:04"),
		DateTime:     stat.ModTime(),
	}, nil
}

func(n *Controller) FilePath(filename string) string {
	return filepath.Join(n.baseDir, filename)
}