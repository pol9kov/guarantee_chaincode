package revoke

import (
	"encoding/xml"
)

type Revoke struct {
	XMLName xml.Name `xml:"revoke"`

	RelationTxId string `xml:"relation_tx_id"`
	Key          string `xml:"key"`
	Id           string `xml:"id"`

	MSPId string `xml:"msp_id"`

	EntityId     string `xml:"entity_id"`
	EntityType   string `xml:"entity_type"`
	OrgSignature string `xml:"org_signature"`
	CreateDate   string `xml:"create_date"`
	CreateUser   string `xml:"create_user"`
	EntitySigned string `xml:"entity_signed"`
}

type RevokeOut struct {
	XMLName xml.Name `xml:"revoke"`

	Id string `xml:"id"`

	EntityId     string `xml:"entity_id"`
	EntityType   string `xml:"entity_type"`
	OrgSignature string `xml:"org_signature"`
	CreateDate   string `xml:"create_date"`
	CreateUser   string `xml:"create_user"`
	EntitySigned string `xml:"entity_signed"`
}

func (revoke Revoke) ToOut() interface{} {
	result := RevokeOut{}

	result.Id = revoke.Id
	result.EntityId = revoke.EntityId
	result.EntityType = revoke.EntityType
	result.OrgSignature = revoke.OrgSignature
	result.CreateDate = revoke.CreateDate
	result.CreateUser = revoke.CreateUser
	result.EntitySigned = revoke.EntitySigned

	return result
}
