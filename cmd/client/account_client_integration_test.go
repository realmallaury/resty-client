//+build integration

package client

import (
	"testing"

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
