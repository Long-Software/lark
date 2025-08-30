package main

import (
	"context"

	"github.com/Long-Software/Bex/apps/cmd/everia/controllers"
	"github.com/Long-Software/Bex/apps/cmd/everia/internal/fetcher"
	"github.com/Long-Software/Bex/apps/cmd/everia/internal/logger"
	"github.com/Long-Software/Bex/apps/cmd/everia/internal/note"
	"github.com/Long-Software/Bex/apps/cmd/everia/models"
	"github.com/Long-Software/Bex/apps/cmd/everia/utils"
	"github.com/Long-Software/Bex/packages/file"
	"github.com/Long-Software/Bex/packages/log"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Notes
var notectrl = note.NewController()

func (a *App) MockNotes() []note.Info {
	return notectrl.Init()
}

func (a *App) ListNotes() []note.Info {
	return notectrl.Notes()
}

func (a *App) ReadNote(filename string) string {
	// return notectrl.Read(filename)
	return a.ReadFile(notectrl.FilePath(filename))
}

func (a *App) WriteNote(filename, content string)(error) {
	notectrl.Write(filename, content)
	return a.WriteFile(notectrl.FilePath(filename), content)
}

func (a *App) SaveNoteDialog() {
	notectrl.ShowSaveDialog(a.ctx)
}

func (a *App) DeleteNote(filename string) {
	notectrl.Delete(a.ctx, filename)
}

// Password

func (a *App) UnlockVault(password string) ([]models.PasswordEntry, error) {
	data, err := utils.LoadEncryptedFile("vault.enc")
	if err != nil {
		return []models.PasswordEntry{}, err
	}

	vData, err := utils.DecryptVault(data, password)
	return vData.Entries, err
}

// Catalog

var catalog_ctrl = controllers.NewCatalogController()

func (a *App) GetCatalog() models.Catalog {
	return catalog_ctrl.GetCatalog()
}

func (a *App) AppendApp(app models.AppCatalog) {
	catalog_ctrl.AppendApp(app)
}
func (a *App) AppendWebsite(website models.WebsiteCatalog) {
	catalog_ctrl.AppendWebsite(website)
}

// Fetcher
var fetchercontroller = fetcher.NewController()

func (a *App) Fetch(url string) (string, error) {
	return fetchercontroller.FetchAndConvertToMarkdown(url)
}

// TODO: use this function to read the file content from the filepath after the fetch is completed
func (a *App) ReadFile(filePath string) string {
	content, err := file.ReadFile(filePath)
	if err != nil {
		logger.NewLog(log.FATAL, err.Error())
	}
	return content
}

func (a *App) WriteFile(filePath, content string) error {
	return file.Write(filePath, content)
}
