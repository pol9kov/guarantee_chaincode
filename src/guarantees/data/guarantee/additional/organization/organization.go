package organization

const (
	// Name of entity (for logs)
	ENTITY_NAME = "Organization"

	// Xml tag name
	XML_TAG = "organizations"

	// Object type names (for storage)
	KEY = "ORGANIZATION"
)

func (organization Organization) CreateValidation() bool {
	return true
}

func (organization Organization) ChangeValidation(newOrganizationInterface interface{}) bool {
	_ = newOrganizationInterface.(*Organization)
	return false
}

func (organization Organization) GetKeyObjectType() string {
	return KEY
}

func (organization Organization) GetIndexes() [][]string {
	return [][]string{
		{"MSPId"},
	}
}

func (organization Organization) GetEntityName() string {
	return ENTITY_NAME
}
func (organization Organization) GetTagName() string {
	return XML_TAG
}
func (organization *Organization) SetRelationTxId(relationTxId string) {
	organization.RelationTxId = relationTxId
}
func (organization *Organization) GetRelationTxId() string {
	return organization.RelationTxId
}
func (organization *Organization) SetId(id string) {
	organization.Id = id
}
func (organization Organization) GetId() string {
	return organization.Id
}
func (organization *Organization) SetKey(key string) {
	organization.Key = key
}
func (organization Organization) GetKey() string {
	return organization.Key
}
func (organization *Organization) SetMSPId(MSPId string) {
	organization.MSPId = MSPId
}
func (organization Organization) GetMSPId() string {
	return organization.MSPId
}
