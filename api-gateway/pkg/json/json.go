package json

import (
	"encoding/json"
)

func Marshal(response any) []byte {
	requestBytes, err := json.Marshal(response)

	if err != nil {
		return nil
	}

	return requestBytes
}
