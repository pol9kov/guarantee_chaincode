package guarantor

import (
	"encoding/xml"
)

// Define  structures.
type Guarantor struct {
	XMLName xml.Name `xml:"guarantor"`

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
	BIK      string `xml:"bik"`
	KorrAcc  string `xml:"korr_acc"`
	License  string `xml:"license"`
}

// Define  structures.
type GuarantorOut struct {
	XMLName xml.Name `xml:"guarantor"`

	Id string `xml:"id"`

	MSPId string `xml:"msp_id"`

	INN      string `xml:"inn"`
	OGRN     string `xml:"ogrn"`
	KPP      string `xml:"kpp"`
	Okved    string `xml:"okved"`
	FullName string `xml:"full_name"`
	Name     string `xml:"name"`
	Address  string `xml:"address"`
	BIK      string `xml:"bik"`
	KorrAcc  string `xml:"korr_acc"`
	License  string `xml:"license"`
}

func (guarantor Guarantor) ToOut() interface{} {
	result := GuarantorOut{}

	result.Id = guarantor.Id
	result.MSPId = guarantor.MSPId
	result.INN = guarantor.INN
	result.OGRN = guarantor.OGRN
	result.KPP = guarantor.KPP
	result.Okved = guarantor.Okved
	result.FullName = guarantor.FullName
	result.Name = guarantor.Name
	result.Address = guarantor.Address
	result.BIK = guarantor.BIK
	result.KorrAcc = guarantor.KorrAcc
	result.License = guarantor.License

	return result
}
