package authentication

import (
	"dbu-api/internal/authorization"
	"dbu-api/internal/logger"
	"dbu-api/internal/middleware"
	"dbu-api/internal/models"
	"dbu-api/pkg/orchestrator/auth"
	"dbu-api/pkg/orchestrator/response_messages"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type handlerAuthentication struct {
	db   *sqlx.DB
	txID string
	msg  response_messages.Message
}

// Login godoc
// @Summary Iniciar sesión
// @Description Método que permite iniciar sesión
// @Tags Login
// @Accept json
// @Produce json
// @Param RequestLogin body RequestLogin true "Datos para realizar el inicio de sesión"
// @Success 200 {object} models.Response{error=boolean,data=JWTToken,code=integer,type=string,msg=string} "Login exitoso"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Router /v1/login [POST]
func (h *handlerAuthentication) Login(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := RequestLogin{}

	if err := c.BodyParser(&req); err != nil {
		logger.Error.Printf("couldn't parse body request, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	isValid, err := req.Valid()
	if err != nil {
		logger.Error.Printf("couldn't validate body request, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	if !isValid {
		logger.Error.Println("couldn't validate body request")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	srv := auth_orchestrator.NewLogin(h.db, h.txID)
	token, refreshToken, code, err := srv.Login(req.Username, req.Password, c.IP(), c.BaseURL(), true)
	if err != nil {
		logger.Error.Printf("couldn't login, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Data = JWTToken{
		Token:        token,
		RefreshToken: refreshToken,
	}

	res.Error = false
	res.Code, res.Type, res.Msg = h.msg.GetByCode(220)
	return c.Status(fiber.StatusOK).JSON(res)
}

// RefreshToken godoc
// @Summary Refrescar token
// @Description Método que permite crear otra sesión a partir del refresh token
// @Tags RefreshToken
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} models.Response{error=boolean,data=JWTToken,code=integer,type=string,msg=string} "Token refrescado exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Failure 403 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de permisos"
// @Router /v1/refresh-token [GET]
func (h *handlerAuthentication) RefreshToken(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(51)
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	err = authorization.ValidPermissions(user, h.db, c)
	if err != nil {
		logger.Error.Printf("User does not have permission to call the api, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(10)
		return c.Status(fiber.StatusForbidden).JSON(res)
	}

	srv := auth_orchestrator.NewLogin(h.db, h.txID)
	token, refreshToken, _, err := srv.Login(user.Username, "", "", c.BaseURL(), false)
	if err != nil {
		logger.Error.Printf("couldn't refresh token, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(58)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Data = JWTToken{
		Token:        token,
		RefreshToken: refreshToken,
	}

	res.Error = false
	res.Code, res.Type, res.Msg = h.msg.GetByCode(221)
	return c.Status(fiber.StatusOK).JSON(res)
}
