package rejtype

const (
	// Name of entity (for logs)
	ENTITY_NAME = "RejType"

	// Xml tag name
	XML_TAG = "rejtypes"

	// Object type names (for storage)
	KEY = "REJTYPE" // default
)

func (rejtype RejType) CanCreate() bool {
	return true
}

func (rejtype RejType) CanBeChangedOn(newRejTypeInterface interface{}) bool {
	_ = newRejTypeInterface.(*RejType)
	return false
}

func (rejtype RejType) GetKeyObjectType() string {
	return KEY
}

func (rejtype RejType) GetIndexes() [][]string {
	return [][]string{
		{"GuarantorId"},
	}
}

func (rejtype RejType) GetEntityName() string {
	return ENTITY_NAME
}
func (rejtype RejType) GetTagName() string {
	return XML_TAG
}
func (rejtype *RejType) SetRelationTxId(relationTxId string) {
	rejtype.RelationTxId = relationTxId
}
func (rejtype *RejType) GetRelationTxId() string {
	return rejtype.RelationTxId
}
func (rejtype *RejType) SetId(id string) {
	rejtype.Id = id
}
func (rejtype RejType) GetId() string {
	return rejtype.Id
}
func (rejtype *RejType) SetKey(key string) {
	rejtype.Key = key
}
func (rejtype RejType) GetKey() string {
	return rejtype.Key
}
func (rejtype *RejType) SetMSPId(MSPId string) {
	rejtype.MSPId = MSPId
}
func (rejtype RejType) GetMSPId() string {
	return rejtype.MSPId
}
