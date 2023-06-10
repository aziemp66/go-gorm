package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Customer Model
type Customer struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;"`
	CreditCard CreditCard
}

// CreditCard Model
type CreditCard struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;"`
	Number     string
	CustomerID uuid.UUID `gorm:"type:uuid;"`
}
