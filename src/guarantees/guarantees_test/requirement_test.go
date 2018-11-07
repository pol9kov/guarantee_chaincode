package guarantees

import (
	"encoding/xml"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"guarantees/guarantees/data/primary/requirement"
	"guarantees/platform/com"
	"guarantees/platform/sc"
	"testing"
)

var requirementObject = requirement.Requirement{}
var requirement1 = readFile("requirement.xml")
var requirement1_id = "AFJFUVVJUkVNRU5UADEA"

var requirement_index_fieldPaths = []com.FieldPath{
	com.FieldPath{FieldPath: []string{"GuaranteeId"}},
	com.FieldPath{FieldPath: []string{"RequirementSigned", "Principal", "Organization", "INN"}},
	com.FieldPath{FieldPath: []string{"RequirementSigned", "Beneficiary", "Organization", "INN"}},
	com.FieldPath{FieldPath: []string{"RequirementSigned", "Guarantor", "INN"}},
}

var requirement_index_values = []string{guarantee1_id, "7706107510", "7708503727", "7707083893"}

func CreateRequirement(t *testing.T, stub *shim.MockStub) {
	CreateGuarantee(t, stub)
	createEntityFromRelationId(t, stub, &requirementObject, requirement1, guarantee1_id, requirement1, requirement1_id)
}

func InitRequirement(t *testing.T) *shim.MockStub {
	sc := new(sc.SmartContract)
	stub := shim.NewMockStub("requirement_cc_test", sc)

	checkInit(t, stub, [][]byte{[]byte("init")})

	CreateRequirement(t, stub)

	return stub
}

func Test_Requirement_Invoke(t *testing.T) {
	InitRequirement(t)
}

func Test_Requirement_QueryIndexes(t *testing.T) {
	stub := InitRequirement(t)

	expectResult := "<" + requirement.XML_TAG + ">" + requirement1 + "</" + requirement.XML_TAG + ">"
	for i, fieldPath := range requirement_index_fieldPaths {
		requestStr, _ := xml.Marshal(com.Request{EntityName: requirement.ENTITY_NAME, Type: "By", TypeAttr: "Index", FieldPaths: []com.FieldPath{fieldPath}, FieldValues: []string{requirement_index_values[i]}})
		checkQuery(t, stub, "query", []string{string(requestStr)}, expectResult)
	}
}

func Test_Requirement_QueryByRelation(t *testing.T) {
	stub := InitRequirement(t)

	requestStr, _ := xml.Marshal(com.Request{EntityName: requirement.ENTITY_NAME, Type: "By", TypeAttr: "RelationId", Args: []string{guarantee1_id}})
	checkQuery(t, stub, "query", []string{string(requestStr)}, requirement1)
}
