package main

import (
	apiclients "api-gateway/internal/api_clients"
	"api-gateway/internal/configs"
	"api-gateway/internal/handler"
	"time"

	"github.com/Levap123/utils/lg"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	log, err := lg.NewLogger()
	if err != nil {
		logrus.Fatalf("fatal in initialize logger: %v", err)
	}

	cfg, err := configs.GetConfigs()
	if err != nil {
		log.Fatalf("fatal in itialize configs: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	userServiceAddress := cfg.UserService.Addr
	conn, err := grpc.DialContext(ctx, userServiceAddress)
	if err != nil {
		log.Fatalf("fatal in conntect to user service: %v", err)
	}

	userServiceClient := apiclients.InitUserClient(conn)
	handler := handler.NewHandler(log, userServiceClient)
}
