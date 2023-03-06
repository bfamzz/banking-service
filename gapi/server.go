package gapi

import (
	"fmt"

	db "github.com/bfamzz/banking-service/db/sqlc"
	"github.com/bfamzz/banking-service/pb"
	"github.com/bfamzz/banking-service/token"
	"github.com/bfamzz/banking-service/util"
)

// Server serves HTTP requests for the banking service
type Server struct {
	pb.UnimplementedBankingServiceServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoV4Maker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
