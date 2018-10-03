package reject

const (
	// Name of entity (for logs)
	ENTITY_NAME = "Reject"

	// Xml tag name
	XML_TAG = "rejects"

	// Object type names (for storage)
	KEY = "REJECT"
)

func CanChangeStatusOn(oldStatus, newStatus string) bool {
	var statusMap = make(map[string][]string)
	statusMap["created"] = []string{"created", "readyToSign", "cancelled"}
	statusMap["readyToSign"] = []string{"created", "readyToSign", "cancelled", "finished"}
	statusMap["cancelled"] = []string{}
	statusMap["finished"] = []string{}

	for _, status := range statusMap[oldStatus] {
		if status == newStatus {
			return true
		}
	}

	return false
}

func (reject Reject) CanBeChangedOn(newRejectInterface interface{}) bool {
	newReject := newRejectInterface.(*Reject)
	valid := true

	valid = valid &&
		CanChangeStatusOn(reject.Status, newReject.Status)

	return valid
}

func (reject Reject) GetKeyObjectType() string {
	return KEY
}

func (reject Reject) GetIndexes() [][]string {
	return [][]string{
		{"GuaranteeId"},
		{"RejectSigned", "Principal", "Organization", "INN"},
		{"RejectSigned", "Beneficiary", "Organization", "INN"},
		{"RejectSigned", "Guarantor", "INN"},
	}
}

func (reject Reject) GetEntityName() string {
	return ENTITY_NAME
}
func (reject Reject) GetTagName() string {
	return XML_TAG
}
func (reject *Reject) SetRelationTxId(relationTxId string) {
	reject.RelationTxId = relationTxId
}
func (reject *Reject) GetRelationTxId() string {
	return reject.RelationTxId
}
func (reject *Reject) SetId(id string) {
	reject.Id = id
}
func (reject Reject) GetId() string {
	return reject.Id
}
func (reject *Reject) SetKey(key string) {
	reject.Key = key
}
func (reject Reject) GetKey() string {
	return reject.Key
}
func (reject *Reject) SetMSPId(MSPId string) {
	reject.MSPId = MSPId
}
func (reject Reject) GetMSPId() string {
	return reject.MSPId
}
