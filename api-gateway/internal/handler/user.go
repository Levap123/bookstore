package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	jsend "github.com/Levap123/api_gateway/pkg/json"
	"github.com/julienschmidt/httprouter"

	"github.com/Levap123/utils/apperror"
)

func (h *Handler) getMe(w http.ResponseWriter, r *http.Request) error {
	authHeader := r.Header.Get("Authorization")
	authHeaderSplit := strings.Split(authHeader, "Bearer ")
	if len(authHeaderSplit) != 2 {
		return apperror.NewError(errors.New("auth header must be: Bearer <token>"), "invalid auth header", http.StatusUnauthorized)
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	authToken := authHeaderSplit[1]
	user, err := h.userClient.GetMe(ctx, authToken)
	if err != nil {
		return err
	}

	bytes := jsend.Marshal(user)
	jsend.SendJSON(w, bytes, http.StatusOK)
	return nil
}

func (h *Handler) getByID(w http.ResponseWriter, r *http.Request) error {
	params := httprouter.ParamsFromContext(r.Context())
	userID, err := strconv.Atoi(params.ByName("user_id"))
	if err != nil {
		return apperror.NewError(errors.New("not found"), "not found", http.StatusNotFound)
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	user, err := h.userClient.GetByID(ctx, uint64(userID))
	if err != nil {
		return err
	}

	bytes := jsend.Marshal(user)
	jsend.SendJSON(w, bytes, http.StatusOK)
	return nil
}
