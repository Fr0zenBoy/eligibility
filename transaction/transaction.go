package transaction

import (
	"github.com/Fr0zenBoy/authoraizer/logic"
)

const (
	maxTransation = 10
	maxTimesLimit = 3
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
		return maxTransation > 0
	}
	return false 
}

func (t Transaction) checkTimeBetweenTransactions(l LastTransactions, maxTimesLimit int) bool {
	dialBreaker := maxTimesLimit
	if lt := len(l); lt > 0 {
		for _, ltTimes := range l {
			if logic.TimeDiff(logic.TimeParse(t.Time, TimeLayoutTransaction), logic.TimeParse(ltTimes.Time, TimeLayoutTransaction)) > 120.0 {
				dialBreaker--	
			}
		}
	}
	return dialBreaker > 0
}
