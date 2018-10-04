package rtype

import (
	"encoding/xml"
	"guarantees/data/guarantee/additional"
)

type RType struct {
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

type RTypeOut struct {
	XMLName xml.Name `xml:"requirement_template"`

	Id string `xml:"id"`

	GuarantorId string           `xml:"guarantor_id"`
	Name        string           `xml:"name"`
	Description string           `xml:"description"`
	Pars        []additional.Par `xml:"pars>par"`
}

func (rtype RType) ToOut() interface{} {
	result := RTypeOut{}

	result.Id = rtype.Id
	result.GuarantorId = rtype.GuarantorId
	result.Name = rtype.Name
	result.Description = rtype.Description
	result.Pars = rtype.Pars

	return result
}
