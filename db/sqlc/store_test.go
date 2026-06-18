package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testdb)

	account1 := createrandomaccount(t)
	account2 := createrandomaccount(t)
	//run n cuncorrent transactions
	n := 10
	errs := make(chan error)
	amount := int64(10)
	for i := 0; i < n; i++ {
		fromaccountID := account1.ID
		toaccountID := account2.ID
		if i%2 == 1 {
			fromaccountID = account2.ID
			toaccountID = account1.ID
		}
		go func() {
			ctx := context.Background()
			_, err := store.TransferTX(ctx, TransferTxParams{
				FromAccountID: fromaccountID,
				ToAccountID:   toaccountID,
				Amount:        amount,
			},
			)
			errs <- err
		}()
	}
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		// result := <-results
		// require.NotEmpty(t, result)
		// //check transfer
		// transfer := result.Transfer
		// require.NotEmpty(t, transfer)
		// require.Equal(t, account1.ID, transfer.FromAccountID.Int64)
		// require.Equal(t, account2.ID, transfer.ToAccountID.Int64)
		// require.NotZero(t, transfer.ID)
		// require.NotZero(t, transfer.Createdat)

		// _, err = store.GetTransfer(context.Background(), transfer.ID)
		// require.NoError(t, err)

		// //check entries
		// fromentry := result.FromEntry
		// require.NotEmpty(t, fromentry)
		// require.Equal(t, account1.ID, fromentry.AccountID)
		// require.Equal(t, -amount, fromentry.Amount)
		// require.NotZero(t, fromentry.ID)
		// require.NotZero(t, fromentry.Createdat)

		// _, err = store.GetEntry(context.Background(), fromentry.ID)
		// require.NoError(t, err)

		// toentry := result.ToEntry
		// require.NotEmpty(t, toentry)
		// require.Equal(t, account2.ID, toentry.AccountID)
		// require.Equal(t, amount, toentry.Amount)
		// require.NotZero(t, toentry.ID)
		// require.NotZero(t, toentry.Createdat)

		// _, err = store.GetEntry(context.Background(), toentry.ID)
		// require.NoError(t, err)
		// //check accounts
		// fromaccount := result.FromAccount
		// require.NotEmpty(t, fromaccount)

		// toaccount := result.ToAccount
		// require.NotEmpty(t, toaccount)

		// //check balance
		// diff1 := account1.Balance - fromaccount.Balance
		// diff2 := toaccount.Balance - account2.Balance
		// require.Equal(t, diff1, diff2)
		// require.True(t, diff1 > 0)
		// require.True(t, diff1%amount == 0)

		// k := int(diff1 / amount)
		// require.True(t, k >= 1 && k <= n)
	}
	updatedaccount1, err := testqueries.GetAccounts(context.Background(), account1.ID)
	require.NoError(t, err)
	updatedaccount2, err := testqueries.GetAccounts(context.Background(), account2.ID)
	require.NoError(t, err)
	require.Equal(t, account1.Balance, updatedaccount1.Balance)
	require.Equal(t, account2.Balance, updatedaccount2.Balance)
}
