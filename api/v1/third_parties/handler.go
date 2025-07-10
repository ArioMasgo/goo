package third_parties

import (
	"dbu-api/internal/authorization"
	"dbu-api/internal/logger"
	"dbu-api/internal/middleware"
	"dbu-api/internal/models"
	"dbu-api/pkg/orchestrator/low_code_submissions"
	"dbu-api/pkg/orchestrator/response_messages"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerThirdParty struct {
	db   *sqlx.DB
	txID string
	msg  response_messages.Message
}

// StudentDebts godoc
// @Summary Obtener alumnos por convocatoria
// @Description Método que permite obtener la lista de alumnos asociados a una convocatoria
// @Tags Convocatorias
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path integer true "ID de la convocatoria"
// @Param page query integer false "Número de página" minimum(1) default(1)
// @Param limit query integer false "Límite de registros por página" minimum(1) maximum(100) default(10)
// @Success 200 {object} models.Response{error=boolean,data=[]models.Student,code=integer,type=string,msg=string} "Lista de alumnos obtenida exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Failure 500 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error interno del servidor"
// @Router /v1/convocatorias/{id}/deuda/alumnos-aceptados [GET]
func (h *handlerThirdParty) StudentDebts(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	err = authorization.ValidPermissions(user, h.db, c)
	if err != nil {
		logger.Error.Printf("User does not have permission to call the api, error: %v", err)
		res.Code, res.Type, res.Msg = 10, "error", "rejected for route permits"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	page := c.QueryInt("page", 1)

	if page < 1 {
		logger.Error.Printf("invalid page number: %d", page)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	limit := c.QueryInt("limit", 10)

	if limit < 1 || limit > 100 {
		logger.Error.Printf("invalid limit number: %d", limit)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	submission, err := c.ParamsInt("id")
	if err != nil {
		logger.Error.Printf("missing submission_id parameter")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	srv := low_code_submissions.NewSubmission(h.db, user, h.txID)
	students, _, code, err := srv.GetStudentsBySubmissionsLowCode(submission, page, limit)
	if err != nil {
		logger.Error.Printf("error getting students: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res.Data = students

	res.Error = false
	res.Code, res.Type, res.Msg = h.msg.GetByCode(210)
	return c.Status(http.StatusOK).JSON(res)
}
