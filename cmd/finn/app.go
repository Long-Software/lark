package main

import (
	"Finn/category"
	"context"
	"fmt"

	"github.com/Long-Software/lark/pkg/log"
)

// App struct
type App struct {
	ctx context.Context
	lg  log.Logger
	// TODO: Implement the logging system here
	catRepo category.Repository
}

// NewApp creates a new App application struct
func NewApp() *App {
	var app App

	app.lg = log.Logger{
		IsProduction: true,
		HasTimestamp: true,
		HasFilepath:  true,
		HasMethod:    true,
	}
	catRepo, err := category.NewGORMRepository("")
	if err != nil {
		app.lg.NewLog(log.FATAL, err.Error())
	}
	// TODO: Implement the transaction repository here
	app.catRepo = catRepo
	return &app
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

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) ListCategories() ([]category.Category, error) {
	return a.catRepo.List()
}
