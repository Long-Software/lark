package transaction

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Repository interface {
	CreateTransaction(title string, amount float64, categoryID uint) error
	ListTransactions() ([]Transaction, error)
	ListTransactionsByCategory(categoryID uint) ([]Transaction, error)
}

type GORMRepository struct {
	db *gorm.DB
}

func NewGORMRepository(dns string) (Repository, error) {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	return &GORMRepository{db: db}, err
}

func (g *GORMRepository) CreateTransaction(title string, amount float64, categoryID uint) error {
	tx := Transaction{
		Title:      title,
		Amount:     amount,
		CategoryID: categoryID,
	}
	return g.db.Create(&tx).Error
}

func (g *GORMRepository) ListTransactions() ([]Transaction, error) {
	var transactions []Transaction
	err := g.db.Order("created_at desc").Find(&transactions).Error
	return transactions, err
}

func (g *GORMRepository) ListTransactionsByCategory(categoryID uint) ([]Transaction, error) {
	var transactions []Transaction
	err := g.db.
		Where("category_id = ?", categoryID).
		Order("created_at desc").
		Find(&transactions).Error
	return transactions, err
}
