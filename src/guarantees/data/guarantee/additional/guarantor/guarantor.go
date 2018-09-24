package guarantor

const (
	// Name of entity (for logs)
	ENTITY_NAME = "Guarantor"

	// Xml tag name
	XML_TAG = "guarantors"

	// Object type names (for storage)
	KEY = "GUARANTOR"
)

func (guarantor Guarantor) CanBeChangedOn(newGuarantorInterface interface{}) bool {
	_ = newGuarantorInterface.(*Guarantor)
	return false
}

func (guarantor *Guarantor) GetKeyObjectType() string {
	return KEY
}

func (guarantor Guarantor) GetIndexes() [][]string {
	return [][]string{
		{"MSPId"},
	}
}

func (guarantor *Guarantor) GetEntityName() string {
	return ENTITY_NAME
}
func (guarantor *Guarantor) GetTagName() string {
	return XML_TAG
}
func (guarantor *Guarantor) SetRelationTxId(relationTxId string) {
	guarantor.RelationTxId = relationTxId
}
func (guarantor *Guarantor) GetRelationTxId() string {
	return guarantor.RelationTxId
}
func (guarantor *Guarantor) SetId(id string) {
	guarantor.Id = id
}
func (guarantor *Guarantor) GetId() string {
	return guarantor.Id
}
func (guarantor *Guarantor) SetKey(key string) {
	guarantor.Key = key
}
func (guarantor *Guarantor) GetKey() string {
	return guarantor.Key
}
func (guarantor *Guarantor) SetMSPId(MSPId string) {
	guarantor.MSPId = MSPId
}
func (guarantor *Guarantor) GetMSPId() string {
	return guarantor.MSPId
}
