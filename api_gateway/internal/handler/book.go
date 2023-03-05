package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Levap123/api_gateway/internal/dto"
	jsend "github.com/Levap123/api_gateway/pkg/json"
	"github.com/Levap123/utils/apperror"
)

func (h *Handler) createBook(w http.ResponseWriter, r *http.Request) error {
	h.log.Debug("create book")

	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Errorf("error in reading request: %v", err)
		return apperror.NewError(err, "incorrect request body", http.StatusBadRequest)
	}

	var createBookDTO dto.CreateBookDTO
	if err := json.Unmarshal(reqBytes, &createBookDTO); err != nil {
		h.log.Errorf("error in unmarshalling request: %v", err)
		return apperror.NewError(err, "incorrect request body", http.StatusBadRequest)
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second)
	defer cancel()

	bookID, err := h.apiClients.BookClient.Create(ctx, createBookDTO)
	if err != nil {
		h.log.Errorf("error in creaing book: %v", err)
		return err
	}

	respBytes := jsend.Marshal(map[string]string{"book_id": bookID})

	jsend.SendJSON(w, respBytes, http.StatusOK)

	return nil
}
