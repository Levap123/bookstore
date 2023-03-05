package handler

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Levap123/api_gateway/internal/dto"
	jsend "github.com/Levap123/api_gateway/pkg/json"
	"github.com/Levap123/utils/apperror"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) error {
	h.log.Debug("user signup handler")
	var dto dto.SignUpDTO

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Errorf("error in reading request: %v", err)
		return apperror.NewError(err, "incorrect request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(bytes, &dto); err != nil {
		h.log.Errorf("error in unmarshaling: %v", err)
		return apperror.NewError(err, "incorrect request body", http.StatusBadRequest)
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	userID, err := h.apiClients.UserClient.SignUp(ctx, &dto)
	if err != nil {
		h.log.Errorf("error in sending request to user service: %v", err)
		return err
	}

	responseBytes := jsend.Marshal(map[string]uint64{"user_id": userID})
	jsend.SendJSON(w, responseBytes, http.StatusOK)
	return nil
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) error {
	h.log.Debug("user signin handler")
	var dto dto.SignInDTO

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Errorf("error in reading request: %v", err)
		return apperror.NewError(err, "incorrect request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(bytes, &dto); err != nil {
		h.log.Error("error in unmarshalling: %v", err)
		return apperror.NewError(err, "incorrect request body", http.StatusBadRequest)
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	tokens, err := h.apiClients.UserClient.SignIn(ctx, &dto)
	if err != nil {
		h.log.Errorf("error in sending request to user service: %v", err)
		return err
	}

	responseBytes := jsend.Marshal(tokens)
	jsend.SendJSON(w, responseBytes, http.StatusOK)
	return nil
}

func (h *Handler) refresh(w http.ResponseWriter, r *http.Request) error {
	h.log.Debug("refresh")

	authHeader := r.Header.Get("Authorization")
	authHeaderSplit := strings.Split(authHeader, "Bearer ")
	if len(authHeaderSplit) != 2 {
		return apperror.NewError(errors.New("auth header must be: Bearer <token>"), "invalid auth header", http.StatusUnauthorized)
	}

	authToken := authHeaderSplit[1]
	dto := &dto.RefreshDTO{
		AccessToken: authToken,
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Errorf("error in reading request: %v", err)
		return apperror.NewError(err, "incorrect request body", http.StatusBadRequest)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(bytes, dto); err != nil {
		h.log.Error("error in unmarshalling: %v", err)
		return apperror.NewError(err, "incorrect request body", http.StatusBadRequest)
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	tokens, err := h.apiClients.UserClient.Rerfresh(ctx, dto)
	if err != nil {
		return err
	}

	responseBytes := jsend.Marshal(tokens)
	jsend.SendJSON(w, responseBytes, http.StatusOK)
	return nil
}
