package revoke

import "guarantees/com"

const (
	// Name of entity (for logs)
	ENTITY_NAME = "Revoke"

	// Xml tag name
	XML_TAG = "revokes"

	// Object type names (for storage)
	KEY = "REVOKE"
)

func (revoke Revoke) CreateValidation() bool {
	return true
}

func (revoke Revoke) ChangeValidation(newRevokeInterface interface{}) bool {
	newRevoke := newRevokeInterface.(*Revoke)
	valid := true

	valid = valid &&
		newRevoke.Id == revoke.Id &&
		newRevoke.RelationTxId == revoke.RelationTxId &&
		newRevoke.EntityId == revoke.EntityId &&
		newRevoke.EntityType == revoke.EntityType &&
		newRevoke.OrgSignature == revoke.OrgSignature &&
		newRevoke.CreateDate == revoke.CreateDate &&
		newRevoke.CreateUser == revoke.CreateUser &&
		newRevoke.EntitySigned == revoke.EntitySigned
	if valid == false {
		com.DebugLogMsg("Fields cann't be changed")
		return false
	}

	return valid
}

func (revoke Revoke) GetKeyObjectType() string {
	return KEY
}

func (revoke Revoke) GetIndexes() [][]string {
	return [][]string{
		{"EntityId"},
	}
}

func (revoke Revoke) GetEntityName() string {
	return ENTITY_NAME
}
func (revoke Revoke) GetTagName() string {
	return XML_TAG
}
func (revoke *Revoke) SetRelationTxId(relationTxId string) {
	revoke.RelationTxId = relationTxId
}
func (revoke *Revoke) GetRelationTxId() string {
	return revoke.RelationTxId
}
func (revoke *Revoke) SetId(id string) {
	revoke.Id = id
}
func (revoke Revoke) GetId() string {
	return revoke.Id
}
func (revoke *Revoke) SetKey(key string) {
	revoke.Key = key
}
func (revoke Revoke) GetKey() string {
	return revoke.Key
}
func (revoke *Revoke) SetMSPId(MSPId string) {
	revoke.MSPId = MSPId
}
func (revoke Revoke) GetMSPId() string {
	return revoke.MSPId
}
