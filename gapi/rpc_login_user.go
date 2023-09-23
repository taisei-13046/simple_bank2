package gapi

import (
	"context"
	"database/sql"

	db "github.com/taisei-13046/simple_bank2/db/sqlc"
	"github.com/taisei-13046/simple_bank2/pb"
	"github.com/taisei-13046/simple_bank2/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := server.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "not found")
		}
		return nil, status.Errorf(codes.Internal, "fail to find user")
	}

	err = util.CheckPassword(req.GetPassword(), user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to find user")
	}

	accessToken, accesPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to find user")
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to find user")
	}

	mtdt := server.extractMetadata(ctx)
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    mtdt.UserAgent,
		ClientIp:     mtdt.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiresAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to find user")
	}

	rsp := &pb.LoginUserResponse{
		User:                  convertUser(user),
		SessionId:             session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  timestamppb.New(accesPayload.ExpiresAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiresAt),
	}

	return rsp, nil
}
