package model

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

	assert.Nil(err, "Error should be nil")
	assert.EqualValues(expectedAcc, acc, "Account should be correctely unmarshalled")
}

func TestMarshallToAccount(t *testing.T) {
	assert := assert.New(t)

	acc := GetMissingTestCreateAccount()
	accountJSON, err := MarshallToAccount(&acc)

	assert.Error(err, "MarshallToAccount(...) should return error")

	acc = GetTestCreateAccount()
	accountJSON, err = MarshallToAccount(&acc)

	assert.Nil(err, "Error should be nil")
	assert.True(compareJSON(accountJSON, []byte(AccountRequestJSON)), "Account should be correctely unmarshalled")
}

func compareJSON(result, expected []byte) bool {
	diff, _ := jsondiff.Compare([]byte(result), []byte(expected), &jsondiff.Options{})
	return diff == jsondiff.FullMatch
}
