package transaction


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
			if tDiff(tParse(t.Time), tParse( ltTimes.Time)) > 120.0 {
				dialBreaker--	
			}
		}
	}
	return dialBreaker > 0
}
