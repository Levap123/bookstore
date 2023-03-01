package handler

import (
	apiclients "api-gateway/internal/api_clients"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	log        *logrus.Logger
	userClient *apiclients.UserClient
}

func NewHandler(log *logrus.Logger, userClient *apiclients.UserClient) *Handler {
	return &Handler{
		log:        log,
		userClient: userClient,
	}
}
