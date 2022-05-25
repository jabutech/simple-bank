package db

import (
	"context"
	"database/sql"
	"github/jabutech/simplebank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Function for create random account
func createRandomAccount(t *testing.T) Accounts {
	// Create argument
	arg := CreateAccountParams{
		// Data obtained from random result from utils/random
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	// Run test
	account, err := testQueries.CreateAccount(context.Background(), arg)
	// Result must be no error
	require.NoError(t, err)
	// Result mus be no empty
	require.NotEmpty(t, account)

	// Test compatibilty
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	// Test not zero data
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	// Return account
	return account
}

// Test create account
func TestCreateAccount(t *testing.T) {
	// Call function createRandomAccount for test create account
	createRandomAccount(t)
}

// Test get account
func TestGetAccount(t *testing.T) {
	// Create new random account
	account1 := createRandomAccount(t)
	// Test get account by id account 1
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	// Result must be no error
	require.NoError(t, err)
	// Result must be get account 2 no empty
	require.NotEmpty(t, account2)

	// Test compatibilty
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)

	// Duration createdAt must be same
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

// Test update account
func TestUpdateAccount(t *testing.T) {
	// Create new random account
	account1 := createRandomAccount(t)

	// Create new object
	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	// Run test update
	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	// Result must be no error
	require.NoError(t, err)
	// Result must be get account 2 no empty
	require.NotEmpty(t, account2)

	// Test compatibilty
	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance) // arg.Balance and account2.Balance must be same
	require.Equal(t, account1.Currency, account2.Currency)

	// Duration createdAt must be same
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

// Test delete account
func TestDeleteAccount(t *testing.T) {
	// Create new random account
	account1 := createRandomAccount(t)
	// Run test delete account
	err := testQueries.DeleteAccounts(context.Background(), account1.ID)

	// Result must be no error
	require.NoError(t, err)

	// Run test get account
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	// Result must be error
	require.Error(t, err)
	// Error must be `sql.Err.NoRows.Error()`
	require.EqualError(t, err, sql.ErrNoRows.Error())
	// Account 2 must be empty
	require.Empty(t, account2)
}

// Test get list account
func TestListAccount(t *testing.T) {
	// Create new ten(10) random account
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	// Create argument list account param, for get only 5 account
	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	// Test run get list account
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	// Result must be no erro
	require.NoError(t, err)
	// Result must be length five(5)
	require.Len(t, accounts, 5)

	// All account must be no empty
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
