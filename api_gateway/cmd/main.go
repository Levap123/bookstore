package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	apiclients "github.com/Levap123/api_gateway/internal/api_clients"
	"github.com/Levap123/api_gateway/internal/configs"
	"github.com/Levap123/api_gateway/internal/handler"
	"github.com/Levap123/api_gateway/pkg/server"

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

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	ctxUsersrv, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	userServiceAddress := cfg.UserService.Addr

	connUsersrv, err := grpc.DialContext(ctxUsersrv, userServiceAddress, opts...)
	if err != nil {
		log.Fatalf("fatal in connect to user service: %v", err)
	}
	defer connUsersrv.Close()

	ctxBooksrv, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	bookServiceAddress := cfg.UserService.Addr

	connBooksrv, err := grpc.DialContext(ctxBooksrv, bookServiceAddress, opts...)
	if err != nil {
		log.Fatalf("fatal in connect to book service: %v", err)
	}
	defer connBooksrv.Close()

	userServiceClient := apiclients.InitUserClient(connUsersrv, log)
	bookServiceClient := apiclients.InitBookClient(connBooksrv, log)

	apiclients := apiclients.NewApiClients(bookServiceClient, userServiceClient)

	handler := handler.NewHandler(log, apiclients)

	server := new(server.Server)

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := server.Run(handler.InitRoutes(), cfg.Server.Addr); err != nil {
			log.Fatalf("fatal in running server: %v", err)
		}
	}()
	log.Info("server started")

	<-quit

	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()

	if err := server.Stop(ctx1); err != nil {
		log.Fatalf("fatal in stopping server: %v", err)
	}

	log.Info("server stopped")
}
