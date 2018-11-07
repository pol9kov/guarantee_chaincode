package requirement

import (
	"encoding/xml"
	"guarantees/guarantees/data/additional"
	"guarantees/guarantees/data/additional/guarantor"
	"guarantees/guarantees/data/additional/rtype"
)

type Requirement struct {
	XMLName xml.Name `xml:"requirement"`

	RelationTxId string `xml:"relation_tx_id"`
	Key          string `xml:"key"`
	Id           string `xml:"id"`

	MSPId string `xml:"msp_id"`

	Status               string `xml:"status"`
	InternalStatus       string `xml:"internal_status"`
	GuaranteeId          string `xml:"guarantee_id"`
	PlanExecutionDate    string `xml:"plan_execution_date"`
	UpdateDate           string `xml:"update_date"`
	PrimaryId            string `xml:"primary_id"`
	BeneficiarySignature string `xml:"beneficiary_signature"`

	RequirementSigned RequirementSigned `xml:"requirement_signed"`
}

type RequirementSigned struct {
	XMLName xml.Name `xml:"requirement_signed"`

	Number              string                 `xml:"number"`
	CreateDate          string                 `xml:"createDate"`
	GuaranteeNumber     string                 `xml:"guarantee_number"`
	GuaranteeDate       string                 `xml:"guarantee_date"`
	Second              string                 `xml:"second"`
	Beneficiary         additional.Beneficiary `xml:"beneficiary"`
	Principal           additional.Principal   `xml:"principal"`
	Guarantor           guarantor.Guarantor    `xml:"guarantor"`
	RequirementTemplate rtype.RType            `xml:"requirement_template"`
}

type RequirementOut struct {
	XMLName xml.Name `xml:"requirement"`

	Id string `xml:"id"`

	Status               string `xml:"status"`
	InternalStatus       string `xml:"internal_status"`
	GuaranteeId          string `xml:"guarantee_id"`
	PlanExecutionDate    string `xml:"plan_execution_date"`
	UpdateDate           string `xml:"update_date"`
	PrimaryId            string `xml:"primary_id"`
	BeneficiarySignature string `xml:"beneficiary_signature"`

	RequirementSigned interface{} `xml:"requirement_signed"`
}

type RequirementSignedOut struct {
	XMLName xml.Name `xml:"requirement_signed"`

	Number              string      `xml:"number"`
	CreateDate          string      `xml:"createDate"`
	GuaranteeNumber     string      `xml:"guarantee_number"`
	GuaranteeDate       string      `xml:"guarantee_date"`
	Second              string      `xml:"second"`
	Beneficiary         interface{} `xml:"beneficiary"`
	Principal           interface{} `xml:"principal"`
	Guarantor           interface{} `xml:"guarantor"`
	RequirementTemplate interface{} `xml:"requirement_template"`
}

func (requirement Requirement) ToOut() interface{} {
	result := RequirementOut{}

	result.Id = requirement.Id
	result.Status = requirement.Status
	result.InternalStatus = requirement.InternalStatus
	result.GuaranteeId = requirement.GuaranteeId
	result.PlanExecutionDate = requirement.PlanExecutionDate
	result.UpdateDate = requirement.UpdateDate
	result.PrimaryId = requirement.PrimaryId
	result.UpdateDate = requirement.UpdateDate
	result.BeneficiarySignature = requirement.BeneficiarySignature
	result.RequirementSigned = requirement.RequirementSigned.ToOut()

	return result
}

func (requirementSigned RequirementSigned) ToOut() interface{} {
	result := RequirementSignedOut{}

	result.Number = requirementSigned.Number
	result.CreateDate = requirementSigned.CreateDate
	result.GuaranteeNumber = requirementSigned.GuaranteeNumber
	result.GuaranteeDate = requirementSigned.GuaranteeDate
	result.Second = requirementSigned.Second
	result.Beneficiary = (requirementSigned.Beneficiary.ToOut())
	result.Principal = (requirementSigned.Principal.ToOut())
	result.Guarantor = (requirementSigned.Guarantor.ToOut())
	result.RequirementTemplate = (requirementSigned.RequirementTemplate.ToOut())

	return result
}
