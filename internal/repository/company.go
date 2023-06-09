package repository

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	ID       string
	Name     string
	Networth float64
}
