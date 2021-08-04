package shorter

import (
	"context"
	"grpcshorter/pkg/api"
)

type GRPCServer struct{}

func (s *GRPCServer) Short(ctx context.Context, req *api.UrlRequest) (*api.UrlResponse, error) {
	return &api.UrlResponse{Result: req.GetUrl()}, nil
}
