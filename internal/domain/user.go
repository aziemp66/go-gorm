package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;"`
	Name     string
	Networth float64
}

// Belong To Association
type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;"`
	Name      string
	Age       uint
	CompanyID uuid.UUID `gorm:"type:uuid;"`
	Company   Company
}
