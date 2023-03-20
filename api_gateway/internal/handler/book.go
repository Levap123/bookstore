package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Levap123/api_gateway/internal/dto"
	"github.com/Levap123/api_gateway/internal/entity"
	jsend "github.com/Levap123/api_gateway/pkg/json"
	"github.com/Levap123/utils/apperror"
	"github.com/julienschmidt/httprouter"
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

func (h *Handler) getAllBoks(w http.ResponseWriter, r *http.Request) error {
	h.log.Debug("get all books")

	params := r.URL.Query()

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()
	var books []entity.Book
	if len(params) != 0 {
		var err error
		books, err = h.apiClients.BookClient.GetByFiltering(ctx, params)
		if err != nil {
			h.log.Errorf("error in getting books by filter: %v", err)
			return err
		}

	} else {
		var err error
		books, err = h.apiClients.BookClient.GetAll(ctx)
		if err != nil {
			h.log.Errorf("error in getting all books: %v", err)
			return err
		}
	}

	reqBytes := jsend.Marshal(books)

	jsend.SendJSON(w, reqBytes, http.StatusOK)

	return nil
}

func (h *Handler) getBookByID(w http.ResponseWriter, r *http.Request) error {
	h.log.Debug(("get book by ID"))

	ctx, cancel := context.WithTimeout(r.Context(), time.Second)
	defer cancel()

	params := httprouter.ParamsFromContext(r.Context())

	bookID := params.ByName("book_id")

	book, err := h.apiClients.BookClient.GetByID(ctx, bookID)
	if err != nil {
		h.log.Errorf("error in getting book by ID: %v", err)
		return err
	}

	reqBytes := jsend.Marshal(book)

	jsend.SendJSON(w, reqBytes, http.StatusOK)

	return nil
}
