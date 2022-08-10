package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/techschool/simplebank/util"
)

func createRandomTransfer(t *testing.T) Transfer {
	args := CreateTransferParams{
		FromAccountID: CreateRandomAccuont(t).ID,
		ToAccountID:   CreateRandomAccuont(t).ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, transfer.ID)
	require.NotEmpty(t, transfer.CreatedAt)
	require.Equal(t, transfer.FromAccountID, args.FromAccountID)
	require.Equal(t, transfer.ToAccountID, args.ToAccountID)
	require.Equal(t, transfer.Amount, args.Amount)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomEntry(t)

}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.Equal(t, transfer2.ID, transfer1.ID)
	require.Equal(t, transfer2.FromAccountID, transfer1.FromAccountID)
	require.Equal(t, transfer2.ToAccountID, transfer1.ToAccountID)
	require.Equal(t, transfer2.Amount, transfer1.Amount)
	require.WithinDuration(t, transfer2.CreatedAt, transfer1.CreatedAt, time.Second)

}

func TestUpdateTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	args := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: util.RandomMoney(),
	}

	transfer2, err := testQueries.UpdateTransfer(context.Background(), args)

	require.NoError(t, err)
	require.Equal(t, transfer2.ID, args.ID)
	require.Equal(t, transfer2.Amount, args.Amount)
	require.Equal(t, transfer2.FromAccountID, transfer1.FromAccountID)
	require.Equal(t, transfer2.ToAccountID, transfer1.ToAccountID)

}

func TestListTransfers(t *testing.T) {

	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
	}

	args := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListEntries(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, transfers, 5)

}

func TestDeleteTransfers(t *testing.T) {
	transfer := createRandomTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)

	require.Empty(t, transfer2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
