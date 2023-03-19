package controller

import (
	"github.com/Fr0zenBoy/authoraizer/account"
	"github.com/Fr0zenBoy/authoraizer/transaction"
)

type authoraize interface {
	Allowed(o OutPut) OutPut
}

type Request struct {
	Account          account.Account              `json:"account" binding:"required"`
	Transaction      transaction.Transaction      `json:"transaction" binding:"required"`
	LastTransactions transaction.LastTransactions `json:"lastTransactions" binding:"required"`
}

type OutPut struct {
	Approved    bool     `json:"approved" binding:"required"`
	NewLimit    float64  `json:"newLimit" binding:"required"`
	DenyReasons []string `json:"denyReasons" binding:"required"`
}

func (r Request) Allowed(o OutPut) OutPut {
	account := r.Account
	transaction := r.Transaction
	lastTransactions := r.LastTransactions

	reasons := map[string]bool{
		"Card is not active!": account.CardIsDisabled(),
		"Amount are more than above the limit!": account.AmountAboveLimit(transaction),
		"The first transaction do not are more than obove 90% of limit!": account.FirstTransactionAreUnsafe(transaction, lastTransactions),
		"The Merchant stay present in the deny list!": account.MerchantInDenyList(transaction),
		"Limit of transaction per merchant exed!": transaction.TransactionLimitPerMerchants(lastTransactions),
		"Time beetween transactions execceded!": transaction.TimeLimitBeetweenTrancastions(lastTransactions),
	}

	for k, v := range reasons {
		if v {
			o.DenyReasons = append(o.DenyReasons, k)
		}
	}

	o.NewLimit = account.Limit

	if len(o.DenyReasons) == 0 {
		o.Approved = true
		o.NewLimit -= transaction.Amount
		o.DenyReasons = []string{}
		return o
	}
	return o
}

