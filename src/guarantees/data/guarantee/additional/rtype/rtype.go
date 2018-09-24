package rtype

const (
	// Name of entity (for logs)
	ENTITY_NAME = "RType"

	// Xml tag name
	XML_TAG = "rtypes"

	// Object type names (for storage)
	KEY = "RTYPE" // default
)

func (rtype RType) CanBeChangedOn(newRTypeInterface interface{}) bool {
	_ = newRTypeInterface.(*RType)
	return false
}

func (rtype RType) GetKeyObjectType() string {
	return KEY
}

func (rtype RType) GetIndexes() [][]string {
	return [][]string{
		{"GuarantorId"},
	}
}

func (rtype RType) GetEntityName() string {
	return ENTITY_NAME
}
func (rtype RType) GetTagName() string {
	return XML_TAG
}
func (rtype *RType) SetRelationTxId(relationTxId string) {
	rtype.RelationTxId = relationTxId
}
func (rtype *RType) GetRelationTxId() string {
	return rtype.RelationTxId
}
func (rtype *RType) SetId(id string) {
	rtype.Id = id
}
func (rtype RType) GetId() string {
	return rtype.Id
}
func (rtype *RType) SetKey(key string) {
	rtype.Key = key
}
func (rtype RType) GetKey() string {
	return rtype.Key
}
func (rtype *RType) SetMSPId(MSPId string) {
	rtype.MSPId = MSPId
}
func (rtype RType) GetMSPId() string {
	return rtype.MSPId
}
