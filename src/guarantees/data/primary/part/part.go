package part

const (
	// Name of entity (for logs)
	ENTITY_NAME = "Part"

	// Xml tag name
	XML_TAG = "parts"

	// Object type names (for storage)
	KEY = "PART"
)

func (part Part) CreateValidation() bool {
	return true
}

func (part Part) ChangeValidation(newPartInterface interface{}) bool {
	//newPart := newPartInterface.(*Part)
	//valid := true

	return false
}

func (part Part) GetKeyObjectType() string {
	return KEY
}

func (part Part) GetIndexes() [][]string {
	return [][]string{
		{"DocumentId"},
	}
}

func (part Part) GetEntityName() string {
	return ENTITY_NAME
}
func (part Part) GetTagName() string {
	return XML_TAG
}
func (part *Part) SetRelationTxId(relationTxId string) {
	part.RelationTxId = relationTxId
}
func (part *Part) GetRelationTxId() string {
	return part.RelationTxId
}
func (part *Part) SetId(id string) {
	part.Id = id
}
func (part Part) GetId() string {
	return part.Id
}
func (part *Part) SetKey(key string) {
	part.Key = key
}
func (part Part) GetKey() string {
	return part.Key
}
func (part *Part) SetMSPId(MSPId string) {
	part.MSPId = MSPId
}
func (part Part) GetMSPId() string {
	return part.MSPId
}
