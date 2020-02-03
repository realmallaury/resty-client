//+build integration

package client

import (
	"testing"

	"github.com/google/uuid"
	"github.com/push-er/resty-client/cmd/model"
	"github.com/stretchr/testify/assert"
)

func TestFetchIntegration(t *testing.T) {
	assert := assert.New(t)

	accountRestClient, err := New("http://localhost:8080")
	assert.Nil(err, "Error should be nil")
	assert.NotNil(accountRestClient, "accountRestClient should not be nil")

	_, err = accountRestClient.Fetch("test")
	assert.Error(err, "Fetch(...) should return error")
}

func TestCreateBadAccountDataIntegration(t *testing.T) {
	assert := assert.New(t)

	accountRestClient, err := New("http://localhost:8080")
	assert.Nil(err, "Error should be nil")

	account := model.GetMissingTestCreateAccount()
	_, err = accountRestClient.Create(account)
	assert.Error(err, "Create(...) should return error")
}

func TestCreateIntegration(t *testing.T) {
	assert := assert.New(t)

	accountRestClient, err := New("http://localhost:8080")
	assert.Nil(err, "Error should be nil")

	account := model.GetTestCreateAccount()
	expectedAccount := model.GetTestAccount()
	createdAccount, err := accountRestClient.Create(account)

	assert.Nil(err, "Error should be nil")
	assert.EqualValues(expectedAccount, createdAccount, "Response should be same")
}

func TestDeleteIntegration(t *testing.T) {
	assert := assert.New(t)

	accountRestClient, err := New("http://localhost:8080")
	assert.Nil(err, "Error should be nil")

	err = accountRestClient.Delete("cd27e265-9605-4b4b-a0e5-3003ea9cc4dc", 0)
	assert.Nil(err, "Error should be nil")
}

func TestLisIntegration(t *testing.T) {
	assert := assert.New(t)

	accountRestClient, err := New("http://localhost:8080")
	assert.Nil(err, "Error should be nil")

	for i := 1; i <= 10; i++ {
		account := model.GetTestCreateAccount()
		account.ID = uuid.New().String()

		_, err := accountRestClient.Create(account)
		assert.Nil(err, "Error should be nil")
	}

	accounts, err := accountRestClient.List(0, 100)
	assert.Nil(err, "Error should be nil")
	assert.Len(accounts, 10, "Accounts should contain 10 elements")

	accounts, err = accountRestClient.List(0, 5)
	assert.Nil(err, "Error should be nil")
	assert.Len(accounts, 5, "Accounts should contain 5 elements")

	accounts, err = accountRestClient.List(1, 5)
	assert.Nil(err, "Error should be nil")
	assert.Len(accounts, 5, "Accounts should contain 5 elements")
}
