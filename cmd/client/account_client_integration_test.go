//+build integration

package client

import (
	"testing"

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
