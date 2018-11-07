package guarantees

import (
	"encoding/xml"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"guarantees/guarantees/data/additional/bankpars"
	"guarantees/platform/com"
	"guarantees/platform/sc"
	"testing"
)

var bankparsObject = bankpars.BankPars{}
var bankpars1 = readFile("bankpars.xml")
var bankpars1_id = "AEJBTktQQVJTADEA"

func CreateBankPars(t *testing.T, stub *shim.MockStub) {
	createEntityFromRelationId(t, stub, &bankparsObject, bankpars1, guarantor1_id, bankpars1, bankpars1_id)
}

func InitBankPars(t *testing.T) *shim.MockStub {
	sc := new(sc.SmartContract)
	stub := shim.NewMockStub("guarantee_cc_test", sc)

	checkInit(t, stub, [][]byte{[]byte("init")})

	CreateBankPars(t, stub)

	return stub
}

func Test_BankPars_Invoke(t *testing.T) {
	InitBankPars(t)
}

func Test_BankPars_QueryByRelation(t *testing.T) {
	stub := InitBankPars(t)

	requestStr, _ := xml.Marshal(com.Request{EntityName: bankpars.ENTITY_NAME, Type: "By", TypeAttr: "RelationId", Args: []string{guarantor1_id}})
	checkQuery(t, stub, "query", []string{string(requestStr)}, bankpars1)

}
