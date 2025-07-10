package auth_orchestrator

import (
	"dbu-api/internal/helpers"
	"dbu-api/internal/logger"
	"dbu-api/internal/middleware"
	"dbu-api/pkg/core"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"time"
)

// Códigos de error específicos del servicio de login
const (
	// Códigos de éxito (1000-1099)
	LoginSuccess           = 1000
	TokenGenerationSuccess = 1001

	// Códigos de error de usuario (1100-1199)
	ErrUserNotFound    = 1100
	ErrInvalidPassword = 1101
	ErrUserInactive    = 1102
	ErrUserBlocked     = 1103

	// Códigos de error de token (1200-1299)
	ErrTokenGeneration        = 1200
	ErrRefreshTokenGeneration = 1201
)

type LoginService struct {
	db   *sqlx.DB
	txID string
}

type PortsServerLogin interface {
	Login(username, password, realIP, baseURL string, isLogin bool) (string, string, int, error)
}

func NewLogin(db *sqlx.DB, txID string) PortsServerLogin {
	return &LoginService{db: db, txID: txID}
}

func (s *LoginService) Login(username, password, realIP, baseURL string, isLogin bool) (string, string, int, error) {
	var token, refreshToken string

	srvCoreService := core.NewServerCore(s.db, nil, s.txID)
	usr, _, err := srvCoreService.SrvUsers.GetUsersByUsername(strings.ToLower(username))
	if err != nil {
		logger.Error.Printf("%s - couldn't get user by username: %v", s.txID, err)
		return "", "", 51, fmt.Errorf("error getting user: %w", err)
	}

	if usr == nil {
		logger.Error.Printf("%s - user not found for username: %s", s.txID, username)
		return "", "", 51, fmt.Errorf("user not found")
	}

	if valid := helpers.CheckPasswordHash(password, usr.Password); !valid && isLogin {
		logger.Error.Printf("%s - password validation failed for user: %s", s.txID, username)
		return "", "", 51, fmt.Errorf("invalid credentials")
	}

	// Generar JWT Token
	token, err = middleware.CreateJWT(usr.ID, baseURL, time.Duration(1))
	if err != nil {
		logger.Error.Printf("%s - error generating JWT token: %v", s.txID, err)
		return "", "", 54, fmt.Errorf("error generating token: %w", err)
	}

	// Generar Refresh Token
	refreshToken, err = middleware.CreateJWT(usr.ID, baseURL, time.Duration(24))
	if err != nil {
		logger.Error.Printf("%s - error generating refresh token: %v", s.txID, err)
		return "", "", 58, fmt.Errorf("error generating refresh token: %w", err)
	}

	return token, refreshToken, 220, nil
}
