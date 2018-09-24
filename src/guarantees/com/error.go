package com

import (
	"github.com/hyperledger/fabric/protos/peer"
	"strconv"
)

const (
	INVALID_FUNCTION_NAME_ERROR              = 501
	INCORRECT_NUMBER_OF_ARGS_ERROR           = 510
	MARSHAL_ERROR                            = 520
	RESPONSE_MARSHAL_ERROR                   = 521
	UNMARSHAL_ERROR                          = 530
	GET_STATE_ERROR                          = 540
	GET_STATE_BY_RANGE_ERROR                 = 541
	GET_STATE_BY_PARTIAL_COMPOSITE_KEY_ERROR = 542
	PUT_STATE_ERROR                          = 550
	PUT_STATE_ON_EXIST_KEY_ERROR             = 551
	EDIT_STATE_ON_NOT_EXIST_KEY_ERROR        = 552
	CREATE_COMPOSITE_KEY_ERROR               = 560
	SPLIT_COMPOSITE_KEY_ERROR                = 561
	ITERATOR_NEXT_ERROR                      = 570
	ENCRYPT_ERROR                            = 580
	DECRYPT_ERROR                            = 590
	DECODE_ERROR                             = 591
	GET_ATTRIBUTE_ERROR                      = 600
	NOT_POSSESS_ATTRIBUTE_ERROR              = 601
	GET_ATTRIBUTE_VALUE_ERROR                = 610
	GET_ATTRIBUTE_VALUE_NOT_OK_ERROR         = 611
	CREATE_KEY_ATTRIBUTES_ERROR              = 620
	CREATE_INDEX_ATTRIBUTES_ERROR            = 621
	GET_MSP_ID_ERROR                         = 630
	CREATE_ADDITIONAL_DATA_ERROR             = 640
	MANY_ORGS_FOR_ONE_MSPID_ERROR            = 650
	NO_ORG_FOR_MSPID_ERROR                   = 651
	INCORECT_ENTITY_NAME_ERROR               = 660
	INCORRECT_NUMBER_OF_FIELDS_ERROR         = 670
	NO_SUCH_CASE_OF_ENTITY_ERROR             = 680
	INCORRECT_ARRAY_PARSE_ARGS_ERROR         = 690
	REGEXP_VALIDATION_ERROR                  = 700
	REGEXP_MATCH_ERROR                       = 701
	ENTITY_VALIDATION_ERROR                  = 710

	STATEMENT_TYPE_DOES_NOT_EXIST_ERROR      = 1000
	GUARANTEE_STATEMENT_DOES_NOT_EXIST_ERROR = 1001
)

type CCError struct {
	error string `protobuf:"bytes,1,opt,name=error" json:"error" xml:"error"`
}

func (error CCError) Error() string {
	return error.error
}

func InvalidFunctionNameForRoleError(function string, role string) peer.Response {
	massage := "Invalid Smart Contract function name: '" + function + "' for role: '" + role + "'"
	return ErrorMessageResponse(CCError{"InvalidFunctionNameForRoleError"}, INVALID_FUNCTION_NAME_ERROR, massage)
}

func IncorrectNumberOfArgsError(args []string, narg int) peer.Response {
	massage := "Incorrect number of arguments. Expecting " + strconv.Itoa(narg)
	return ErrorMessageResponse(CCError{"IncorrectNumberOfArgsError"}, INCORRECT_NUMBER_OF_ARGS_ERROR, massage)
}

func MarshalError(err error) peer.Response {
	massage := "Error on Marshal data"
	return ErrorMessageResponse(err, MARSHAL_ERROR, massage)
}

func ResponseMarshalError(err error) peer.Response {
	massage := "Error on Marshal response"
	return ErrorMessageResponse(err, RESPONSE_MARSHAL_ERROR, massage)
}

func UnmarshalError(err error, data string) peer.Response {
	massage := "Error on Unmarshal data: " + data
	return ErrorMessageResponse(err, UNMARSHAL_ERROR, massage)
}

func GetStateError(err error, data string) peer.Response {
	massage := "Error on GetState: " + data
	return ErrorMessageResponse(err, GET_STATE_ERROR, massage)
}

func GetStateByRangeError(err error, startKey, endKey string) peer.Response {
	massage := "Error on GetStateByRange from startKey = " + startKey + "; to endKey = " + endKey
	return ErrorMessageResponse(err, GET_STATE_BY_RANGE_ERROR, massage)
}

