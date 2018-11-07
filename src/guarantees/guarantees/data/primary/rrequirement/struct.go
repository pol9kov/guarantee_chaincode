package rrequirement

import (
	"encoding/xml"
	"guarantees/guarantees/data/additional"
	"guarantees/guarantees/data/additional/guarantor"
	"guarantees/guarantees/data/additional/rtype"
)

type RRequirement struct {
	XMLName xml.Name `xml:"rrequirement"`

	RelationTxId string `xml:"relation_tx_id"`
	Key          string `xml:"key"`
	Id           string `xml:"id"`

	MSPId string `xml:"msp_id"`

	Status               string `xml:"status"`
	GuaranteeId          string `xml:"guarantee_id"`
	PlanExecutionDate    string `xml:"plan_execution_date"`
	UpdateDate           string `xml:"update_date"`
	PrimaryId            string `xml:"primary_id"`
	BeneficiarySignature string `xml:"beneficiary_signature"`

	RequirementSigned RRequirementSigned `xml:"requirement_signed"`
}

type RRequirementSigned struct {
	XMLName xml.Name `xml:"rrequirement_signed"`

	Number              string                 `xml:"number"`
	CreateDate          string                 `xml:"createDate"`
	Second              string                 `xml:"second"`
	Beneficiary         additional.Beneficiary `xml:"beneficiary"`
	Principal           additional.Principal   `xml:"principal"`
	Guarantor           guarantor.Guarantor    `xml:"guarantor"`
	RequirementTemplate rtype.RType            `xml:"requirement_template"`
}

type RRequirementOut struct {
	XMLName xml.Name `xml:"rrequirement"`

	Id string `xml:"id"`

	Status               string `xml:"status"`
	GuaranteeId          string `xml:"guarantee_id"`
	PlanExecutionDate    string `xml:"plan_execution_date"`
	UpdateDate           string `xml:"update_date"`
	PrimaryId            string `xml:"primary_id"`
	BeneficiarySignature string `xml:"beneficiary_signature"`

	RequirementSigned interface{} `xml:"requirement_signed"`
}

type RRequirementSignedOut struct {
	XMLName xml.Name `xml:"rrequirement_signed"`

	Number              string      `xml:"number"`
	CreateDate          string      `xml:"createDate"`
	Second              string      `xml:"second"`
	Beneficiary         interface{} `xml:"beneficiary"`
	Principal           interface{} `xml:"principal"`
	Guarantor           interface{} `xml:"guarantor"`
	RequirementTemplate interface{} `xml:"requirement_template"`
}

func (rrequirement RRequirement) ToOut() interface{} {
	result := RRequirementOut{}

	result.Id = rrequirement.Id
	result.Status = rrequirement.Status
	result.PlanExecutionDate = rrequirement.PlanExecutionDate
	result.UpdateDate = rrequirement.UpdateDate
	result.PrimaryId = rrequirement.PrimaryId
	result.UpdateDate = rrequirement.UpdateDate
	result.BeneficiarySignature = rrequirement.BeneficiarySignature
	result.RequirementSigned = rrequirement.RequirementSigned.ToOut()

	return result
}

func (rrequirementSigned RRequirementSigned) ToOut() interface{} {
	result := RRequirementSignedOut{}

	result.Number = rrequirementSigned.Number
	result.CreateDate = rrequirementSigned.CreateDate
	result.Second = rrequirementSigned.Second
	result.Beneficiary = (rrequirementSigned.Beneficiary.ToOut())
	result.Principal = (rrequirementSigned.Principal.ToOut())
	result.Guarantor = (rrequirementSigned.Guarantor.ToOut())
	result.RequirementTemplate = (rrequirementSigned.RequirementTemplate.ToOut())

	return result
}
