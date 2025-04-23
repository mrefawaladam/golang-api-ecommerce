package domain

import "time"

type Order struct {
	ID        uint  `gorm:"primaryKey"`
	UserID    uint
	Total     float64
	Status    string  
	CreatedAt time.Time
	UpdatedAt time.Time
}
