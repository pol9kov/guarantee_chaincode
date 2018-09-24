package com

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

var Logger = shim.NewLogger("myChaincode")

func InfoLogMsg(message string) {
	Logger.Info("Message: ", FPath.getPath()+message)
}

func DebugLogMsg(message string) {
	Logger.Debug("Message: ", FPath.getPath()+message)
}

func ErrorLogMsg(error error, status int32, message string) {
	Logger.Error("Code: ", strconv.Itoa(int(status)),
		"; Message: ", message,
		"; Error: ", error.Error())
}
