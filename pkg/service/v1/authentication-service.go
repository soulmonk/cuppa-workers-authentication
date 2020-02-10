package v1

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// authenticationServiceServer is implementation of v1.AuthenticationServiceServer proto interface
type authenticationServiceServer struct {
	db *sqlx.DB
}

// NewAuthenticationServiceServer creates Authentication service
func NewAuthenticationServiceServer(db *sqlx.DB) v1.AuthenticationServiceServer {
	return &authenticationServiceServer{db: db}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *authenticationServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

func (s *authenticationServiceServer) SignUp(ctx context.Context, req *v1.SignUpRequest) (*v1.SignUpResponse, error) {
	panic("implement me")
}

func (s *authenticationServiceServer) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	panic("implement me")
}

func (s *authenticationServiceServer) Validate(ctx context.Context, req *v1.ValidateRequest) (*v1.ValidateResponse, error) {
	panic("implement me")
}
