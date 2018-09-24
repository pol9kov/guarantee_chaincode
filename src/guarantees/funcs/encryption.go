package funcs

import (
	"encoding/base64"
	"github.com/hyperledger/fabric/protos/peer"
	"guarantees/com"
)

func IdToKey(id string) (string, peer.Response) {
	element := com.FPath.Path.PushBack("IdToKey")
	defer com.FPath.Path.Remove(element)

	result, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		com.DecodeError(err, "Data: "+id)
	}

	return string(result), com.SuccessMessageResponse("String was decrypted successfully.")
}

func KeyToId(key string) (id string) {
	return base64.StdEncoding.EncodeToString([]byte(key))
}
