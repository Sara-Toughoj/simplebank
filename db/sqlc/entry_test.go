package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/techschool/simplebank/util"
)

func createRandomEntry(t *testing.T) Entry {
	account := CreateRandomAccuont(t)
	args := CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, entry.AccountID, args.AccountID)
	require.Equal(t, entry.Amount, args.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry

}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestUpdateEntry(t *testing.T) {
	entry1 := createRandomEntry(t)

	args := UpdateEntrtyParams{
		ID:     entry1.ID,
		Amount: util.RandomMoney(),
	}

	entry2, err := testQueries.UpdateEntrty(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry2)
	require.Equal(t, entry2.ID, entry1.ID)
	require.Equal(t, entry2.AccountID, entry1.AccountID)
	require.Equal(t, entry2.Amount, args.Amount)
	require.WithinDuration(t, entry2.CreatedAt, entry1.CreatedAt, time.Second)

}

func TestDeleteEntry(t *testing.T) {
	entry1 := createRandomEntry(t)

	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}

func TestListEntries(t *testing.T) {
	args := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	entries2, err2 := testQueries.ListEntries(context.Background(), args)
	require.NoError(t, err2)
	require.Len(t, entries2, 5)
}
