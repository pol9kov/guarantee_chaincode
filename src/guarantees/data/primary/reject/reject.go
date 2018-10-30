package reject

import (
	"platform/com"
	"regexp"
)

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
	statusMap["readyToSign"] = []string{"created", "cancelled", "finished"}
	statusMap["cancelled"] = []string{}
	statusMap["finished"] = []string{}

	for _, status := range statusMap[oldStatus] {
		if status == newStatus {
			return true
		}
	}

	return false
}

func (reject Reject) regExpCheck() bool {
	valid := true

	for _, par := range reject.Rejtype.Pars {
		if par.Value != "" {
			match, err := regexp.MatchString(par.RegularExpression, par.Value)
			if err != nil {
				return false
			}
			valid = valid && match
			if valid == false {
				com.DebugLogMsg("RegularExpression is not valid for par: " + par.Name)
				return false
			}
		}
	}

	return valid
}

func (reject Reject) NeedCreateValidation() bool {
	if reject.Status == "validationErr" {
		return false
	}
	return true
}

func (reject Reject) NeedChangeValidation() bool {
	return reject.NeedCreateValidation()
}

func (reject Reject) CreateValidation() bool {
	if !reject.NeedCreateValidation() {
		return true
	}
	return reject.regExpCheck()
}

func (reject Reject) ChangeValidation(newRejectInterface interface{}) bool {
	if !reject.NeedChangeValidation() {
		return true
	}
	newReject := newRejectInterface.(*Reject)
	valid := true

	valid = valid &&
		CanChangeStatusOn(reject.Status, newReject.Status)
	if valid == false {
		com.DebugLogMsg("Status cann't be changed from " + reject.Status + " to " + newReject.Status)
		return false
	}

	valid = valid &&
		newReject.Id == reject.Id &&
		newReject.RelationTxId == reject.RelationTxId &&
		newReject.EntityId == reject.EntityId &&
		newReject.EntityType == reject.EntityType &&
		newReject.CreateDate == reject.CreateDate &&
		newReject.CreateUser == reject.CreateUser &&
		newReject.Number == reject.Number &&
		newReject.Rejtype.Id == reject.Rejtype.Id
	if valid == false {
		com.DebugLogMsg("Fields cann't be changed")
		return false
	}

	for i, par := range reject.Rejtype.Pars {
		if par.Name == "conclusion" ||
			par.Name == "attachment" ||
			par.Name == "delegateFio" ||
			par.Name == "delegatePosition" ||
			par.Name == "docDelegate" ||
			par.Name == "docDelegateNumber" ||
			par.Name == "docDelegateDate" {
			if reject.Status != "created" && reject.Status != "validationErr" {
				valid = valid &&
					par.Value == newReject.Rejtype.Pars[i].Value &&
					par.Name == newReject.Rejtype.Pars[i].Name
				if valid == false {
					com.DebugLogMsg("Par " + par.Name + " cann't be changed in this status!")
					return false
				}
			}
		} else {
			valid = valid &&
				par.Value == newReject.Rejtype.Pars[i].Value &&
				par.Name == newReject.Rejtype.Pars[i].Name
			if valid == false {
				com.DebugLogMsg("Par " + par.Name + " cann't be never changed!")
				return false
			}
		}
	}

	valid = valid &&
		newReject.regExpCheck()

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
