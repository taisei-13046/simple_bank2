package gapi

import (
	"fmt"

	db "github.com/taisei-13046/simple_bank2/db/sqlc"
	"github.com/taisei-13046/simple_bank2/pb"
	"github.com/taisei-13046/simple_bank2/token"
	"github.com/taisei-13046/simple_bank2/util"
	"github.com/taisei-13046/simple_bank2/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributer worker.TaskDistributor
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store, taskDistributer worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributer: taskDistributer,
	}

	return server, nil
}
