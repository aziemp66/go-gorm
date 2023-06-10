package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Gamer Model
type Gamer struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Username   string    `gorm:"type:varchar(255)"`
	BoughtGame []Game    `gorm:"many2many:gamer_bought_game"`
}

// Game Model
type Game struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Name       string    `gorm:"type:varchar(255)"`
	GamePlayer []Gamer   `gorm:"many2many:gamer_bought_game"`
}
