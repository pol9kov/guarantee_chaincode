package guarantees

import (
	"encoding/xml"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"guarantees/guarantees/data/additional/guarantor"
	"guarantees/platform/com"
	"guarantees/platform/funcs"
	"guarantees/platform/sc"
	"testing"
)

var guarantorObject = guarantor.Guarantor{}
var guarantor1 = readFile("guarantor.xml")
var guarantor1_id = "AEdVQVJBTlRPUgAxAA=="
var guarantor_index_fieldPaths = []com.FieldPath{
	com.FieldPath{FieldPath: []string{"MSPId"}},
}

var guarantor_index_values = []string{funcs.MSPId}

func CreateGuarantor(t *testing.T, stub *shim.MockStub) {
	createEntity(t, stub, &guarantorObject, guarantor1, guarantor1, guarantor1_id)
}

func InitGuarantor(t *testing.T) *shim.MockStub {
	sc := new(sc.SmartContract)
	stub := shim.NewMockStub("guarantee_cc_test", sc)

	checkInit(t, stub, [][]byte{[]byte("init")})

	CreateGuarantor(t, stub)

	return stub
}

func Test_Guarantor_Invoke(t *testing.T) {
	InitGuarantor(t)
}

func Test_Guarantor_QueryIndexes(t *testing.T) {
	stub := InitGuarantor(t)

	expectResult := "<" + guarantor.XML_TAG + ">" + guarantor1 + "</" + guarantor.XML_TAG + ">"
	for i, fieldPath := range guarantor_index_fieldPaths {
		requestStr, _ := xml.Marshal(com.Request{EntityName: guarantor.ENTITY_NAME, Type: "By", TypeAttr: "Index", FieldPaths: []com.FieldPath{fieldPath}, FieldValues: []string{guarantor_index_values[i]}})
		checkQuery(t, stub, "query", []string{string(requestStr)}, expectResult)
	}
}
