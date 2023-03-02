package json

import "net/http"

func SendJSON(w http.ResponseWriter, responseBytes []byte, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(responseBytes)
}
