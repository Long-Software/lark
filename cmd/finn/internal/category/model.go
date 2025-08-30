package category
type Category struct {
    ID   uint   `gorm:"primaryKey" json:"id"`
    Name string `gorm:"uniqueIndex;size:100;not null" json:"name"`
}