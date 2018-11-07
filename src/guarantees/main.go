package guarantees

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"platform/sc"
)

// The main function is for instantiate chaincode
func main() {

	logger := shim.NewLogger("mainLogger")

	//test.GuaranteesTest()

	// Put a new Smart Contract
	smartContract := new(sc.SmartContract)
	err := shim.Start(smartContract)
	if err != nil {
		logger.Error("Error creating new Smart Contract: ", err)
	}
}
