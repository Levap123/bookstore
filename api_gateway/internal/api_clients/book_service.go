package apiclients

import (
	"github.com/Levap123/api_gateway/proto"
	"github.com/sirupsen/logrus"
)

type BookClient struct {
	cl  proto.BookClient
	log *logrus.Logger
}

