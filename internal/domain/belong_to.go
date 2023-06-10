// Package domain is where the Database tables are modeled
package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Company Struct Model
type Company struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name     string
	Networth float64
}

// User Struct Model
type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string
	Age       uint
	CompanyID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Company   Company
}
