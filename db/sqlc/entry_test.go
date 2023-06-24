package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/dasotd/gobank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T , account Account) Entry {
	// account := createRandomAccount(t)
	arg:= CreateEntryParams {
		Amount: util.RandomMoney(),
		AccountID: account.ID,

	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T)  {
	account := createRandomAccount(t)
	entry := createRandomEntry(t, account)

	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry.ID, entry2.ID)

	// return entry2
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	// entry := createRandomEntry(t, account)
	for i := 0; i < 10; i++{
		createRandomEntry(t, account)
	}
	arg := ListEntriesParams {
		AccountID: account.ID,
		Limit: 5,
		Offset: 2,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}

}

func TestUpdateEntry(t *testing.T){
	account := createRandomAccount(t)
	entry := createRandomEntry(t, account)

	args := UpdateEntryParams {
		ID: entry.ID,
		Amount: util.RandomMoney(),
	}
	entry, err := testQueries.UpdateEntry(context.Background(), args)
	require.NotEmpty(t, entry)
	require.NoError(t, err)
	
}

func TestDeleteEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry := createRandomEntry(t, account)

	err := testQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)

}