package sc

import (
	"encoding/xml"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"guarantees/guarantees/data/additional/gtype"
	"guarantees/guarantees/data/primary/statement"
	"guarantees/platform/com"
	"guarantees/platform/data"
)

//todo change this trash
func QueryStatementWithGType(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	element := com.FPath.Path.PushBack("s.queryStatementWithGType")
	defer com.FPath.Path.Remove(element)

	if len(args) != 1 {
		return com.IncorrectNumberOfArgsError(args, 1)
	}

	gtypeId := args[0]

	gtypeObject := gtype.GType{Id: gtypeId}
	shimResponse := data.QueryById(&gtypeObject, APIstub)
	if shimResponse.Status >= com.ERRORTHRESHOLD {
		return shimResponse
	}
	response := com.Response{}
	err := xml.Unmarshal(shimResponse.Payload, &response)
	if err != nil {
		return com.UnmarshalError(err, string(shimResponse.Payload))
	}
	err = xml.Unmarshal(response.Payload, &gtypeObject)
	if err != nil {
		return com.UnmarshalError(err, string(response.Payload))
	}

	statementObject := statement.Statement{}
	statementObject.StatementSigned.GType = gtypeObject

	return com.SuccessPayloadResponse(&statementObject)
}
