package gtype

const (
	// Name of entity (for logs)
	ENTITY_NAME = "GType"

	// Xml tag name
	XML_TAG = "gtypes"

	// Object type names (for storage)
	KEY = "GTYPE" // default
)

func (gtype GType) CreateValidation() bool {
	return true
}

func (gtype GType) ChangeValidation(newGTypeInterface interface{}) bool {
	_ = newGTypeInterface.(*GType)
	return false
}

func (gtype GType) GetKeyObjectType() string {
	return KEY
}

func (gtype GType) GetIndexes() [][]string {
	return [][]string{
		{"GuarantorId"},
	}
}

func (gtype GType) GetEntityName() string {
	return ENTITY_NAME
}
func (gtype GType) GetTagName() string {
	return XML_TAG
}
func (gtype *GType) SetRelationTxId(relationTxId string) {
	gtype.RelationTxId = relationTxId
}
func (gtype *GType) GetRelationTxId() string {
	return gtype.RelationTxId
}
func (gtype *GType) SetId(id string) {
	gtype.Id = id
}
func (gtype GType) GetId() string {
	return gtype.Id
}
func (gtype *GType) SetKey(key string) {
	gtype.Key = key
}
func (gtype GType) GetKey() string {
	return gtype.Key
}
func (gtype *GType) SetMSPId(MSPId string) {
	gtype.MSPId = MSPId
}
func (gtype GType) GetMSPId() string {
	return gtype.MSPId
}
