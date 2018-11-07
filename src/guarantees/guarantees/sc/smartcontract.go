package sc

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"guarantees/platform/com"
)

func CallBuisnessFunc(APIstub shim.ChaincodeStubInterface, function string, args []string, simulate string, role string) peer.Response {

	if function == "createGuaranteeByStatementId" {
		return CreateGuaranteeByStatementId(APIstub, args, simulate)
	}

	if function == "queryStatementWithGType" {
		return QueryStatementWithGType(APIstub, args)
	}

	return com.InvalidFunctionNameForRoleError(function, role)
}
