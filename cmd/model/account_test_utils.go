package model

// GetMissingTestCreateAccount returns test account for creation with missing data.
func GetMissingTestCreateAccount() Account {
	return Account{
		Data: Data{
			ID:             "cd27e265-9605-4b4b-a0e5-3003ea9cc4dc",
			OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Type:           "accounts",
			Version:        0,
			Attributes: Attributes{
				Country: "GB",
				Bic:     "NWBKGB22",
				BankID:  "400300",
			},
		},
	}
}

// GetTestCreateAccount returns test account for creation.
func GetTestCreateAccount() Account {
	return Account{
		Data: Data{
			ID:             "cd27e265-9605-4b4b-a0e5-3003ea9cc4dc",
			OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Type:           "accounts",
			Version:        0,
			Attributes: Attributes{
				Country:                     "GB",
				BaseCurrency:                "GBP",
				AccountNumber:               "41426819",
				Bic:                         "NWBKGB22",
				BankIDCode:                  "GBDSC",
				BankID:                      "400300",
				AccountClassification:       "Personal",
				Iban:                        "GB11NWBK40030041426819",
				Title:                       "Ms",
				FirstName:                   "Samantha",
				BankAccountName:             "Samantha Holder",
				AlternativeBankAccountNames: []string{"Sam Holder"},
				JointAccount:                false,
				AccountMatchingOptOut:       false,
				SecondaryIdentification:     "A1B2C3D4",
				PrivateIdentification: PrivateIdentification{
					Title:          "Ms",
					FirstName:      "Samantha",
					LastName:       "Holder",
					BirthDate:      "2017-07-23",
					BirthCountry:   "GB",
					DocumentNumber: "13YH458762",
					Address:        "[10 Avenue des Champs]",
					City:           "London",
					Country:        "GB",
				},
			},
		},
	}
}

// AccountRequestJSON represents account api create resquest.
const AccountRequestJSON = `{
	"data": {
		"id": "cd27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		"type": "accounts",
		"version": 0,
		"attributes": {
			"account_classification": "Personal",
			"account_matching_opt_out": false,
			"account_number": "41426819",
			"alternative_bank_account_names": ["Sam Holder"],
			"bank_account_name": "Samantha Holder",
			"bank_id": "400300",
			"bank_id_code": "GBDSC",
			"base_currency": "GBP",
			"bic": "NWBKGB22",
			"country": "GB",
			"first_name": "Samantha",
			"iban": "GB11NWBK40030041426819",
			"joint_account": false,
			"organisation_identification": {
				"representative": {}
			},
			"private_identification": {
				"address": "[10 Avenue des Champs]",
				"city": "London",
				"birth_country": "GB",
				"birth_date": "2017-07-23",
				"country": "GB",
				"document_number": "13YH458762",
				"first_name": "Samantha",
				"last_name": "Holder",
				"title": "Ms"
			},
			"secondary_identification": "A1B2C3D4",
			"title": "Ms"
		},
		"relationships": {
			"master_account": {}
		}
	}
}`

// GetTestAccount returns test account model.
func GetTestAccount() Account {
	return Account{
		Data: Data{
			ID:             "cd27e265-9605-4b4b-a0e5-3003ea9cc4dc",
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
				SecondaryIdentification:     "A1B2C3D4",
				Title:                       "Ms",
			},
		},
	}
}

// AccountResponseJSON represents account api fetch response.
const AccountResponseJSON = `{
	"data": {
		"attributes": {
			"account_classification": "Personal",
			"account_matching_opt_out": false,
			"account_number": "41426819",
			"alternative_bank_account_names": [
				"Sam Holder"
			],
			"bank_account_name": "Samantha Holder",
			"bank_id": "400300",
			"bank_id_code": "GBDSC",
			"base_currency": "GBP",
			"bic": "NWBKGB22",
			"country": "GB",
			"first_name": "Samantha",
			"iban": "GB11NWBK40030041426819",
			"joint_account": false,
			"secondary_identification": "A1B2C3D4",
			"title": "Ms"
		},
		"created_on": "2020-01-31T12:41:40.509Z",
		"id": "cd27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		"modified_on": "2020-01-31T12:41:40.509Z",
		"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		"type": "accounts",
		"version": 0
	},
	"links": {
		"self": "/v1/organisation/accounts/cd27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	}
}`
