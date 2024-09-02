package apiservice

import (
	"context"

	"connectrpc.com/connect"
	apiv1 "github.com/geoah/go-skeleton/proto/api/v1"
	apiv1connect "github.com/geoah/go-skeleton/proto/api/v1/apiv1connect"
)

// APIService implements the APIService interface defined in the proto file
type APIService struct {
	apiv1connect.UnimplementedAPIServiceHandler
}

// NewAPIService creates a new instance of APIService
func NewAPIService() *APIService {
	return &APIService{}
}

// Ping implements the Ping RPC method
func (s *APIService) Ping(ctx context.Context, req *connect.Request[apiv1.PingRequest]) (res *connect.Response[apiv1.PingResponse], err error) {
	return connect.NewResponse(&apiv1.PingResponse{
		Message: "Pong: " + req.Msg.Message,
	}), nil
}
