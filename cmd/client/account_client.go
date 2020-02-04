package client

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	resty "github.com/go-resty/resty/v2"

	"github.com/realmallaury/resty-client/cmd/model"
	"github.com/realmallaury/resty-client/internal/validation"
)

// AccountEndpoint is path resource for account
const AccountEndpoint = "/v1/organisation/accounts/{account_id}"

// AccountsEndpoint is path resource for accounts
const AccountsEndpoint = "/v1/organisation/accounts"

// AccountRestClient represent account rest client
type AccountRestClient struct {
	client *resty.Client
	logger *log.Logger
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
		client: c,
		logger: logger,
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

	resp, err := r.client.R().
		SetPathParams(map[string]string{
			"account_id": accountID,
		}).
		Get(AccountEndpoint)
	if err != nil {
		r.logger.Printf("error fetching account: %v", err)
		return account, err
	}

	if resp.IsError() {
		return account, fmt.Errorf("error fetching account: %v", string(resp.Body()))
	}

	account, err = model.UnmarshallToAccount(resp.Body())
	if err != nil {
		r.logger.Printf("error parsing get account response: %v", err)
	}

	return account, nil
}

// Create creates new account.
// If account data is not valid or account there is error while creating account it returns error.
func (r *AccountRestClient) Create(account model.Account) (model.Account, error) {
	accountJSON, err := model.MarshallToAccount(&account)
	if err != nil {
		r.logger.Printf("error marshalling account: %+v to JSON: %v", account, err)
		return account, err
	}

	resp, err := r.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(accountJSON).
		SetResult(&account).
		Post(AccountsEndpoint)
	if err != nil {
		r.logger.Printf("error fetching account: %v", err)
		return account, err
	}

	created := make(map[string]interface{})

	err = json.Unmarshal(resp.Body(), &created)
	if err != nil {
		r.logger.Printf("error parsing account creted response: %v", err)
		return account, err
	}

	account, err = model.UnmarshallToAccount(resp.Body())
	if err != nil {
		r.logger.Printf("error parsing create account response: %v", err)
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

	resp, err := r.client.R().
		SetPathParams(map[string]string{
			"account_id": accountID,
		}).
		SetQueryParam("version", strconv.Itoa(version)).
		Delete(AccountEndpoint)
	if err != nil {
		r.logger.Printf("error deleting account: %v", err)
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("error deleting account: %v", string(resp.Body()))
	}

	return nil
}

// List returns existing accounts.
// Page number and page size are optional values with deafault values of 0 and 100.
func (r *AccountRestClient) List(pageNumber, pageSize int) ([]model.Account, error) {
	var accounts []model.Account

	if pageSize == 0 {
		pageSize = 100
	}

	resp, err := r.client.R().
		SetQueryParams(map[string]string{
			"page[number]:": strconv.Itoa(pageNumber),
			"page[size]":    strconv.Itoa(pageSize),
		}).
		Get(AccountsEndpoint)
	if err != nil {
		r.logger.Printf("error lising account: %v", err)
		return accounts, err
	}

	if resp.IsError() {
		return accounts, fmt.Errorf("error lisitng account: %v", string(resp.Body()))
	}

	accounts, err = model.UnmarshallToAccounts(resp.Body())
	if err != nil {
		r.logger.Printf("error parsing get accounts response: %v", err)
	}

	return accounts, nil
}
