package model

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
)

type (
	// Account type represents palceholder type for account.
	Account struct {
		Data `json:"data"`
	}

	// Data type represents general type account data.
	Data struct {
		ID             string `json:"id" valid:"uuid"`
		OrganisationID string `json:"organisation_id" valid:"uuid"`
		Type           string `json:"type,omitempty" valid:"alpha"`
		Version        int    `json:"version,omitempty" valid:"alphanum"`
		Attributes     `json:"attributes"`
	}

	// Attributes represents specific account data.
	Attributes struct {
		AccountClassification       string   `json:"account_classification,omitempty" valid:"in(Personal|Business|)"`
		AccountMatchingOptOut       bool     `json:"account_matching_opt_out,omitempty"`
		AccountNumber               string   `json:"account_number,omitempty" valid:"alphanum, length(0|64)"`
		AlternativeBankAccountNames []string `json:"alternative_bank_account_names"`
		BankAccountName             string   `json:"bank_account_name,omitempty" valid:"length(1|140)"`
		BankID                      string   `json:"bank_id,omitempty" valid:"alphanum, length(0|16)"`
		BankIDCode                  string   `json:"bank_id_code,omitempty" valid:"alpha, length(0|16)"`
		BaseCurrency                string   `json:"base_currency,omitempty" valid:"alpha, length(3|3)"`
		Bic                         string   `json:"bic,omitempty" valid:"matches(^([A-Z]{6}[A-Z0-9]{2}|[A-Z]{6}[A-Z0-9]{5})$)"`
		Country                     string   `json:"country" valid:"matches(^[A-Z]{2}$)"`
		CustomerID                  string   `json:"customer_id,omitempty" valid:"alphanum, length(0|256)"`
		FirstName                   string   `json:"first_name,omitempty" valid:"length(1|40)"`
		Iban                        string   `json:"iban,omitempty"`
		JointAccount                bool     `json:"joint_account,omitempty"`
		OrganisationIdentification  `json:"organisation_identification,omitempty"`
	}

	// OrganisationIdentification ...
	OrganisationIdentification struct {
		Address string `json:"address,omitempty" valid:"length(1|140)"`
		City    string `json:"city,omitempty" valid:"length(1|35)"`
		Country string `json:"country,omitempty" valid:"matches(^[A-Z]{2}$)"`
	}
)

// UnmarshallToAccount parses JSON response to account type.
func UnmarshallToAccount(accountJSONResponse []byte) (Account, error) {
	var account Account
	err := json.Unmarshal(accountJSONResponse, &account)
	if err != nil {
		return account, err
	}

	govalidator.SetFieldsRequiredByDefault(false)
	if v, err := govalidator.ValidateStruct(&account); !v || err != nil {
		return account, err
	}

	return account, nil
}
