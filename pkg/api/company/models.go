package company

import (
	common2 "github.com/erply/api-go-wrapper/pkg/api/common"
)

type (
	//CompanyInfos ..
	Infos []Info
	//CompanyInfo ..
	Info struct {
		ID                 string `json:"id"`
		Name               string `json:"name"`
		Code               string `json:"code"`
		VAT                string `json:"VAT"`
		Phone              string `json:"phone"`
		Mobile             string `json:"mobile"`
		Fax                string `json:"fax"`
		Email              string `json:"email"`
		Web                string `json:"web"`
		BankAccountNumber  string `json:"bankAccountNumber"`
		BankName           string `json:"bankName"`
		BankSWIFT          string `json:"bankSWIFT"`
		BankIBAN           string `json:"bankIBAN"`
		BankAccountNumber2 string `json:"bankAccountNumber2"`
		BankName2          string `json:"bankName2"`
		BankSWIFT2         string `json:"bankSWIFT2"`
		BankIBAN2          string `json:"bankIBAN2"`
		Address            string `json:"address"`
		Country            string `json:"country"`

		//field for ConfParameters
		ConfParameters ConfParameter
	} //GetCompanyInfoResponse ...
	GetCompanyInfoResponse struct {
		Status       common2.Status `json:"status"`
		CompanyInfos Infos          `json:"records"`
	}
)

type (
	ConfParameter struct {
		Announcement         string `json:"invoice_announcement_eng"`
		InvoiceClientIsPayer string `json:"invoice_client_is_payer"`
		DefaultLanguage      string `json:"default_language"`
		Country              string `json:"country"`
		CSVSeparator         string `json:"csv_field_separator"`
	}
	//GetConfParametersResponse ...
	GetConfParametersResponse struct {
		Status         common2.Status  `json:"status"`
		ConfParameters []ConfParameter `json:"records"`
	}
)
