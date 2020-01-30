package client

import (
	"log"
	"os"

	resty "github.com/go-resty/resty/v2"
	"github.com/push-er/resty-client/cmd/model"
	"github.com/push-er/resty-client/internal/validation"
)

// FetchAccountEndpoint is path resource for account
const FetchAccountEndpoint = "/v1/organisation/accounts/{account_id}"

// AccountRestClient represent account rest client
type AccountRestClient struct {
	Client *resty.Client
	Logger *log.Logger
}

// New returns new AccountRestClient instance. Host url should resolve to valid url otherwise error is returned.
func New(hostURL string) (*AccountRestClient, error) {
	logger := log.New(os.Stdout, "Rest client: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	c := resty.New()

	if v, err := validation.ValidateURL(hostURL); !v || err != nil {
		return nil, err
	}
	c.SetHostURL(hostURL)

	rc := AccountRestClient{
		Client: c,
		Logger: logger,
	}

	return &rc, nil
}

// Fetch gets accound data for provided account id.
// If account id is not valid UUID or account is not found it returns error.
func (r *AccountRestClient) Fetch(accountID string) (model.Account, error) {
	var account model.Account
	if v, err := validation.ValidateUUID(accountID); !v || err != nil {
		return account, err
	}

	resp, err := r.Client.R().
		SetPathParams(map[string]string{
			"account_id": accountID,
		}).
		Get(FetchAccountEndpoint)
	if err != nil {
		r.Logger.Printf("Error fetching account: %v", err)
		return account, err
	}

	account, err = model.UnmarshallToAccount(resp.Body())
	if err != nil {
		r.Logger.Printf("Error parsing get account response: %v", err)
	}

	return account, err
}
