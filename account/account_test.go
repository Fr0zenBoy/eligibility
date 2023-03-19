package account

import (
	"testing"

	"github.com/Fr0zenBoy/authoraizer/transaction"
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

func TestAccountIsDisable(t *testing.T) {

	ActiveCard := Account{
		CardIsActive: newTrue(),
	}

	DesactiveCard := Account{
		CardIsActive: newFalse(),
	}

	t.Run("Test account is disabled and returens false", func(t *testing.T) {
		actual := ActiveCard.CardIsDisabled()
		assert.Equal(t, false, actual)
	})

	t.Run("Test account is disabled and returens true", func(t *testing.T) {
		actual := DesactiveCard.CardIsDisabled()
		assert.Equal(t, true, actual)
	})
}

func TestAmountAboveLimit(t *testing.T) {
	someClient := Account{
		Limit: 5000.0,
	}

	smallValue := transaction.Transaction{
		Merchant: "Moes",
		Amount: 300.0,
	}

	bigValue := transaction.Transaction{
		Merchant: "Moes",
		Amount: 6000.0,
	}

	t.Run("Test Amount is less than transaction ammount and returns false", func(t *testing.T) {
		result := someClient.AmountAboveLimit(smallValue)

		assert.Equal(t, false, result)
	})

	t.Run("Test Amount is bigger than transaction amount and returns true", func(t *testing.T) {
		result := someClient.AmountAboveLimit(bigValue)

		assert.Equal(t, true, result)
	})
}

func TestFirstTransactionAreUnsafe(t *testing.T) {

	validClient := Account{
		Limit: 5000.0,
	}

	emptyLastTransactions := transaction.LastTransactions{}

	lowAmountTransaction := transaction.Transaction{
		Amount: 50.5,
	}

	bigAmmountTransaction := transaction.Transaction{
		Amount: 50000.0,
	}

	t.Run("Test the first transaction has an amount less than the threshold and return false", func(t *testing.T) {
		actual := validClient.FirstTransactionAreUnsafe(lowAmountTransaction, emptyLastTransactions)
		assert.Equal(t, false, actual)
	})

	t.Run("Test the first transaction has an amount greater than the threshold and return true", func(t *testing.T) {
		actual := validClient.FirstTransactionAreUnsafe(bigAmmountTransaction, emptyLastTransactions)
		assert.Equal(t, true, actual)
	})
}

func TestMerchantInDenyList(t *testing.T) {

	someAccount := Account{
		DenyList: []string{"Moes", "The Crab Shed", "Central Perk"},
	}

	validTransaction := transaction.Transaction{
		Merchant: "The Cheesecake Factory",
	}

	invalidTransaction := transaction.Transaction{
		Merchant: "The Crab Shed",
	}

	t.Run("Test the merchant is on the denylist and return true", func(t *testing.T) {
		actual := someAccount.MerchantInDenyList(invalidTransaction)
		assert.Equal(t, true, actual)
	})

	t.Run("Test the merchant is not on the denylist and return false", func(t *testing.T) {
		actual := someAccount.MerchantInDenyList(validTransaction)
		assert.Equal(t, false, actual)
	})

}
