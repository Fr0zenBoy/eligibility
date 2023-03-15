package account

import (
	"github.com/Fr0zenBoy/authoraizer/transaction"
	"github.com/Fr0zenBoy/authoraizer/logic"
)

type Account struct {
	CardIsActive bool  `json:"cardIsActive"`
	IsWhiteListed bool `json:"isWhiteListed"`
	Limit float64      `json:"limit"`
	DenyList []string  `json:"denyList"`
}

func (a Account) checkAmountAboveLimit(t transaction.Transaction) bool {
	return a.Limit > t.Amount
}

func (a Account) checkCardIsActive() bool {
	return a.CardIsActive
}

func (a Account) checkFirstTransactionSafe(t transaction.Transaction, l transaction.LastTransactions) bool {
	areEmptyList := len(l) == 0
	if areEmptyList {
		return logic.GetPercentege(t.Amount, a.Limit) < 90
	}
	return false
}

func (a Account) checkDenyList(t transaction.Transaction) bool {
	allowed := true
	for _, merchant := range a.DenyList {
		if merchant == t.Merchant {
			allowed = false
			break
		}
	}
	return allowed
}
