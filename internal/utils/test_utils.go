package utils

// AccountResponseJSON represents account api fetch response.
const AccountResponseJSON = `
{
	"data": {
		"type": "accounts",
		"id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		"organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		"version": 0,
		"attributes": {
			"country": "GB",
			"base_currency": "GBP",
			"account_number": "41426819",
			"bank_id": "400300",
			"bank_id_code": "GBDSC",
			"bic": "NWBKGB22",
			"iban": "GB11NWBK40030041426819",
			"title": "Ms",
			"first_name": "Samantha",
			"bank_account_name": "Samantha Holder",
			"alternative_bank_account_names": [
				"Sam Holder"
			],
			"account_classification": "Personal",
			"joint_account": false,
			"account_matching_opt_out": false,
			"secondary_identification": "A1B2C3D4",
			"private_identification": {
				"title": "Ms",
				"first_name": "Samantha",
				"last_name": "Holder",
				"birth_date": "2017-07-23",
				"birth_country": "GB",
				"document_number": "13YH458762",
				"address": "[10 Avenue des Champs]",
				"city": "London",
				"country": "GB"
			},
			"organisation_identification": {
				"name": "form3",
				"tax_residency": "GB",
				"registration_number": "123654",
				"representative": {
					"name": "Jeff Page",
					"birth_date": "1970-01-01",
					"residency": "GB"
				},
				"address": "[10 Avenue des Champs]",
				"city": "London",
				"country": "GB"
			}
		},
		"relationships": {
			"master_account": {
				"data": [{
					"type": "accounts",
					"id": "a52d13a4-f435-4c00-cfad-f5e7ac5972df"
				}]
			}
		}
	}
}`
