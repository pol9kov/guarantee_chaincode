package bankpars

import (
	"encoding/xml"
	"guarantees/guarantees/data/additional"
)

type BankPars struct {
	XMLName xml.Name `xml:"bank_pars"`

	RelationTxId string `xml:"relation_tx_id"`
	Key          string `xml:"key"`
	Id           string `xml:"id"`

	MSPId string `xml:"msp_id"`

	GuarantorINN string `xml:"guarantor_inn"`

	Pars []additional.Par `xml:"par,omitempty"`
}

type BankParsOut struct {
	XMLName xml.Name `xml:"bank_pars"`

	Pars []additional.Par `xml:"par,omitempty"`
}

func (bankPars BankPars) ToOut() interface{} {
	result := BankParsOut{}

	result.Pars = bankPars.Pars

	return result
}
