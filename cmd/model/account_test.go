package model

import (
	"testing"

	"github.com/push-er/resty-client/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshallToAccount(t *testing.T) {
	assert := assert.New(t)

	byteJSON := []byte(utils.AccountResponseJSON)
	expectedAcc := getTestAccount()
	acc, err := UnmarshallToAccount(byteJSON)

	assert.Nil(err, "Error should be nil")
	assert.EqualValues(expectedAcc, acc, "Account should be correctely unmarshalled")
}

// GetTestAccount returns test account model.
func getTestAccount() Account {
	return Account{
		Data: Data{
			ID:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
			OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Type:           "accounts",
			Version:        0,
			Attributes: Attributes{
				AccountClassification:       "Personal",
				AccountMatchingOptOut:       false,
				AccountNumber:               "41426819",
				AlternativeBankAccountNames: []string{"Sam Holder"},
				BankAccountName:             "Samantha Holder",
				BankID:                      "400300",
				BankIDCode:                  "GBDSC",
				BaseCurrency:                "GBP",
				Bic:                         "NWBKGB22",
				Country:                     "GB",
				FirstName:                   "Samantha",
				Iban:                        "GB11NWBK40030041426819",
				JointAccount:                false,
				OrganisationIdentification: OrganisationIdentification{
					Address: "[10 Avenue des Champs]",
					City:    "London",
					Country: "GB",
				},
			},
		},
	}
}
