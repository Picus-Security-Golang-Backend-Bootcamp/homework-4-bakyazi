package model

import (
	"gorm.io/gorm"
)

type Author struct {
	//gorm.Model
	ID      uint           `gorm:"primaryKey" json:"id"`
	Name    string         `json:"name"`
	Books   []Book         `gorm:"foreignKey:AuthorID" json:"books,omitempty"`
	Deleted gorm.DeletedAt `json:"deleted"`
}
