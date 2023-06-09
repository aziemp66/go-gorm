package test

import (
	"testing"

	"github.com/google/uuid"

	domain "github.com/aziemp66/go-gorm/internal/domain"
	DB "github.com/aziemp66/go-gorm/pkg/postgres"
)

func TestCreateUser(t *testing.T) {
	db := DB.NewDB()

	userID, err := uuid.NewRandom()
	if err != nil {
		t.Error(err)
	}

	companyID, err := uuid.Parse("3f365a60-1cf6-4bcf-93f2-a90c6e3038a1")
	if err != nil {
		t.Error(err)
	}

	user := domain.User{
		ID:        userID,
		Name:      "Azie Melza Pratama",
		CompanyID: companyID,
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
		ID: uuid.MustParse("42d0ed09-d41f-47d3-81bf-e59dea7ac9ac"),
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
		Joins("left join companies on companies.id = users.company_id").Limit(1).Scan(&result)

	if tx.Error != nil {
		t.Error(tx.Error)
	}

	t.Log(result)
}
