package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

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

	url, server := MockHTTPServerEndpoint(
		http.MethodGet,
		"/v1/organisation/accounts/cd27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		model.AccountResponseJSON,
		http.StatusOK,
	)
	defer server.Close()

	accountRestClient, err := New(url)
	assert.Nil(err, "Error should be nil")

	_, err = accountRestClient.Fetch("test")
	assert.Error(err, "Fetch(...) should return error")

	acc, err := accountRestClient.Fetch("cd27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	assert.Nil(err, "Error should be nil")
	assert.EqualValues(model.GetTestAccount(), acc, "Response should be same")
}

func TestCreateBadAccountData(t *testing.T) {
	assert := assert.New(t)

	accountRestClient, err := New("http://test")
	assert.Nil(err, "Error should be nil")

	account := model.GetMissingTestCreateAccount()
	_, err = accountRestClient.Create(account)
	assert.Error(err, "Creaate(...) should return error")
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	url, server := MockHTTPServerEndpoint(
		http.MethodPost,
		"/v1/organisation/accounts",
		model.AccountResponseJSON,
		http.StatusOK,
	)
	defer server.Close()

	accountRestClient, err := New(url)
	assert.Nil(err, "Error should be nil")

	account := model.GetTestCreateAccount()
	expectedAccount := model.GetTestAccount()
	createdAccount, err := accountRestClient.Create(account)

	assert.Nil(err, "Error should be nil")
	assert.EqualValues(expectedAccount, createdAccount, "Response should be same")
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)

	url, server := MockHTTPServerEndpoint(
		http.MethodDelete,
		"/v1/organisation/accounts/cd27e265-9605-4b4b-a0e5-3003ea9cc4dc?version=0",
		"",
		http.StatusOK,
	)
	defer server.Close()

	accountRestClient, err := New(url)
	assert.Nil(err, "Error should be nil")

	err = accountRestClient.Delete("cd27e265-9605-4b4b-a0e5-3003ea9cc4dc", 0)
	assert.Nil(err, "Error should be nil")
}

func TestList(t *testing.T) {
	assert := assert.New(t)

	url, server := MockHTTPServerEndpoint(
		http.MethodGet,
		"/v1/organisation/accounts?page%5Bnumber%5D%3A=0&page%5Bsize%5D=100",
		model.AccountsResponseJSON,
		http.StatusOK,
	)
	defer server.Close()

	accountRestClient, err := New(url)
	assert.Nil(err, "Error should be nil")

	accounts, err := accountRestClient.List(0, 100)

	assert.Nil(err, "Error should be nil")
	assert.Len(accounts, 10, "Accounts should contain 10 elements")
}

func MockHTTPServerEndpoint(method, path, body string, statusCode int) (string, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == method && req.URL.String() == path {
			rw.WriteHeader(statusCode)
			_, _ = rw.Write([]byte(body))
		} else {
			http.Error(rw, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	}))

	return server.URL, server
}
