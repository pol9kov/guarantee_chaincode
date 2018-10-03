package part

import (
	"encoding/xml"
)

type Part struct {
	XMLName xml.Name `xml:"part"`

	RelationTxId string `xml:"relation_tx_id"`
	Key          string `xml:"key"`
	Id           string `xml:"id"`

	MSPId string `xml:"msp_id"`

	DocumentId string `xml:"document_id"`
	Index      string `xml:"index"`
	Data       string `xml:"data"`
}

type PartOut struct {
	XMLName xml.Name `xml:"part"`

	Id string `xml:"id"`

	DocumentId string `xml:"document_id"`
	Index      string `xml:"index"`
	Data       string `xml:"data"`
}

func (part Part) ToOut() interface{} {
	result := PartOut{}

	result.Id = part.Id
	result.DocumentId = part.DocumentId
	result.Index = part.Index
	result.Data = part.Data

	return result
}
