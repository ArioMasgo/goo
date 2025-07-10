package visita_general

import (
	"dbu-api/internal/authorization"
	"dbu-api/internal/logger"
	"dbu-api/internal/middleware"
	"dbu-api/internal/models"
	"dbu-api/pkg/general_visit"
	"dbu-api/pkg/orchestrator/response_messages"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type handlerVisitaGeneral struct {
	db   *sqlx.DB
	txID string
	msg  response_messages.Message
}

// CreateVisitaGeneral godoc
// @Summary Crear una instancia de Visita General
// @Description Método que permite crear una instancia del objeto Visita General en la base de datos
// @Tags Visita General
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param VisitaGeneral body models.VisitaGeneral true "Datos para crear Visita General"
// @Success 201 {object} models.Response{error=boolean,data=models.VisitaGeneral,code=integer,type=string,msg=string} "Visita General creada exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/visita-general [POST]
func (h *handlerVisitaGeneral) CreateVisitaGeneral(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := models.VisitaGeneral{}

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

	err = c.BodyParser(&req)
	if err != nil {
		logger.Error.Printf("Error parsing request body: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	service := general_visit.NewGeneralVisitService(h.db, user, h.txID)
	visitaGeneral, code, err := service.VisitaGeneral().CreateVisitaGeneral(
		req.ID,
		req.TipoUsuario,
		req.CodigoEstudiante,
		req.DNI,
		req.NombreCompleto,
		req.Genero,
		req.Edad,
		req.Escuela,
		req.Area,
		req.MotivoAtencion,
		req.DescripcionMotivo,
		req.URLImagen,
		req.Departamento,
		req.Provincia,
		req.Distrito,
		req.LugarAtencion,
	)

	if err != nil {
		logger.Error.Printf("Error creating visita general: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = visitaGeneral
	res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
	return c.Status(fiber.StatusCreated).JSON(res)
}

// GetAllVisitaGeneral godoc
// @Summary Obtener todas las Visitas Generales
// @Description Método que permite obtener todas las Visitas Generales registradas en la base de datos
// @Tags Visita General
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Response{error=boolean,data=[]models.VisitaGeneral,code=integer,type=string,msg=string} "Visitas Generales obtenidas exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/visita-general [GET]
func (h *handlerVisitaGeneral) GetAllVisitaGeneral(c *fiber.Ctx) error {
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

	service := general_visit.NewGeneralVisitService(h.db, user, h.txID)
	visitasGenerales, err := service.VisitaGeneral().GetAllVisitaGeneral()

	if err != nil {
		logger.Error.Printf("Error getting all visitas generales: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(3)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = visitasGenerales
	res.Code, res.Type, res.Msg = h.msg.GetByCode(29)
	return c.Status(fiber.StatusOK).JSON(res)
}

// UpdateVisitaGeneral godoc
// @Summary Actualizar una Visita General
// @Description Método que permite actualizar una instancia del objeto Visita General en la base de datos
// @Tags Visita General
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param VisitaGeneral body models.VisitaGeneral true "Datos para actualizar Visita General"
// @Success 200 {object} models.Response{error=boolean,data=models.VisitaGeneral,code=integer,type=string,msg=string} "Visita General actualizada exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/visita-general [PUT]
func (h *handlerVisitaGeneral) UpdateVisitaGeneral(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := models.VisitaGeneral{}

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

	err = c.BodyParser(&req)
	if err != nil {
		logger.Error.Printf("Error parsing request body: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	if req.ID == "" {
		logger.Error.Printf("ID must not be empty")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	service := general_visit.NewGeneralVisitService(h.db, user, h.txID)
	visitaGeneral, code, err := service.VisitaGeneral().UpdateVisitaGeneral(
		req.ID,
		req.TipoUsuario,
		req.CodigoEstudiante,
		req.DNI,
		req.NombreCompleto,
		req.Genero,
		req.Edad,
		req.Escuela,
		req.Area,
		req.MotivoAtencion,
		req.DescripcionMotivo,
		req.URLImagen,
		req.Departamento,
		req.Provincia,
		req.Distrito,
		req.LugarAtencion,
	)

	if err != nil {
		logger.Error.Printf("Error updating visita general: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = visitaGeneral
	res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
	return c.Status(fiber.StatusOK).JSON(res)
}

// GetVisitaGeneralByID godoc
// @Summary Obtener una Visita General por ID
// @Description Método que permite obtener una instancia del objeto Visita General por su ID
// @Tags Visita General
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID de la Visita General a obtener"
// @Success 200 {object} models.Response{error=boolean,data=models.VisitaGeneral,code=integer,type=string,msg=string} "Visita General obtenida exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/visita-general/{id} [GET]
func (h *handlerVisitaGeneral) GetVisitaGeneralByID(c *fiber.Ctx) error {
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

	id := c.Params("id")
	if id == "" {
		logger.Error.Printf("ID must not be empty")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	service := general_visit.NewGeneralVisitService(h.db, user, h.txID)
	visitaGeneral, code, err := service.VisitaGeneral().GetVisitaGeneralByID(id)

	if err != nil {
		logger.Error.Printf("Error getting visita general by ID: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = visitaGeneral
	res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
	return c.Status(fiber.StatusOK).JSON(res)
}

// DeleteVisitaGeneral godoc
// @Summary Eliminar una Visita General
// @Description Método que permite eliminar una instancia del objeto Visita General de la base de datos
// @Tags Visita General
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID de la Visita General a eliminar"
// @Success 200 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Visita General eliminada exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/visita-general/{id} [DELETE]
func (h *handlerVisitaGeneral) DeleteVisitaGeneral(c *fiber.Ctx) error {
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

	id := c.Params("id")
	if id == "" {
		logger.Error.Printf("ID must not be empty")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	service := general_visit.NewGeneralVisitService(h.db, user, h.txID)
	code, err := service.VisitaGeneral().DeleteVisitaGeneral(id)

	if err != nil {
		logger.Error.Printf("Error deleting visita general: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
	return c.Status(fiber.StatusOK).JSON(res)
}

// GetAllDepartments godoc
// @Summary Obtener todos los departamentos
// @Description Método que permite obtener todos los departamentos disponibles
// @Tags Visita General - Ubicación
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Response{error=boolean,data=[]models.Departamento,code=integer,type=string,msg=string} "Departamentos obtenidos exitosamente"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/visita-general/departamentos [GET]
func (h *handlerVisitaGeneral) GetAllDepartments(c *fiber.Ctx) error {
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

	service := general_visit.NewGeneralVisitService(h.db, user, h.txID)
	departments, err := service.VisitaGeneral().GetAllDepartments()

	if err != nil {
		logger.Error.Printf("Error getting departments: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(3)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = departments
	res.Code, res.Type, res.Msg = h.msg.GetByCode(200)
	return c.Status(fiber.StatusOK).JSON(res)
}

// GetProvincesByDepartment godoc
// @Summary Obtener provincias por departamento
// @Description Método que permite obtener todas las provincias de un departamento específico
// @Tags Visita General - Ubicación
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param departmentId path string true "ID del departamento"
// @Success 200 {object} models.Response{error=boolean,data=[]models.Provincia,code=integer,type=string,msg=string} "Provincias obtenidas exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/visita-general/departamentos/{departmentId}/provincias [GET]
func (h *handlerVisitaGeneral) GetProvincesByDepartment(c *fiber.Ctx) error {
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

	departmentID := c.Params("departmentId")
	if departmentID == "" {
		logger.Error.Printf("Department ID must not be empty")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	service := general_visit.NewGeneralVisitService(h.db, user, h.txID)
	provinces, err := service.VisitaGeneral().GetProvincesByDepartment(departmentID)

	if err != nil {
		logger.Error.Printf("Error getting provinces: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(3)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = provinces
	res.Code, res.Type, res.Msg = h.msg.GetByCode(200)
	return c.Status(fiber.StatusOK).JSON(res)
}

// GetDistrictsByProvince godoc
// @Summary Obtener distritos por provincia
// @Description Método que permite obtener todos los distritos de una provincia específica
// @Tags Visita General - Ubicación
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param provinceId path string true "ID de la provincia"
// @Success 200 {object} models.Response{error=boolean,data=[]models.Distrito,code=integer,type=string,msg=string} "Distritos obtenidos exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/visita-general/provincias/{provinceId}/distritos [GET]
func (h *handlerVisitaGeneral) GetDistrictsByProvince(c *fiber.Ctx) error {
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

	provinceID := c.Params("provinceId")
	if provinceID == "" {
		logger.Error.Printf("Province ID must not be empty")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	service := general_visit.NewGeneralVisitService(h.db, user, h.txID)
	districts, err := service.VisitaGeneral().GetDistrictsByProvince(provinceID)

	if err != nil {
		logger.Error.Printf("Error getting districts: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(3)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = districts
	res.Code, res.Type, res.Msg = h.msg.GetByCode(200)
	return c.Status(fiber.StatusOK).JSON(res)
}

// GetLocationHierarchy godoc
// @Summary Obtener jerarquía completa de ubicaciones
// @Description Método que permite obtener toda la jerarquía de ubicaciones (departamentos con sus provincias y distritos)
// @Tags Visita General - Ubicación
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Response{error=boolean,data=models.LocationResponse,code=integer,type=string,msg=string} "Jerarquía de ubicaciones obtenida exitosamente"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/visita-general/ubicaciones [GET]
func (h *handlerVisitaGeneral) GetLocationHierarchy(c *fiber.Ctx) error {
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

	service := general_visit.NewGeneralVisitService(h.db, user, h.txID)
	hierarchy, err := service.VisitaGeneral().GetLocationHierarchy()

	if err != nil {
		logger.Error.Printf("Error getting location hierarchy: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(3)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = hierarchy
	res.Code, res.Type, res.Msg = h.msg.GetByCode(200)
	return c.Status(fiber.StatusOK).JSON(res)
}
