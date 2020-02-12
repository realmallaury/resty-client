package account

import (
	"testing"

	"github.com/nsf/jsondiff"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshallToAccount(t *testing.T) {
	assert := assert.New(t)

	byteJSON := []byte(AccountResponseJSON)
	expectedAcc := GetTestAccount()
	acc, err := UnmarshallToAccount(byteJSON)

	assert.NoError(err, "Error should be nil")
	assert.EqualValues(expectedAcc, acc, "Account should be correctely unmarshalled")
}

func TestUnmarshallToAccounts(t *testing.T) {
	assert := assert.New(t)

	byteJSON := []byte(AccountsResponseJSON)
	accounts, err := UnmarshallToAccounts(byteJSON)

	assert.NoError(err, "Error should be nil")
	assert.Len(accounts, 10, "Accounts should contain 10 elements")
}

func TestMarshallToAccount(t *testing.T) {
	assert := assert.New(t)

	acc := GetMissingTestCreateAccount()
	_, err := MarshallToAccount(&acc)

	assert.Error(err, "MarshallToAccount(...) should return error")

	acc = GetTestCreateAccount()
	accountJSON, err := MarshallToAccount(&acc)

	assert.NoError(err, "Error should be nil")
	assert.True(compareJSON(accountJSON, []byte(AccountRequestJSON)), "Account should be correctely unmarshalled")
}

func compareJSON(result, expected []byte) bool {
	diff, _ := jsondiff.Compare(result, expected, &jsondiff.Options{})
	return diff == jsondiff.FullMatch
}
