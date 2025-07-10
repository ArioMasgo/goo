package patients

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

// CreatePatients godoc
// @Summary Crear una instancia de paciente
// @Description Método que permite crear una instancia del objeto paciente en la base de datos
// @tags Pacientes
// @Accept json
// @Produce json
// @Param models.Patients body models.Patients true "Datos para crear patients"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/paciente [POST]
func (h *handlerMedicalArea) CreatePatients(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := models.Patients{}

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

	isValid, err := req.ValidPatients()
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

	srv := low_code_medical_area.NewPatient(h.db, user, h.txID)
	code, err := srv.CreatePatientLowCode(&req)
	if err != nil {
		logger.Error.Printf("couldn't create patient, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", err.Error()
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Data = req
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", "Procesado correctamente"
	return c.Status(http.StatusCreated).JSON(res)
}

// UpdatePatients godoc
// @Summary Actualiza una instancia de paciente
// @Description Método que permite actualizar una instancia del objeto paciente en la base de datos
// @tags Pacientes
// @Accept json
// @Produce json
// @Param models.Patients body models.Patients true "Datos para actualizar patients"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/paciente [PUT]
func (h *handlerMedicalArea) UpdatePatients(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := models.Patients{}

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

	isValid, err := req.ValidPatients()
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

	srv := low_code_medical_area.NewPatient(h.db, user, h.txID)
	code, err := srv.UpdatePatientLowCode(&req)
	if err != nil {
		logger.Error.Printf("couldn't update patient, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", ""
	return c.Status(http.StatusOK).JSON(res)
}

// DeletePatients godoc
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
func (h *handlerMedicalArea) DeletePatients(c *fiber.Ctx) error {
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

	srv := low_code_medical_area.NewPatient(h.db, user, h.txID)
	code, err := srv.DeletePatientLowCode(idStr)
	if err != nil {
		logger.Error.Printf("couldn't delete patient, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", ""
	return c.Status(http.StatusOK).JSON(res)
}

// GetPatientsByID godoc
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
func (h *handlerMedicalArea) GetPatientsByID(c *fiber.Ctx) error {
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

	srv := low_code_medical_area.NewPatient(h.db, user, h.txID)
	data, code, err := srv.GetPatientByIdLowCode(idStr)
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

// GetAllPatients godoc
// @Summary Obtiene todas las instancias de pacientes
// @Description Método que permite obtener todas las instancias del objeto paciente en la base de datos
// @tags Pacientes
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/patients [GET]
func (h *handlerMedicalArea) GetAllPatients(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	srv := low_code_medical_area.NewPatient(h.db, user, h.txID)
	data, err := srv.GetAllPatientLowCode()
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

func (h *handlerMedicalArea) GetPatients(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	queryParams := &QueryParamsSearchPatients{
		Limit:  20,
		Offset: 0,
	}
	err = c.QueryParser(queryParams)
	if err != nil {
		logger.Error.Printf("bad request, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "bad request"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if queryParams.Offset < 0 || queryParams.Limit <= 0 {
		logger.Error.Printf("bad request, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "bad request"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	rq := new(RequestSearchPatients)
	err = c.BodyParser(&rq)
	if err != nil {
		logger.Error.Printf("bad request, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "bad request"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := low_code_medical_area.NewPatient(h.db, user, h.txID)
	data, err := srv.GetPatients(rq.Dni, rq.Nombres, rq.Apellidos, queryParams.Limit, queryParams.Offset)
	if err != nil {
		logger.Error.Printf("couldn't get all patient, error: %v", err)
		res.Code, res.Type, res.Msg = 23, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	metadata, err := srv.GetMetadata(rq.Dni, rq.Nombres, rq.Apellidos, queryParams.Limit, queryParams.Offset)
	if err != nil {
		logger.Error.Printf("couldn't get pagination patient, error: %v", err)
		res.Code, res.Type, res.Msg = 23, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Metadata = metadata
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", ""
	return c.Status(http.StatusOK).JSON(res)
}

// GetPatientsByID godoc
// @Summary Obtiene una instancia de paciente por su id
// @Description Método que permite obtener una instancia del objeto paciente en la base de datos por su DNI
// @tags Pacientes
// @Accept json
// @Produce json
// @Param	id	path int true "Paciente ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/paciente/:id [GET]
func (h *handlerMedicalArea) GetPatientsByDNI(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	dniStr := c.Params("dni")
	if dniStr == "" {
		logger.Error.Println("couldn't parse id request")
		res.Code, res.Type, res.Msg = 1, "", ""
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := low_code_medical_area.NewPatient(h.db, user, h.txID)
	data, code, err := srv.GetPatientByDNILowCode(dniStr)
	if err != nil {
		logger.Error.Printf("couldn't get patient by DNI, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", ""
	return c.Status(http.StatusOK).JSON(res)
}
