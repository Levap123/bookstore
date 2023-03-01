package handler

import (
	"net/http"

	"api-gateway/internal/middlwares"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) InitRoutes() http.Handler {
	r := httprouter.New()
	r.Handler(http.MethodPost, "/auth/sign-up", middlwares.CheckErrorMiddlware(h.signUp))
	return r
}
