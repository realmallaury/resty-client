package client

import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/push-er/resty-client/cmd/model"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)

	accountRestClient, err := New("http://test")
	assert.Nil(err, "Error should be nil")
	assert.NotNil(accountRestClient, "accountRestClient should not be nil")

	accountRestClient, err = New("http:test")
	assert.Error(err, "New(...) should return error")
	assert.Nil(accountRestClient, "accountRestClient should be nil")
}

func TestFetch(t *testing.T) {
	assert := assert.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"GET",
		`/v1/organisation/accounts/cd27e265-9605-4b4b-a0e5-3003ea9cc4dc`,
		httpmock.NewStringResponder(http.StatusOK, model.AccountResponseJSON),
	)

	accountRestClient, err := New("http://test")
	assert.Nil(err, "Error should be nil")
	httpmock.ActivateNonDefault(accountRestClient.Client.GetClient())

	_, err = accountRestClient.Fetch("test")
	assert.Error(err, "Fetch(...) should return error")

	acc, err := accountRestClient.Fetch("cd27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	assert.Nil(err, "Error should be nil")
	assert.EqualValues(model.GetTestAccount(), acc, "Response should be same")
}

func TestCreateBadAccountData(t *testing.T) {
	assert := assert.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	accountRestClient, err := New("http://test")
	assert.Nil(err, "Error should be nil")
	httpmock.ActivateNonDefault(accountRestClient.Client.GetClient())

	account := model.GetMissingTestCreateAccount()
	_, err = accountRestClient.Create(account)
	assert.Error(err, "Creaate(...) should return error")
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"POST",
		`/v1/organisation/accounts`,
		httpmock.NewStringResponder(http.StatusOK, model.AccountResponseJSON),
	)

	accountRestClient, err := New("http://test")
	assert.Nil(err, "Error should be nil")
	httpmock.ActivateNonDefault(accountRestClient.Client.GetClient())

	account := model.GetTestCreateAccount()
	expectedAccount := model.GetTestAccount()
	createdAccount, err := accountRestClient.Create(account)

	assert.Nil(err, "Error should be nil")
	assert.EqualValues(expectedAccount, createdAccount, "Response should be same")
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		"DELETE",
		`/v1/organisation/accounts/cd27e265-9605-4b4b-a0e5-3003ea9cc4dc?version=0`,
		httpmock.NewStringResponder(http.StatusNoContent, ""),
	)

	accountRestClient, err := New("http://test")
	assert.Nil(err, "Error should be nil")
	httpmock.ActivateNonDefault(accountRestClient.Client.GetClient())

	err = accountRestClient.Delete("cd27e265-9605-4b4b-a0e5-3003ea9cc4dc", 0)
	assert.Nil(err, "Error should be nil")
}
