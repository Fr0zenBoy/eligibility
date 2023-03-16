package transaction

import (
	"github.com/Fr0zenBoy/authoraizer/logic"
)

const (
	maxTransation = 10
	MaxNumberTransactionsIn2M = 3
	TimeLayoutTransaction = "2006-01-02 15:04:05"
)

type Transaction struct {
	Merchant string `json:"merchant"`
	Amount float64  `json:"amount"`
	Time string     `json:"time"`
}

type LastTransactions []Transaction

func (t Transaction)checkLimitTransactionPerMerchant(l LastTransactions, maxTransation int) bool {
	dealBreaker := maxTransation

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

	return logic.TimeDiff(logic.TimeParse(ftime, TimeLayoutTransaction),
		                    logic.TimeParse(stime, TimeLayoutTransaction))
}

func (t Transaction) checkTimeBetweenTransactions(l LastTransactions, maxTimesLimit int) bool {
	dialBreaker := MaxNumberTransactionsIn2M
	if len(l) > 0 {
		for i := 0; i < MaxNumberTransactionsIn2M; i++ {
			if diffBeetweenTwoTime(t.Time, l[i].Time) <= 120.0 {
				dialBreaker--	
			}
		}
		return dialBreaker > 0
	}
	return true
}
