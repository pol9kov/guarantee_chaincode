package guarantees

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"guarantees/guarantees/data/additional/organization"
	"guarantees/platform/sc"
	"testing"
)

var organizationObject = organization.Organization{}
var organization1 = readFile("organization.xml")
var organization1_id = "AE9SR0FOSVpBVElPTgAxAA=="

func CreateOrganization(t *testing.T, stub *shim.MockStub) {
	createEntity(t, stub, &organizationObject, organization1, organization1, organization1_id)
}

func InitOrganization(t *testing.T) *shim.MockStub {
	sc := new(sc.SmartContract)
	stub := shim.NewMockStub("guarantee_cc_test", sc)

	checkInit(t, stub, [][]byte{[]byte("init")})

	CreateOrganization(t, stub)

	return stub
}

func Test_Organization_Invoke(t *testing.T) {
	InitOrganization(t)
}
