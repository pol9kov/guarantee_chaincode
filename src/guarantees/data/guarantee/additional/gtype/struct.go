package gtype

import (
	"encoding/xml"
	"guarantees/data/guarantee/additional"
)

type GType struct {
	XMLName xml.Name `xml:"gtype"`

	RelationTxId string `xml:"relation_tx_id"`
	Key          string `xml:"key"`
	Id           string `xml:"id"`

	MSPId string `xml:"msp_id"`

	GuarantorId string           `xml:"guarantor_id"`
	Name        string           `xml:"name"`
	Description string           `xml:"description"`
	Pars        []additional.Par `xml:"pars>par"`
}

type GTypeOut struct {
	XMLName xml.Name `xml:"gtype"`

	Id string `xml:"id"`

	GuarantorId string           `xml:"guarantor_id"`
	Name        string           `xml:"name"`
	Description string           `xml:"description"`
	Pars        []additional.Par `xml:"pars>par"`
}

func (gtype GType) ToOut() interface{} {
	result := GTypeOut{}

	result.Id = gtype.Id
	result.GuarantorId = gtype.GuarantorId
	result.Name = gtype.Name
	result.Description = gtype.Description
	result.Pars = gtype.Pars

	return result
}
