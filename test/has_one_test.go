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

func TestReadCustomer(t *testing.T) {
	var customers []domain.Customer
	err := DB.Preload("CreditCard").Find(&customers).Error
	if err != nil {
		t.Error(err)
	}
	for _, v := range customers {
		t.Log(v.CreditCard)
	}
}

func TestReadCreditCard(t *testing.T) {
	var creditCards []domain.CreditCard

	err := DB.Find(&creditCards).Error
	if err != nil {
		t.Error(err)
	}

	t.Log(creditCards)
}
