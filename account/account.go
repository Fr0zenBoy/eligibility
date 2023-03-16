package account

import (
	"fmt"
	"regexp"

	"github.com/Fr0zenBoy/authoraizer/logic"
	"github.com/Fr0zenBoy/authoraizer/transaction"
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
	if len(l) == 0 {
		return logic.GetPercentege(t.Amount, a.Limit) < 90.0
	}
	return true
}

func matchMerchantsInDenylist(merchantInTransaction, merchantInDenyList string) bool {
	match, err := regexp.Compile(fmt.Sprintf("(?i)%s", merchantInDenyList))
	if err != nil {
		fmt.Println(err)
	}
	return match.MatchString(merchantInTransaction)
}

func (a Account) checkDenyList(t transaction.Transaction) bool {
	allowed := true

	for _, merchants := range a.DenyList {
		if matchMerchantsInDenylist(t.Merchant, merchants) {
			allowed = false
			break
		}
	}
	return allowed
}
