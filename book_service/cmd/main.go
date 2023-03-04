package main

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Levap123/book_service/internal/book"
	"github.com/Levap123/book_service/internal/book/mongo"
	"github.com/Levap123/book_service/internal/configs"
	"github.com/Levap123/book_service/proto"
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	DB, err := mongo.InitDB(ctx, cfg)
	if err != nil {
		log.Fatalf("fatal in initializing DB: %v", err)
	}

	repo := mongo.NewBookRepo(DB)
	service := book.NewBookService(repo)
	handler := book.NewBookHandler(service, log)

	listener, err := net.Listen("tcp", cfg.Server.Addr)
	if err != nil {
		log.Fatalf("error in starting listener: %v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterBookServer(srv, handler)
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

	ctxShutdown, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := mongo.ShutDownDB(ctxShutdown, DB); err != nil {
		log.Fatalf("error in db shutdown: %v", err)
	}

	log.Info("server stopped")
}
