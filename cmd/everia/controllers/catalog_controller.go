package controllers

import (
	"encoding/json"
	"path/filepath"
	"time"

	"github.com/Long-Software/Bex/apps/cmd/everia/constants"
	"github.com/Long-Software/Bex/apps/cmd/everia/models"
	"github.com/Long-Software/Bex/apps/cmd/everia/utils"
	"github.com/Long-Software/Bex/packages/file"
	"github.com/Long-Software/Bex/packages/log"
)

type CatalogController struct {
	baseDir  string         `json:"-"`
	filePath string         `json:"-"`
	catalog  models.Catalog `json:"-"`
}

func NewCatalogController() *CatalogController {
	var catalog_ctrl CatalogController
	rootDir, err := file.GetExecDir()
	if err != nil {
		utils.Logger.NewLog(log.FATAL, err.Error())
		return &catalog_ctrl
	}
	catalog_ctrl.baseDir = filepath.Join(rootDir, constants.CatalogDirName)
	err = file.MkdirAll(catalog_ctrl.baseDir)
	if err != nil {
		utils.Logger.NewLog(log.FATAL, err.Error())
		return &catalog_ctrl
	}
	catalog_ctrl.filePath = filepath.Join(catalog_ctrl.baseDir, constants.CatalogFile)
	catalog_ctrl.catalog = models.Catalog{
		Authors:  []string{},
		Apps:     []models.AppCatalog{},
		Websites: []models.WebsiteCatalog{},
	}
	err = catalog_ctrl.saveToFile()
	if err != nil {
		utils.Logger.NewLog(log.FATAL, err.Error())
		return &catalog_ctrl
	}
	return &catalog_ctrl
}

func (c *CatalogController) GetCatalog() models.Catalog {
	data, err := file.ReadFile(c.filePath)
	if err != nil {
		utils.Logger.NewLog(log.FATAL, err.Error())
		return models.Catalog{}
	}
	var catalog models.Catalog
	if err := json.Unmarshal([]byte(data), &catalog); err != nil {
		utils.Logger.NewLog(log.FATAL, err.Error())
		return models.Catalog{}
	}
	return catalog
}

func (c *CatalogController) AppendApp(app models.AppCatalog) {
	c.catalog.Apps = append(c.catalog.Apps, app)
	c.catalog.LastModified = time.Now().Format("2006-01-02 15:04")
	err := c.saveToFile()
	if err != nil {
		utils.Logger.NewLog(log.FATAL, err.Error())
	}
}

func (c *CatalogController) AppendWebsite(website models.WebsiteCatalog) {
	c.catalog.Websites = append(c.catalog.Websites, website)
	c.catalog.LastModified = time.Now().Format("2006-01-02 15:04")
	err := c.saveToFile()
	if err != nil {
		utils.Logger.NewLog(log.FATAL, err.Error())
	}
}

func (c *CatalogController) saveToFile() error {
	data, err := json.Marshal(c.catalog)
	if err != nil {
		return err
	}
	return file.Write(c.filePath, string(data))
}
