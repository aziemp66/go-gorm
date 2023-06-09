package test

import (
	"fmt"
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
		ID: "0ba40e37-778e-4371-9b8b-6628376406a1",
	}

	db.Joins("companies").First(&user)

	fmt.Println(user.Company)
}
