package domain

import "time"

type Cart struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint
	ProductID uint
	Product   Product   `gorm:"foreignKey:ProductID"` 
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
