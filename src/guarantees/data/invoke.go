package data

import (
	"encoding/xml"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"guarantees/com"
	"guarantees/funcs"
	"reflect"
	"strconv"
)

func Put(entity Entity, APIstub shim.ChaincodeStubInterface) peer.Response {
	element := com.FPath.Path.PushBack(entity.GetEntityName() + ".Put")
	defer com.FPath.Path.Remove(element)

	response := putValue(entity, APIstub)
	if response.Status >= com.ERRORTHRESHOLD {
		return response
	}

	response = putIndexes(entity, APIstub)
	if response.Status >= com.ERRORTHRESHOLD {
		return response
	}

	com.InfoLogMsg(entity.GetEntityName() + " was created.")
	return com.SuccessPayloadResponse(entity.ToOut())
}

func putValue(entity Entity, APIstub shim.ChaincodeStubInterface) peer.Response {
	element := com.FPath.Path.PushBack("putValue")
	defer com.FPath.Path.Remove(element)

	response := createKey(entity, APIstub)
	if response.Status >= com.ERRORTHRESHOLD {
		return response
	}

	entity.SetId(funcs.KeyToId(entity.GetKey()))

	com.DebugLogMsg("Id was created. Id: " + entity.GetId())

	entityAsBytes, err := xml.Marshal(entity)
	if err != nil {
		return com.MarshalError(err)
	}

	if IsStateExist(APIstub, entity.GetKey()) {
		return com.PutStateOnExistKeyError(entity.GetKey(), entityAsBytes)
	}

	err = APIstub.PutState(entity.GetKey(), entityAsBytes)
	if err != nil {
		return com.PutStateError(err, entity.GetKey(), entityAsBytes)
	}
	com.DebugLogMsg("State was put. Key: " + entity.GetKey() + "; Value: " + string(entityAsBytes))

	return com.SuccessMessageResponse("State for " + entity.GetEntityName() + " was put.")
}

func putIndexes(entity Entity, APIstub shim.ChaincodeStubInterface) peer.Response {
	element := com.FPath.Path.PushBack("putIndexes")
	defer com.FPath.Path.Remove(element)

	for i, indexFieldPath := range entity.GetIndexes() {

		field := reflect.ValueOf(entity).Elem()

		for _, fieldName := range indexFieldPath {
			field = field.FieldByName(fieldName)
		}

		fieldValue := field.String()

		attributes := append(indexFieldPath, fieldValue)
		fullAttributes := append(attributes, entity.GetRelationTxId())
		indexKey, err := APIstub.CreateCompositeKey(getIndexObjectType(entity), fullAttributes)
		if err != nil {
			return com.CreateCompositeKeyError(err, getIndexObjectType(entity), fullAttributes)
		}

		err = APIstub.PutState(indexKey, []byte(entity.GetKey()))
		if err != nil {
			return com.PutStateError(err, indexKey, []byte(entity.GetKey()))
		}

		com.DebugLogMsg("Index key №" + strconv.Itoa(i) + ": " + indexKey)
	}

	return com.SuccessMessageResponse("Indexes for Guarantor was created.")
}

func EditAll(entity Entity, APIstub shim.ChaincodeStubInterface) peer.Response {
	element := com.FPath.Path.PushBack(entity.GetEntityName() + ".EditAll")
	defer com.FPath.Path.Remove(element)

	key, response := funcs.IdToKey(entity.GetId())
	if response.Status >= com.ERRORTHRESHOLD {
		return response
	}

	entity.SetKey(key)

	entityAsBytes, err := xml.Marshal(entity)
	if err != nil {
		return com.MarshalError(err)
	}

	if !IsStateExist(APIstub, entity.GetKey()) {
		return com.EditStateOnNotExistKeyError(entity.GetKey(), entityAsBytes)
	}

	bytes, _ := APIstub.GetState(key)
	oldEntity, response := GetEntityByName(entity.GetEntityName())
	if response.Status >= com.ERRORTHRESHOLD {
		return response
	}
	err = xml.Unmarshal(bytes, oldEntity)
	if err != nil {
		return com.UnmarshalError(err, string(bytes))
	}

	if !oldEntity.ChangeValidation(entity) {
		return com.EntityValidationError()
	}

	error := APIstub.PutState(entity.GetKey(), entityAsBytes)
	if error != nil {
		return com.PutStateError(err, entity.GetKey(), entityAsBytes)
	}
	com.DebugLogMsg("State was edited. Key: " + entity.GetKey() + "; Value: " + string(entityAsBytes))

	return com.SuccessMessageResponse(entity.GetEntityName() + " was edited.")
}

func EditFieldsById(entity Entity, APIstub shim.ChaincodeStubInterface, fieldPaths []com.FieldPath, fieldValues []string) peer.Response {
	element := com.FPath.Path.PushBack(entity.GetEntityName() + "EditFieldsById")
	defer com.FPath.Path.Remove(element)

	entity, response := queryById(entity, APIstub)
	if response.Status >= com.ERRORTHRESHOLD {
		return response
	}

	for i, fieldPath := range fieldPaths {
		fieldValue := fieldValues[i]
		entity, response = editFieldValueByFieldPath(entity, fieldPath.FieldPath, fieldValue)
		if response.Status >= com.ERRORTHRESHOLD {
			return response
		}

	}
	return EditAll(entity, APIstub)
}

func putAdditionalData(entity Entity, APIstub shim.ChaincodeStubInterface, additionalDataArr []string) peer.Response {
	element := com.FPath.Path.PushBack("putAdditionalData")
	defer com.FPath.Path.Remove(element)

	return com.SuccessMessageResponse("There is no additional data for " + entity.GetEntityName() + " for creation.")
}
