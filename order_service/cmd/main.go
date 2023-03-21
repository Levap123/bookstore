package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Levap123/order_service/internal/configs"
	"github.com/Levap123/order_service/internal/order"
	"github.com/Levap123/order_service/internal/order/repository/postgres"
	"github.com/Levap123/order_service/proto"
	"github.com/Levap123/utils/lg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log, err := lg.NewLogger()
	if err != nil {
		log.Fatalf("fatal in creating logger: %v", err)
	}

	cfg, err := configs.GetConfigs()
	if err != nil {
		log.Fatalf("fatal in getting configs: %v", err)
	}

	DB, err := postgres.InitDB(cfg)
	if err != nil {
		log.Fatalf("fatal in init db: %v", err)
	}
	defer DB.Close()

	ctxPing, cancelPing := context.WithTimeout(context.Background(), time.Second)
	defer cancelPing()

	if err := DB.PingContext(ctxPing); err != nil {
		log.Fatalf("fatal in pinging db: %v", err)
	}

	repo := postgres.NewOrderRepoPostgres(DB)
	service := order.NewService(repo)
	handler := order.NewOrderHandler(service, log)

	listener, err := net.Listen("tcp", cfg.Server.Addr)
	if err != nil {
		log.Fatalf("error in starting listener: %v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterOrdersServer(srv, handler)
	reflection.Register(srv)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	log.Info("server started")
	go func() {
		if err := srv.Serve(listener); err != nil {
			log.Fatalf("error in starting serving: %v", err)
		}
	}()

	<-quit

	srv.GracefulStop()
}
