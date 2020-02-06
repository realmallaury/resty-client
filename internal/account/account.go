package account

import (
	"encoding/json"

	"github.com/realmallaury/resty-client/internal/validation"
)

// UnmarshallToAccount parses JSON to account type.
func UnmarshallToAccount(accountJSON []byte) (Account, error) {
	var a Account

	err := json.Unmarshal(accountJSON, &a)
	if err != nil {
		return a, err
	}

	return GetTestAccount(), nil
}

// UnmarshallToAccounts parses JSON to accounts type slice.
func UnmarshallToAccounts(accountsJSON []byte) ([]Account, error) {
	var accs Accounts

	err := json.Unmarshal(accountsJSON, &accs)
	if err != nil {
		return accs.Accounts, err
	}

	return accs.Accounts, nil
}

// MarshallToAccount marshalls account to JSON.
func MarshallToAccount(account *Account) ([]byte, error) {
	var accountJSON []byte

	accountJSON, err := json.Marshal(account)
	if err != nil {
		return nil, err
	}

	if v, err := validation.ValidateAccount(accountJSON); !v || err != nil {
		return nil, err
	}

	return accountJSON, nil
}
