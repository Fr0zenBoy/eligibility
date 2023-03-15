package account

import (
	"testing"

	"github.com/Fr0zenBoy/authoraizer/transaction"
)

func TestAccount(t *testing.T) {
	someClient := Account{
		CardIsActive: true,
		IsWhiteListed: true,
		Limit: 5000.0,
		DenyList: []string{"moes"},
	}

	otherClient := Account{
		CardIsActive: false,
		IsWhiteListed: false,
		Limit: 50.0,
		DenyList: []string{},
	}

	someTransaction := transaction.Transaction{
		Merchant: "Moes",
		Amount: 300.0,
		Time: "2023-03-14 18:04:00",
	}
}
