package json

import "net/http"

func SendJSON(w http.ResponseWriter, responseBytes []byte) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBytes)
}
