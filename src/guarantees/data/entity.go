package data

import (
	"encoding/xml"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"guarantees/com"
	"guarantees/data/guarantee/additional/bankpars"
	"guarantees/data/guarantee/additional/gtype"
	"guarantees/data/guarantee/additional/guarantor"
	"guarantees/data/guarantee/additional/organization"
	"guarantees/data/guarantee/primary/guarantee"
	"guarantees/data/guarantee/primary/requirement"
	"guarantees/data/guarantee/primary/rrequirement"
	"guarantees/data/guarantee/primary/statement"
	"guarantees/funcs"
	"reflect"
)

type Entity interface {
	// Keys types and attributes
	GetKeyObjectType() string
	GetIndexes() [][]string

	// Standard getters/setters
	GetEntityName() string
	GetTagName() string
	SetRelationTxId(string)
	GetRelationTxId() string
	SetId(string)
	GetId() string
	SetKey(string)
	GetKey() string
	SetMSPId(string)
	GetMSPId() string
	CanBeChangedOn(interface{}) bool

	// Transform entity to entityOut
	ToOut() interface{}
}

func getTxIdByRelationId(APIstub shim.ChaincodeStubInterface, relationId string) (string, peer.Response) {
	element := com.FPath.Path.PushBack("getTxIdByRelationId")
	defer com.FPath.Path.Remove(element)

	statementKey, response := funcs.IdToKey(relationId)
	if response.Status >= com.ERRORTHRESHOLD {
		return "", response
	}
	_, relationTxIdArr, err := APIstub.SplitCompositeKey(statementKey)
	if err != nil || len(relationTxIdArr) != 1 {
		return "", com.SplitCompositeKeyError(err, statementKey)
	}

	return relationTxIdArr[0], com.SuccessMessageResponse("Tx id was gotten from relation id.")
}

func SetIdByRelationId(entity Entity, APIstub shim.ChaincodeStubInterface, relationId string) peer.Response {
	element := com.FPath.Path.PushBack("SetIdByRelationId")
	defer com.FPath.Path.Remove(element)

	txId, response := getTxIdByRelationId(APIstub, relationId)
	if response.Status >= com.ERRORTHRESHOLD {
		return response
	}
	entity.SetRelationTxId(txId)

	key, err := APIstub.CreateCompositeKey(entity.GetKeyObjectType(), []string{txId})
	if err != nil {
		return com.CreateCompositeKeyError(err, entity.GetKeyObjectType(), []string{txId})
	}

	id := funcs.KeyToId(key)
	entity.SetId(id)

	return com.SuccessMessageResponse("Id was set by relation Id.")
}

func editFieldValueByFieldPath(entity Entity, fieldPath []string, fieldValue string) (Entity, peer.Response) {
	element := com.FPath.Path.PushBack("editFieldValueByFieldPath")
	defer com.FPath.Path.Remove(element)

	field := reflect.ValueOf(entity).Elem()

	for _, fieldName := range fieldPath {
		if field.Type().Kind() == reflect.Array || field.Type().Kind() == reflect.Slice {
			break
		} else {
			field = field.FieldByName(fieldName)
		}
	}

	switch field.Type().Kind() {
	case reflect.Array, reflect.Slice:
		len := len(fieldPath)
		if len > 3 {
			if field.Len() > 0 {
				com.DebugLogMsg("Go throw array ny name" + fieldPath[len-3])
				com.DebugLogMsg("Try to find " + fieldPath[len-2] + " and change " + fieldPath[len-1])
				for i := 0; i < field.Len(); i++ {
					com.DebugLogMsg("Array elem: " + field.Index(i).FieldByName(fieldPath[len-3]).String())
					if field.Index(i).FieldByName(fieldPath[len-3]).String() == fieldPath[len-2] { // Name, PlanIssueDate
						field.Index(i).FieldByName(fieldPath[len-1]).SetString(fieldValue) // Value
					}
				}
			}
		} else {
			return nil, com.IncorrectArrayParseArgsError()
		}
	case reflect.String:
		field.SetString(fieldValue)
	case reflect.Struct:
		entity, response := GetEntityByName(field.Type().Name())
		if response.Status >= com.ERRORTHRESHOLD {
			return nil, response
		}
		err := xml.Unmarshal([]byte(fieldValue), entity)
		if err != nil {
			return nil, com.UnmarshalError(err, fieldValue)
		}
		field.Set(reflect.ValueOf(entity).Elem())
	}

	return entity, com.SuccessMessageResponse("FieldPath was edited.")
}

