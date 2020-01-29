//+build integration

package client

import (
	"testing"

	"github.com/push-er/resty-client/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestFetchIntegration(t *testing.T) {
	assert := assert.New(t)

	accountRestClient, err := New("http://localhost:8080")
	assert.Nil(err, "Error should be nil")
	assert.NotNil(accountRestClient, "accountRestClient should not be nil")

	res, err := accountRestClient.Fetch("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	assert.Nil(err, "Error should be nil")
	assert.EqualValues(utils.AccountResponseJSON, string(res.Body()), "Response should be same")
}
