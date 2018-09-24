package sc

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"guarantees/com"
	"guarantees/login/user"
)


func (s *SmartContract) queryUser(APIstub shim.ChaincodeStubInterface) (peer.Response) {
	element := com.FPath.Path.PushBack("s.queryUser")
	defer com.FPath.Path.Remove(element)

	user, response := user.GetUser(APIstub)
	if response.Status >= com.ERRORTHRESHOLD {return response}
	
	return com.SuccessPayloadResponse(&user)
}