func IsStateExist(APIstub shim.ChaincodeStubInterface, key string) bool {
	bytes, err := APIstub.GetState(key)
	if err == nil && len(bytes) > 0 {
		return true
	} else {
		return false
	}
}

func createKey(entity Entity, APIstub shim.ChaincodeStubInterface) peer.Response {
	element := com.FPath.Path.PushBack("createKey")
	defer com.FPath.Path.Remove(element)

	var attributes []string
	if entity.GetRelationTxId() != "" {
		attributes = []string{entity.GetRelationTxId()}
	} else {
		attributes = []string{APIstub.GetTxID()}
		entity.SetRelationTxId(APIstub.GetTxID())
	}
	objectType := entity.GetKeyObjectType()

	key, err := APIstub.CreateCompositeKey(objectType, attributes)
	if err != nil {
		return com.CreateCompositeKeyError(err, objectType, attributes)
	}

	entity.SetKey(key)

	return com.SuccessMessageResponse("Key was created.")
}

func GetEntityByName(name string) (Entity, peer.Response) {
	element := com.FPath.Path.PushBack("GetEntityByName")
	defer com.FPath.Path.Remove(element)

	var entity Entity

	switch name {
	case guarantor.ENTITY_NAME:
		entity = (Entity)(&guarantor.Guarantor{})
		break
	case organization.ENTITY_NAME:
		entity = (Entity)(&organization.Organization{})
		break
	case gtype.ENTITY_NAME:
		entity = (Entity)(&gtype.GType{})
		break
	case bankpars.ENTITY_NAME:
		entity = (Entity)(&bankpars.BankPars{})
		break
	case statement.ENTITY_NAME:
		entity = (Entity)(&statement.Statement{})
		break
	case guarantee.ENTITY_NAME:
		entity = (Entity)(&guarantee.Guarantee{})
		break
	case requirement.ENTITY_NAME:
		entity = (Entity)(&requirement.Requirement{})
		break
	case rrequirement.ENTITY_NAME:
		entity = (Entity)(&rrequirement.RRequirement{})
		break
	default:
		return nil, com.NoSuchCaseOfEntityError(name)
	}

	return entity, com.SuccessMessageResponse(name + " instance was created.")

}

func XmlBytesToEntitiesArr(xmlBytes []byte) ([]Entity, peer.Response) {
	type Entities struct {
		XMLName xml.Name //`xml:",any"`

		Entities []Entity `xml:",any"`
	}

	entities := Entities{}
	err := xml.Unmarshal(xmlBytes, &entities)
	if err != nil {
		return nil, com.UnmarshalError(err, string(xmlBytes))
	}

	return entities.Entities, com.SuccessMessageResponse("XML was unmarshaled in array of entities!")
}

func EntitiesToOut(entity Entity, entities []Entity) interface{} {
	element := com.FPath.Path.PushBack("EntitiesToOut")
	defer com.FPath.Path.Remove(element)

	type EntityArr struct {
		XMLName     xml.Name
		EntitiesOut []interface{}
	}

	entitiesOut := make([]interface{}, len(entities))
	entityArr := EntityArr{
		XMLName:     xml.Name{Local: entity.GetTagName()}, //+"s" },
		EntitiesOut: entitiesOut}

	for i, entity := range entities {
		entityArr.EntitiesOut[i] = entity.ToOut()
	}

	return entityArr
}

func iteratorToArray(iterator shim.StateQueryIteratorInterface) ([]string, peer.Response) {
	element := com.FPath.Path.PushBack("iteratorToArray")
	defer com.FPath.Path.Remove(element)

	defer iterator.Close()

	var result []string

	for iterator.HasNext() {
		queryResponse, err := iterator.Next()
		if err != nil {
			return nil, com.IteratorNextError(err)
		}

		result = append(result, string(queryResponse.Value))
	}

	return result, com.SuccessMessageResponse("Iterator was converted to array")
}

func getIndexObjectType(entity Entity) string {
	return entity.GetKeyObjectType() + "_INDEX"
}
