package guarantee

import (
	"encoding/xml"
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

func (guarantee Guarantee) CanBeChangedOn(newGuaranteeInterface interface{}) bool {
	newGuarantee := newGuaranteeInterface.(*Guarantee)
	valid := true

	valid = valid &&
		newGuarantee.Id == guarantee.Id &&
		newGuarantee.RelationTxId == guarantee.RelationTxId

	if guarantee.Status != "created" && guarantee.Status != "validationErr" {
		valid = valid &&
			newGuarantee.GuaranteeSigned.IssueDate == guarantee.GuaranteeSigned.IssueDate
	}

	if guarantee.Status != "created" && guarantee.Status != "validationErr" {
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
	}

	if guarantee.Status == "created" && guarantee.Status == "validationErr" {
		for i, par := range guarantee.GuaranteeSigned.StatementFields.GType.Pars {
			if par.Name != "expirationDate" {
				valid = valid &&
					par.Value == newGuarantee.GuaranteeSigned.StatementFields.GType.Pars[i].Value &&
					par.Name == newGuarantee.GuaranteeSigned.StatementFields.GType.Pars[i].Name
			}
		}
	}

	//todo status changes

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
