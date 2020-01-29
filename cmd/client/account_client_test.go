package client

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/push-er/resty-client/internal/utils"
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
		`/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc`,
		httpmock.NewStringResponder(200, utils.AccountResponseJSON),
	)

	accountRestClient, err := New("http://test")
	httpmock.ActivateNonDefault(accountRestClient.Client.GetClient())

	res, err := accountRestClient.Fetch("test")
	assert.Error(err, "Fetch(...) should return error")
	assert.Nil(res, "res should be nil")

	res, err = accountRestClient.Fetch("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	assert.Nil(err, "Error should be nil")
	assert.EqualValues(utils.AccountResponseJSON, string(res.Body()), "Response should be same")
}
