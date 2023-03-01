package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Levap123/user_service/internal/configs"
	"github.com/Levap123/user_service/internal/user"
	"github.com/Levap123/user_service/internal/user/postgres"
	"github.com/Levap123/user_service/internal/validator"
	"github.com/Levap123/user_service/proto"

	"github.com/Levap123/utils/jwt"
	"github.com/Levap123/utils/lg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lg, err := lg.NewLogger()
	if err != nil {
		lg.Fatalf("error in creating logger: %v", err)
	}

	cfg, err := configs.GetConfigs()
	if err != nil {
		lg.Fatalf("error in geting configs %v", err)
	}

	DB, err := postgres.InitDB(cfg)
	if err != nil {
		lg.Fatalf("error in initializing db: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	if err := DB.PingContext(ctx); err != nil {
		lg.Fatalf("ping db error: %v", err)
	}

	listener, err := net.Listen("tcp", cfg.Server.Addr)
	if err != nil {
		log.Fatalf("error in starting listener: %v", err)
	}

	repo := postgres.NewUserRepo(DB, lg)

	jwt := jwt.NewJWT(cfg.JWTSign)
	service := user.NewUserService(repo, jwt)

	validator := validator.NewValidator(cfg)
	handler := user.NewUserHandler(service, lg, validator)

	srv := grpc.NewServer()
	proto.RegisterUserServer(srv, handler)
	reflection.Register(srv)
	lg.Info("server is started")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		if err := srv.Serve(listener); err != nil {
			log.Fatalln(err)
		}
	}()
	<-quit
	srv.Stop()
	lg.Info("server is stoped")

}
