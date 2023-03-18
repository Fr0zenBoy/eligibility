package account

import (
	"testing"

	"github.com/Fr0zenBoy/authoraizer/transaction"
	"github.com/stretchr/testify/assert"
)

func TestCheckAmountIsAvaliable(t *testing.T) {
	someClient := Account{
		CardIsActive: true,
		IsWhiteListed: true,
		Limit: 5000.0,
		DenyList: []string{"moes"},
	}

	someTransaction := transaction.Transaction{
		Merchant: "Moes",
		Amount: 300.0,
	}

	smillerAmmount := transaction.Transaction{
		Merchant: "Moes",
		Amount: 6000.0,
	}

	t.Run("Test Amount is bigger than transaction ammount", func(t *testing.T) {
		result := someClient.CheckAmountAboveLimit(someTransaction)
		expected := true

		assert.Equal(t, result, expected)
	})

	t.Run("Test Amount is smaller than transaction ammount", func(t *testing.T) {
		result := someClient.CheckAmountAboveLimit(smillerAmmount)
		expected := false

		assert.Equal(t, result, expected)
	})
}

func TestAccountIsActive(t *testing.T) {
	ActiveClient := Account{
		CardIsActive: true,
	}

	DesactiveClient := Account{
		CardIsActive: false,
	}

	t.Run("Test active account", func(t *testing.T) {
		result := ActiveClient.CheckCardIsActive()
		expected := true
		assert.Equal(t, expected, result)
	})

	t.Run("Test desactive account", func(t *testing.T) {
		result := DesactiveClient.CheckCardIsActive()
		expected := false
		assert.Equal(t, expected, result)
	})
}

func TestFisrtTransactionIsSecure(t *testing.T) {

	validClient := Account{
		CardIsActive: true,
		IsWhiteListed: true,
		Limit: 5000.0,
		DenyList: []string{"moes"},
	}

	lowAmountTransaction := transaction.Transaction{
		Merchant: "Moes",
		Amount: 50.5,
		Time: "20023-02-02 10:08:05",
	}

	bigAmmountTransaction := transaction.Transaction{
		Merchant: "Moes",
		Amount: 50000.0,
		Time: "20023-02-02 10:08:05",
	}

	emptyLastTransactions := transaction.LastTransactions{}

	t.Run("Test a valid first transaction", func(t *testing.T) {
		result := validClient.CheckFirstTransactionSafe(lowAmountTransaction, emptyLastTransactions)
		expected := true
		assert.Equal(t, expected, result)
	})

	t.Run("Test a invalid first transaction", func(t *testing.T) {
		result := validClient.CheckFirstTransactionSafe(bigAmmountTransaction, emptyLastTransactions)
		expected := false
		assert.Equal(t, expected, result)
	})
}

func TestMatchMerchantsInDenylist(t *testing.T) {
	moes := "moes"
	moesUpperCase := "Moes"
	theCrab := "The Crab Shed"  
	theCrabLowerCase := "the crab shed"
	centralPerk := "Central Perk"
	centralPerkRandomCase := "cEnTraL PERk" 
	assert.Equal(t, matchMerchantsInDenylist(moes, moesUpperCase), true)
	assert.Equal(t, matchMerchantsInDenylist(theCrab, theCrabLowerCase), true)
	assert.Equal(t, matchMerchantsInDenylist(centralPerk, centralPerkRandomCase), true)
}

func TestCheckDenyList(t *testing.T) {

	validAccount := Account{
		CardIsActive: true,
		IsWhiteListed: true,
		Limit: 5000.0,
		DenyList: []string{"Moes", "The Crab Shed", "Central Perk"},
	}

	validTransaction := transaction.Transaction{
		Merchant: "The Cheesecake Factory",
		Amount: 50000.0,
		Time: "20023-02-02 10:08:05",
	}

	invalidTransaction := transaction.Transaction{
		Merchant: "The Crab Shed",
		Amount: 50000.0,
		Time: "20023-02-02 05:00:49",
	}
	t.Run("Test the merchant in transaction do not stay in the black list", func(t *testing.T){
		result := validAccount.CheckDenyList(validTransaction)
		expected := true
		assert.Equal(t, expected, result)
	})

	t.Run("Test the merchant in transaction stay in the black list", func(t *testing.T){
		result := validAccount.CheckDenyList(invalidTransaction)
		expected := false
		assert.Equal(t, expected, result)
	})
}
