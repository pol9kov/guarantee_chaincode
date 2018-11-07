package guarantees

import (
	"encoding/xml"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"guarantees/guarantees/data/primary/statement"
	"guarantees/platform/com"
	"guarantees/platform/sc"
	"testing"
)

var statementObject = statement.Statement{}
var statement1 = readFile("statement.xml")
var statement1_edit = readFile("statement_edit.xml")
var statement1_id = "AFNUQVRFTUVOVAAxAA=="

var statement_index_fieldPaths = []com.FieldPath{
	com.FieldPath{FieldPath: []string{"StatementSigned", "Principal", "Organization", "INN"}},
	com.FieldPath{FieldPath: []string{"StatementSigned", "Beneficiary", "Organization", "INN"}},
	com.FieldPath{FieldPath: []string{"StatementSigned", "Guarantor", "INN"}},
}

var statement_index_values = []string{"7706107510", "7708503727", "7707083893"}

func CreateStatement(t *testing.T, stub *shim.MockStub) {
	createEntity(t, stub, &statementObject, statement1, statement1, statement1_id)
}

func InitStatement(t *testing.T) *shim.MockStub {
	sc := new(sc.SmartContract)
	stub := shim.NewMockStub("guarantee_cc_test", sc)

	checkInit(t, stub, [][]byte{[]byte("init")})

	CreateStatement(t, stub)

	return stub
}

func Test_Statement_Invoke(t *testing.T) {
	InitStatement(t)
}

func Test_Statement_EditAll(t *testing.T) {
	stub := InitStatement(t)

	requestStr, _ := xml.Marshal(com.Request{EntityName: statement.ENTITY_NAME, Type: "AllFields", Args: []string{statement1_edit}})
	checkInvoke(t, stub, "edit", []string{string(requestStr)})

	expectResult := "<" + statement.XML_TAG + ">" + statement1_edit + "</" + statement.XML_TAG + ">"
	requestStr, _ = xml.Marshal(com.Request{EntityName: statement.ENTITY_NAME, Type: "All", Args: []string{}})
	checkQuery(t, stub, "query", []string{string(requestStr)}, expectResult)

}

func Test_Statement_EditField(t *testing.T) {
	stub := InitStatement(t)

	requestStr, _ := xml.Marshal(com.Request{EntityName: statement.ENTITY_NAME, Type: "Fields", TypeAttr: "ById", Args: []string{statement1_id}, FieldPaths: []com.FieldPath{com.FieldPath{FieldPath: []string{"StatementSigned", "GType", "Pars", "Name", "obligations", "Value"}}}, FieldValues: []string{"MyNew Value 777"}})

	requestStr, _ = xml.Marshal(com.Request{EntityName: statement.ENTITY_NAME, Type: "Fields", TypeAttr: "ById", Args: []string{statement1_id}, FieldPaths: []com.FieldPath{com.FieldPath{FieldPath: []string{"Status"}}}, FieldValues: []string{"validationErr"}})
	checkInvoke(t, stub, "edit", []string{string(requestStr)})

	expectResult := "<" + statement.XML_TAG + ">" + statement1_edit + "</" + statement.XML_TAG + ">"
	requestStr, _ = xml.Marshal(com.Request{EntityName: statement.ENTITY_NAME, Type: "All", Args: []string{}})
	checkQuery(t, stub, "query", []string{string(requestStr)}, expectResult)

}

func Test_Statement_QueryIndexes(t *testing.T) {
	stub := InitStatement(t)

	expectResult := "<" + statement.XML_TAG + ">" + statement1 + "</" + statement.XML_TAG + ">"
	for i, fieldPath := range statement_index_fieldPaths {
		requestStr, _ := xml.Marshal(com.Request{EntityName: statement.ENTITY_NAME, Type: "By", TypeAttr: "Index", FieldPaths: []com.FieldPath{fieldPath}, FieldValues: []string{statement_index_values[i]}})
		checkQuery(t, stub, "query", []string{string(requestStr)}, expectResult)
	}
}

func Test_Statement_QueryByRelation(t *testing.T) {
	stub := InitGuarantee(t)

	requestStr, _ := xml.Marshal(com.Request{EntityName: statement.ENTITY_NAME, Type: "By", TypeAttr: "RelationId", Args: []string{guarantee1_id}})
	checkQuery(t, stub, "query", []string{string(requestStr)}, statement1)

}
