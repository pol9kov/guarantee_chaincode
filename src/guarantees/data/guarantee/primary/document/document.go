package document

const (
	// Name of entity (for logs)
	ENTITY_NAME = "Document"

	// Xml tag name
	XML_TAG = "documents"

	// Object type names (for storage)
	KEY = "DOCUMENT"
)

func (document Document) CanBeChangedOn(newDocumentInterface interface{}) bool {
	//newDocument := newDocumentInterface.(*Document)
	//valid := true

	return false
}

func (document Document) GetKeyObjectType() string {
	return KEY
}

func (document Document) GetIndexes() [][]string {
	return [][]string{
		{"EntityId"},
	}
}

func (document Document) GetEntityName() string {
	return ENTITY_NAME
}
func (document Document) GetTagName() string {
	return XML_TAG
}
func (document *Document) SetRelationTxId(relationTxId string) {
	document.RelationTxId = relationTxId
}
func (document *Document) GetRelationTxId() string {
	return document.RelationTxId
}
func (document *Document) SetId(id string) {
	document.Id = id
}
func (document Document) GetId() string {
	return document.Id
}
func (document *Document) SetKey(key string) {
	document.Key = key
}
func (document Document) GetKey() string {
	return document.Key
}
func (document *Document) SetMSPId(MSPId string) {
	document.MSPId = MSPId
}
func (document Document) GetMSPId() string {
	return document.MSPId
}
