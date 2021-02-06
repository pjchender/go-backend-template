package grpc

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/pjchender/go-backend-template/internal/database"
	"github.com/pjchender/go-backend-template/internal/grpc/api"
	juboxPB "github.com/pjchender/go-backend-template/internal/grpc/proto/jubox"
	"net"
)

func Serve(db *database.GormDatabase, grpcAddress, certFile, keyFile string) error {
	var err error

	conn, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}

	grpcServer := grpc.NewServer()

	juboxHandler := api.NewJuboxHandler(db)
	juboxPB.RegisterJuboxServer(grpcServer, juboxHandler)

	reflection.Register(grpcServer)

	err = grpcServer.Serve(conn)
	if err != nil {
		log.Fatal("grpc.Serve Error: ", err)
		return err
	}

	return nil
}
