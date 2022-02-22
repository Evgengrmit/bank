package db

import (
	"bank/utils"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateAccount(t *testing.T) {
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
}
