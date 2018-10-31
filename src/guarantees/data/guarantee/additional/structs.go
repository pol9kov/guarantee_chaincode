package additional

import (
	"encoding/xml"
	"guarantees/data/guarantee/additional/organization"
)

//todo Я хз куда это деть. Дублировать в заявках и гарантиях я не хочу

type Principal struct {
	Organization organization.Organization `xml:"organization"`
	Delegate     Delegate                  `xml:"delegate"`
}

type Delegate struct {
	FIO      string `xml:"fio"`
	Post     string `xml:"post"`
	UserName string `xml:"username"`
}

type Beneficiary struct {
	Organization organization.Organization `xml:"organization"`
}

type PrincipalOut struct {
	Organization interface{} `xml:"organization`
	Delegate     Delegate    `xml:"delegate"`
}

type BeneficiaryOut struct {
	Organization interface{} `xml:"organization"`
}

func (principal Principal) ToOut() PrincipalOut {
	result := PrincipalOut{}

	result.Organization = principal.Organization.ToOut()
	result.Delegate = principal.Delegate

	return result
}

func (beneficiary Beneficiary) ToOut() BeneficiaryOut {
	result := BeneficiaryOut{}

	result.Organization = beneficiary.Organization.ToOut()

	return result
}

type Par struct {
	XMLName xml.Name `xml:"par"`

	Name                   string                `xml:"name,attr,omitempty"`
	Value                  string                `xml:"value,attr"`
	RegularExpression      string                `xml:"regular_expression,attr"`
	Caption                string                `xml:"caption,attr"`
	DisplayType            string                `xml:"display_type,attr"`
	Placeholder            string                `xml:"placeholder,attr,omitempty"`
	VisibleDefaultValue    string                `xml:"visible_default_value,attr,omitempty"`
	DefaultValue           string                `xml:"default_value,attr,omitempty"`
	Disabled               string                `xml:"disabled,attr,omitempty"`
	Required               string                `xml:"required,attr,omitempty"`
	Group                  string                `xml:"group,attr,omitempty"`
	Subgroup               string                `xml:"subgroup,attr,omitempty"`
	GroupOrder             string                `xml:"group_order,attr,omitempty"`
	Options                []Option              `xml:"option,omitempty"`
	DisabledConditionLogic string                `xml:"disabled_condition_logic,attr,omitempty"`
	DisabledConditions     []DisabledCondition   `xml:"disabled_condition,omitempty"`
	VisibleConditionLogic  string                `xml:"visible_condition_logic,attr,omitempty"`
	VisibleConditions      []VisibleCondition    `xml:"visible_condition,omitempty"`
	ValidationConditions   []ValidationCondition `xml:"validation_condition,omitempty"`
}

type Option struct {
	XMLName xml.Name `xml:"option"`

	Code              string             `xml:"code,attr,omitempty"`
	Value             string             `xml:"value,attr,omitempty"`
	Chosen            string             `xml:"chosen,attr,omitempty"`
	Detailed          string             `xml:"detailed,attr,omitempty"`
	DetailedValue     string             `xml:"detailed_value,attr,omitempty"`
	VisibleConditions []VisibleCondition `xml:"visible_condition,omitempty"`
}

type DisabledCondition struct {
	XMLName xml.Name `xml:"disabled_condition"`

	Name   string `xml:"name,attr,omitempty"`
	Value  string `xml:"value,attr,omitempty"`
	Chosen string `xml:"chosen,attr,omitempty"`
	Type   string `xml:"type,attr,omitempty"`
}

type VisibleCondition struct {
	XMLName xml.Name `xml:"visible_condition"`

	Name   string `xml:"name,attr,omitempty"`
	Value  string `xml:"value,attr,omitempty"`
	Chosen string `xml:"chosen,attr,omitempty"`
	Type   string `xml:"type,attr,omitempty"`
}

type ValidationCondition struct {
	XMLName xml.Name `xml:"validation_condition"`

	Condition string `xml:"condition,attr,omitempty"`
	Message   string `xml:"message,attr,omitempty"`
}
