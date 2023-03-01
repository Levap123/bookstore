package middlwares

import (
	"errors"
	"net/http"

	"github.com/Levap123/user_service/internal/domain"

	"github.com/Levap123/utils/apperror"
)

func CheckErrorMiddlware(prev func(w http.ResponseWriter, r *http.Request) error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := prev(w, r)
		if err != nil {
			if errors.Is(err, domain.ErrPasswordLengthIncorrect) {
				appErr := apperror.MakeBadRequestErr(domain.ErrPasswordLengthIncorrect, "password length should be from 8 to 20")
			}
		}
	})
}
