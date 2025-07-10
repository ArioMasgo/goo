package room_assignment

import (
	"dbu-api/internal/authorization"
	"dbu-api/internal/logger"
	"dbu-api/internal/middleware"
	"dbu-api/internal/models"
	"dbu-api/pkg/orchestrator/low_code_residences"
	"dbu-api/pkg/orchestrator/response_messages"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerRoomAssignment struct {
	db   *sqlx.DB
	txID string
	msg  response_messages.Message
}

// ExecuteRoomAssignment godoc
// @Summary Asignación automática de cuartos en residencias
// @Description Endpoint para ejecutar la asignación automática de cuartos en una residencia específica
// @Tags Automatización
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param RequestRoomAssignment body RequestRoomAssignment true "ID de la residencia para asignación"
// @Success 200 {object} models.Response{error=boolean,data=ResponseAutomaticAssignationResidence,code=integer,type=string,msg=string} "Asignación exitosa"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Failure 500 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error interno del servidor"
// @Router /v1/automatizacion/asignacion-cuartos [POST]
func (h *handlerRoomAssignment) ExecuteRoomAssignment(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := RequestRoomAssignment{}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(9)
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	err = authorization.ValidPermissions(user, h.db, c)
	if err != nil {
		logger.Error.Printf("User does not have permission to call the api, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(10)
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	if err := c.BodyParser(&req); err != nil {
		logger.Error.Printf("couldn't parse body request, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	isValid, err := req.valid()
	if err != nil {
		logger.Error.Printf("couldn't validate body request, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if !isValid {
		logger.Error.Println("couldn't validate body request")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := low_code_residences.NewRobotResidence(h.db, nil, h.txID)
	residenceRobot, _, err := srv.AssignationResidenceLowCode(req.ResidenceId)
	if err != nil {
		logger.Error.Printf("couldn't execute room assignment, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(91)
		return c.Status(http.StatusInternalServerError).JSON(res)
	}

	res.Error = false

	if residenceRobot == "" {
		logger.Error.Printf("no room assignment was generated")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(92)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = ResponseAutomaticAssignationResidence{
		RegisterID: residenceRobot,
	}

	res.Code, res.Type, res.Msg = h.msg.GetByCode(206)
	return c.Status(http.StatusOK).JSON(res)
}
