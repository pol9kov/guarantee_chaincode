package guarantees

import (
	"encoding/xml"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"guarantees/platform/com"
	"guarantees/platform/data"
	"testing"
)

func createEntity(t *testing.T, stub *shim.MockStub, entity data.Entity, entityXml string, entityOutXml string, entityId string) {
	requestStr, _ := xml.Marshal(com.Request{EntityName: entity.GetEntityName(), Args: []string{entityXml}, Type: "OutOfRelation"})
	checkInvoke(t, stub, "create", []string{string(requestStr)})

	requestStr, _ = xml.Marshal(com.Request{EntityName: entity.GetEntityName(), Type: "By", TypeAttr: "Id", Args: []string{entityId}})
	checkQuery(t, stub, "query", []string{string(requestStr)}, entityOutXml)
}

func createEntityFromRelationId(t *testing.T, stub *shim.MockStub, entity data.Entity, entityXml, relationId, entityOutXml, entityId string) {
	requestStr, _ := xml.Marshal(com.Request{EntityName: entity.GetEntityName(), Args: []string{entityXml, relationId}, Type: "FromRelationId"})
	checkInvoke(t, stub, "create", []string{string(requestStr)})

	requestStr, _ = xml.Marshal(com.Request{EntityName: entity.GetEntityName(), Type: "By", TypeAttr: "Id", Args: []string{entityId}})
	checkQuery(t, stub, "query", []string{string(requestStr)}, entityOutXml)
}
