package domain

import "time"

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Price       float64
	Stock       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
