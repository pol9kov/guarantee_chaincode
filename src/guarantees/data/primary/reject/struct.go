package reject

import (
	"encoding/xml"
	"guarantees/data/additional/rejtype"
)

type Reject struct {
	XMLName xml.Name `xml:"reject"`

	RelationTxId string `xml:"relation_tx_id"`
	Key          string `xml:"key"`
	Id           string `xml:"id"`

	MSPId string `xml:"msp_id"`

	Status             string          `xml:"status"`
	EntityId           string          `xml:"entity_id"`
	EntityType         string          `xml:"entity_type"`
	GuarantorSignature string          `xml:"guarantor_signature"`
	CreateDate         string          `xml:"create_date"`
	CreateUser         string          `xml:"create_user"`
	Number             string          `xml:"number"`
	Rejtype            rejtype.RejType `xml:"rejtype"`
}

type RejectOut struct {
	XMLName xml.Name `xml:"reject"`

	Id string `xml:"id"`

	Status             string      `xml:"status"`
	EntityId           string      `xml:"entity_id"`
	EntityType         string      `xml:"entity_type"`
	GuarantorSignature string      `xml:"guarantor_signature"`
	CreateDate         string      `xml:"create_date"`
	CreateUser         string      `xml:"create_user"`
	Number             string      `xml:"number"`
	Rejtype            interface{} `xml:"rejtype"`
}

func (reject Reject) ToOut() interface{} {
	result := RejectOut{}

	result.Id = reject.Id
	result.Status = reject.Status
	result.EntityId = reject.EntityId
	result.EntityType = reject.EntityType
	result.GuarantorSignature = reject.GuarantorSignature
	result.CreateDate = reject.CreateDate
	result.CreateUser = reject.CreateUser
	result.Number = reject.Number
	result.Rejtype = reject.Rejtype.ToOut()

	return result
}
