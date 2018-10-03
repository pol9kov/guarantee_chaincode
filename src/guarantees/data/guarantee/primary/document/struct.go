package document

import (
	"encoding/xml"
)

type Document struct {
	XMLName xml.Name `xml:"document"`

	RelationTxId string `xml:"relation_tx_id"`
	Key          string `xml:"key"`
	Id           string `xml:"id"`

	MSPId string `xml:"msp_id"`

	EntityId      string `xml:"entity_id"`
	DocName       string `xml:"doc_name"`
	FieldName     string `xml:"field_name"`
	FieldDocIndex string `xml:"field_doc_index"`
	CreationDate  string `xml:"creation_date"`
	CreationUser  string `xml:"creation_uate"`
	MIMEType      string `xml:"mime_type"`
	Hash          string `xml:"hash"`
	Size          string `xml:"size"`
	CountParts    string `xml:"count_parts"`
	PartSize      string `xml:"part_size"`
}

type DocumentOut struct {
	XMLName xml.Name `xml:"document"`

	Id string `xml:"id"`

	EntityId      string `xml:"entity_id"`
	DocName       string `xml:"doc_name"`
	FieldName     string `xml:"field_name"`
	FieldDocIndex string `xml:"field_doc_index"`
	CreationDate  string `xml:"creation_date"`
	CreationUser  string `xml:"creation_uate"`
	MIMEType      string `xml:"mime_type"`
	Hash          string `xml:"hash"`
	Size          string `xml:"size"`
	CountParts    string `xml:"count_parts"`
	PartSize      string `xml:"part_size"`
}

func (document Document) ToOut() interface{} {
	result := DocumentOut{}

	result.Id = document.Id
	result.EntityId = document.EntityId
	result.DocName = document.DocName
	result.FieldName = document.FieldName
	result.FieldDocIndex = document.FieldDocIndex
	result.CreationDate = document.CreationDate
	result.CreationUser = document.CreationUser
	result.MIMEType = document.MIMEType
	result.Hash = document.Hash
	result.Size = document.Size
	result.CountParts = document.CountParts
	result.PartSize = document.PartSize

	return result
}
