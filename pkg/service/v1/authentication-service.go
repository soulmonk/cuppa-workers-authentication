package v1

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/api/v1"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/db/pg"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/db/pg/domain"
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
	User *domain.User
	jwt.StandardClaims
}

// authenticationServiceServer is implementation of v1.AuthenticationServiceServer proto interface
type authenticationServiceServer struct {
	dao *pg.Dao
}

// NewAuthenticationServiceServer creates Authentication service
func NewAuthenticationServiceServer(dao *pg.Dao) v1.AuthenticationServiceServer {
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

	if _, err := s.dao.UserDao.FindByName(req.Username); err == nil {
		return nil, status.Errorf(codes.AlreadyExists,
			"user '%s' already exists", req.Username)
	}

	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	var user = domain.User{Name: req.Username, Email: req.Email, Password: string(hashedPass)}
	if err := s.dao.UserDao.Create(&user); err != nil {
		return nil, err
	}
	return &v1.SignUpResponse{
		Api: apiVersion,
		Id:  user.ID,
	}, nil
}

func (s *authenticationServiceServer) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	user, err := s.dao.UserDao.FindByName(req.Username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, "")
	}

	//token := jwt.New(jwt.SigningMethodRS256)
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    "auth.service",
		IssuedAt:  time.Now().Unix(),
		Subject:   strconv.FormatUint(user.ID, 10),
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
	token, err := jwt.ParseWithClaims(req.Token, &jwt.StandardClaims{}, func(token *jwt.Token) (any, error) {
		return key, nil
	})

	// Validate the token and return the custom claims
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok || !token.Valid || claims.Subject == "" {
		return nil, err
	}

	//uid, err := strconv.ParseUint(claims.Subject, 10, 64)
	//if err != nil {
	//	// todo log
	//	logger.Log.Error("ParseUint")
	//	return nil, status.Error(codes.Internal, "")
	//}

	//user, err := s.dao.UserDao.FindById(strconv.FormatUint(claims.Subject, 10))
	user, err := s.dao.UserDao.FindById(claims.Subject)
	if err != nil {
		logger.Log.Error("find by id")
		return nil, status.Error(codes.Internal, "")
	}

	if !user.Enabled {
		return nil, status.Error(codes.PermissionDenied, "")
	}

	return &v1.ValidateResponse{
		Api: apiVersion,
		Id:  user.ID,
	}, nil
}
