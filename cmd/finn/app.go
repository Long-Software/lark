package main

import (
	"finn/internal/category"
	"finn/internal/transaction"
	"context"
	"fmt"

	"github.com/Long-Software/lark/pkg/env"
	"github.com/Long-Software/lark/pkg/log"
)

type Config struct {
	DBURL string `mapstructure:"DB_URL"`
}

// App struct
type App struct {
	ctx context.Context
	lg  log.Logger
	// TODO: Implement the logging system here
	catRepo  category.Repository
	tranRepo transaction.Repository
}

// NewApp creates a new App application struct
func NewApp() *App {
	var cfg Config
	var app App

	app.lg = log.Logger{
		IsProduction: true,
		HasTimestamp: true,
		HasFilepath:  true,
		HasMethod:    true,
	}

	err := env.Load(&cfg, ".env")
	if err != nil {
		app.lg.NewLog(log.FATAL, err.Error())
	}
	catRepo, err := category.NewGORMRepository(cfg.DBURL)
	if err != nil {
		app.lg.NewLog(log.FATAL, err.Error())
	}
	app.catRepo = catRepo

	tranRepo, err := transaction.NewGORMRepository(cfg.DBURL)
	if err != nil {
		app.lg.NewLog(log.FATAL, err.Error())
	}
	app.tranRepo = tranRepo
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

func (a *App) CreateCategory(name string) error {
	return a.catRepo.Create(name)
}

func (a *App) ListTransaction() ([]transaction.Transaction, error) {
	return a.tranRepo.ListTransactions()
}

func (a *App) CreateTransaction(title string, amount float64, categoryID uint) error {
	return a.tranRepo.CreateTransaction(title, amount, categoryID)
}
