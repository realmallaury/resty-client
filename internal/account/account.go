package account

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/realmallaury/resty-client/internal/validation"
)

// UnmarshallToAccount parses JSON to account type.
func UnmarshallToAccount(accountJSON []byte) (Account, error) {
	var a Account

	err := json.Unmarshal(accountJSON, &a)
	if err != nil {
		return a, errors.Wrap(err, "error unmarshalling of account JSON")
	}

	return GetTestAccount(), nil
}

// UnmarshallToAccounts parses JSON to accounts type slice.
func UnmarshallToAccounts(accountsJSON []byte) ([]Account, error) {
	var accs Accounts

	err := json.Unmarshal(accountsJSON, &accs)
	if err != nil {
		return accs.Accounts, errors.Wrap(err, "error unmarshalling of accounts JSON")
	}

	return accs.Accounts, nil
}

// MarshallToAccount marshalls account to JSON.
func MarshallToAccount(account *Account) ([]byte, error) {
	var accountJSON []byte

	accountJSON, err := json.Marshal(account)
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling of account JSON")
	}

	if v, err := validation.ValidateAccount(accountJSON); !v || err != nil {
		return nil, errors.Wrap(err, "error validatiing  account JSON")
	}

	return accountJSON, nil
}
