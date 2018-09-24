package statement

import (
	"encoding/xml"
	"guarantees/data/guarantee/additional"
	"guarantees/data/guarantee/additional/gtype"
	"guarantees/data/guarantee/additional/guarantor"
)

type Statement struct {
	XMLName xml.Name `xml:"statement"`

	RelationTxId string `xml:"relation_tx_id"`
	Key          string `xml:"key"`
	Id           string `xml:"id"`

	MSPId string `xml:"msp_id"`

	Status             string `xml:"status"`
	PrincipalSignature string `xml:"principal_signature"`

	StatementSigned StatementSigned `xml:"statement_signed"`
}

type StatementSigned struct {
	Number      string                 `xml:"number"`
	CreateDate  string                 `xml:"createDate"`
	Principal   additional.Principal   `xml:"principal"`
	Beneficiary additional.Beneficiary `xml:"beneficiary"`
	Guarantor   guarantor.Guarantor    `xml:"guarantor"`
	GType       gtype.GType            `xml:"gtype"`
}

type StatementOut struct {
	XMLName xml.Name `xml:"statement"`

	Id string `xml:"id"`

	Status             string `xml:"status"`
	PrincipalSignature string `xml:"principal_signature"`

	StatementSigned interface{} `xml:"statement_signed"`
}

type StatementSignedOut struct {
	Number      string      `xml:"number"`
	CreateDate  string      `xml:"createDate"`
	Principal   interface{} `xml:"principal"`
	Beneficiary interface{} `xml:"beneficiary"`
	Guarantor   interface{} `xml:"guarantor"`
	GType       interface{} `xml:"gtype"`
}

func (statement Statement) ToOut() interface{} {
	result := StatementOut{}

	result.Id = statement.Id
	result.Status = statement.Status
	result.PrincipalSignature = statement.PrincipalSignature
	result.StatementSigned = statement.StatementSigned.ToOut()

	return result
}

func (statementSigned StatementSigned) ToOut() interface{} {
	result := StatementSignedOut{}

	result.Number = statementSigned.Number
	result.CreateDate = statementSigned.CreateDate
	result.Principal = statementSigned.Principal.ToOut()
	result.Beneficiary = statementSigned.Beneficiary.ToOut()
	result.Guarantor = statementSigned.Guarantor.ToOut()
	result.GType = statementSigned.GType.ToOut()

	return result
}
