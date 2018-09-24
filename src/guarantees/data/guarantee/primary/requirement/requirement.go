package requirement

const (
	// Name of entity (for logs)
	ENTITY_NAME = "Requirement"

	// Xml tag name
	XML_TAG = "requirements"

	// Object type names (for storage)
	KEY = "REQUIREMENT"
)

func (requirement Requirement) CanBeChangedOn(newRequirementInterface interface{}) bool {
	_ = newRequirementInterface.(*Requirement)
	return true

	//todo status changes
}

func (requirement Requirement) GetKeyObjectType() string {
	return KEY
}

func (requirement Requirement) GetIndexes() [][]string {
	return [][]string{
		{"GuaranteeId"},
		{"RequirementSigned", "Principal", "Organization", "INN"},
		{"RequirementSigned", "Beneficiary", "Organization", "INN"},
		{"RequirementSigned", "Guarantor", "INN"},
	}
}

func (requirement Requirement) GetEntityName() string {
	return ENTITY_NAME
}
func (requirement Requirement) GetTagName() string {
	return XML_TAG
}
func (requirement *Requirement) SetRelationTxId(relationTxId string) {
	requirement.RelationTxId = relationTxId
}
func (requirement *Requirement) GetRelationTxId() string {
	return requirement.RelationTxId
}
func (requirement *Requirement) SetId(id string) {
	requirement.Id = id
}
func (requirement Requirement) GetId() string {
	return requirement.Id
}
func (requirement *Requirement) SetKey(key string) {
	requirement.Key = key
}
func (requirement Requirement) GetKey() string {
	return requirement.Key
}
func (requirement *Requirement) SetMSPId(MSPId string) {
	requirement.MSPId = MSPId
}
func (requirement Requirement) GetMSPId() string {
	return requirement.MSPId
}
