// Package test : Test Package
package test

import (
	"testing"

	"github.com/google/uuid"

	domain "github.com/aziemp66/go-gorm/internal/domain"
)

func TestCreateCompany(t *testing.T) {
	companyID, err := uuid.NewRandom()
	if err != nil {
		t.Error(err)
	}

	company := domain.Company{
		ID:       companyID,
		Name:     "Azie's Company",
		Networth: 3000000,
	}

	result := DB.Create(&company)
	if result.Error != nil {
		t.Error(result.Error)
		return
	}
}

func TestCreateUser(t *testing.T) {
	companyID, err := uuid.Parse("5d732723-9b94-436f-877d-983dd998cd85")
	if err != nil {
		t.Error(err)
		return
	}

	user := domain.User{
		Name:      "Azie Melza Pratama",
		CompanyID: companyID,
		Age:       20,
	}

	result := DB.Create(&user)

	if result.Error != nil {
		t.Error(result.Error)
		return
	}
}

func TestReadUser(t *testing.T) {
	user := domain.User{
		ID: uuid.MustParse("42d0ed09-d41f-47d3-81bf-e59dea7ac9ac"),
	}

	DB.Where(&user).First(&user)

	t.Log(user)
}

func TestJoinUserAndCompany(t *testing.T) {
	type Result struct {
		Name     string
		Networth string
	}

	result := Result{}

	tx := DB.Model(&domain.User{}).
		Select("users.name", "companies.networth").
		Joins("left join companies on companies.id = users.company_id").Limit(1).Scan(&result)

	if tx.Error != nil {
		t.Error(tx.Error)
		return
	}

	t.Log(result)
}

func TestBelongToEagerLoading(t *testing.T) {
	var user domain.User

	tx := DB.Preload("Company").Find(&user)
	if tx.Error != nil {
		t.Error(tx.Error)
		return
	}

	t.Log(user.Company)
}
