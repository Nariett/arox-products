package main

import (
	"arox-products/internal/handler"
	"context"
	"github.com/Nariett/arox-pkg/config"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartServer(lc fx.Lifecycle, h handler.Handler, cfg *config.Config) {
	protocol, port := cfg.Protocol, cfg.LPort
	listener, err := net.Listen(protocol, port)
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	proto.RegisterProductsServiceServer(server, h)

	reflection.Register(server)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("gRPC-сервер запущен")
			go func() {
				if err := server.Serve(listener); err != nil {
					log.Fatalf("Ошибка запуска gRPC-сервера: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.GracefulStop()
			log.Println("gRPC-сервер остановлен")
			return nil
		},
	})
}

func main() {
	application := fx.New(
		fx.Provide(
			config.New,
			//db.NewPostgres,
			//store.Construct,
			handler.NewHandler,
		),
		fx.Invoke(
			//schema.Migrate,
			StartServer),
	)
	application.Run()
}
