package main

// 1. The transaction amount should not be above limit
// 2. No transaction should be approved when the card is blocked
// 3. The first transaction shouldn't be above 90% of the limit
// 4. There should not be more than 10 transactions on the same merchant
// 5. Merchant blacklist
// 6. There should not be more than 3 transactions on a 2 minutes interval

import (
	"fmt"
	"time"
	"encoding/json"
)

type RequestTransaction struct {
	Account          `json:"account"`
	Transaction      `json:"transaction"`
	LastTransactions `json:"lastTransactions"`
}

type ResultOfTransaction struct {
	Approved bool        `json:"approved"`
	NewLimit float64     `json:"newLimit"`
	DenyReasons []string `json:"denyReasons"`
}

const (
	maxTransation = 10
	maxTimesLimit = 3
	TimeLayoutTransaction = "2006-01-02 15:04:05"
)



// 4. There should not be more than 10 transactions on the same merchant
func checkLimitTransactionPerMerchant(t Transaction, l LastTransactions) bool {
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


func tParse(d string) (t time.Time) {
	t, err := time.Parse(TimeLayoutTransaction, d)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

func tDiff(actual, previous time.Time) float64 {
	diff := actual.Sub(previous)
	infloat := diff.Abs().Hours()

	return infloat
}

func (t Transaction) checkTimeBetweenTransactions(l LastTransactions) bool {

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


func main() {
	request1 := `{
  "account": {
    "cardIsActive": "true",
    "limit": "5000",
    "danyList": ["Moes"],
    "isWhitelisted": "true"
  },
  "transaction": {
    "merchant": "MacLarens",
    "amount": "2000",
    "time": "2019-06-19 21:04:00"
  },
  "lastTransactions": [
    {
      "merchant": "MacLarens",
      "amount": "1000",
      "time": "2019-06-19 21:01:00"
    }
  ]
}`
	var transation map[string]any
	json.Unmarshal([]byte(request1), &transation)
	fmt.Println("unmarchal", transation)

	dat := transation["account"].(map[string]any)
	fmt.Println("ok isso é uma conta: ", dat)
	
}
