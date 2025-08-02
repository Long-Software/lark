package transaction

import (
	"Finn/category"
	"time"
)

type Transaction struct {
	ID         uint              `gorm:"primaryKey" json:"id"`
	Title      string            `gorm:"size:255;not null" json:"title"`
	Amount     float64           `gorm:"not null" json:"amount"`
	CategoryID uint              `gorm:"index;not null" json:"category_id"`
	Category   category.Category `gorm:"foreignKey:CategoryID" json:"-"`
	CreatedAt  time.Time         `json:"created_at"`
}
