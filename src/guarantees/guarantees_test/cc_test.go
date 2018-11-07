package guarantees

import (
	"encoding/xml"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"guarantees/platform/com"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"unicode"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}

	checkInvoke(t, stub, "setDebugLogLevel", []string{})
	//checkInvoke(t, stub, "setInfoLogLevel", []string{})
}

func checkState(t *testing.T, stub *shim.MockStub, name string, value string) {
	bytes := stub.State[name]
	if bytes == nil {
		fmt.Println("State", name, "failed to get value")
		t.FailNow()
	}
	if string(bytes) != value {
		fmt.Println("State value", name, "was not", value, "as expected")
		t.FailNow()
	}
}

func checkQuery(t *testing.T, stub *shim.MockStub, functionName string, args []string, value string) {

	argsAsBytes := [][]byte{[]byte(functionName)}
	for _, arg := range args {
		argsAsBytes = append(argsAsBytes, []byte(arg))
	}

	response := stub.MockInvoke("1", argsAsBytes)
	if response.Status != shim.OK {
		fmt.Println("\nQuery\n", functionName, "\nfailed\n", string(response.Message))
		t.FailNow()
	}
	if response.Payload == nil {
		fmt.Println("\nQuery\n", functionName, "\nfailed to get value")
		t.FailNow()
	}

	res := com.Response{}
	err := xml.Unmarshal(response.Payload, &res)
	if err != nil {
		fmt.Println("\nError on unmarshal response\n error : ", err.Error(), "\n response: ", string(response.Payload))
		t.FailNow()
	}

	removeSpacesValue := removeSpaces(value)
	removeSpacesPayload := removeSpaces(string(res.Payload))
	if removeSpaces(removeSpacesPayload) != removeSpacesValue {
		fmt.Println("\n"+functionName+" value \n", removeSpacesPayload, "\n was not\n", removeSpacesValue, "\n as expected ")
		t.FailNow()
	} else {
		fmt.Println("\nQueried ", functionName, " successfully!")
	}
}

func checkInvoke(t *testing.T, stub *shim.MockStub, functionName string, args []string) {
	argsAsBytes := [][]byte{[]byte(functionName)}
	for _, arg := range args {
		argsAsBytes = append(argsAsBytes, []byte(arg))
	}

	res := stub.MockInvoke("1", argsAsBytes)
	if res.Status != shim.OK {
		fmt.Println("\nInvoke\n", args, "\n failed\n", string(res.Message))
		t.FailNow()
	} else {
		fmt.Println("\nInvoked ", functionName, " successfully!")
	}
}

func readFile(fileName string) string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Print(err)
	}

	fileContent, err := ioutil.ReadFile(dir + "/testdata/entities/" + fileName)
	if err != nil {
		fmt.Print(err)
	}

	return string(fileContent)
}

func removeSpaces(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
