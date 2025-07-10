package nursing_consultation_assignment

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/middleware"
	"dbu-api/internal/models"
	"dbu-api/pkg/medical_area"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerMedicalArea struct {
	db   *sqlx.DB
	txID string
}

// CreateConsultationAssignment godoc
// @Summary Crear una instancia de paciente
// @Description Método que permite crear una instancia del objeto paciente en la base de datos
// @tags Pacientes
// @Accept json
// @Produce json
// @Param models.ConsultationAssignment body models.ConsultationAssignment true "Datos para crear ConsultationAssignment"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/paciente [POST]
func (h *handlerMedicalArea) CreateConsultationAssignment(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := models.ConsultationAssignment{}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	if err := c.BodyParser(&req); err != nil {
		logger.Error.Printf("couldn't parse body request, error: %v", err)
		res.Code, res.Type, res.Msg = 1, "", ""
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	isValid, err := req.ValidConsultationAssignment()
	if err != nil {
		logger.Error.Printf("couldn't validate body request, error: %v", err)
		res.Code, res.Type, res.Msg = 1, "", ""
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if !isValid {
		logger.Error.Println("couldn't validate body request")
		res.Code, res.Type, res.Msg = 1, "", ""
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := medical_area.NewServerMedicalArea(h.db, user, h.txID)
	data, code, err := srv.SrvConsultationAssignment.CreateConsultationAssignment("", req.IDConsulta, req.AreaAsignada)
	if err != nil {
		logger.Error.Printf("couldn't create patient, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", err.Error()
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", ""
	return c.Status(http.StatusCreated).JSON(res)
}

// UpdateConsultationAssignment godoc
// @Summary Actualiza una instancia de paciente
// @Description Método que permite actualizar una instancia del objeto paciente en la base de datos
// @tags Pacientes
// @Accept json
// @Produce json
// @Param models.ConsultationAssignment body models.ConsultationAssignment true "Datos para actualizar ConsultationAssignment"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/paciente [PUT]
func (h *handlerMedicalArea) UpdateConsultationAssignment(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := models.ConsultationAssignment{}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	if err := c.BodyParser(&req); err != nil {
		logger.Error.Printf("couldn't parse body request, error: %v", err)
		res.Code, res.Type, res.Msg = 1, "", ""
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	isValid, err := req.ValidConsultationAssignment()
	if err != nil {
		logger.Error.Printf("couldn't validate body request, error: %v", err)
		res.Code, res.Type, res.Msg = 1, "", ""
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if !isValid {
		logger.Error.Println("couldn't validate body request")
		res.Code, res.Type, res.Msg = 1, "", ""
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := medical_area.NewServerMedicalArea(h.db, user, h.txID)
	data, code, err := srv.SrvConsultationAssignment.UpdateConsultationAssignmentByIDConsultation("", req.IDConsulta, req.AreaAsignada)
	if err != nil {
		logger.Error.Printf("couldn't update patient, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "1", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}

// DeleteConsultationAssignment godoc
// @Summary Elimina una instancia de paciente
// @Description Método que permite eliminar una instancia del objeto paciente en la base de datos
// @tags Pacientes
// @Accept json
// @Produce json
// @Param	id	path int true "Paciente ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/paciente [DELETE]
func (h *handlerMedicalArea) DeleteConsultationAssignment(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	idStr := c.Params("id")
	if idStr == "" {
		logger.Error.Println("couldn't parse id request")
		res.Code, res.Type, res.Msg = 1, "", ""
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := medical_area.NewServerMedicalArea(h.db, user, h.txID)
	code, err := srv.SrvConsultationAssignment.DeleteConsultationAssignment(idStr)
	if err != nil {
		logger.Error.Printf("couldn't delete patient, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", ""
	return c.Status(http.StatusOK).JSON(res)
}

// GetConsultationAssignmentByID godoc
// @Summary Obtiene una instancia de paciente por su id
// @Description Método que permite obtener una instancia del objeto paciente en la base de datos por su id
// @tags Pacientes
// @Accept json
// @Produce json
// @Param	id	path int true "Paciente ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/paciente/:id [GET]
func (h *handlerMedicalArea) GetConsultationAssignmentByID(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	idStr := c.Params("id")
	if idStr == "" {
		logger.Error.Println("couldn't parse id request")
		res.Code, res.Type, res.Msg = 1, "", ""
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := medical_area.NewServerMedicalArea(h.db, user, h.txID)
	data, code, err := srv.SrvConsultationAssignment.GetConsultationAssignmentByID(idStr)
	if err != nil {
		logger.Error.Printf("couldn't get patient by id, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", ""
	return c.Status(http.StatusOK).JSON(res)
}

// GetAllConsultationAssignment godoc
// @Summary Obtiene todas las instancias de pacientes
// @Description Método que permite obtener todas las instancias del objeto paciente en la base de datos
// @tags Pacientes
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/ConsultationAssignment [GET]
func (h *handlerMedicalArea) GetAllConsultationAssignment(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	srv := medical_area.NewServerMedicalArea(h.db, user, h.txID)
	data, err := srv.SrvConsultationAssignment.GetAllConsultationAssignment()
	if err != nil {
		logger.Error.Printf("couldn't get all patient, error: %v", err)
		res.Code, res.Type, res.Msg = 23, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", ""
	return c.Status(http.StatusOK).JSON(res)
}
