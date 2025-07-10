package medical_consultation

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/middleware"
	"dbu-api/internal/models"
	"dbu-api/pkg/orchestrator/low_code_medical_area"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerMedicalArea struct {
	db   *sqlx.DB
	txID string
}

// CreateMedicalConsultation godoc
// @Summary Crear una instancia de consulta medicina
// @Description Método que permite crear una instancia del objeto consulta medicina en la base de datos
// @tags consulta medicinas
// @Accept json
// @Produce json
// @Param models.RequestMedicalConsultation body models.RequestMedicalConsultation true "Datos para crear MedicalConsultation"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/consulta_medicina [POST]
func (h *handlerMedicalArea) CreateMedicalConsultation(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := models.RequestMedicalConsultation{}

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

	isValid, err := req.ValidMedicalConsultation()
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

	srv := low_code_medical_area.NewMedicalConsultation(h.db, user, h.txID)
	code, err := srv.CreateMedicalConsultationLowCode(&req)
	if err != nil {
		logger.Error.Printf("couldn't create medical consultation, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", ""
	return c.Status(http.StatusCreated).JSON(res)
}

// UpdateMedicalConsultation godoc
// @Summary Actualiza una instancia de consulta medicina
// @Description Método que permite Actualiza una instancia del objeto consulta medicina en la base de datos
// @tags consulta medicinas
// @Accept json
// @Produce json
// @Param models.RequestMedicalConsultation body models.RequestMedicalConsultation true "Datos para actualizar MedicalConsultation"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/consulta_medicina [PUT]
func (h *handlerMedicalArea) UpdateMedicalConsultation(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := models.RequestMedicalConsultation{}

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

	isValid, err := req.ValidMedicalConsultation()
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

	srv := low_code_medical_area.NewMedicalConsultation(h.db, user, h.txID)
	code, err := srv.UpdateMedicalConsultationLowCode(&req)
	if err != nil {
		logger.Error.Printf("couldn't update medical consultation, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", ""
	return c.Status(http.StatusOK).JSON(res)
}

// DeleteMedicalConsultation godoc
// @Summary Elimina una instancia de consulta medicina
// @Description Método que permite eliminar una instancia del objeto consulta medicina en la base de datos
// @tags consulta medicinas
// @Accept json
// @Produce json
// @Param	id	path int true "consulta medicina ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/consulta_medicina/:id [DELETE]
func (h *handlerMedicalArea) DeleteMedicalConsultation(c *fiber.Ctx) error {
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

	srv := low_code_medical_area.NewMedicalConsultation(h.db, user, h.txID)
	code, err := srv.DeleteMedicalConsultationLowCode(idStr)
	if err != nil {
		logger.Error.Printf("couldn't delete medical consultation, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", ""
	return c.Status(http.StatusOK).JSON(res)
}

// GetMedicalConsultationByID godoc
// @Summary Obtiene una instancia de consulta medicina por su id
// @Description Método que permite obtener una instancia del objeto consulta medicina en la base de datos por su id
// @tags consulta medicinas
// @Accept json
// @Produce json
// @Param	id	path int true "consulta medicina ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/consulta_medicina/:id [GET]
func (h *handlerMedicalArea) GetMedicalConsultationByID(c *fiber.Ctx) error {
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

	srv := low_code_medical_area.NewMedicalConsultation(h.db, user, h.txID)
	data, code, err := srv.GetMedicalConsultationByIdConsultationLowCode(idStr)
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

// GetMedicalConsultationByID godoc
// @Summary Obtiene una instancia de consulta medicina por su id
// @Description Método que permite obtener una instancia del objeto consulta medicina en la base de datos por su id
// @tags consulta medicinas
// @Accept json
// @Produce json
// @Param	id	path int true "consulta medicina ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/consultas_medicina/paciente/:id [GET]
func (h *handlerMedicalArea) GetMedicalConsultationByIDPatient(c *fiber.Ctx) error {
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

	srv := low_code_medical_area.NewMedicalConsultation(h.db, user, h.txID)
	data, code, err := srv.GetMedicalConsultationByIdPatientLowCode(idStr)
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

func (h *handlerMedicalArea) GetMedicalConsultationByDNIPatient(c *fiber.Ctx) error {
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

	srv := low_code_medical_area.NewMedicalConsultation(h.db, user, h.txID)
	data, code, err := srv.GetMedicalConsultationByDNILowCode(idStr)
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

// GetAllMedicalConsultation godoc
// @Summary Obtiene todas las instancias de MedicalConsultation
// @Description Método que permite obtener todas las instancias del objeto consulta medicina en la base de datos
// @tags consulta medicinas
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/consultas_medicina [GET]
func (h *handlerMedicalArea) GetAllMedicalConsultation(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	srv := low_code_medical_area.NewMedicalConsultation(h.db, user, h.txID)
	data, err := srv.GetAllMedicalConsultationLowCode()
	if err != nil {
		logger.Error.Printf("couldn't get all medical consultation, error: %v", err)
		res.Code, res.Type, res.Msg = 23, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", ""
	return c.Status(http.StatusOK).JSON(res)
}
