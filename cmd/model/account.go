package model

import (
	"encoding/json"

	"github.com/push-er/resty-client/internal/validation"
)

type (
	// Account type represents placeholder type for account.
	Account struct {
		Data `json:"data"`
	}

	// Data type represents general type account data.
	Data struct {
		ID             string `json:"id"`
		OrganisationID string `json:"organisation_id"`
		Type           string `json:"type,omitempty"`
		Version        int    `json:"version"`
		Attributes     `json:"attributes"`
		Relationships  `json:"relationships"`
	}

	// Attributes represents specific account data.
	Attributes struct {
		AccountClassification       string   `json:"account_classification,omitempty"`
		AccountMatchingOptOut       bool     `json:"account_matching_opt_out"`
		AccountNumber               string   `json:"account_number,omitempty"`
		AlternativeBankAccountNames []string `json:"alternative_bank_account_names"`
		BankAccountName             string   `json:"bank_account_name,omitempty"`
		BankID                      string   `json:"bank_id,omitempty"`
		BankIDCode                  string   `json:"bank_id_code,omitempty"`
		BaseCurrency                string   `json:"base_currency,omitempty"`
		Bic                         string   `json:"bic,omitempty"`
		Country                     string   `json:"country"`
		CustomerID                  string   `json:"customer_id,omitempty"`
		FirstName                   string   `json:"first_name,omitempty"`
		Iban                        string   `json:"iban,omitempty"`
		JointAccount                bool     `json:"joint_account"`
		OrganisationIdentification  `json:"organisation_identification,omitempty"`
		PrivateIdentification       `json:"private_identification,omitempty"`
		SecondaryIdentification     string `json:"secondary_identification,omitempty"`
		Switched                    bool   `json:"switched,omitempty"`
		Title                       string `json:"title,omitempty"`
	}

	// OrganisationIdentification represents organisation data.
	OrganisationIdentification struct {
		Address            string `json:"address,omitempty"`
		City               string `json:"city,omitempty"`
		Country            string `json:"country,omitempty"`
		Name               string `json:"name,omitempty"`
		RegistrationNumber string `json:"registration_number,omitempty"`
		Representative     `json:"representative,omitempty"`
		TaxResidency       string `json:"tax_residency,omitempty"`
	}

	// Representative represents organisation representative data.
	Representative struct {
		BirthDate string `json:"birth_date,omitempty"`
		Name      string `json:"name,omitempty"`
		Residency string `json:"residency,omitempty"`
	}

	// PrivateIdentification represents private person data.
	PrivateIdentification struct {
		Address        string `json:"address,omitempty"`
		City           string `json:"city,omitempty"`
		BirthCountry   string `json:"birth_country,omitempty"`
		BirthDate      string `json:"birth_date,omitempty"`
		Country        string `json:"country,omitempty"`
		DocumentNumber string `json:"document_number,omitempty"`
		FirstName      string `json:"first_name,omitempty"`
		LastName       string `json:"last_name,omitempty"`
		Title          string `json:"title,omitempty"`
	}

	// Relationships placeholder type for master account.
	Relationships struct {
		MasterAccount `json:"master_account,omitempty"`
	}

	// MasterAccount respresents master accounts relationship data.
	MasterAccount struct {
		Data []RelationshipData `json:"data,omitempty"`
	}

	// RelationshipData respresents master account data.
	RelationshipData struct {
		ID   string `json:"id"`
		Type string `json:"type,omitempty"`
	}
)

// UnmarshallToAccount parses JSON to account type.
func UnmarshallToAccount(accountJSON []byte) (Account, error) {
	var account Account

	if v, err := validation.ValidateAccount(accountJSON); !v || err != nil {
		return account, err
	}

	err := json.Unmarshal(accountJSON, &account)
	if err != nil {
		return account, err
	}

	return account, nil
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
