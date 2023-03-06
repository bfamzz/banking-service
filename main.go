package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/bfamzz/banking-service/api"
	db "github.com/bfamzz/banking-service/db/sqlc"
	"github.com/bfamzz/banking-service/gapi"
	"github.com/bfamzz/banking-service/pb"
	"github.com/bfamzz/banking-service/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the db: ", err)
	}

	store := db.NewStore(conn)

	rungRPCServer(config, store)
}

func rungRPCServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create grpc server:", err)
	}

	grpcSever := grpc.NewServer()
	pb.RegisterBankingServiceServer(grpcSever, server)
	reflection.Register(grpcSever)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create grpc listener:", err)
	}

	log.Printf("starting a gRPC server at %s", listener.Addr().String())
	err = grpcSever.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server:", err)
	}
}

func runHTTPServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create http server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
