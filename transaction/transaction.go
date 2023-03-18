package transaction

import "github.com/Fr0zenBoy/authoraizer/logic"

const (
	maxTransactionsPeerMerchant = 10
	safeTransactionTimeInSeconds = 120.0
	maxTransactionUnderTimeLimit = 3
	timeLayoutTransaction = "2006-01-02 15:04:05"
)

type Transaction struct {
	Merchant string `json:"merchant" binding:"required"`
	Amount float64  `json:"amount" binding:"required"`
	Time string     `json:"time" binding:"required"`
}

type LastTransactions []Transaction

func (t Transaction) CheckLimitTransactionPerMerchant(l LastTransactions) bool {
	dealBreaker := maxTransactionsPeerMerchant

	if transations := len(l); transations > 0 {
		for i := 0; i < transations; i++ {
			if t.Merchant == l[i].Merchant {
				dealBreaker--
			}
		}
		return dealBreaker >= 0
	}
	return true 
}

func diffBeetweenTwoTime(ftime, stime string) float64 {

	return logic.TimeDiff(logic.TimeParse(ftime, timeLayoutTransaction),
		                    logic.TimeParse(stime, timeLayoutTransaction))
}

func (t Transaction) CheckTimeBetweenTransactions(l LastTransactions) bool {
	dialBreaker := 0
	if transaction := len(l); transaction > 0 {
		for i := 0; i < transaction; i++ {
			if diffBeetweenTwoTime(t.Time, l[i].Time) <= safeTransactionTimeInSeconds {
				dialBreaker++
			}
		}
		return dialBreaker <= maxTransactionUnderTimeLimit
	}
	return true
}
