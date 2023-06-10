package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Pawn Model
type Pawn struct {
	gorm.Model
	ID                   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Usename              string    `gorm:"type:varchar(225)"`
	Exp                  int64
	LinkedPersonaWeapons []PersonaWeapon `gorm:"foreignKey:LinkedPawnID"`
}

// PersonaWeapon Model
type PersonaWeapon struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name         string    `gorm:"type:varchar(225);unique;not null"`
	Description  string    `gorm:"not null"`
	Damage       float64
	Enchantment  []string `gorm:"type:varchar(255)[]"`
	LinkedPawnID uuid.UUID
}
