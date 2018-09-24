package sc

import (
	"encoding/xml"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"guarantees/com"
	"guarantees/data"
	"guarantees/login/user"
)

type SmartContract struct {
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) peer.Response {
	return com.SuccessMessageResponse("Starting Guarantee chaincode!")
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) peer.Response {
	element := com.FPath.Path.PushBack("Invoke")
	defer com.FPath.Path.Remove(element)

	function, invokeArgs := APIstub.GetFunctionAndParameters()

	com.DebugLogMsg("Invoke of " + function)

	if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "queryUser" {
		return s.queryUser(APIstub)
	}

	switch function {
	case "setDebugLogLevel":
		com.Logger.SetLevel(shim.LogDebug)
		return com.SuccessMessageResponse("Debug log level was set.")
	case "setInfoLogLevel":
		com.Logger.SetLevel(shim.LogInfo)
		return com.SuccessMessageResponse("Info log level was set.")
	case "setNoticeLogLevel":
		com.Logger.SetLevel(shim.LogNotice)
		return com.SuccessMessageResponse("Notice log level was set.")
	case "setWarningLogLevel":
		com.Logger.SetLevel(shim.LogWarning)
		return com.SuccessMessageResponse("Warning log level was set.")
	case "setErrorLogLevel":
		com.Logger.SetLevel(shim.LogError)
		return com.SuccessMessageResponse("Error log level was set.")
	case "setCriticalLogLevel":
		com.Logger.SetLevel(shim.LogCritical)
		return com.SuccessMessageResponse("Critical log level was set.")
	}

	if len(invokeArgs) != 1 {
		return com.IncorrectNumberOfArgsError(invokeArgs, 1)
	}
	com.DebugLogMsg("Request: " + invokeArgs[0])

	// Get mspID
	//mspID, err := cid.GetMSPID(APIstub)
	//if err != nil {
	//	return com.GetMSPIDError()
	//}

	//Get user's role
	//role, ok, err := cid.GetAttributeValue(APIstub, user.ROLE)
	//if err != nil {
	//	// There was an error trying to retrieve the attribute
	//	return com.GetAttributeValueError(user.ROLE)
	//}
	//if !ok {
	//	// The client identity does not possess the attribute
	//	return com.GetAttributeValueNotOkError(user.ROLE)
	//}
	//
	role := user.ROLE_ADMIN

	request := com.Request{}
	err := xml.Unmarshal([]byte(invokeArgs[0]), &request)
	if err != nil {
		return com.UnmarshalError(err, invokeArgs[0])
	} else {
		com.SuccessMessageResponse("Request successfully unmarshaled!")
	}
	com.InfoLogMsg("Request function: " + function + request.EntityName + request.Type + request.TypeAttr)

	entity, response := data.GetEntityByName(request.EntityName)
	if response.Status >= com.ERRORTHRESHOLD {
		return response
	}

	args := request.Args

	if len(request.Channels) > 0 {
		requestWithoutChannels := request
		requestWithoutChannels.Channels = []string{}
		requestWithoutChannelsAsBytes, err := xml.Marshal(requestWithoutChannels)
		if err != nil {
			return com.MarshalError(err)
		}
		var entities []data.Entity
		for _, channelName := range request.Channels {
			response := APIstub.InvokeChaincode("guarantees_cc", [][]byte{requestWithoutChannelsAsBytes}, channelName)
			if response.Status >= com.ERRORTHRESHOLD {
				com.ErrorLogMsg(nil, response.Status, "Error on invoke channel "+channelName)
				return response
			}
			channelEntities, response := data.XmlBytesToEntitiesArr(response.Payload)
			if response.Status >= com.ERRORTHRESHOLD {
				return response
			}
			entities = append(entities, channelEntities...)
		}

		return com.SuccessPayloadResponse(data.EntitiesToOut(entity, entities))
	}

	if function == "create" {
		if request.Type == "OutOfRelation" {
			if len(args) != 1 {
				return com.IncorrectNumberOfArgsError(args, 1)
			}
			err = xml.Unmarshal([]byte(args[0]), entity)
			if err != nil {
				return com.UnmarshalError(err, args[0])
			}

			return data.Put(entity, APIstub)
		}
		if request.Type == "FromRelationId" {
			if len(args) != 2 {
				return com.IncorrectNumberOfArgsError(args, 2)
			}
			err = xml.Unmarshal([]byte(args[0]), entity)
			if err != nil {
				return com.UnmarshalError(err, args[0])
			}

			response := data.SetIdByRelationId(entity, APIstub, args[1])
			if response.Status >= com.ERRORTHRESHOLD {
				return response
			}

			return data.Put(entity, APIstub)

		}
	}

	if function == "edit" {
		if request.Type == "AllFields" {
			if len(args) != 1 {
				return com.IncorrectNumberOfArgsError(args, 1)
			}

			err = xml.Unmarshal([]byte(args[0]), entity)
			if err != nil {
				return com.UnmarshalError(err, args[0])
			}
			return data.EditAll(entity, APIstub)
		}
		if request.Type == "Fields" {
			if request.TypeAttr == "ById" {
				if len(args) != 1 {
					return com.IncorrectNumberOfArgsError(args, 1)
				}
				entity.SetId(args[0])

				if len(request.FieldPaths) != len(request.FieldValues) {
					return com.IncorectNumberOfFieldsError(len(request.FieldPaths), len(request.FieldValues))
				}
				return data.EditFieldsById(entity, APIstub, request.FieldPaths, request.FieldValues)
			}
			if request.TypeAttr == "ByRelationId" {
				if len(args) != 1 {
					return com.IncorrectNumberOfArgsError(args, 1)
				}
				data.SetIdByRelationId(entity, APIstub, args[0])

				if len(request.FieldPaths) != len(request.FieldValues) {
					return com.IncorectNumberOfFieldsError(len(request.FieldPaths), len(request.FieldValues))
				}
				return data.EditFieldsById(entity, APIstub, request.FieldPaths, request.FieldValues)
			}
		}
	}

	if function == "query" {
		if request.Type == "All" {
			return data.QueryAll(entity, APIstub)
		}
		if request.Type == "By" {
			switch request.TypeAttr {
			case "Id":
				if len(args) != 1 {
					return com.IncorrectNumberOfArgsError(args, 1)
				}
				entity.SetId(args[0])
				return data.QueryById(entity, APIstub)
			case "Key":
				if len(args) != 1 {
					return com.IncorrectNumberOfArgsError(args, 1)
				}
				return data.QueryByKey(entity, APIstub, args[0])
			case "Index":
				com.DebugLogMsg(com.ConcatArrStr(request.FieldPaths[0].FieldPath)) //todo remove
				return data.QueryByFieldPathAndFieldValue(entity, APIstub, request.FieldPaths[0].FieldPath, request.FieldValues[0])
			case "RelationId":
				if len(args) != 1 {
					return com.IncorrectNumberOfArgsError(args, 1)
				}
				data.SetIdByRelationId(entity, APIstub, args[0])

				return data.QueryById(entity, APIstub)
			}
		}
	}

	function = function + request.EntityName + request.Type + request.TypeAttr

	if function == "createGuaranteeByStatementId" {
		return s.createGuaranteeByStatementId(APIstub, args)
	}

	if function == "queryStatementWithGType" {
		return s.queryStatementWithGType(APIstub, args)
	}

	return com.InvalidFunctionNameForRoleError(function, role)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) peer.Response {
	return com.SuccessMessageResponse("Starting Guarantee chaincode!")
}
