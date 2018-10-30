package data

import (
	"encoding/xml"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"platform/com"
	"platform/funcs"
	"strconv"
)

func QueryAll(entity Entity, APIstub shim.ChaincodeStubInterface) peer.Response {
	element := com.FPath.Path.PushBack(entity.GetEntityName() + ".QueryAll")
	defer com.FPath.Path.Remove(element)

	resultsIterator, err := APIstub.GetStateByPartialCompositeKey(entity.GetKeyObjectType(), []string{})
	if err != nil {
		return com.GetStateError(err, entity.GetKeyObjectType())
	}

	bytesArray, response := iteratorToArray(resultsIterator)
	if response.Status >= com.ERRORTHRESHOLD {
		return response
	}

	entities, response := valuesToEntities(entity, bytesArray)
	if response.Status >= com.ERRORTHRESHOLD {
		return response
	}

	return com.SuccessPayloadResponse(EntitiesToOut(entity, entities))
}

func valuesToEntities(entity Entity, values []string) ([]Entity, peer.Response) {
	element := com.FPath.Path.PushBack("valuesToEntities")
	defer com.FPath.Path.Remove(element)

	var result []Entity
	for i, value := range values {
		com.DebugLogMsg(strconv.Itoa(i) + " element: " + string(value))
		entityNew, response := GetEntityByName(entity.GetEntityName())
		if response.Status >= com.ERRORTHRESHOLD {
			return nil, response
		}

		err := xml.Unmarshal([]byte(value), entityNew)
		if err != nil {
			return nil, com.UnmarshalError(err, value)
		}
		result = append(result, entityNew)
	}
	return result, com.SuccessMessageResponse("Values were unmarshaled to " + entity.GetEntityName() + "s succesfully.")
}

func QueryById(entity Entity, APIstub shim.ChaincodeStubInterface) peer.Response {
	element := com.FPath.Path.PushBack(entity.GetEntityName() + ".QueryById")
	defer com.FPath.Path.Remove(element)

	entity, response := queryById(entity, APIstub)
	if response.Status >= com.ERRORTHRESHOLD {
		return response
	}

	return com.SuccessPayloadResponse(entity.ToOut())
}

func queryById(entity Entity, APIstub shim.ChaincodeStubInterface) (Entity, peer.Response) {
	element := com.FPath.Path.PushBack("queryById")
	defer com.FPath.Path.Remove(element)

	key, response := funcs.IdToKey(entity.GetId())
	if response.Status >= com.ERRORTHRESHOLD {
		return nil, response
	}

	entity, response = queryByKey(entity, APIstub, key)
	if response.Status >= com.ERRORTHRESHOLD {
		return nil, response
	}

	return entity, com.SuccessMessageResponse(entity.GetEntityName() + " was queried by id.")
}

func QueryByKey(entity Entity, APIstub shim.ChaincodeStubInterface, key string) peer.Response {
	element := com.FPath.Path.PushBack(entity.GetEntityName() + ".QueryByKey")
	defer com.FPath.Path.Remove(element)

	entity, response := queryByKey(entity, APIstub, key)
	if response.Status >= com.ERRORTHRESHOLD {
		return response
	}

	return com.SuccessPayloadResponse(entity.ToOut())
}

func queryByKey(entity Entity, APIstub shim.ChaincodeStubInterface, key string) (Entity, peer.Response) {
	element := com.FPath.Path.PushBack("queryByKey")
	defer com.FPath.Path.Remove(element)

	entityAsBytes, err := APIstub.GetState(key)
	if err != nil {
		return nil, com.GetStateError(err, key)
	}
	com.DebugLogMsg("State was gotten. Key: " + key + "; Value: " + string(entityAsBytes))
	err = xml.Unmarshal(entityAsBytes, entity)
	if err != nil {
		return nil, com.UnmarshalError(err, string(entityAsBytes))
	}
	return entity, com.SuccessMessageResponse(entity.GetEntityName() + " was queried by key")

}

func QueryByKeys(entity Entity, APIstub shim.ChaincodeStubInterface, keys []string) peer.Response {
	element := com.FPath.Path.PushBack(entity.GetEntityName() + ".queryByKeys")
	defer com.FPath.Path.Remove(element)

	entities, response := queryByKeys(entity, APIstub, keys)
	if response.Status >= com.ERRORTHRESHOLD {
		return response
	}

	return com.SuccessPayloadResponse(EntitiesToOut(entity, entities))
}

