package validation

// AddAccountSchema ...
var AddAccountSchema = string(`
{
	"$schema": "http://json-schema.org/draft-07/schema#",
	"type": "object",
	"properties": {
		"data": {
			"type": "object",
			"properties": {
				"attributes": {
					"type": "object",
					"properties": {
						"bank_id": {
							"type": "string",
							"pattern": "^[A-Z0-9]{0,16}$"
						},
						"bank_id_code": {
							"type": "string",
							"pattern": "^[A-Z]{0,16}$"
						},
						"base_currency": {
							"type": "string",
							"pattern": "^[A-Z]{3}$"
						},
						"bic": {
							"type": "string",
							"pattern": "^([A-Z]{6}[A-Z0-9]{2}|[A-Z]{6}[A-Z0-9]{5})$"
						},
						"country": {
							"type": "string",
							"pattern": "^[A-Z]{2}$"
						}
					},
					"anyOf": [{
							"properties": {
								"country": {
									"const": "GB"
								},
								"bank_id_code": {
									"const": "GBDSC"
								}
							},
							"required": ["bic", "bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "AU"
								},
								"bank_id_code": {
									"const": "AUBSB"
								}
							},
							"required": ["bic","bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "BE"
								},
								"bank_id_code": {
									"const": "BE"
								}
							},
							"required": ["bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "CA"
								},
								"bank_id_code": {
									"const": "CACPA"
								}
							},
							"required": ["bic"]
						},
						{
							"properties": {
								"country": {
									"const": "FR"
								},
								"bank_id_code": {
									"const": "FR"
								}
							},
							"required": ["bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "DE"
								},
								"bank_id_code": {
									"const": "DEBLZ"
								}
							},
							"required": ["bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "GR"
								},
								"bank_id_code": {
									"const": "GRBIC"
								}
							},
							"required": ["bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "HK"
								},
								"bank_id_code": {
									"const": "HKNCC"
								}
							},
							"required": ["bic"]
						},
						{
							"properties": {
								"country": {
									"const": "IT"
								},
								"bank_id_code": {
									"const": "ITNCC"
								}
							},
							"required": ["bank_id"]
						},
						{
							"properties": {
								"country": {
									"const": "LU"
								},
								"bank_id_code": {
									"const": "LULUX"
								}
							},
							"required": ["bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "NL"
								}
							},
							"required": ["bic"]
						},
						{
							"properties": {
								"country": {
									"const": "PL"
								},
								"bank_id_code": {
									"const": "PLKNR"
								}
							},
							"required": ["bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "PT"
								},
								"bank_id_code": {
									"const": "PTNCC"
								}
							},
							"required": ["bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "ES"
								},
								"bank_id_code": {
									"const": "ESNCC"
								}
							},
							"required": ["bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "CH"
								},
								"bank_id_code": {
									"const": "CHBCC"
								}
							},
							"required": ["bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "US"
								},
								"bank_id_code": {
									"const": "USABA"
								}
							},
							"required": ["bic", "bank_id", "bank_id_code"]
						}
					]
				}
			}
		}
	},
	"required": ["data"]
}
`)

