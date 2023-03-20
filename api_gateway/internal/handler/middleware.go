package handler

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/Levap123/api_gateway/pkg/json"
	"github.com/Levap123/utils/apperror"
)

func (h *Handler) UserIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.log.Debug("refresh")

		authHeader := r.Header.Get("Authorization")
		authHeaderSplit := strings.Split(authHeader, "Bearer ")
		if len(authHeaderSplit) != 2 {
			err := apperror.NewError(errors.New("auth header must be: Bearer <token>"), "invalid auth header", http.StatusUnauthorized)
			bytes := json.Marshal(err)
			json.SendJSON(w, bytes, http.StatusUnauthorized)
			return
		}

		authToken := authHeaderSplit[1]

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*1)
		defer cancel()

		userID, err := h.apiClients.UserClient.Validate(ctx, authToken)
		if err != nil {
			err := apperror.NewError(err, "error in validating token", http.StatusUnauthorized)
			bytes := json.Marshal(err)
			json.SendJSON(w, bytes, http.StatusUnauthorized)
			return
		}

		ctxWithValue := context.WithValue(r.Context(), "user_id", userID)

		next.ServeHTTP(w, r.WithContext(ctxWithValue))
	})
}

func (h *Handler) AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.log.Debug("admin middleware")

		authHeader := r.Header.Get("Authorization")
		authHeaderSplit := strings.Split(authHeader, "Bearer ")
		if len(authHeaderSplit) != 2 {
			err := apperror.NewError(errors.New("auth header must be: Bearer <token>"), "invalid auth header", http.StatusUnauthorized)
			bytes := json.Marshal(err)
			json.SendJSON(w, bytes, http.StatusUnauthorized)
			return
		}

		authToken := authHeaderSplit[1]

		ctx, cancel := context.WithTimeout(r.Context(), time.Second*1)
		defer cancel()

		userID, err := h.apiClients.UserClient.Validate(ctx, authToken)
		if err != nil {
			err := apperror.NewError(err, "error in validating token", http.StatusUnauthorized)
			bytes := json.Marshal(err)
			json.SendJSON(w, bytes, http.StatusUnauthorized)
			return
		}

		if userID != 1 {
			err := apperror.NewError(errors.New("invalid user"), "you are not admin", http.StatusUnauthorized)
			bytes := json.Marshal(err)
			json.SendJSON(w, bytes, http.StatusUnauthorized)
			return
		}

		ctxWithValue := context.WithValue(r.Context(), "user_id", userID)

		next.ServeHTTP(w, r.WithContext(ctxWithValue))
	})
}
