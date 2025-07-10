package submissions

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

type handlerResidencias struct {
	db   *sqlx.DB
	txID string
	msg  response_messages.Message
}

// StudentsBySubmissions godoc
// @Summary Obtener alumnos por convocatoria
// @Description Método que permite obtener la lista de alumnos asociados a una convocatoria
// @Tags Convocatorias
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path integer true "ID de la convocatoria"
// @Param page query integer false "Número de página" minimum(0) default(0)
// @Param limit query integer false "Límite de registros por página" minimum(0) maximum(100) default(0)
// @Param gender query string false "Género del estudiante (masculino/femenino)" Enums(masculino,femenino)
// @Success 200 {object} models.Response{error=boolean,data=[]models.Student,code=integer,type=string,msg=string} "Lista de alumnos obtenida exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Failure 500 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error interno del servidor"
// @Router /v1/convocatorias/{id}/alumnos-aceptados [GET]
func (h *handlerResidencias) StudentsBySubmissions(c *fiber.Ctx) error {
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

	page := c.QueryInt("page", 0)

	limit := c.QueryInt("limit", 0)

	gender := c.Query("gender")

	if page < 1 {
		limit = 0
	}

	if page >= 1 && (limit < 1 || limit > 100) {
		logger.Error.Printf("invalid limit number: %d", limit)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	if gender != "" && (gender != "masculino" && gender != "femenino") {
		logger.Error.Printf("invalid gender: %s", gender)
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
	students, code, total, err := srv.GetStudentsBySubmissionsLowCode(submission, page, limit, gender)
	if err != nil {
		logger.Error.Printf("error getting students: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res.Data = ResponseStudentsSubmission{
		Total:    total,
		Students: students,
	}

	res.Error = false
	res.Code, res.Type, res.Msg = h.msg.GetByCode(210)
	return c.Status(http.StatusOK).JSON(res)
}

// StudentReportsBySubmissions godoc
// @Summary Generar reporte Excel de alumnos por convocatoria
// @Description Método que permite generar un reporte Excel en formato base64 de los alumnos asociados a una convocatoria
// @Tags Convocatorias
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path integer true "ID de la convocatoria"
// @Success 200 {object} models.Response{error=boolean,data=string,code=integer,type=string,msg=string} "Reporte Excel generado exitosamente en formato base64"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Failure 500 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error interno del servidor"
// @Router /v1/convocatorias/{id}/reporte-residencias [GET]
func (h *handlerResidencias) StudentReportsBySubmissions(c *fiber.Ctx) error {
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

	submission, err := c.ParamsInt("id")
	if err != nil {
		logger.Error.Printf("missing submission_id parameter")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	srv := low_code_submissions.NewSubmission(h.db, user, h.txID)
	base64, code, err := srv.GetReportBySubmissionsLowCode(submission)
	if err != nil {
		logger.Error.Printf("error generate report: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res.Data = base64

	res.Error = false
	res.Code, res.Type, res.Msg = h.msg.GetByCode(223)
	return c.Status(http.StatusOK).JSON(res)
}