func queryByKeys(entity Entity, APIstub shim.ChaincodeStubInterface, keys []string) ([]Entity, peer.Response) {
	element := com.FPath.Path.PushBack("queryByKeys")
	defer com.FPath.Path.Remove(element)

	var result []Entity

	for _, key := range keys {

		entityNew, response := GetEntityByName(entity.GetEntityName())
		if response.Status >= com.ERRORTHRESHOLD {
			return nil, response
		}

		entityNew, response = queryByKey(entityNew, APIstub, key)
		if response.Status >= com.ERRORTHRESHOLD {
			return nil, response
		}

		result = append(result, entityNew)
	}

	return result, com.SuccessMessageResponse(entity.GetEntityName() + "s were queried by keys")
}

func QueryByFieldPathAndFieldValue(entity Entity, APIstub shim.ChaincodeStubInterface, fieldPath []string, fieldValue string) peer.Response {
	element := com.FPath.Path.PushBack("QueryByFieldPathAndFieldValue")
	defer com.FPath.Path.Remove(element)

	com.DebugLogMsg(com.ConcatArrStr(fieldPath)) //todo remove
	entities, response := queryByFieldPathAndFieldValue(entity, APIstub, fieldPath, fieldValue)
	if response.Status >= com.ERRORTHRESHOLD {
		return response
	}

	return com.SuccessPayloadResponse(EntitiesToOut(entity, entities))
}

func queryByFieldPathAndFieldValue(entity Entity, APIstub shim.ChaincodeStubInterface, fieldPath []string, fieldValue string) ([]Entity, peer.Response) {
	element := com.FPath.Path.PushBack("queryByFieldPathAndFieldValue")
	defer com.FPath.Path.Remove(element)

	//entity, response := editFieldValueByFieldPath(entity, fieldPath, fieldValue)
	//if response.Status >= com.ERRORTHRESHOLD {
	//	return nil, response
	//}

	indexField := append(fieldPath, fieldValue)
	com.DebugLogMsg(com.ConcatArrStr(indexField)) //todo remove

	return queryByIndexAttr(entity, APIstub, indexField)
}

func queryByIndexAttr(entity Entity, APIstub shim.ChaincodeStubInterface, indexAttr []string) ([]Entity, peer.Response) {
	element := com.FPath.Path.PushBack("queryByIndexAttr")
	defer com.FPath.Path.Remove(element)

	iterator, response := queryKeysByIndexAttr(entity, APIstub, indexAttr)
	if response.Status >= com.ERRORTHRESHOLD {
		return nil, response
	}

	array, response := queryByKeys(entity, APIstub, iterator)
	if response.Status >= com.ERRORTHRESHOLD {
		return nil, response
	}

	return array, com.SuccessMessageResponse(entity.GetEntityName() + " were gotten")
}

func queryKeysByIndexAttr(entity Entity, APIstub shim.ChaincodeStubInterface, indexAttr []string) ([]string, peer.Response) {
	element := com.FPath.Path.PushBack("queryKeysByIndexAttr")
	defer com.FPath.Path.Remove(element)

	com.DebugLogMsg(com.ConcatArrStr(indexAttr)) //todo remove
	iterator, err := APIstub.GetStateByPartialCompositeKey(getIndexObjectType(entity), indexAttr)
	if err != nil {
		return nil, com.GetStateByPartialCompositeKeyError(err, getIndexObjectType(entity), indexAttr)
	}

	array, response := iteratorToArray(iterator)
	if response.Status >= com.ERRORTHRESHOLD {
		return nil, response
	}

	return array, com.SuccessMessageResponse(entity.GetEntityName() + "'s keys was gotten")
}

func QuerySomethingById(entity Entity, APIstub shim.ChaincodeStubInterface, getSomethingFromEntity func(interface{}) peer.Response) peer.Response {
	element := com.FPath.Path.PushBack(entity.GetEntityName() + "QuerySomethingById")
	defer com.FPath.Path.Remove(element)

	entity, response := queryById(entity, APIstub)
	if response.Status >= com.ERRORTHRESHOLD {
		return response
	}

	return getSomethingFromEntity(entity)
}
