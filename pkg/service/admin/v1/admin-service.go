package admin_v1

import (
	"context"
	admin_v1 "github.com/soulmonk/cuppa-workers-authentication/pkg/api/admin"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/db"
)

const (
	adminSecret = "adminSecret"
	apiVersion  = "v1"
)

// adminServiceServer is implementation of admin_v1.adminServiceServer proto interface
type adminServiceServer struct {
	dao *db.Dao
}

func (a *adminServiceServer) ResetPassword(ctx context.Context, request *admin_v1.ResetPasswordRequest) (*admin_v1.ResetPasswordResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *adminServiceServer) Activate(ctx context.Context, request *admin_v1.ActivateRequest) (*admin_v1.ActivateResponse, error) {
	//TODO implement me
	panic("implement me")
}

// NewAuthenticationServiceServer creates Authentication service
func NewAuthenticationServiceServer(dao *db.Dao) admin_v1.AdminServiceServer {
	return &adminServiceServer{dao: dao}
}
