package server

import (
	"context"

	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server
	ctx context.Context
}
