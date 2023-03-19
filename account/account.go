package account

import (
	"github.com/Fr0zenBoy/authoraizer/logic"
	"github.com/Fr0zenBoy/authoraizer/transaction"
)

type Account struct {
	CardIsActive *bool  `json:"cardIsActive" binding:"required"`
	IsWhiteListed *bool `json:"isWhiteListed" binding:"required"`
	Limit float64      `json:"limit" binding:"required"`
	DenyList []string  `json:"denyList" binding:"required"`
}

func (a Account) CardIsDisabled() bool {
	return *a.CardIsActive == false
}

func (a Account) AmountAboveLimit(t transaction.Transaction) bool {
	return t.Amount > a.Limit
}

func (a Account) FirstTransactionAreUnsafe(t transaction.Transaction, l transaction.LastTransactions) bool {
	if len(l) == 0 {
		return logic.GetPercentege(t.Amount, a.Limit) > 90.0
	}
	return false
}

func (a Account) MerchantInDenyList(t transaction.Transaction) bool {
	allowed := false

	for _, merchants := range a.DenyList {
		if logic.RegexMatchMenchantName(t.Merchant, merchants) {
			allowed = true
			break
		}
	}
	return allowed
}
