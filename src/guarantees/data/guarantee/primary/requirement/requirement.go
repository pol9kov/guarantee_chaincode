package requirement

const (
	// Name of entity (for logs)
	ENTITY_NAME = "Requirement"

	// Xml tag name
	XML_TAG = "requirements"

	// Object type names (for storage)
	KEY = "REQUIREMENT"
)

func CanChangeStatusOn(oldStatus, newStatus string) bool {
	var statusMap = make(map[string][]string)
	statusMap["validationErr"] = []string{"validationErr", "created", "cancelled"}
	statusMap["created"] = []string{"validationErr", "created", "cancelled", "sent"}
	statusMap["cancelled"] = []string{}
	statusMap["sent"] = []string{"inProgress", "revoked"}
	statusMap["inProgress"] = []string{"inProgress", "paid", "rejected", "paused", "revoked"}
	statusMap["paid"] = []string{}
	statusMap["rejected"] = []string{}
	statusMap["paused"] = []string{"paid", "rejected"}
	statusMap["revoked"] = []string{}

	for _, status := range statusMap[oldStatus] {
		if status == newStatus {
			return true
		}
	}

	return false
}

func CanChangeInternalStatusOn(oldInternalStatus, newInternalStatus string) bool {
	var internalStatusMap = make(map[string][]string)
	internalStatusMap["validationErr"] = []string{"validationErr", "created", "cancelled"}
	internalStatusMap["created"] = []string{"validationErr", "created", "cancelled", "sent"}
	internalStatusMap["cancelled"] = []string{}
	internalStatusMap["sent"] = []string{"inProgress", "revoked"}
	internalStatusMap["inProgress"] = []string{"paid", "rejectCreated", "paused", "revoked"}
	internalStatusMap["paid"] = []string{}
	internalStatusMap["rejectCreated"] = []string{"rejectToSign", "revoked"}
	internalStatusMap["rejectToSign"] = []string{"rejected", "revoked"}
	internalStatusMap["rejected"] = []string{}
	internalStatusMap["paused"] = []string{"paid", "rejected"}
	internalStatusMap["revoked"] = []string{}

	for _, internalStatus := range internalStatusMap[oldInternalStatus] {
		if internalStatus == newInternalStatus {
			return true
		}
	}

	return false
}

func (requirement Requirement) CanBeChangedOn(newRequirementInterface interface{}) bool {
	newRequirement := newRequirementInterface.(*Requirement)
	valid := true

	valid = valid &&
		CanChangeStatusOn(requirement.Status, newRequirement.Status)
	valid = valid &&
		CanChangeInternalStatusOn(requirement.InternalStatus, newRequirement.InternalStatus)

	return valid
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
