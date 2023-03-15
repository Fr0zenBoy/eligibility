package transaction

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)


func TestCheckLimitTransactionsPeerMerchant(t *testing.T) {
	
	someTransaction := Transaction{
		Merchant: "Moes",
		Amount: 300.0,
		Time: "2023-03-14 18:04:00",
	}

	emptylastTransactions := LastTransactions{}



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
	}

	t.Run("Test Check the last transaction list is empty and return true", func(t *testing.T){
		isEmpty := someTransaction.checkLimitTransactionPerMerchant(emptylastTransactions, 3)

		assert.Equal(t, true, isEmpty)
	})

	t.Run("Test Check Limit Transaction peer Merchant is greater than limit and return false", func(t *testing.T){
		isSame := someTransaction.checkLimitTransactionPerMerchant(sameMerchants, 3)

		assert.Equal(t, false, isSame)
	})
}

func TestCheckTimeBeetweenTransactions(t *testing.T){

	someTransaction := Transaction{
		Merchant: "Moes",
		Amount: 300.0,
		Time: "2023-03-14 18:04:00",
	}
	emptyLastTransactions := LastTransactions{}

	litterTime := LastTransactions{
		Transaction{
			Time: "2023-03-14 18:03:00",
		},
		Transaction{
			Time: "2023-03-14 18:02:00",
		},
		Transaction{
			Time: "2023-03-14 18:01:00",
		},
		Transaction{
			Time: "2023-03-14 18:00:00",
		},
		Transaction{
			Time: "2023-03-14 17:59:00",
		},

	}

	t.Run("Check the last transactions are empty list", func(t *testing.T){
		withEmpty := someTransaction.checkTimeBetweenTransactions(emptyLastTransactions, 3)

		assert.Equal(t, true, withEmpty)
	})

	t.Run("Check the last transactions have more transactions than limit", func(t *testing.T){
		limitExceeded := someTransaction.checkTimeBetweenTransactions(litterTime, 3)

		assert.Equal(t, false, limitExceeded)
	})
}
