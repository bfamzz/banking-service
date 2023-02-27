package db

import (
	"context"
	"testing"
	"time"

	"github.com/bfamzz/banking-service/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID: toAccount.ID,
		Amount: util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, fromAccount.ID, transfer.FromAccountID)
	require.Equal(t, toAccount.ID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func createRandomTransfers(t *testing.T, n int) (Account, Account) {
	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID: toAccount.ID,
		Amount: util.RandomMoney(),
	}

	for i := 0; i < n; i++ {
		testQueries.CreateTransfer(context.Background(), arg)
	}

	return fromAccount, toAccount
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer1 := createRandomTransfer(t)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)

	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	createRandomTransfers(t, 10)

	arg := ListTransfersParams{
		Limit: 5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)

	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

func TestListTransfersByFromAccountID(t *testing.T) {
	fromAccount, _ := createRandomTransfers(t, 10)

	arg := ListTransfersByFromAccountIDParams{
		FromAccountID: fromAccount.ID,
		Limit: 5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfersByFromAccountID(context.Background(), arg)
	
	require.NoError(t, err)

	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, fromAccount.ID, transfer.FromAccountID)
	}
}

func TestListTransfersByFromToID(t *testing.T) {
	_, toAccount := createRandomTransfers(t, 10)

	arg := ListTransfersByToAccountIDParams{
		ToAccountID: toAccount.ID,
		Limit: 5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfersByToAccountID(context.Background(), arg)
	
	require.NoError(t, err)

	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, toAccount.ID, transfer.ToAccountID)
	}
}