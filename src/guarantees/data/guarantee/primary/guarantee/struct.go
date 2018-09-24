package guarantee

import (
	"encoding/xml"
	"guarantees/data/guarantee/additional/bankpars"
	"guarantees/data/guarantee/primary/statement"
)

type Guarantee struct {
	XMLName xml.Name `xml:"guarantee"`

	RelationTxId string `xml:"relation_tx_id"`
	Key          string `xml:"key"`
	Id           string `xml:"id"`

	MSPId string `xml:"msp_id"`

	Status             string `xml:"status"`
	PrincipalSignature string `xml:"principal_signature"`
	BankSignature      string `xml:"bank_signature"`

	GuaranteeSigned GuaranteeSigned `xml:"guarantee_signed"`
}

type GuaranteeSigned struct {
	XMLName xml.Name `xml:"guarantee_signed"`

	Number          string                    `xml:"number"`
	CreateDate      string                    `xml:"createDate"`
	IssueDate       string                    `xml:"issueDate"`
	StatementFields statement.StatementSigned `xml:"statement_fields"`
	BankPars        bankpars.BankPars         `xml:"bank_pars"`
}

type GuaranteeOut struct {
	XMLName xml.Name `xml:"guarantee"`

	Id string `xml:"id"`

	Status             string `xml:"status"`
	PrincipalSignature string `xml:"principal_signature"`
	BankSignature      string `xml:"bank_signature"`

	GuaranteeSigned interface{} `xml:"guarantee_signed"`
}

type GuaranteeSignedOut struct {
	XMLName xml.Name `xml:"guarantee_signed"`

	Number          string      `xml:"number"`
	CreateDate      string      `xml:"createDate"`
	IssueDate       string      `xml:"issueDate"`
	StatementFields interface{} `xml:"statement_fields"`
	BankPars        interface{} `xml:"bank_pars"`
}

func (guarantee Guarantee) ToOut() interface{} {
	result := GuaranteeOut{}

	result.Id = guarantee.Id
	result.Status = guarantee.Status
	result.PrincipalSignature = guarantee.PrincipalSignature
	result.BankSignature = guarantee.BankSignature
	result.GuaranteeSigned = guarantee.GuaranteeSigned.ToOut()

	return result
}

func (guaranteeSigned GuaranteeSigned) ToOut() interface{} {
	result := GuaranteeSignedOut{}

	result.Number = guaranteeSigned.Number
	result.CreateDate = guaranteeSigned.CreateDate
	result.IssueDate = guaranteeSigned.IssueDate
	result.StatementFields = guaranteeSigned.StatementFields.ToOut()
	result.BankPars = guaranteeSigned.BankPars.ToOut()

	return result
}
