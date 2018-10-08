package guarantee

import (
	"encoding/xml"
	"guarantees/com"
)

const (
	// Name of entity (for logs)
	ENTITY_NAME = "Guarantee"

	// Xml tag name
	XML_TAG = "guarantees"

	// Object type names (for storage)
	KEY = "GUARANTEE"
)

func (guarantee Guarantee) GetKeyObjectType() string {
	return KEY
}

func (guarantee Guarantee) GetIndexes() [][]string {
	return [][]string{
		{"GuaranteeSigned", "StatementFields", "Principal", "Organization", "INN"},
		{"GuaranteeSigned", "StatementFields", "Beneficiary", "Organization", "INN"},
		{"GuaranteeSigned", "StatementFields", "Guarantor", "INN"},
	}
}

func CanChangeStatusOn(oldStatus, newStatus string) bool {
	var statusMap = make(map[string][]string)
	statusMap["validationErr"] = []string{"validationErr", "created", "readyToSign"}
	statusMap["created"] = []string{"validationErr", "created", "readyToSign", "issued"}
	statusMap["readyToSign"] = []string{"issued"}
	statusMap["issued"] = []string{"requirementReceived"}
	statusMap["requirementReceived"] = []string{"issued", "closed", "wavedOfRights"}
	statusMap["closed"] = []string{}
	statusMap["wavedOfRights"] = []string{}

	for _, status := range statusMap[oldStatus] {
		if status == newStatus {
			return true
		}
	}

	return false
}

func (guarantee Guarantee) CanBeChangedOn(newGuaranteeInterface interface{}) bool {
	newGuarantee := newGuaranteeInterface.(*Guarantee)
	valid := true

	valid = valid &&
		CanChangeStatusOn(guarantee.Status, newGuarantee.Status)
	if valid == false {
		com.DebugLogMsg("Status cann't be changed from " + guarantee.Status + " to " + newGuarantee.Status)
		return false
	}

	valid = valid &&
		newGuarantee.Id == guarantee.Id &&
		newGuarantee.RelationTxId == guarantee.RelationTxId
	if valid == false {
		com.DebugLogMsg("Fields cann't be changed")
		return false
	}

	if guarantee.Status != "created" && guarantee.Status != "validationErr" && guarantee.Status != "readyToSign" {
		newGS, err := xml.Marshal(newGuarantee.GuaranteeSigned)
		if err != nil {
			return false
		}
		oldGS, err := xml.Marshal(guarantee.GuaranteeSigned)
		if err != nil {
			return false
		}

		valid = valid &&
			string(newGS) == string(oldGS)
		if valid == false {
			com.DebugLogMsg("GuaranteeSigned cann't be changed")
			return false
		}
	}

	if guarantee.Status == "readyToSign" {
		for i, par := range guarantee.GuaranteeSigned.StatementFields.GType.Pars {
			if par.Name != "expirationDate" || par.Name != "planIssueDate" {
				valid = valid &&
					par.Value == newGuarantee.GuaranteeSigned.StatementFields.GType.Pars[i].Value &&
					par.Name == newGuarantee.GuaranteeSigned.StatementFields.GType.Pars[i].Name
				if valid == false {
					com.DebugLogMsg("Par " + par.Name + " cann't be changed")
					return false
				}
			}
		}
	}

	return valid
}

func (guarantee Guarantee) GetEntityName() string {
	return ENTITY_NAME
}
func (guarantee Guarantee) GetTagName() string {
	return XML_TAG
}
func (guarantee *Guarantee) SetRelationTxId(relationTxId string) {
	guarantee.RelationTxId = relationTxId
}
func (guarantee *Guarantee) GetRelationTxId() string {
	return guarantee.RelationTxId
}
func (guarantee *Guarantee) SetId(id string) {
	guarantee.Id = id
}
func (guarantee Guarantee) GetId() string {
	return guarantee.Id
}
func (guarantee *Guarantee) SetKey(key string) {
	guarantee.Key = key
}
func (guarantee Guarantee) GetKey() string {
	return guarantee.Key
}
func (guarantee *Guarantee) SetMSPId(MSPId string) {
	guarantee.MSPId = MSPId
}
func (guarantee Guarantee) GetMSPId() string {
	return guarantee.MSPId
}
