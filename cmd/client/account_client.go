package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"

	"github.com/realmallaury/resty-client/internal/account"
	"github.com/realmallaury/resty-client/internal/validation"
)

// AccountEndpoint is path resource for account
const AccountEndpoint = "/v1/organisation/accounts/{account_id}"

// AccountsEndpoint is path resource for accounts
const AccountsEndpoint = "/v1/organisation/accounts"

// AccountRestClient represent account rest client
type AccountRestClient struct {
	client *resty.Client
}

// New returns new AccountRestClient instance.
// Host url should resolve to valid url otherwise error is returned.
func New(hostURL string) (*AccountRestClient, error) {
	c := resty.New()

	if v, err := validation.ValidateURL(hostURL); !v || err != nil {
		return nil, errors.Wrap(err, "validate host URL")
	}

	c.SetHostURL(hostURL)

	rc := AccountRestClient{
		client: c,
	}

	return &rc, nil
}

// Fetch gets accound data for provided account id.
// If account id is not valid UUID or account is not found it returns error.
func (r *AccountRestClient) Fetch(accountID string) (account.Account, error) {
	var acc account.Account

	if v, err := validation.ValidateUUID(accountID); !v || err != nil {
		return acc, err
	}

	resp, err := r.client.R().
		SetPathParams(map[string]string{
			"account_id": accountID,
		}).
		Get(AccountEndpoint)
	if err != nil {
		return acc, errors.Wrapf(err, "error fetching account: %v", accountID)
	}

	if resp.IsError() {
		return acc, fmt.Errorf("error fetching account: %v", string(resp.Body()))
	}

	acc, err = account.UnmarshallToAccount(resp.Body())
	if err != nil {
		return acc, errors.Wrap(err, "error parsing account response")
	}

	return acc, nil
}

// Create creates new account.
// If account data is not valid or account there is error while creating account it returns error.
func (r *AccountRestClient) Create(acc account.Account) (account.Account, error) {
	accountJSON, err := account.MarshallToAccount(&acc)
	if err != nil {
		return acc, errors.Wrap(err, "error marshalling account to JSON")
	}

	resp, err := r.client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(accountJSON).
		SetResult(&acc).
		Post(AccountsEndpoint)
	if err != nil {
		return acc, errors.Wrap(err, "error creating account")
	}

	created := make(map[string]interface{})

	err = json.Unmarshal(resp.Body(), &created)
	if err != nil {
		return acc, errors.Wrap(err, "error unmarshalling account response")
	}

	acc, err = account.UnmarshallToAccount(resp.Body())
	if err != nil {
		return acc, errors.Wrap(err, "error parsing create account response")
	}

	if resp.IsError() {
		return acc, fmt.Errorf("error creating account: %v", string(resp.Body()))
	}

	return acc, nil
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
		return errors.Wrapf(err, "error deleting account: %v", accountID)
	}

	if resp.IsError() {
		return fmt.Errorf("error deleting account: %v", string(resp.Body()))
	}

	return nil
}

// List returns existing accounts.
// Page number and page size are optional values with deafault values of 0 and 100.
func (r *AccountRestClient) List(pageNumber, pageSize int) ([]account.Account, error) {
	var accs []account.Account

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
		return accs, errors.Wrapf(err, "error lising accounts")
	}

	if resp.IsError() {
		return accs, fmt.Errorf("error lisitng accounts: %v", string(resp.Body()))
	}

	accs, err = account.UnmarshallToAccounts(resp.Body())
	if err != nil {
		return accs, errors.Wrap(err, "error parsing accounts response")
	}

	return accs, nil
}
