package db

import (
	"bank/utils"
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Account {
	req := require.New(t)
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	req.NoError(err)
	req.NotEmpty(account)

	req.Equal(arg.Owner, account.Owner)
	req.Equal(arg.Balance, account.Balance)
	req.Equal(arg.Currency, account.Currency)

	req.NotZero(account.ID)
	req.NotZero(account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	req := require.New(t)

	req.NoError(err)
	req.NotEmpty(account2)

	req.Equal(account1.ID, account2.ID)
	req.Equal(account1.Owner, account2.Owner)
	req.Equal(account1.Balance, account2.Balance)
	req.Equal(account1.Currency, account2.Currency)

	req.WithinDuration(account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: utils.RandomMoney(),
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	req := require.New(t)
	req.NoError(err)
	req.NotEmpty(account2)

	req.Equal(account1.ID, account2.ID)
	req.Equal(account1.Owner, account2.Owner)
	req.Equal(arg.Balance, account2.Balance)
	req.Equal(account1.Currency, account2.Currency)

	req.WithinDuration(account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	req := require.New(t)
	req.NoError(err)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	req.Error(err)
	req.EqualError(err, sql.ErrNoRows.Error())
	req.Empty(account2)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}
