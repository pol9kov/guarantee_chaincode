package com

import (
	"container/list"
	"encoding/xml"
	"github.com/hyperledger/fabric/protos/peer"
)

const (
	// OK_MESSAGE constant - status code less than 400, endorser will endorse it.
	// OK_MESSAGE means init or invoke successfully.
	OK         = 200
	OK_MESSAGE = 0
	OK_PAYLOAD = 0

	// ERRORTHRESHOLD constant - status code greater than or equal to 400 will be considered an error and rejected by endorser.
	ERRORTHRESHOLD = 400

	// ERROR constant - default error value
	ERROR = 500
)

var FPath_obj = FuncPath{Path: list.New()}
var FPath *FuncPath = &FPath_obj

type FuncPath struct {
	Path *list.List `protobuf:"bytes,1,opt,name=path" json:"path" xml:"path"`
}

type Response struct {
	XMLName xml.Name `xml:"payload"`

	// A status code that should follow the HTTP status codes.
	Status int32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty" xml:"status"`
	// A error message that may be kept.
	Error string `protobuf:"varint,1,opt,name=error" json:"error,omitempty" xml:"error,omitempty"`
	// A message associated with the response code.
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty" xml:"message"`
	// A payload that can be used to include metadata with this response.
	Payload []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty" xml:"payload"`
}

func (FPath FuncPath) getPath() string {
	if FPath.Path == nil {
		return ""
	}

	result := " '"

	for e := FPath.Path.Front(); e != nil; e = e.Next() {
		result += "/" + (e.Value).(string)
	}

	result += "': "

	return result
}

// todo: Кроме самописного сообщения выводить также ошибку err, которую получили. Она может помочь.

func ErrorMessageResponse(error error, status int32, message string) peer.Response {
	ErrorLogMsg(error, status, message)

	response := Response{
		Status:  status,
		Error:   error.Error(),
		Message: message,
	}

	payload, err := xml.Marshal(response)
	if err != nil {
		ResponseMarshalError(err)
	}

	return peer.Response{
		Status:  ERROR,
		Message: message,
		Payload: payload,
	}
}

func SuccessMessageResponse(message string) peer.Response {
	DebugLogMsg(message)

	response := Response{
		Status:  OK_MESSAGE,
		Message: message,
	}

	payload, err := xml.Marshal(response)
	if err != nil {
		ResponseMarshalError(err)
	}

	return peer.Response{
		Status:  OK,
		Message: message,
		Payload: payload,
	}
}

func SuccessPayloadResponse(data interface{}) peer.Response {

	dataXML, err := xml.Marshal(data)
	if err != nil {
		ResponseMarshalError(err)
	}

	resp := Response{
		Status:  OK_PAYLOAD,
		Payload: dataXML,
	}

	payload, err := xml.Marshal(resp)
	if err != nil {
		ResponseMarshalError(err)
	}

	DebugLogMsg(string(payload))

	return peer.Response{
		Status:  OK,
		Message: "Successfully got data.",
		Payload: payload,
	}
}
