package dentistry_consultation

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

// CreateDentistryConsultation godoc
// @Summary Crear una instancia de consulta odontología
// @Description Método que permite crear una instancia del objeto consulta odontología en la base de datos
// @tags Consultas odontología
// @Accept json
// @Produce json
// @Param models.RequestDentistryConsultation body models.RequestDentistryConsultation true "Datos para crear consulta odontología"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/consulta_odontologia [POST]
func (h *handlerMedicalArea) CreateDentistryConsultation(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := models.RequestDentistryConsultation{}

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

	isValid, err := req.ValidDentistryConsultation()
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

	srv := low_code_medical_area.NewDentistryConsultation(h.db, user, h.txID)
	code, err := srv.CreateDentistryConsultationLowCode(&req)
	if err != nil {
		logger.Error.Printf("couldn't create dentistry consultation, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "1", "Procesado correctamente"
	return c.Status(http.StatusCreated).JSON(res)
}

// UpdateDentistryConsultation godoc
// @Summary Actualiza una instancia de consulta odontología
// @Description Método que permite actualizar una instancia del objeto consulta odontología en la base de datos
// @tags Consultas odontología
// @Accept json
// @Produce json
// @Param models.RequestDentistryConsultation body models.RequestDentistryConsultation true "Datos para actualizar consulta enfermería"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/consulta_odontologia [PUT]
func (h *handlerMedicalArea) UpdateDentistryConsultation(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := models.RequestDentistryConsultation{}

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

	isValid, err := req.ValidDentistryConsultation()
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

	srv := low_code_medical_area.NewDentistryConsultation(h.db, user, h.txID)
	code, err := srv.UpdateDentistryConsultationLowCode(&req)
	if err != nil {
		logger.Error.Printf("couldn't update dentistry consultation, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "1", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}

// DeleteDentistryConsultation godoc
// @Summary Elimina una instancia de consulta odontología
// @Description Método que permite eliminar una instancia del objeto consulta odontología en la base de datos
// @tags Consultas odontología
// @Accept json
// @Produce json
// @Param	id	path int true "consulta odontología ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/consulta_odontologia/:id [DELETE]
func (h *handlerMedicalArea) DeleteDentistryConsultation(c *fiber.Ctx) error {
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

	srv := low_code_medical_area.NewDentistryConsultation(h.db, user, h.txID)
	code, err := srv.DeleteDentistryConsultationLowCode(idStr)
	if err != nil {
		logger.Error.Printf("couldn't delete dentistry consultation, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = 29, "1", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}

// GetDentistryConsultationByID godoc
// @Summary Obtiene una instancia de consulta odontología por su id
// @Description Método que permite obtener una instancia del objeto consulta odontología en la base de datos por su id
// @tags Consultas odontología
// @Accept json
// @Produce json
// @Param	id	path uuid true "consulta odontología ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/consulta_odontologia/:id [GET]
func (h *handlerMedicalArea) GetDentistryConsultationByID(c *fiber.Ctx) error {
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

	srv := low_code_medical_area.NewDentistryConsultation(h.db, user, h.txID)
	data, code, err := srv.GetDentistryConsultationByIdConsultationLowCode(idStr)
	if err != nil {
		logger.Error.Printf("couldn't get dentistry consultation by id, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", ""
	return c.Status(http.StatusOK).JSON(res)
}

// GetDentistryConsultationByID godoc
// @Summary Obtiene una instancia de consulta odontología por su id
// @Description Método que permite obtener todas las instancias del objeto consulta odontología en la base de datos por su id de paciente
// @tags Consultas odontología
// @Accept json
// @Produce json
// @Param	id	path int true "consulta odontología ID paciente"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/consultas_odontologia/paciente/:id [GET]
func (h *handlerMedicalArea) GetDentistryConsultationByIDPatient(c *fiber.Ctx) error {
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

	srv := low_code_medical_area.NewDentistryConsultation(h.db, user, h.txID)
	data, code, err := srv.GetDentistryConsultationByIdPatientLowCode(idStr)
	if err != nil {
		logger.Error.Printf("couldn't get dentistry consultation by id, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "1", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerMedicalArea) GetDentistryConsultationByDNIPatient(c *fiber.Ctx) error {
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

	srv := low_code_medical_area.NewDentistryConsultation(h.db, user, h.txID)
	data, code, err := srv.GetDentistryConsultationByDNILowCode(idStr)
	if err != nil {
		logger.Error.Printf("couldn't get dentistry consultation by id, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "1", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}

// GetAllDentistryConsultation godoc
// @Summary Obtiene todas las instancias de DentistryConsultation
// @Description Método que permite obtener todas las instancias del objeto consulta odontología en la base de datos
// @tags Consultas odontología
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/consultas_odontologia [GET]
func (h *handlerMedicalArea) GetAllDentistryConsultation(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	srv := low_code_medical_area.NewDentistryConsultation(h.db, user, h.txID)
	data, err := srv.GetAllDentistryConsultationLowCode()
	if err != nil {
		logger.Error.Printf("couldn't get all dentistry consultation, error: %v", err)
		res.Code, res.Type, res.Msg = 23, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "1", "Procesado correctamente"
	return c.Status(http.StatusOK).JSON(res)
}
