package nursing_consultation

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/middleware"
	"dbu-api/internal/models"
	"dbu-api/pkg/medical_area"
	"dbu-api/pkg/orchestrator/low_code_medical_area"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerMedicalArea struct {
	db   *sqlx.DB
	txID string
}

// CreateNursingConsultation godoc
// @Summary Crear una instancia de consulta enfermería
// @Description Método que permite crear una instancia del objeto consulta enfermería en la base de datos
// @tags Consulta enfermería
// @Accept json
// @Produce json
// @Param models.RequestNursingConsultation body models.RequestNursingConsultation true "Datos para crear una consulta enfermería"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/consulta_enfermeria [POST]
func (h *handlerMedicalArea) CreateNursingConsultation(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := models.RequestNursingConsultation{}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	if err := c.BodyParser(&req); err != nil {
		logger.Error.Printf("couldn't parse body request, error: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", err.Error()
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	isValid, err := req.ValidNursingConsultation()
	if err != nil {
		logger.Error.Printf("couldn't validate body request, error: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", err.Error()
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if !isValid {
		logger.Error.Println("couldn't validate body request")
		res.Code, res.Type, res.Msg = 1, "", err.Error()
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := low_code_medical_area.NewNursingConsultation(h.db, user, h.txID)
	code, err := srv.CreateNursingConsultationLowCode(&req, false)
	if err != nil {
		logger.Error.Printf("couldn't create nursing consultation, error: %v", err)
		res.Code, res.Type, res.Msg = code, "error", err.Error()
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", "Procesado correctamente"
	return c.Status(http.StatusCreated).JSON(res)
}

// UpdateNursingConsultation godoc
// @Summary Crear una instancia de consulta enfermería
// @Description Método que permite actualizar una instancia del objeto consulta enfermería en la base de datos
// @tags Consulta enfermería
// @Accept json
// @Produce json
// @Param models.RequestNursingConsultation body models.RequestNursingConsultation true "Datos para actualizar una consulta enfermería"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/consulta_enfermeria [PUT]
func (h *handlerMedicalArea) UpdateNursingConsultation(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := models.RequestNursingConsultation{}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	if err := c.BodyParser(&req); err != nil {
		logger.Error.Printf("couldn't parse body request, error: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", err.Error()
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	isValid, err := req.ValidNursingConsultation()
	if err != nil {
		logger.Error.Printf("couldn't validate body request, error: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", err.Error()
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if !isValid {
		logger.Error.Println("couldn't validate body request")
		res.Code, res.Type, res.Msg = 1, "error", err.Error()
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := low_code_medical_area.NewNursingConsultation(h.db, user, h.txID)
	code, err := srv.UpdateNursingConsultationLowCode(&req)
	if err != nil {
		logger.Error.Printf("couldn't update nursing consultation, error: %v", err)
		res.Code, res.Type, res.Msg = code, "error", err.Error()
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}

// DeleteNursingConsultation godoc
// @Summary Elimina una instancia de consulta enfermería
// @Description Método que permite eliminar una instancia del objeto consulta enfermería en la base de datos
// @tags Consulta enfermería
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/consulta_enfermeria/:id [DELETE]
func (h *handlerMedicalArea) DeleteNursingConsultation(c *fiber.Ctx) error {
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
		res.Code, res.Type, res.Msg = 1, "error", err.Error()
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := low_code_medical_area.NewNursingConsultation(h.db, user, h.txID)
	code, err := srv.DeleteNursingConsultationLowCode(idStr, false)
	if err != nil {
		logger.Error.Printf("couldn't delete nursing consultation, error: %v", err)
		res.Code, res.Type, res.Msg = code, "error", err.Error()
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}

// GetNursingConsultationByID godoc
// @Summary Obtiene una instancia de consulta enfermería por su id
// @Description Método que permite obtener una instancia del objeto consulta enfermería en la base de datos por su id
// @tags Consulta enfermería
// @Accept json
// @Produce json
// @Param	id	path int true "consulta enfermería ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/consulta_enfermeria/:id [GET]
func (h *handlerMedicalArea) GetNursingConsultationByID(c *fiber.Ctx) error {
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
		res.Code, res.Type, res.Msg = 1, "error", err.Error()
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := low_code_medical_area.NewNursingConsultation(h.db, user, h.txID)
	data, code, err := srv.GetNursingConsultationByIdConsultationLowCode(idStr)
	if err != nil {
		logger.Error.Printf("couldn't get nursing consultation by id, error: %v", err)
		res.Code, res.Type, res.Msg = code, "error", err.Error()
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}

// GetAllNursingConsultation godoc
// @Summary Obtiene todas las instancias de NursingConsultation
// @Description Método que permite obtener todas las instancias del objeto consulta enfermería en la base de datos
// @tags Consulta enfermería
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/consultas_enfermeria [GET]
func (h *handlerMedicalArea) GetAllNursingConsultation(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	srv := low_code_medical_area.NewNursingConsultation(h.db, user, h.txID)
	data, err := srv.GetAllNursingConsultationLowCode()
	if err != nil {
		logger.Error.Printf("couldn't get all nursing consultation, error: %v", err)
		res.Code, res.Type, res.Msg = 23, "error", err.Error()
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}

// GetAllNursingConsultation godoc
// @Summary Obtiene todas las instancias de NursingConsultation
// @Description Método que permite obtener todas las instancias del objeto consulta enfermería en la base de datos por id del paciente
// @tags Consulta enfermería
// @Accept json
// @Produce json
// @Param	id	path int true "consulta enfermería ID paciente"
// @Success 200 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/consultas_enfermeria/paciente/:id [GET]
func (h *handlerMedicalArea) GetNursingConsultationByIDPatient(c *fiber.Ctx) error {
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
		res.Code, res.Type, res.Msg = 1, "error", err.Error()
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := low_code_medical_area.NewNursingConsultation(h.db, user, h.txID)
	data, code, err := srv.GetNursingConsultationByIdPatientLowCode(idStr)
	if err != nil {
		logger.Error.Printf("couldn't get nursing consultation by id, error: %v", err)
		res.Code, res.Type, res.Msg = code, "error", err.Error()
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerMedicalArea) GetNursingConsultationByDNIPatient(c *fiber.Ctx) error {
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
		res.Code, res.Type, res.Msg = 1, "error", err.Error()
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := low_code_medical_area.NewNursingConsultation(h.db, user, h.txID)
	data, code, err := srv.GetNursingConsultationByDNILowCode(idStr)
	if err != nil {
		logger.Error.Printf("couldn't get nursing consultation by id, error: %v", err)
		res.Code, res.Type, res.Msg = code, "error", err.Error()
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerMedicalArea) GetAllTypesVaccines(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	srv := low_code_medical_area.NewNursingConsultation(h.db, user, h.txID)
	data, err := srv.GetAllTypesVaccinesLowCode()
	if err != nil {
		logger.Error.Printf("couldn't get all nursing consultation, error: %v", err)
		res.Code, res.Type, res.Msg = 23, "error", err.Error()
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerMedicalArea) GetTypesVaccineRequired(c *fiber.Ctx) error {
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
		res.Code, res.Type, res.Msg = 1, "error", err.Error()
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := low_code_medical_area.NewNursingConsultation(h.db, user, h.txID)
	data, code, err := srv.GetTypesVaccineRequiredLowCode(idStr)
	if err != nil {
		logger.Error.Printf("couldn't get type vaccine required by id, error: %v", err)
		res.Code, res.Type, res.Msg = code, "error", err.Error()
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerMedicalArea) GetAllVaccinesByPatientDni(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	dni := c.Params("dni")
	if dni == "" {
		logger.Error.Println("couldn't parse dni request")
		res.Code, res.Type, res.Msg = 1, "error", err.Error()
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := medical_area.NewServerMedicalArea(h.db, user, h.txID)
	data, err := srv.SrvNursingConsultationVaccine.GetAllVaccinesByPatientDni(dni)
	if err != nil {
		logger.Error.Printf("couldn't get type vaccine required by id, error: %v", err)
		res.Code, res.Type, res.Msg = 0, "error", err.Error()
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}
