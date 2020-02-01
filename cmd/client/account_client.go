package client

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	resty "github.com/go-resty/resty/v2"

	"github.com/push-er/resty-client/cmd/model"
	"github.com/push-er/resty-client/internal/validation"
)

// FetchAccountEndpoint is path resource for getting account
const FetchAccountEndpoint = "/v1/organisation/accounts/{account_id}"

// CreateAccountEndpoint is path resource for creating account
const CreateAccountEndpoint = "/v1/organisation/accounts"

// DeleteAccountEndpoint is path resource for deleting account
const DeleteAccountEndpoint = "/v1/organisation/accounts/{account_id}?version={version}"

// AccountRestClient represent account rest client
type AccountRestClient struct {
	Client *resty.Client
	Logger *log.Logger
}

// New returns new AccountRestClient instance.
// Host url should resolve to valid url otherwise error is returned.
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
		r.Logger.Printf("error fetching account: %v", err)
		return account, err
	}

	if resp.IsError() {
		return account, fmt.Errorf("error fetching account: %v", string(resp.Body()))
	}

	account, err = model.UnmarshallToAccount(resp.Body())
	if err != nil {
		r.Logger.Printf("error parsing get account response: %v", err)
	}

	return account, nil
}

// Create creates new account.
// If account data is not valid or account there is error while creating account it returns error.
func (r *AccountRestClient) Create(account model.Account) (model.Account, error) {
	accountJSON, err := model.MarshallToAccount(&account)
	if err != nil {
		r.Logger.Printf("error marshalling account: %+v to JSON: %v", account, err)
		return account, err
	}

	resp, err := r.Client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(accountJSON).
		SetResult(&account).
		Post(CreateAccountEndpoint)
	if err != nil {
		r.Logger.Printf("error fetching account: %v", err)
		return account, err
	}

	created := make(map[string]interface{})

	err = json.Unmarshal(resp.Body(), &created)
	if err != nil {
		r.Logger.Printf("error parsing account creted response: %v", err)
		return account, err
	}

	account, err = model.UnmarshallToAccount(resp.Body())
	if err != nil {
		r.Logger.Printf("error parsing create account response: %v", err)
	}

	if resp.IsError() {
		return account, fmt.Errorf("error creating account: %v", string(resp.Body()))
	}

	return account, nil
}

// Delete deletes existing account.
// If account id is not valid UUID or account is not found it returns error.
func (r *AccountRestClient) Delete(accountID string, version int) error {
	if v, err := validation.ValidateUUID(accountID); !v || err != nil {
		return err
	}

	resp, err := r.Client.R().
		SetPathParams(map[string]string{
			"account_id": accountID,
			"version":    strconv.Itoa(version),
		}).
		Delete(DeleteAccountEndpoint)
	if err != nil {
		r.Logger.Printf("Error deleting account: %v", err)
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("rror deleting account: %v", string(resp.Body()))
	}

	return nil
}
