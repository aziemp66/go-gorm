package repository

import "gorm.io/gorm"

// Belong To Association
type User struct {
	gorm.Model
	ID        string
	Name      string
	Age       uint
	CompanyID string
	Company
}
