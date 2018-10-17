package statement

import (
	"encoding/xml"
	"guarantees/com"
	"guarantees/funcs"
	"regexp"
)

const (
	// Name of entity (for logs)
	ENTITY_NAME = "Statement"

	// Xml tag name
	XML_TAG = "statements"

	// Object type names (for storage)
	KEY = "STATEMENT"
)

func CanChangeStatusOn(oldStatus, newStatus string) bool {
	var statusMap = make(map[string][]string)
	statusMap["validationErr"] = []string{"validationErr", "created", "cancelled"}
	statusMap["created"] = []string{"validationErr", "created", "cancelled", "sent"}
	statusMap["cancelled"] = []string{}
	statusMap["sent"] = []string{"inProgress", "revoked"}
	statusMap["inProgress"] = []string{"unsigned", "rejected", "finished", "revoked"}
	statusMap["unsigned"] = []string{"validationErr", "created", "revoked"}
	statusMap["rejected"] = []string{}
	statusMap["finished"] = []string{}
	statusMap["revoked"] = []string{}

	for _, status := range statusMap[oldStatus] {
		if status == newStatus {
			return true
		}
	}

	return false
}

func (statement Statement) regExpCheck() bool {
	valid := true

	for _, par := range statement.StatementSigned.GType.Pars {
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

func (statement Statement) NeedCreateValidation() bool {
	if statement.Status == "validationErr" {
		return false
	}
	return true
}

func (statement Statement) NeedChangeValidation() bool {
	return statement.NeedCreateValidation()
}

func (statement Statement) CreateValidation() bool {
	if !statement.NeedCreateValidation() {
		return true
	}
	return statement.regExpCheck()
}

func (statement Statement) ChangeValidation(newStatementInterface interface{}) bool {
	newStatement := newStatementInterface.(*Statement)

	if !newStatement.NeedChangeValidation() {
		return true
	}

	valid := true

	valid = valid &&
		CanChangeStatusOn(statement.Status, newStatement.Status)
	if valid == false {
		com.DebugLogMsg("Status cann't be changed from " + statement.Status + " to " + newStatement.Status)
		return false
	}

	valid = valid &&
		newStatement.Id == statement.Id &&
		newStatement.StatementSigned.Number == statement.StatementSigned.Number &&
		funcs.IsFieldsEqual(&newStatement.StatementSigned.Guarantor, &statement.StatementSigned.Guarantor) &&
		funcs.IsFieldsEqual(&newStatement.StatementSigned.Beneficiary, &statement.StatementSigned.Beneficiary) &&
		funcs.IsFieldsEqual(&newStatement.StatementSigned.Principal, &statement.StatementSigned.Principal)
	if valid == false {
		com.DebugLogMsg("Fields cann't be changed")
		return false
	}

	if statement.Status != "created" && statement.Status != "validationErr" {
		newSS, err := xml.Marshal(newStatement.StatementSigned)
		if err != nil {
			return false
		}
		oldSS, err := xml.Marshal(statement.StatementSigned)
		if err != nil {
			return false
		}

		valid = valid &&
			string(newSS) == string(oldSS)

		if valid == false {
			com.DebugLogMsg("StatementSigned cann't be changed")
			return false
		}
	}

	newStatement.regExpCheck()

	return valid
}

func (statement Statement) GetKeyObjectType() string {
	return KEY
}

func (statement Statement) GetIndexes() [][]string {
	return [][]string{
		{"StatementSigned", "Principal", "Organization", "INN"},
		{"StatementSigned", "Beneficiary", "Organization", "INN"},
		{"StatementSigned", "Guarantor", "INN"},
	}
}

func (statement Statement) GetEntityName() string {
	return ENTITY_NAME
}
func (statement Statement) GetTagName() string {
	return XML_TAG
}
func (statement *Statement) SetRelationTxId(relationTxId string) {
	statement.RelationTxId = relationTxId
}
func (statement *Statement) GetRelationTxId() string {
	return statement.RelationTxId
}
func (statement *Statement) SetId(id string) {
	statement.Id = id
}
func (statement Statement) GetId() string {
	return statement.Id
}
func (statement *Statement) SetKey(key string) {
	statement.Key = key
}
func (statement Statement) GetKey() string {
	return statement.Key
}
func (statement *Statement) SetMSPId(MSPId string) {
	statement.MSPId = MSPId
}
func (statement Statement) GetMSPId() string {
	return statement.MSPId
}
