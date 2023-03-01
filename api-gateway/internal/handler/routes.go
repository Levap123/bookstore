package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Handler) InitRoutes() http.Handler {
	r := httprouter.New()
	r.POST("/auth/sign-up", h.signUp)
}
