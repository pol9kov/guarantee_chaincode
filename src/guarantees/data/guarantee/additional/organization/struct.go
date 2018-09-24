package organization

import "encoding/xml"

// Define  structures.
type Organization struct {
	XMLName xml.Name `xml:"organization"`

	RelationTxId string `xml:"relation_tx_id"`
	Key          string `xml:"key"`
	Id           string `xml:"id"`

	MSPId string `xml:"msp_id"`

	INN      string `xml:"inn"`
	OGRN     string `xml:"ogrn"`
	KPP      string `xml:"kpp"`
	Okved    string `xml:"okved"`
	FullName string `xml:"full_name"`
	Name     string `xml:"name"`
	Address  string `xml:"address"`
}

// Define  structures.
type OrganizationOut struct {
	XMLName xml.Name `xml:"organization"`

	Id string `xml:"id"`

	MSPId string `xml:"msp_id"`

	INN      string `xml:"inn"`
	OGRN     string `xml:"ogrn"`
	KPP      string `xml:"kpp"`
	Okved    string `xml:"okved"`
	FullName string `xml:"full_name"`
	Name     string `xml:"name"`
	Address  string `xml:"address"`
}

func (organization Organization) ToOut() interface{} {
	result := OrganizationOut{}

	result.Id = organization.Id
	result.MSPId = organization.MSPId
	result.INN = organization.INN
	result.OGRN = organization.OGRN
	result.KPP = organization.KPP
	result.Okved = organization.Okved
	result.FullName = organization.FullName
	result.Name = organization.Name
	result.Address = organization.Address

	return result
}
