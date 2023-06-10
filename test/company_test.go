package test

import (
	"testing"

	"github.com/google/uuid"

	domain "github.com/aziemp66/go-gorm/internal/domain"
	DB "github.com/aziemp66/go-gorm/pkg/postgres"
)

func TestCreateCompany(t *testing.T) {
	db := DB.NewDB()

	companyID, err := uuid.NewRandom()
	if err != nil {
		t.Error(err)
	}

	company := domain.Company{
		ID:       companyID,
		Name:     "Azie's Company",
		Networth: 3000000,
	}

	result := db.Create(&company)
	if result.Error != nil {
		t.Error(result.Error)
		return
	}
}
