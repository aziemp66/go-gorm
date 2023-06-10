package test

import (
	"testing"

	"github.com/google/uuid"

	"github.com/aziemp66/go-gorm/internal/domain"
)

func TestCreateCustomer(t *testing.T) {
	result := DB.Create(&domain.Customer{
		Name: "Jambai-kun",
	})

	if result.Error != nil {
		t.Error(result.Error)
		return
	}

	t.Log("Success")
}

func TestCreateCreditCard(t *testing.T) {
	err := DB.Create(&domain.CreditCard{
		Number:     "082141-189024",
		Points:     300,
		CustomerID: uuid.MustParse("c619c4f3-1220-48bc-95df-4c6f2e320fc4"),
	}).Error
	if err != nil {
		t.Error(err)
		return
	}
}
