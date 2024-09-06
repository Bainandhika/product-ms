package views

import "time"

type Product struct {
    ID          string     `json:"id" gorm:"type:varchar(36);primaryKey"`
    Name        string     `json:"name" gorm:"type:varchar(255);not null"`
    Description string     `json:"description" gorm:"type:text"`
    Price       float64    `json:"price" gorm:"type:decimal(10,2);not null"`
    Variety     string     `json:"variety" gorm:"type:varchar(100)"`
    Rating      float32    `json:"rating" gorm:"type:decimal(3,2)"`
    Stock       int        `json:"stock" gorm:"not null"`
    CreatedAt   *time.Time `json:"created_at" gorm:"type:datetime"`
    UpdatedAt   *time.Time `json:"updated_at" gorm:"type:datetime"`
}
