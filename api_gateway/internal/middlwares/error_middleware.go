package middlwares

import (
	"errors"
	"net/http"

	"github.com/Levap123/api_gateway/pkg/json"

	"github.com/Levap123/utils/apperror"
)

func CheckErrorMiddlware(prev func(w http.ResponseWriter, r *http.Request) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := prev(w, r)
		if err != nil {
			var appError *apperror.AppError
			if errors.As(err, &appError) {
				err := err.(*apperror.AppError)
				respBytes := json.Marshal(err)
				json.SendJSON(w, respBytes, err.Status)
				return
			}
			json.SendJSON(w, []byte(err.Error()), 418)
		}
	})
}
