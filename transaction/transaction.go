package transaction

import "github.com/Fr0zenBoy/authoraizer/logic"

const (
	maxTransactionsPerMerchant = 10
	safeTransactionTimeInSeconds = 120.0
	maxTransactionsUnderTimeLimit = 3
	timeLayoutTransaction = "2006-01-02 15:04:05"
)

type Transaction struct {
	Merchant string `json:"merchant" binding:"required"`
	Amount float64  `json:"amount" binding:"required"`
	Time string     `json:"time" binding:"required"`
}

type LastTransactions []Transaction

func (t Transaction) TransactionLimitPerMerchants(l LastTransactions) bool {
	var dealBreaker int
	if transations := len(l); transations > 0 {
		for i := 0; i < transations; i++ {
			if t.Merchant == l[i].Merchant {
				dealBreaker++
			}
		}
		return dealBreaker > maxTransactionsPerMerchant
	}
	return false
}

func (t Transaction) TimeLimitBeetweenTrancastions(l LastTransactions) bool {
	var dealBreaker int
	if transaction := len(l); transaction > 0 {
		for i := 0; i < transaction; i++ {
			if logic.ParseAndGetTimeDiff(t.Time, l[i].Time, timeLayoutTransaction) <= safeTransactionTimeInSeconds {
				dealBreaker++
			}
		}
		return dealBreaker > maxTransactionsUnderTimeLimit
	}
	return false
}
