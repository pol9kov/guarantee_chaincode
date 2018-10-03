package statement

import (
	"encoding/xml"
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

func (statement Statement) CanBeChangedOn(newStatementInterface interface{}) bool {
	newStatement := newStatementInterface.(*Statement)
	valid := true

	valid = valid &&
		CanChangeStatusOn(statement.Status, newStatement.Status)

	valid = valid &&
		newStatement.Id == statement.Id &&
		newStatement.StatementSigned.Number == statement.StatementSigned.Number &&
		newStatement.StatementSigned.Guarantor == statement.StatementSigned.Guarantor &&
		newStatement.StatementSigned.Beneficiary == statement.StatementSigned.Beneficiary &&
		newStatement.StatementSigned.Principal == statement.StatementSigned.Principal

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
	}

	//todo status changes

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
