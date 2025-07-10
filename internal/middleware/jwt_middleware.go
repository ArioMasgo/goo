package middleware

import (
	"crypto/rand"
	"dbu-api/internal/env"
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"dbu-api/pkg/core"
	"dbu-api/pkg/core/users"
	"encoding/base64"
	"errors"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

// TokenMetadata struct to describe metadata in JWT.
type TokenMetadata struct {
	Expires int64
	Ip      string
}

type JwtCustomClaims struct {
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   int64  `json:"sub,omitempty"`
	jwt.StandardClaims
}

func JWTProtected() fiber.Handler {
	e := env.NewConfiguration()
	config := jwtware.Config{
		ErrorHandler:  jwtError,
		SigningKey:    e.Key.PublicKey,
		SigningMethod: "RS256",
	}
	return jwtware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {

	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusUnauthorized).
			JSON(models.Response{Error: true, Data: nil, Code: 0, Msg: "Missing or malformed JWT", Type: "error"})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(models.Response{Error: true, Data: nil, Code: 0, Msg: "Invalid or expired JWT", Type: "error"})
}

func generateJTI() (string, error) {
	bytes := make([]byte, 12)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(bytes)[:16], nil
}

// CreateJWT crea un token JWT firmado utilizando la clave privada RSA.
func CreateJWT(userID int64, url string, expireTime time.Duration) (string, error) {
	jti, err := generateJTI()
	if err != nil {
		return "", err
	}
	e := env.NewConfiguration()
	claims := JwtCustomClaims{
		ExpiresAt: time.Now().Add(time.Hour * expireTime).Unix(), // El token expira en 24 horas.
		Issuer:    url,                                           // URL del api
		IssuedAt:  time.Now().Unix(),                             // El token fue generado ahora
		NotBefore: time.Now().Unix(),                             // El token solo funciona a partir de ahora
		Id:        jti,                                           // Identificador unico
		Subject:   userID,                                        // ID del usuario
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(e.Key.PrivateKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetUser(bearer string, db *sqlx.DB) (*models.User, error) {
	e := env.NewConfiguration()

	if bearer == "" {
		logger.Warning.Printf("Token vacío")
		return nil, errors.New("token no Valido")
	}

	if len(bearer) < 7 {
		logger.Warning.Printf("Token inválido")
		return nil, errors.New("token no Valido")
	}

	tkn := bearer[7:]

	verifyFunction := func(tkn *jwt.Token) (interface{}, error) {
		return e.Key.PublicKey, nil
	}

	token, err := jwt.ParseWithClaims(tkn, &JwtCustomClaims{}, verifyFunction)
	if err != nil {
		var validationError *jwt.ValidationError
		if !errors.As(err, &validationError) {
			logger.Warning.Printf("Error al procesar el token: %v", err)
			return nil, err
		}

		if validationError.Errors == jwt.ValidationErrorExpired {
			logger.Warning.Printf("token expirado: %v", err)
			return nil, err
		}

		logger.Warning.Printf("Error de validacion del token: %v", err)
		return nil, err
	}

	if !token.Valid {
		logger.Warning.Printf("Token no Valido: %v", err)
		return nil, errors.New("Token no Valido")
	}

	userID := token.Claims.(*JwtCustomClaims).Subject
	if userID == 0 {
		logger.Warning.Printf("Token mal formado: %v", err)
		return nil, errors.New("Token mal formado")
	}

	srv := core.NewServerCore(db, nil, uuid.New().String())
	user, _, err := srv.SrvUsers.GetUsersByID(userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		logger.Error.Println("Usuario no encontrado")
		return nil, errors.New("token no valido")
	}

	return users.MapperUser(user), nil
}
