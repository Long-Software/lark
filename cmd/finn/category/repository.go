package category

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository interface {
	Create(string) error
	List() ([]Category, error)
}

type GORMRepository struct {
	db *gorm.DB
}

func NewGORMRepository(dns string) (Repository, error) {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	return &GORMRepository{db: db}, err
}

func (g *GORMRepository) Create(name string) error {
	return g.db.Create(&Category{Name: name}).Error
}
func (g *GORMRepository) List() ([]Category, error) {
	var categories []Category
	err := g.db.Find(&categories).Error
	return categories, err
}