func GetStateByPartialCompositeKeyError(err error, objectType string, keys []string) peer.Response {
	massage := "Error on GetStateByPartialCompositeKey for objectType = " + objectType + "; keys = " + ConcatArrStr(keys)
	return ErrorMessageResponse(err, GET_STATE_BY_PARTIAL_COMPOSITE_KEY_ERROR, massage)
}

func PutStateError(err error, key string, data []byte) peer.Response {
	str := "Key: " + key + "; data: " + string(data)
	massage := "Error on PutState: " + str
	return ErrorMessageResponse(err, PUT_STATE_ERROR, massage)
}

func PutStateOnExistKeyError(key string, data []byte) peer.Response {
	str := "Key: " + key + "; data: " + string(data)
	massage := "Error on PutState, the key already exist! " + str
	return ErrorMessageResponse(CCError{"PutStateOnExistKeyError"}, PUT_STATE_ON_EXIST_KEY_ERROR, massage)
}

func EditStateOnNotExistKeyError(key string, data []byte) peer.Response {
	str := "Key: " + key + "; data: " + string(data)
	massage := "Error on PutState, the key does not exist! " + str
	return ErrorMessageResponse(CCError{"EditStateOnNotExistKeyError"}, EDIT_STATE_ON_NOT_EXIST_KEY_ERROR, massage)
}

func CreateCompositeKeyError(err error, objectType string, attributes []string) peer.Response {
	data := "objectType: " + objectType + "; attributes: " + ConcatArrStr(attributes)
	massage := "Error on CreateCompositeKey: " + data
	return ErrorMessageResponse(err, CREATE_COMPOSITE_KEY_ERROR, massage)
}

func SplitCompositeKeyError(err error, compositeKey string) peer.Response {
	data := "compositeKey: " + compositeKey
	massage := "Error on SplitCompositeKey: " + data
	return ErrorMessageResponse(err, SPLIT_COMPOSITE_KEY_ERROR, massage)
}

func IteratorNextError(err error) peer.Response {
	massage := "Error on get next item from iterator."
	return ErrorMessageResponse(err, ITERATOR_NEXT_ERROR, massage)
}

func EncryptError(err error, data string) peer.Response {
	massage := "Error on encrypt: " + data
	return ErrorMessageResponse(err, ENCRYPT_ERROR, massage)
}

func DecryptError(err error, data string) peer.Response {
	massage := "Error on decrypt: " + data
	return ErrorMessageResponse(err, DECRYPT_ERROR, massage)
}

func DecodeError(err error, data string) peer.Response {
	massage := "Error on decode: " + data
	return ErrorMessageResponse(err, DECODE_ERROR, massage)
}

func GetAttributeError(err error, data string) peer.Response {
	massage := "Error on get attribute: " + data
	return ErrorMessageResponse(err, GET_ATTRIBUTE_ERROR, massage)
}

func NotPossessAttributeError(data string) peer.Response {
	massage := "Error on getting attribute. Not possess attribute: " + data
	return ErrorMessageResponse(CCError{"NotPossessAttributeError"}, NOT_POSSESS_ATTRIBUTE_ERROR, massage)
}

func GetAttributeValueError(err error, attrName string) peer.Response {
	massage := "There was an error trying to retrieve the attribute: " + attrName
	return ErrorMessageResponse(err, GET_ATTRIBUTE_VALUE_ERROR, massage)
}

func GetAttributeValueNotOkError(attrName string) peer.Response {
	massage := "The client identity does not possess the attribute: " + attrName
	return ErrorMessageResponse(CCError{"GetAttributeValueNotOkError"}, GET_ATTRIBUTE_VALUE_NOT_OK_ERROR, massage)
}

func StatementTypeDoesNotExistError(gtypeKey string) peer.Response {
	massage := "Statement's type does not exist! Type's key: " + gtypeKey
	return ErrorMessageResponse(CCError{"StatementTypeDoesNotExistError"}, STATEMENT_TYPE_DOES_NOT_EXIST_ERROR, massage)
}

func GuaranteeStatementDoesNotExistError(statementKey string) peer.Response {
	massage := "Guarantee's statement does not exist! Statement's key: " + statementKey
	return ErrorMessageResponse(CCError{"GuaranteeStatementDoesNotExistError"}, GUARANTEE_STATEMENT_DOES_NOT_EXIST_ERROR, massage)
}

