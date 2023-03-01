package middlwares

import (
	"api-gateway/pkg/json"
	"errors"
	"net/http"

	"github.com/Levap123/utils/apperror"
)

func CheckErrorMiddlware(prev func(w http.ResponseWriter, r *http.Request) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var appError *apperror.AppError
		err := prev(w, r)
		if err != nil {
			if errors.As(err, &appError) {
				err := err.(*apperror.AppError)
				responseBytes := json.Marshal(err)
				json.SendJSON(w, responseBytes)
				return
			}
			w.WriteHeader(418)
			w.Write([]byte(err.Error()))
		}
	})
}
