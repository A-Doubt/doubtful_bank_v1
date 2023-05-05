package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/A-Doubt/doubtful_bank_v1/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:   utils.RandomOwner(),
		Balance: utils.RandomMoney(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)

	require.NotEmpty(t, account)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1, account2)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	newBalance := utils.RandomMoney()

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: newBalance,
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, newBalance, account2.Balance)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{Limit: 5, Offset: 0}
	arg2 := ListAccountsParams{Limit: 5, Offset: 5}
	arg3 := ListAccountsParams{Limit: 10, Offset: 0}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.Len(t, accounts, 5)

	accounts2, err := testQueries.ListAccounts(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, accounts2)
	require.Len(t, accounts2, 5)

	accounts3, err := testQueries.ListAccounts(context.Background(), arg3)
	require.NoError(t, err)
	for _, acc := range accounts3 {
		require.NotEmpty(t, acc)
	}
	require.Len(t, accounts3, 10)

	require.NotEqual(t, accounts, accounts2)
}
