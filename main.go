package main

// 1. The transaction amount should not be above limit
// 2. No transaction should be approved when the card is blocked
// 3. The first transaction shouldn't be above 90% of the limit
// 4. There should not be more than 10 transactions on the same merchant
// 5. Merchant blacklist
// 6. There should not be more than 3 transactions on a 2 minutes interval

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Fr0zenBoy/authoraizer/transaction"
	"github.com/Fr0zenBoy/authoraizer/account"
)

type Request struct {
	Account          account.Account              `json:"account"`
	Transaction      transaction.Transaction      `json:"transaction"`
	LastTransactions transaction.LastTransactions `json:"lastTransactions"`
}

type ResultOfTransaction struct {
	Approved bool        `json:"approved"`
	NewLimit float64     `json:"newLimit"`
	DenyReasons []string `json:"denyReasons"`
}

var requests = []Request{}

func postRequest(c *gin.Context) {
	var newResquest Request
	if err := c.BindJSON(&newResquest); err != nil {
		return
	}

	requests = append(requests, newResquest)
	c.IndentedJSON(http.StatusCreated, newResquest)
}

func main() {
	router := gin.Default()
	router.POST("/api/authoraizer", postRequest)
	router.Run("localhost:8080")
// 	request1 := `{
//   "account": {
//     "cardIsActive": "true",
//     "limit": "5000",
//     "danyList": ["Moes"],
//     "isWhitelisted": "true"
//   },
//   "transaction": {
//     "merchant": "MacLarens",
//     "amount": "2000",
//     "time": "2019-06-19 21:04:00"
//   },
//   "lastTransactions": [
//     {
//       "merchant": "MacLarens",
//       "amount": "1000",
//       "time": "2019-06-19 21:01:00"
//     }
//   ]
// }`
	// var transation map[string]any
	// json.Unmarshal([]byte(request1), &transation)
	// fmt.Println("unmarchal", transation)

	// dat := transation["account"].(map[string]any)
	// fmt.Println("ok isso é uma conta: ", dat)

}
