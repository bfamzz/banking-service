package gapi

import (
	"context"
	"database/sql"
	"time"

	db "github.com/bfamzz/banking-service/db/sqlc"
	"github.com/bfamzz/banking-service/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {

	arg := db.GetVerifyEmailParams{
		ID:         int64(req.GetId()),
		SecretCode: req.GetSecretCode(),
	}

	if arg.ID <= 0 || arg.SecretCode == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Bad request data")
	}

	verifyEmail, err := server.store.GetVerifyEmail(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "email verification data not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "email verification failed: %s", err)
	}

	if verifyEmail.IsUsed {
		return &pb.VerifyEmailResponse{
			VerifyEmailStatus: "Email has been verified",
		}, nil
	}

	if time.Now().After(verifyEmail.ExpiresAt) {
		return nil, status.Errorf(codes.Unauthenticated, "email verification link has expired. Please request a new one")
	}

	argVerifyEmailTx := db.VerifyEmailTxParams{
		VerifyEmailParams: db.VerifyEmailParams{
			ID:         arg.ID,
			SecretCode: arg.SecretCode,
			IsUsed:     true,
		},
		AfterVerify: func(verifyEmail db.VerifyEmail) error {
			_, err = server.store.UpdateUserVerifyEmail(ctx, db.UpdateUserVerifyEmailParams{
				Username:        verifyEmail.Username,
				IsEmailVerified: true,
			})

			return err
		},
	}

	_, err = server.store.VerifyEmailTx(ctx, argVerifyEmailTx)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "email verification failed: %s", err)
	}

	res := &pb.VerifyEmailResponse{
		VerifyEmailStatus: "Email has been verified",
	}
	return res, nil
}
