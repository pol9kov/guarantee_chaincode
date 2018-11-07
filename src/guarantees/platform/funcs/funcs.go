package funcs

import (
	"encoding/xml"
	"guarantees/platform/com"
)

func IsFieldsEqual(field1, field2 interface{}) bool {
	fieldBytes1, err := xml.Marshal(field1)
	if err != nil {
		com.DebugLogMsg("Cann't marshal field 1.")
		return false
	}
	fieldBytes2, err := xml.Marshal(field1)
	if err != nil {
		com.DebugLogMsg("Cann't marshal field 2.")
		return false
	}

	if string(fieldBytes1) != string(fieldBytes2) {
		com.DebugLogMsg("Field1 is not equal field2. Field1: " + string(fieldBytes1) + ". Field2: " + string(fieldBytes2))
		return false
	}

	return true
}
