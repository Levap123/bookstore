package handler

import (
	apiclients "github.com/Levap123/api_gateway/internal/api_clients"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	log        *logrus.Logger
	apiClients *apiclients.ApiClients
}

func NewHandler(log *logrus.Logger, apiClients *apiclients.ApiClients) *Handler {
	return &Handler{
		log:        log,
		apiClients: apiClients,
	}
}
