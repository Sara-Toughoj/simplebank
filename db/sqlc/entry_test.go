package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/techschool/simplebank/util"
)

func TestCreateEntry(t *testing.T) {
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

}