func CreateKeyAttributesError(data string) peer.Response {
	massage := "Error on creating key's attributes! One of the fields is empty: " + data
	return ErrorMessageResponse(CCError{"CreateKeyAttributesError"}, CREATE_KEY_ATTRIBUTES_ERROR, massage)
}

func CreateIndexAttributesError(data string) peer.Response {
	massage := "Error on creating index's attributes! One of the fields is empty: " + data
	return ErrorMessageResponse(CCError{"CreateIndexAttributesError"}, CREATE_INDEX_ATTRIBUTES_ERROR, massage)
}

func CreateIndexesError() peer.Response {
	massage := "Error on creating indexes. Object's name doesn't exist. "
	return ErrorMessageResponse(CCError{"CreateIndexesError"}, CREATE_INDEX_ATTRIBUTES_ERROR, massage)
}

func CreateAdditionalDataError() peer.Response {
	massage := "Error on creating additional data! Lengths of additional data array and additional data object type array are not equal! "
	return ErrorMessageResponse(CCError{"CreateAdditionalDataError"}, CREATE_ADDITIONAL_DATA_ERROR, massage)
}

func GetMSPIDError(err error) peer.Response {
	massage := "Error on getting MSPId!"
	return ErrorMessageResponse(err, GET_MSP_ID_ERROR, massage)
}

func ManyOrgsForOneMSPIdError(mspId string) peer.Response {
	massage := "There are more than one organizations for one MSP id! mspId: " + mspId
	return ErrorMessageResponse(CCError{"ManyOrgsForOneMSPIdError"}, MANY_ORGS_FOR_ONE_MSPID_ERROR, massage)
}

func NoOrgForMSPIdError(mspId string) peer.Response {
	massage := "There are no organization for this MSP id! mspId: " + mspId
	return ErrorMessageResponse(CCError{"NoOrgForMSPIdError"}, NO_ORG_FOR_MSPID_ERROR, massage)
}

func IncorectEntityNameError(entityName string) peer.Response {
	massage := "Incorrect entity name in request! name: " + entityName
	return ErrorMessageResponse(CCError{"IncorectEntityNameError"}, INCORECT_ENTITY_NAME_ERROR, massage)
}

func IncorectNumberOfFieldsError(nFieldPaths, nFieldValues int) peer.Response {
	massage := "Incorrect number of fields values and fields paths. They are not equal. FieldPath paths: " +
		strconv.Itoa(nFieldPaths) + "; FieldPath values: " + strconv.Itoa(nFieldValues)
	return ErrorMessageResponse(CCError{"IncorrectNumberOfArgsError"}, INCORRECT_NUMBER_OF_FIELDS_ERROR, massage)
}

func NoSuchCaseOfEntityError(entityName string) peer.Response {
	massage := "No such case of entity. Entity name: " + entityName
	return ErrorMessageResponse(CCError{"NoSuchCaseOfEntityError"}, NO_SUCH_CASE_OF_ENTITY_ERROR, massage)
}

func IncorrectArrayParseArgsError() peer.Response {
	massage := "Incorrect array parse arguments."
	return ErrorMessageResponse(CCError{"NoSuchCaseOfEntityError"}, INCORRECT_ARRAY_PARSE_ARGS_ERROR, massage)
}

func RegExpValidationError(s, regexp string) peer.Response {
	massage := "Regular expression validation failed. String \"" + s +
		"\" does not contain any match of the regular expression pattern: \"" + regexp + "\""
	return ErrorMessageResponse(CCError{"RegExpValidationError"}, REGEXP_VALIDATION_ERROR, massage)
}

func RegExpMatchError(err error, s, regexp string) peer.Response {
	massage := "Regular expression validation failed. String \"" + s +
		"\" does not contain any match of the regular expression pattern: \"" + regexp + "\""
	return ErrorMessageResponse(err, REGEXP_MATCH_ERROR, massage)
}

func EntityValidationError() peer.Response {
	massage := "Cann't edit entity because no permission for editing some field."
	return ErrorMessageResponse(CCError{"EntityValidationError"}, ENTITY_VALIDATION_ERROR, massage)
}

func ConcatArrStr(arr []string) string {
	result := "["
	for index, element := range arr {
		result += "Index " + strconv.Itoa(index) + ", Value " + element + ";"
	}
	result += "]"

	return result
}
