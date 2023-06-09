package test

import (
	"testing"

	"github.com/google/uuid"

	"github.com/aziemp66/go-gorm/internal/repository"
	DB "github.com/aziemp66/go-gorm/pkg/postgres"
)

func TestCreateCompany(t *testing.T) {
	db := DB.NewDB()

	company := repository.Company{
		ID:       uuid.NewString(),
		Name:     "Azie's Company",
		Networth: 3000000,
	}

	result := db.Create(&company)
	if result.Error != nil {
		t.Error(result.Error)
	}
}
