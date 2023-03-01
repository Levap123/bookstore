package handler

import (
	"api-gateway/internal/dto"
	"encoding/json"
	"io"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) error {
	var dto dto.SignUpDTO

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, &dto); err != nil {
		return err
	}

	h.userClient
}
