package model

import (
	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `gorm:"primary_key" json:"id"`
	Name        string         `json:"name"`
	AuthorID    int            `json:"author_id"`
	StockCode   string         `json:"stock_code"`
	ISBN        string         `json:"isbn"`
	PageCount   int            `json:"page_count"`
	Price       int            `json:"price"`
	StockAmount int            `json:"stock_amount"`
	Deleted     gorm.DeletedAt `json:"deleted"`
	Author      *Author        `gorm:"foreignKey:AuthorID" json:"author,omitempty"`
}
