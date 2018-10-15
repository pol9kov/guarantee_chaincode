package rrequirement

import (
	"guarantees/com"
	"regexp"
)

const (
	// Name of entity (for logs)
	ENTITY_NAME = "RRequirement"

	// Xml tag name
	XML_TAG = "rrequirements"

	// Object type names (for storage)
	KEY = "RREQUIREMENT"
)

func (rrequirement RRequirement) regExpCheck() bool {
	valid := true

	for _, par := range rrequirement.RequirementSigned.RequirementTemplate.Pars {
		if par.Value != "" {
			match, err := regexp.MatchString(par.RegularExpression, par.Value)
			if err != nil {
				return false
			}
			valid = valid && match
			if valid == false {
				com.DebugLogMsg("RegularExpression is not valid")
				return false
			}
		}
	}

	return valid
}

func (rrequirement RRequirement) CanCreate() bool {
	return rrequirement.regExpCheck()
}

func (rrequirement RRequirement) CanBeChangedOn(newRRequirementInterface interface{}) bool {
	newRRequirement := newRRequirementInterface.(*RRequirement)
	valid := true

	valid = valid &&
		newRRequirement.regExpCheck()

	return valid

	//todo status changes
}

func (rrequirement RRequirement) GetKeyObjectType() string {
	return KEY
}

func (rrequirement RRequirement) GetIndexes() [][]string {
	return [][]string{
		{"GuaranteeId"},
		{"RRequirementSigned", "Principal", "Organization", "INN"},
		{"RRequirementSigned", "Beneficiary", "Organization", "INN"},
		{"RRequirementSigned", "Guarantor", "INN"},
	}
}

func (rrequirement RRequirement) GetEntityName() string {
	return ENTITY_NAME
}
func (rrequirement RRequirement) GetTagName() string {
	return XML_TAG
}
func (rrequirement *RRequirement) SetRelationTxId(relationTxId string) {
	rrequirement.RelationTxId = relationTxId
}
func (rrequirement *RRequirement) GetRelationTxId() string {
	return rrequirement.RelationTxId
}
func (rrequirement *RRequirement) SetId(id string) {
	rrequirement.Id = id
}
func (rrequirement RRequirement) GetId() string {
	return rrequirement.Id
}
func (rrequirement *RRequirement) SetKey(key string) {
	rrequirement.Key = key
}
func (rrequirement RRequirement) GetKey() string {
	return rrequirement.Key
}
func (rrequirement *RRequirement) SetMSPId(MSPId string) {
	rrequirement.MSPId = MSPId
}
func (rrequirement RRequirement) GetMSPId() string {
	return rrequirement.MSPId
}
