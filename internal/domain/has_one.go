package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Customer Model
type Customer struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name       string
	CreditCard CreditCard
}

// CreditCard Model
type CreditCard struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Number     string
	Points     uint64
	IsBlocked  bool      `gorm:"default:false"`
	CustomerID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}
