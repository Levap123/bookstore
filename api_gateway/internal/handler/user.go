package handler

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Levap123/api_gateway/internal/dto"
	jsend "github.com/Levap123/api_gateway/pkg/json"
	"github.com/julienschmidt/httprouter"

	"github.com/Levap123/utils/apperror"
)

func (h *Handler) getMe(w http.ResponseWriter, r *http.Request) error {
	h.log.Debug("get me")

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

func (h *Handler) getUserByID(w http.ResponseWriter, r *http.Request) error {
	h.log.Debug("get user by ID")

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

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) error {
	h.log.Debug("update user")

	request, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Errorf("error in reading request: %v", err)
		return apperror.NewError(err, "incorrect request body", http.StatusBadRequest)
	}

	var dto dto.UpdateUserDTO
	if err := json.Unmarshal(request, &dto); err != nil {
		h.log.Errorf("error in unmarshalling request: %v", err)
		return apperror.NewError(err, "incorrect request body", http.StatusBadRequest)
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	userID, err := h.userClient.Update(ctx, &dto)
	if err != nil {
		return err
	}

	bytes := jsend.Marshal(map[string]uint64{"user_id": userID})
	jsend.SendJSON(w, bytes, http.StatusOK)
	return nil
}
