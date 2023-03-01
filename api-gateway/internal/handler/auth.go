package handler

import (
	"github.com/Levap123/api_gateway/internal/dto"
	jsend "github.com/Levap123/api_gateway/pkg/json"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) error {
	h.log.Debug("user signup handler")
	var dto dto.SignUpDTO

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Errorf("error in reading request: %v", err)
		return err
	}
	defer r.Body.Close()

	if err := json.Unmarshal(bytes, &dto); err != nil {
		h.log.Errorf("error in unmarshaling: %v", err)
		return err
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	userID, err := h.userClient.SignUp(ctx, &dto)
	if err != nil {
		h.log.Errorf("error in sending request to user service: %v", err)
		return err
	}

	responseBytes := jsend.Marshal(map[string]uint64{"user_id": userID})
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBytes)
	return nil
}
