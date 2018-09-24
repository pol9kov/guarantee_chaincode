package funcs

import (
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
)

const (
	MSPId = "Org1MSP"
)

func GetMSPID(stub cid.ChaincodeStubInterface) (string, error) {
	//return cid.GetMSPID(stub)
	return MSPId, nil
}
