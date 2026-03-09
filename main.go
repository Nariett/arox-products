package main

import (
	"arox-products/internal/handler"
	"arox-products/internal/stores"
	"arox-products/schema"
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Nariett/arox-pkg/config"
	"github.com/Nariett/arox-pkg/db"
	proto "github.com/Nariett/arox-pkg/grpc/pb/products"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartServer(lc fx.Lifecycle, h handler.Handler, cfg *config.Config) {
	protocol, port := cfg.Protocol, cfg.LPort
	listener, err := net.Listen(protocol, port)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)

		return
	}
	server := grpc.NewServer()
	proto.RegisterProductsServiceServer(server, h)

	reflection.Register(server)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("gRPC-server start")
			go func() {
				if err := server.Serve(listener); err != nil {
					log.Fatalf("error gRPC-server: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.GracefulStop()
			log.Println("gRPC-server stop")
			return nil
		},
	})
}

func main() {
	fs := schema.DB
	application := fx.New(
		fx.Supply(&fs),
		fx.Provide(
			config.New,
			db.NewPostgres,
			handler.NewHandler,
		),
		stores.Construct(),
		fx.Invoke(
			db.Migrate,
			StartServer),
	)
	application.Run()
}
