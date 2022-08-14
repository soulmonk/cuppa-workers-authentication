package v1

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/soulmonk/cuppa-workers-authentication/db/user"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/api/v1"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/db"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"time"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

var (

	// TODO
	// Define a secure key string used
	// as a salt when hashing our tokens.
	// Please make your own way more secure than this,
	// use a randomly generated md5 hash or something.
	key = []byte("mySuperSecretKeyLol")
)

type CustomClaims struct {
	User *user.User
	jwt.RegisteredClaims
}

// authenticationServiceServer is implementation of v1.AuthenticationServiceServer proto interface
type authenticationServiceServer struct {
	dao *db.Dao
}

// NewAuthenticationServiceServer creates Authentication service
func NewAuthenticationServiceServer(dao *db.Dao) v1.AuthenticationServiceServer {
	return &authenticationServiceServer{dao: dao}
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
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	if _, err := s.dao.UserQuerier.FindByName(ctx, req.Username); err == nil {
		return nil, status.Errorf(codes.AlreadyExists,
			"user '%s' already exists", req.Username)
	}

	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	arg := user.CreateParams{Name: req.Username, Email: req.Email, Password: string(hashedPass)}
	userCreated, err := s.dao.UserQuerier.Create(ctx, arg)
	if err != nil {
		return nil, err
	}
	return &v1.SignUpResponse{
		Api: apiVersion,
		Id:  userCreated.ID,
	}, nil
}

func (s *authenticationServiceServer) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	u, err := s.dao.UserQuerier.FindByName(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password))
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "")
	}

	//token := jwt.New(jwt.SigningMethodRS256)
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		Issuer:    "auth.service",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   strconv.FormatInt(u.ID, 10),
	}
	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(key)
	if err != nil {
		return nil, status.Error(codes.Internal, "")
	}
	// Sign token and return
	return &v1.LoginResponse{
		Api:   apiVersion,
		Token: signedToken,
	}, nil
}

func (s *authenticationServiceServer) Validate(ctx context.Context, req *v1.ValidateRequest) (*v1.ValidateResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}
	// Parse the token
	token, err := jwt.ParseWithClaims(req.Token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		return key, nil
	})

	// Validate the token and return the custom claims
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid || claims.Subject == "" {
		return nil, err
	}

	uid, err := strconv.ParseInt(claims.Subject, 10, 64)
	if err != nil {
		// todo log
		logger.Log.Error("ParseUint")
		return nil, status.Error(codes.Internal, "")
	}

	//u, err := s.dao.UserDao.FindById(strconv.FormatUint(claims.Subject, 10))
	u, err := s.dao.UserQuerier.FindById(ctx, uid)
	if err != nil {
		logger.Log.Error("find by id")
		return nil, status.Error(codes.Internal, "")
	}

	if !u.Enabled {
		return nil, status.Error(codes.PermissionDenied, "")
	}

	return &v1.ValidateResponse{
		Api: apiVersion,
		Id:  u.ID,
	}, nil
}
