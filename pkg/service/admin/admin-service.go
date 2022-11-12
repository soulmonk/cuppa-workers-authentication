package admin

import (
	"context"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/api/admin"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	adminSecret = "adminSecret"
	apiVersion  = "v1"
)

// adminServiceServer is implementation of admin.adminServiceServer proto interface
type adminServiceServer struct {
	dao *db.Dao
}

// checkAPI checks if the API version requested by client is supported by server
func (s *adminServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

func (s *adminServiceServer) ResetPassword(ctx context.Context, request *admin.ResetPasswordRequest) (*admin.ResetPasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *adminServiceServer) Activate(ctx context.Context, req *admin.ActivateRequest) (*admin.ActivateResponse, error) {
	//TODO implement me
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	if req.Secret != adminSecret {
		return nil, status.Error(codes.PermissionDenied, "I do not know you")
	}

	u, err := s.dao.UserQuerier.FindById(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "")
	}

	if err := s.dao.UserQuerier.Activate(ctx, u.ID); err != nil {
		return nil, err
	}

	return &admin.ActivateResponse{
		Api: apiVersion,
		Id:  u.ID,
	}, nil
}

// NewAuthenticationServiceServer creates Authentication service
func NewAuthenticationServiceServer(dao *db.Dao) admin.AdminServiceServer {
	return &adminServiceServer{dao: dao}
}
