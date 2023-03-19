package transaction

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestTransactionLimitPerMerchants(t *testing.T) {
	
	someTransaction := Transaction{
		Merchant: "Moes",
	}

	emptylastTransactions := LastTransactions{}

	lessMerchants := LastTransactions{
		Transaction{
			Merchant: "Moes",
		},
		Transaction{
			Merchant: "Moes",
		},
		Transaction{
			Merchant: "Moes",
		},
		Transaction{
			Merchant: "Moes",
		},
		Transaction{
			Merchant: "Moes",
		},
	}

	sameMerchants := LastTransactions{
		Transaction{
			Merchant: "Moes",
		},
		Transaction{
			Merchant: "Moes",
		},
		Transaction{
			Merchant: "Moes",
		},
		Transaction{
			Merchant: "Moes",
		},
		Transaction{
			Merchant: "Moes",
		},
		Transaction{
			Merchant: "Moes",
		},
		Transaction{
			Merchant: "Moes",
		},
		Transaction{
			Merchant: "Moes",
		},
		Transaction{
			Merchant: "Moes",
		},
		Transaction{
			Merchant: "Moes",
		},
		Transaction{
			Merchant: "Moes",
		},
	}

	t.Run("Test Check the last transaction list is empty and return false", func(t *testing.T) {
		isEmpty := someTransaction.TransactionLimitPerMerchants(emptylastTransactions)

		assert.Equal(t, false, isEmpty)
	})


	t.Run("Test Threshold Check Merchant of the same name is less than threshold and returns false", func(t *testing.T) {
		isEmpty := someTransaction.TransactionLimitPerMerchants(lessMerchants)

		assert.Equal(t, false,isEmpty)
	})

	t.Run("Test limit check Merchant of the same name is greater than threshold and returns true", func(t *testing.T) {
		isSame := someTransaction.TransactionLimitPerMerchants(sameMerchants)

		assert.Equal(t, true, isSame)
	})
}


func TestTimeLimitBeetweenTrancastions(t *testing.T){

	someTransaction := Transaction{
		Time: "2023-03-14 18:04:00",
	}

	emptyLastTransactions := LastTransactions{}

	greaterTime := LastTransactions{
		Transaction{
			Time: "2023-03-14 18:03:00",
		},
		Transaction{
			Time: "2023-03-14 17:02:50",
		},
		Transaction{
			Time: "2023-03-14 16:02:30",
		},
		Transaction{
			Time: "2023-03-14 15:02:28",
		},
	}

	quickTime := LastTransactions{
		Transaction{
			Time: "2023-03-14 18:03:00",
		},
		Transaction{
			Time: "2023-03-14 18:02:50",
		},
		Transaction{
			Time: "2023-03-14 18:02:40",
		},
		Transaction{
			Time: "2023-03-14 18:02:30",
		},
	}

	t.Run("Test the last transactions are empty list and returns false", func(t *testing.T) {
		withEmpty := someTransaction.TimeLimitBeetweenTrancastions(emptyLastTransactions)

		assert.Equal(t, withEmpty, false)
	})

	t.Run("Test the last transactions have a time greater than threshold between them and returns false", func(t *testing.T) {
		limitExceeded := someTransaction.TimeLimitBeetweenTrancastions(greaterTime)

		assert.Equal(t, limitExceeded, false)
	})

	t.Run("Test the last transactions have a time less than threshold between them and returns true", func(t *testing.T) {
		limitExceeded := someTransaction.TimeLimitBeetweenTrancastions(quickTime)

		assert.Equal(t, limitExceeded, true)
	})
}
