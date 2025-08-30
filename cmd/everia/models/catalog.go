package models

type Catalog struct {
	Name         string    `json:"name"`
	Authors      []string  `json:"authors"`
	LastModified string    `json:"last_modified"`
	Apps         []AppCatalog     `json:"apps"`
	Websites     []WebsiteCatalog `json:"websites"`
}
type AppCatalog struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	ImageUrl string `json:"image_url"`
}

type WebsiteCatalog struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Version  string `json:"version"`
	ImageUrl string `json:"image_url"`
}
