package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/Fr0zenBoy/authoraizer/account"
	"github.com/Fr0zenBoy/authoraizer/controller"
	"github.com/Fr0zenBoy/authoraizer/transaction"
	"github.com/gin-gonic/gin"

	"testing"

	"github.com/stretchr/testify/assert"
)

func newTrue() *bool {
	b := true
	return &b
}

func newFalse() *bool {
	b := false
	return &b
}
func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestAuthoraizerHandler(t *testing.T) {
	t.Run("Test Succesful Request", func(t *testing.T) {

		payload := controller.Request{
			account.Account{
				CardIsActive: newTrue(),
				Limit: 5000,
				DenyList: []string{"Moes"},
				IsWhiteListed: newTrue(),
			},
			transaction.Transaction{
				Merchant: "Maclarens",
				Amount: 20.0,
				Time: "2019-06-19 21:04:00",
			},
			transaction.LastTransactions{
				transaction.Transaction{
					Merchant: "Maclarens",
					Amount: 20.0,
					Time: "2019-06-19 21:04:00",
				},
			},
		}

		result := controller.OutPut{
			Approved: true,
			NewLimit: 4980,
			DenyReasons: []string{},
		}

		r := SetUpRouter()
		r.POST("/api/authoraizer", AuthoraizerHandler)
		jsonValue, _ := json.Marshal(payload)
		resultValue, _ := json.Marshal(result)
		req, _ := http.NewRequest("POST", "/api/authoraizer", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t , http.StatusAccepted, w.Code)
		assert.Equal(t, resultValue, w.Body.Bytes())
	})

	t.Run("Test erro on json", func(t *testing.T) {
		payload := controller.Request{
			account.Account{
				Limit: 5000,
				DenyList: []string{"Moes"},
				IsWhiteListed: newTrue(),
			},
			transaction.Transaction{
				Merchant: "Maclarens",
				Amount: 20.0,
				Time: "2019-06-19 21:04:00",
			},
			transaction.LastTransactions{
				transaction.Transaction{
					Merchant: "Maclarens",
					Amount: 20.0,
					Time: "2019-06-19 21:04:00",
				},
			},
		}

		error := gin.H{"error":  "Key: 'Request.Account.CardIsActive' Error:Field validation for 'CardIsActive' failed on the 'required' tag",
			             "message": "Invalid inputs. please check your inputs"}

		r := SetUpRouter()
		r.POST("/api/authoraizer", AuthoraizerHandler)
		jsonValue, _ := json.Marshal(payload)
		errorMap, _ := json.Marshal(error)
		req, _ := http.NewRequest("POST", "/api/authoraizer", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t , http.StatusBadRequest, w.Code)
		assert.Equal(t, errorMap, w.Body.Bytes())
	})

	t.Run("Test handle some errors in request", func(t *testing.T) {
		
		payload := controller.Request{
			account.Account{
				CardIsActive: newFalse(),
				Limit: 5000,
				DenyList: []string{"Maclarens"},
				IsWhiteListed: newTrue(),
			},
			transaction.Transaction{
				Merchant: "Maclarens",
				Amount: 20000.0,
				Time: "2019-06-19 21:04:00",
			},
			transaction.LastTransactions{
				transaction.Transaction{
					Merchant: "Maclarens",
					Amount: 20.0,
					Time: "2019-06-19 21:04:00",
				},

				transaction.Transaction{
					Merchant: "Maclarens",
					Amount: 20.0,
					Time: "2019-06-19 21:04:00",
				},
				transaction.Transaction{
					Merchant: "Maclarens",
					Amount: 20.0,
					Time: "2019-06-19 21:04:00",
				},
				transaction.Transaction{
					Merchant: "Maclarens",
					Amount: 20.0,
					Time: "2019-06-19 21:04:00",
				},
				transaction.Transaction{
					Merchant: "Maclarens",
					Amount: 20.0,
					Time: "2019-06-19 21:04:00",
				},
				transaction.Transaction{
					Merchant: "Maclarens",
					Amount: 20.0,
					Time: "2019-06-19 21:04:00",
				},
				transaction.Transaction{
					Merchant: "Maclarens",
					Amount: 20.0,
					Time: "2019-06-19 21:04:00",
				},
				transaction.Transaction{
					Merchant: "Maclarens",
					Amount: 20.0,
					Time: "2019-06-19 21:04:00",
				},
				transaction.Transaction{
					Merchant: "Maclarens",
					Amount: 20.0,
					Time: "2019-06-19 21:04:00",
				},
				transaction.Transaction{
					Merchant: "Maclarens",
					Amount: 20.0,
					Time: "2019-06-19 21:04:00",
				},
				transaction.Transaction{
					Merchant: "Maclarens",
					Amount: 20.0,
					Time: "2019-06-19 21:04:00",
				},
			},
		}

		result := controller.OutPut{
			Approved: false,
			NewLimit: 5000,
			DenyReasons: []string{"The Merchant stay present in the deny list!","Limit of transaction per merchant exed!","Time beetween transactions execceded!","Card is not active!","Amount are more than above the limit!"},
		}

		r := SetUpRouter()
		r.POST("/api/authoraizer", AuthoraizerHandler)
		jsonValue, _ := json.Marshal(payload)
		resultValue, _ := json.Marshal(result)
		req, _ := http.NewRequest("POST", "/api/authoraizer", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t , http.StatusAccepted, w.Code)
		assert.ElementsMatch(t, resultValue, w.Body.Bytes())
	})
}
