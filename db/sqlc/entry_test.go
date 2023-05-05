package db

import (
	"context"
	"testing"

	"github.com/A-Doubt/doubtful_bank_v1/utils"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount:    utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotEmpty(t, entry.ID)
	require.NotEmpty(t, entry.Amount)
	require.NotEmpty(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t, account)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)
}

func TestListEntry(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg1 := ListEntriesParams{AccountID: account.ID, Limit: 5, Offset: 0}
	arg2 := ListEntriesParams{AccountID: account.ID, Limit: 5, Offset: 5}

	entries, err := testQueries.ListEntries(context.Background(), arg1)
	require.NoError(t, err)
	require.NotEmpty(t, entries)
	require.Len(t, entries, 5)

	entries2, err := testQueries.ListEntries(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, entries2)
	require.Len(t, entries2, 5)

	require.NotEqual(t, entries, entries2)
}
