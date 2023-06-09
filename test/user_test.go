package test

import (
	"testing"

	"github.com/google/uuid"

	domain "github.com/aziemp66/go-gorm/internal/domain"
	DB "github.com/aziemp66/go-gorm/pkg/postgres"
)

func TestCreateUser(t *testing.T) {
	db := DB.NewDB()

	user := domain.User{
		ID:        uuid.NewString(),
		Name:      "Azie Melza Pratama",
		CompanyID: "60d59c4c-c209-452d-8df9-c418f60bed25",
		Age:       20,
	}

	result := db.Create(&user)

	if result.Error != nil {
		t.Error(result.Error)
	}
}

func TestReadUser(t *testing.T) {
	db := DB.NewDB()

	user := domain.User{
		ID: "42d0ed09-d41f-47d3-81bf-e59dea7ac9ac",
	}

	db.Where(&user).First(&user)

	t.Log(user)
}

func TestJoinUserAndCompany(t *testing.T) {
	db := DB.NewDB()

	type Result struct {
		Name     string
		Networth string
	}

	result := Result{}

	tx := db.Model(&domain.User{}).
		Select("users.name", "companies.networth").
		Joins("left join companies on companies.id = users.company_id").Scan(&result)

	if tx.Error != nil {
		t.Error(tx.Error)
	}

	t.Log(result)
}
