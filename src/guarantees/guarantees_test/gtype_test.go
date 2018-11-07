package guarantees

import (
	"encoding/xml"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"guarantees/guarantees/data/additional/gtype"
	"guarantees/platform/com"
	"guarantees/platform/sc"
	"testing"
)

var gtypeObject = gtype.GType{}
var gtype1 = readFile("gtype.xml")
var gtype1_id = "AEdUWVBFADEA"

var gtype_index_fieldPaths = []com.FieldPath{
	com.FieldPath{FieldPath: []string{"GuarantorId"}},
}

var gtype_index_values = []string{guarantor1_id}

func CreateGType(t *testing.T, stub *shim.MockStub) {
	createEntity(t, stub, &gtypeObject, gtype1, gtype1, gtype1_id)
}

func InitGType(t *testing.T) *shim.MockStub {
	sc := new(sc.SmartContract)
	stub := shim.NewMockStub("guarantee_cc_test", sc)

	checkInit(t, stub, [][]byte{[]byte("init")})

	CreateGType(t, stub)

	return stub
}

func Test_GType_Invoke(t *testing.T) {
	InitGType(t)
}

func Test_GType_QueryByRelation(t *testing.T) {
	stub := InitGType(t)

	requestStr, _ := xml.Marshal(com.Request{EntityName: gtype.ENTITY_NAME, Type: "By", TypeAttr: "RelationId", Args: []string{guarantor1_id}})
	checkQuery(t, stub, "query", []string{string(requestStr)}, gtype1)
}

func Test_GType_QueryIndexes(t *testing.T) {
	stub := InitGType(t)

	expectResult := "<" + gtype.XML_TAG + ">" + gtype1 + "</" + gtype.XML_TAG + ">"
	for i, fieldPath := range gtype_index_fieldPaths {
		requestStr, _ := xml.Marshal(com.Request{EntityName: gtype.ENTITY_NAME, Type: "By", TypeAttr: "Index", FieldPaths: []com.FieldPath{fieldPath}, FieldValues: []string{gtype_index_values[i]}})
		checkQuery(t, stub, "query", []string{string(requestStr)}, expectResult)
	}
}
