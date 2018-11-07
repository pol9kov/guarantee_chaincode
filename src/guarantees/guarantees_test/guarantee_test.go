package guarantees

import (
	"encoding/xml"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"guarantees/guarantees/data/primary/guarantee"
	"guarantees/platform/com"
	"guarantees/platform/sc"
	"testing"
)

var guaranteeObject = guarantee.Guarantee{}
var guarantee1 = readFile("guarantee.xml")
var guarantee1_id = "AEdVQVJBTlRFRQAxAA=="

var guarantee_index_fieldPaths = []com.FieldPath{
	com.FieldPath{FieldPath: []string{"GuaranteeSigned", "StatementFields", "Principal", "Organization", "INN"}},
	com.FieldPath{FieldPath: []string{"GuaranteeSigned", "StatementFields", "Beneficiary", "Organization", "INN"}},
	com.FieldPath{FieldPath: []string{"GuaranteeSigned", "StatementFields", "Guarantor", "INN"}},
}

var guarantee_index_values = []string{"7706107510", "7708503727", "7707083893"}

func CreateGuarantee(t *testing.T, stub *shim.MockStub) {
	CreateStatement(t, stub)
	createEntityFromRelationId(t, stub, &guaranteeObject, guarantee1, statement1_id, guarantee1, guarantee1_id)
}

func InitGuarantee(t *testing.T) *shim.MockStub {
	sc := new(sc.SmartContract)
	stub := shim.NewMockStub("guarantee_cc_test", sc)

	checkInit(t, stub, [][]byte{[]byte("init")})

	CreateGuarantee(t, stub)

	return stub
}

func Test_Guarantee_Invoke(t *testing.T) {
	InitGuarantee(t)
}

func Test_Guarantee_EditField(t *testing.T) {
	stub := InitGuarantee(t)

	requestStr, _ := xml.Marshal(com.Request{EntityName: guarantee.ENTITY_NAME, Type: "Fields", TypeAttr: "ById", Args: []string{guarantee1_id}, FieldPaths: []com.FieldPath{com.FieldPath{FieldPath: []string{"GuaranteeSigned", "BankPars"}}}, FieldValues: []string{"<bank_pars></bank_pars>"}})
	checkInvoke(t, stub, "edit", []string{string(requestStr)})

}

func Test_Guarantee_QueryIndexes(t *testing.T) {
	stub := InitGuarantee(t)

	expectResult := "<" + guarantee.XML_TAG + ">" + guarantee1 + "</" + guarantee.XML_TAG + ">"
	for i, fieldPath := range guarantee_index_fieldPaths {
		requestStr, _ := xml.Marshal(com.Request{EntityName: guarantee.ENTITY_NAME, Type: "By", TypeAttr: "Index", FieldPaths: []com.FieldPath{fieldPath}, FieldValues: []string{guarantee_index_values[i]}})
		checkQuery(t, stub, "query", []string{string(requestStr)}, expectResult)
	}
}

func Test_Guarantee_QueryByRelation(t *testing.T) {
	stub := InitGuarantee(t)

	requestStr, _ := xml.Marshal(com.Request{EntityName: guarantee.ENTITY_NAME, Type: "By", TypeAttr: "RelationId", Args: []string{statement1_id}})
	checkQuery(t, stub, "query", []string{string(requestStr)}, guarantee1)
}
