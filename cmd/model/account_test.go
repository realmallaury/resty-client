package model

import (
	"testing"

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

	acc := GetTestCreateAccount()
	_, err := MarshallToAccount(&acc)

	assert.Nil(err, "Error should be nil")
}
