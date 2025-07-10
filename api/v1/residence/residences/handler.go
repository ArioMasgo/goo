package residences

import (
	"dbu-api/internal/authorization"
	"dbu-api/internal/logger"
	"dbu-api/internal/middleware"
	"dbu-api/internal/models"
	"dbu-api/pkg/orchestrator/low_code_residences"
	"dbu-api/pkg/orchestrator/response_messages"
	"dbu-api/pkg/residence"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type handlerResidences struct {
	db   *sqlx.DB
	txID string
	msg  response_messages.Message
}

// CreateResidences godoc
// @Summary Crear una instancia de Residencias
// @Description Método que permite crear una instancia del objeto Residencias en la base de datos
// @Tags Residencias
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Residence body models.Residence true "Datos para crear Residencias"
// @Success 201 {object} models.Response{error=boolean,data=models.Residence,code=integer,type=string,msg=string} "Residencia creada exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/residencias [POST]
func (h *handlerResidences) CreateResidences(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := models.Residence{}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(53)
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	err = authorization.ValidPermissions(user, h.db, c)
	if err != nil {
		logger.Error.Printf("User does not have permission to call the api, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(10)
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	if err := c.BodyParser(&req); err != nil {
		logger.Error.Printf("couldn't parse body request, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	isValid, err := req.ValidResidence()
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

	srv := low_code_residences.NewResidence(h.db, user, h.txID)
	code, err := srv.CreateResidenceLowCode(&req)
	if err != nil {
		logger.Error.Printf("couldn't create Residencias, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false
	res.Code, res.Type, res.Msg = h.msg.GetByCode(211)
	return c.Status(fiber.StatusCreated).JSON(res)
}

// UpdateResidences godoc
// @Summary Actualizar una instancia de Residencias
// @Description Método que permite actualizar una instancia del objeto Residencias en la base de datos
// @Tags Residencias
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Residence body models.Residence true "Datos para actualizar Residencias"
// @Success 200 {object} models.Response{error=boolean,data=models.Residence,code=integer,type=string,msg=string} "Residencia actualizada exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/residencias [PUT]
func (h *handlerResidences) UpdateResidences(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := models.Residence{}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(53)
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	err = authorization.ValidPermissions(user, h.db, c)
	if err != nil {
		logger.Error.Printf("User does not have permission to call the api, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(10)
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	if err := c.BodyParser(&req); err != nil {
		logger.Error.Printf("couldn't parse body request, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	isValid, err := req.ValidResidence()
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

	srv := residence.NewServerResidence(h.db, user, h.txID)
	_, code, err := srv.SrvResidence.UpdateResidence(req.ID, req.Name, req.Description, req.Gender, req.Address, req.Status)
	if err != nil {
		logger.Error.Printf("couldn't update Residencias, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false
	res.Code, res.Type, res.Msg = h.msg.GetByCode(212)
	return c.Status(fiber.StatusOK).JSON(res)
}

// DeleteResidences godoc
// @Summary Eliminar una instancia de Residencias
// @Description Método que permite eliminar una instancia del objeto Residencias en la base de datos
// @Tags Residencias
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID de la Residencia" format(uuid)
// @Success 200 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Residencia eliminada exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/residencias [DELETE]
func (h *handlerResidences) DeleteResidences(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(53)
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	err = authorization.ValidPermissions(user, h.db, c)
	if err != nil {
		logger.Error.Printf("User does not have permission to call the api, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(10)
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	idStr := c.Params("id")
	if idStr == "" {
		logger.Error.Println("couldn't parse id request")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	srv := residence.NewServerResidence(h.db, nil, h.txID)
	code, err := srv.SrvResidence.DeleteResidence(idStr)
	if err != nil {
		logger.Error.Printf("couldn't delete Residencias, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = h.msg.GetByCode(213)
	return c.Status(fiber.StatusOK).JSON(res)
}

// GetResidenciasByID godoc
// @Summary Obtener una instancia de Residencias por ID
// @Description Método que permite obtener una instancia del objeto Residencias por su ID
// @Tags Residencias
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID de la Residencia" format(uuid)
// @Success 200 {object} models.Response{error=boolean,data=models.Residence,code=integer,type=string,msg=string} "Residencia obtenida exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/residencias/:id [GET]
func (h *handlerResidences) GetResidenciasByID(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(53)
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	err = authorization.ValidPermissions(user, h.db, c)
	if err != nil {
		logger.Error.Printf("User does not have permission to call the api, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(10)
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	idStr := c.Params("id")
	if idStr == "" {
		logger.Error.Println("couldn't parse id request")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	srv := residence.NewServerResidence(h.db, nil, h.txID)
	data, code, err := srv.SrvResidence.GetResidenceByID(idStr)
	if err != nil {
		logger.Error.Printf("couldn't get Residencias by id, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = h.msg.GetByCode(210)
	return c.Status(fiber.StatusOK).JSON(res)
}

// GetAllResidences godoc
// @Summary Obtener todas las instancias de Residencias
// @Description Método que permite obtener todas las instancias del objeto Residencias
// @Tags Residencias
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Response{error=boolean,data=[]models.Residence,code=integer,type=string,msg=string} "Residencias obtenidas exitosamente"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/residencias [GET]
func (h *handlerResidences) GetAllResidences(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(53)
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	err = authorization.ValidPermissions(user, h.db, c)
	if err != nil {
		logger.Error.Printf("User does not have permission to call the api, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(10)
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	srv := low_code_residences.NewResidence(h.db, user, h.txID)
	residences, code, err := srv.GetResidenceLowCode()
	if err != nil {
		logger.Error.Printf("couldn't get all Residencias, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusAccepted).JSON(res)
	}

	res.Data = residences
	res.Error = false
	res.Code, res.Type, res.Msg = h.msg.GetByCode(210)
	return c.Status(fiber.StatusOK).JSON(res)
}

// GetAllStudentsByResidence godoc
// @Summary Obtener todos los alumnos de una Residencia
// @Description Método que permite obtener todos los alumnos asociados a una Residencia
// @Tags Residencias
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID de la Residencia" format(uuid)
// @Param page query integer false "Número de página" minimum(1) default(1)
// @Param limit query integer false "Límite de registros por página" minimum(1) maximum(100) default(10)
// @Param submission_id query integer true "ID de la solicitud"
// @Success 200 {object} models.Response{error=boolean,data=[]models.Student,code=integer,type=string,msg=string} "Alumnos obtenidos exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/residencias/:id/alumnos [GET]
func (h *handlerResidences) GetAllStudentsByResidence(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(53)
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	err = authorization.ValidPermissions(user, h.db, c)
	if err != nil {
		logger.Error.Printf("User does not have permission to call the api, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(10)
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	residenceID := c.Params("id")
	if residenceID == "" {
		logger.Error.Printf("missing residence_id parameter")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(fiber.StatusBadRequest).JSON(res)
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

	filter := c.Query("filter")

	submissionHeader := c.QueryInt("submission_id", 0)
	if submissionHeader == 0 {
		logger.Error.Printf("missing submission_id parameter")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	srv := low_code_residences.NewResidence(h.db, user, h.txID)
	students, code, total, err := srv.GetStudentsByResidenceLowCode(residenceID, int64(submissionHeader), page, limit, filter)
	if err != nil {
		logger.Error.Printf("error getting students: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res.Data = ResponseStudentsResidence{
		Students: students,
		Total:    total,
	}

	res.Error = false
	res.Code, res.Type, res.Msg = h.msg.GetByCode(210)
	return c.Status(fiber.StatusOK).JSON(res)
}

// UpdateConfigResidence godoc
// @Summary Actualizar configuración de una Residencia
// @Description Método que permite actualizar la configuración de una residencia específica
// @Tags Residencias
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID de la Residencia" format(uuid)
// @Param body body models.Configuration true "Datos de configuración de la residencia"
// @Success 200 {object} models.Response{error=boolean,data=models.Configuration,code=integer,type=string,msg=string} "Configuración actualizada exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/residencias/{id}/configuracion [PUT]
func (h *handlerResidences) UpdateConfigResidence(c *fiber.Ctx) error {
	req := models.Configuration{}
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(53)
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	err = authorization.ValidPermissions(user, h.db, c)
	if err != nil {
		logger.Error.Printf("User does not have permission to call the api, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(10)
		return c.Status(fiber.StatusUnauthorized).JSON(res)
	}

	if err := c.BodyParser(&req); err != nil {
		logger.Error.Printf("couldn't parse body request, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	isValid, err := req.ValidConfig()
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

	residenceID := c.Params("id")
	if residenceID == "" {
		logger.Error.Printf("missing residence_id parameter")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	srv := residence.NewServerResidence(h.db, user, h.txID)
	residenceData, code, err := srv.SrvResidence.GetResidenceByID(residenceID)
	if err != nil {
		logger.Error.Printf("error getting residence: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	if residenceData == nil {
		logger.Error.Printf("no found residence")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(34)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	_, codErr, err := srv.SrvResidenceConfiguration.UpdateResidenceConfigurationByResidenceID(req.PercentageFcea,
		req.PercentageEngineering, req.MinimumGradeFcea, req.MinimumGradeEngineering, residenceData.ID, req.IsNewbie)

	if err != nil {
		logger.Error.Printf("error update residence config: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(codErr)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res.Data = req
	res.Error = false
	res.Code, res.Type, res.Msg = h.msg.GetByCode(212)
	return c.Status(fiber.StatusOK).JSON(res)
}
