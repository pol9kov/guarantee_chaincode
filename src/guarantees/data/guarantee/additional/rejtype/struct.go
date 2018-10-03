package rejtype

import (
	"encoding/xml"
	"guarantees/data/guarantee/additional"
)

type RejType struct {
	XMLName xml.Name `xml:"requirement_template"`

	RelationTxId string `xml:"relation_tx_id"`
	Key          string `xml:"key"`
	Id           string `xml:"id"`

	MSPId string `xml:"msp_id"`

	GuarantorId string           `xml:"guarantor_id"`
	Name        string           `xml:"name"`
	Description string           `xml:"description"`
	Pars        []additional.Par `xml:"pars>par"`
}

type RejTypeOut struct {
	XMLName xml.Name `xml:"requirement_template"`

	Name        string           `xml:"name"`
	Description string           `xml:"description"`
	Pars        []additional.Par `xml:"pars>par"`
}

func (rejtype RejType) ToOut() interface{} {
	result := RejTypeOut{}

	result.Name = rejtype.Name
	result.Description = rejtype.Description
	result.Pars = rejtype.Pars

	return result
}
