package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model           // Adds some metadata fields to the table
	ID         uuid.UUID `gorm:"type:uuid"` // Explicitly specify the type to be uuid
	Name       string
	Price      int
	Amount     int
	Category   string
	Evaluation string
	Note       string
	ImageUrl   string
	Link       string
}