// AccountSchema is a schema that represents api account response.
var AccountSchema = string(`
{
	"$schema": "http://json-schema.org/draft-07/schema#",
	"type": "object",
	"properties": {
		"data": {
			"type": "object",
			"properties": {
				"id": {
					"type": "string",
					"format": "uuid"
				},
				"organisation_id": {
					"type": "string",
					"format": "uuid"
				},
				"type": {
					"type": "string",
					"pattern": "^[A-Za-z]*$"
				},
				"version": {
					"type": "integer",
					"minimum": 0
				},
				"attributes": {
					"type": "object",
					"properties": {
						"account_classification": {
							"type": "string",
							"enum": ["Personal", "Business"]
						},
						"account_matching_opt_out": {
							"type": "boolean"
						},
						"account_number": {
							"type": "string",
							"pattern": "^[A-Z0-9]{0,64}$"
						},
						"alternative_bank_account_names": {
							"type": [ "array", "null" ],
							"contains": {
								"type": "string",
								"minLength": 1,
								"maxLength": 140
							},
							"maxItems": 3
						},
						"bank_account_name": {
							"type": "string",
							"minLength": 1,
							"maxLength": 140
						},
						"bank_id": {
							"type": "string",
							"pattern": "^[A-Z0-9]{0,16}$"
						},
						"bank_id_code": {
							"type": "string",
							"pattern": "^[A-Z]{0,16}$"
						},
						"base_currency": {
							"type": "string",
							"pattern": "^[A-Z]{3}$"
						},
						"bic": {
							"type": "string",
							"pattern": "^([A-Z]{6}[A-Z0-9]{2}|[A-Z]{6}[A-Z0-9]{5})$"
						},
						"country": {
							"type": "string",
							"pattern": "^[A-Z]{2}$"
						},
						"customer_id": {
							"type": "string",
							"pattern": "^[a-zA-Z0-9-$@.,]{0,256}$"
						},
						"first_name": {
							"type": "string",
							"minLength": 1,
							"maxLength": 40
						},
						"iban": {
							"type": "string",
							"pattern": "^[A-Z]{2}[0-9]{2}[A-Z0-9]{0,64}$"
						},
						"joint_account": {
							"type": "boolean"
						},
						"organisation_identification": {
							"type": "object",
							"properties": {
								"address": {
									"type": "string",
									"minLength": 1,
									"maxLength": 140
								},
								"city": {
									"type": "string",
									"minLength": 1,
									"maxLength": 35
								},
								"country": {
									"type": "string",
									"pattern": "^[A-Z]{2}$"
								},
								"name": {
									"type": "string",
									"minLength": 1,
									"maxLength": 40
								},
								"registration_number": {
									"type": "string"
								},
								"representative": {
									"type": "object",
									"properties": {
										"birth_date": {
											"type": "string",
											"format": "date"
										},
										"name": {
											"type": "string",
											"minLength": 1,
											"maxLength": 40
										},
										"residency": {
											"type": "string"
										}
									}
								},
								"tax_residency": {
									"type": "string",
									"pattern": "^[A-Z]{2}$"
								}
							}
						},
						"private_identification": {
							"type": "object",
							"properties": {
								"address": {
									"type": "string",
									"minLength": 1,
									"maxLength": 140
								},
								"birth_country": {
									"type": "string",
									"pattern": "^[A-Z]{2}$"
								},
								"birth_date": {
									"type": "string",
									"format": "date"
								},
								"city": {
									"type": "string",
									"minLength": 1,
									"maxLength": 35
								},
								"country": {
									"type": "string",
									"pattern": "^[A-Z]{2}$"
								},
								"document_number": {
									"type": "string"
								},
								"first_name": {
									"type": "string",
									"minLength": 1,
									"maxLength": 40
								},
								"last_name": {
									"type": "string",
									"minLength": 1,
									"maxLength": 40
								},
								"title": {
									"type": "string",
									"minLength": 1,
									"maxLength": 40
								}
							}
						},
						"secondary_identification": {
							"type": "string",
							"minLength": 1,
							"maxLength": 140
						},
						"switched": {
							"type": "boolean"
						},
						"title": {
							"type": "string",
							"minLength": 1,
							"maxLength": 40
						}
					},
					"anyOf": [{
							"properties": {
								"country": {
									"const": "GB"
								},
								"bank_id_code": {
									"const": "GBDSC"
								}
							},
							"required": ["bic", "bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "AU"
								},
								"bank_id_code": {
									"const": "AUBSB"
								}
							},
							"required": ["bic", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "BE"
								},
								"bank_id_code": {
									"const": "BE"
								}
							},
							"required": ["bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "CA"
								},
								"bank_id_code": {
									"const": "CACPA"
								}
							},
							"required": ["bic"]
						},
						{
							"properties": {
								"country": {
									"const": "FR"
								},
								"bank_id_code": {
									"const": "FR"
								}
							},
							"required": ["bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "DE"
								},
								"bank_id_code": {
									"const": "DEBLZ"
								}
							},
							"required": ["bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "GR"
								},
								"bank_id_code": {
									"const": "GRBIC"
								}
							},
							"required": ["bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "HK"
								},
								"bank_id_code": {
									"const": "HKNCC"
								}
							},
							"required": ["bic"]
						},
						{
							"properties": {
								"country": {
									"const": "IT"
								},
								"bank_id_code": {
									"const": "ITNCC"
								}
							},
							"required": ["bank_id"]
						},
						{
							"properties": {
								"country": {
									"const": "LU"
								},
								"bank_id_code": {
									"const": "LULUX"
								}
							},
							"required": ["bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "NL"
								}
							},
							"required": ["bic"]
						},
						{
							"properties": {
								"country": {
									"const": "PL"
								},
								"bank_id_code": {
									"const": "PLKNR"
								}
							},
							"required": ["bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "PT"
								},
								"bank_id_code": {
									"const": "PTNCC"
								}
							},
							"required": ["bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "ES"
								},
								"bank_id_code": {
									"const": "ESNCC"
								}
							},
							"required": ["bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "CH"
								},
								"bank_id_code": {
									"const": "CHBCC"
								}
							},
							"required": ["bank_id", "bank_id_code"]
						},
						{
							"properties": {
								"country": {
									"const": "US"
								},
								"bank_id_code": {
									"const": "USABA"
								}
							},
							"required": ["bic", "bank_id", "bank_id_code"]
						}
					]
				},
				"relationships": {
					"type": "object",
					"properties": {
						"master_account": {
							"type": "object",
							"properties": {
								"data": {
									"type": "array",
									"contains": {
										"type": "object",
										"properties": {
											"id": {
												"type": "string",
												"format": "uuid"
											},
											"type": {
												"type": "string"
											}
										}
									}
								}
							}
						}
					}
				}
			},
			"required": ["id", "organisation_id", "attributes"]
		}
	},
	"required": ["data"]
}
`)
