//+build integration

package client

import (
	"testing"

	"github.com/google/uuid"
	"github.com/realmallaury/resty-client/internal/account"
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

	acc := account.GetMissingTestCreateAccount()
	_, err = accountRestClient.Create(acc)
	assert.Error(err, "Create(...) should return error")
}

func TestCreateIntegration(t *testing.T) {
	assert := assert.New(t)

	accountRestClient, err := New("http://localhost:8080")
	assert.Nil(err, "Error should be nil")

	acc := account.GetTestCreateAccount()
	expectedAccount := account.GetTestAccount()
	createdAccount, err := accountRestClient.Create(acc)

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
		acc := account.GetTestCreateAccount()
		account.ID = uuid.New().String()

		_, err := accountRestClient.Create(acc)
		assert.Nil(err, "Error should be nil")
	}

	accs, err := accountRestClient.List(0, 100)
	assert.Nil(err, "Error should be nil")
	assert.Len(accs, 10, "Accounts should contain 10 elements")

	accs, err = accountRestClient.List(0, 5)
	assert.Nil(err, "Error should be nil")
	assert.Len(accs, 5, "Accounts should contain 5 elements")

	accs, err = accountRestClient.List(1, 5)
	assert.Nil(err, "Error should be nil")
	assert.Len(accs, 5, "Accounts should contain 5 elements")
}
