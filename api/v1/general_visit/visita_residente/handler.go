package visita_residente

import (
	"dbu-api/internal/authorization"
	"dbu-api/internal/logger"
	"dbu-api/internal/middleware"
	"dbu-api/internal/models"
	"dbu-api/pkg/general_visit"
	"dbu-api/pkg/general_visit/visita_residente"
	"dbu-api/pkg/orchestrator/response_messages"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type handlerVisitaResidente struct {
	db   *sqlx.DB
	txID string
	msg  response_messages.Message
}

// CreateVisitaResidente godoc
// @Summary Crear una instancia de Visita Domiciliaria
// @Description Método que permite crear una instancia del objeto Visita Domiciliaria en la base de datos
// @Tags Visita Residente
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param VisitaResidente body models.VisitaDomiciliaria true "Datos para crear Visita Domiciliaria"
// @Success 201 {object} models.Response{error=boolean,data=models.VisitaDomiciliaria,code=integer,type=string,msg=string} "Visita Domiciliaria creada exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/visita-residente [POST]
func (h *handlerVisitaResidente) CreateVisitaResidente(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := models.VisitaDomiciliaria{}

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
	visitaResidente, code, err := service.VisitaResidente().CreateVisitaResidente(
		fmt.Sprintf("%d", req.AlumnoID),
		req.Estado,
		func() string {
			if req.Comentario != nil {
				return *req.Comentario
			} else {
				return ""
			}
		}(),
		func() string {
			if req.ImagenURL != nil {
				return *req.ImagenURL
			} else {
				return ""
			}
		}(),
	)

	if err != nil {
		logger.Error.Printf("Error creating visita residente: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = visitaResidente
	res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
	return c.Status(fiber.StatusCreated).JSON(res)
}

// GetAllVisitaResidente godoc
// @Summary Obtener todas las Visitas Domiciliarias
// @Description Método que permite obtener todas las Visitas Domiciliarias registradas en la base de datos
// @Tags Visita Residente
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Response{error=boolean,data=[]models.VisitaDomiciliaria,code=integer,type=string,msg=string} "Visitas Domiciliarias obtenidas exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/visita-residente [GET]
func (h *handlerVisitaResidente) GetAllVisitaResidente(c *fiber.Ctx) error {
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
	visitasResidente, err := service.VisitaResidente().GetAllVisitaResidente()

	if err != nil {
		logger.Error.Printf("Error getting all visita residente: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(3)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = visitasResidente
	res.Code, res.Type, res.Msg = h.msg.GetByCode(29)
	return c.Status(fiber.StatusOK).JSON(res)
}

// GetVisitaResidenteByID godoc
// @Summary Obtener Visita Domiciliaria por ID
// @Description Método que permite obtener una Visita Domiciliaria específica por su ID
// @Tags Visita Residente
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID de la Visita Domiciliaria"
// @Success 200 {object} models.Response{error=boolean,data=models.VisitaDomiciliaria,code=integer,type=string,msg=string} "Visita Domiciliaria obtenida exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Failure 404 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Visita Domiciliaria no encontrada"
// @Router /v1/visita-residente/{id} [GET]
func (h *handlerVisitaResidente) GetVisitaResidenteByID(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	id := c.Params("id")

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
	visitaResidente, code, err := service.VisitaResidente().GetVisitaResidenteByID(id)

	if err != nil {
		logger.Error.Printf("Error getting visita residente by ID: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = visitaResidente
	res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
	return c.Status(fiber.StatusOK).JSON(res)
}

// UpdateVisitaResidente godoc
// @Summary Actualizar una Visita Domiciliaria
// @Description Método que permite actualizar una Visita Domiciliaria existente
// @Tags Visita Residente
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID de la Visita Domiciliaria"
// @Param VisitaResidente body models.VisitaDomiciliaria true "Datos para actualizar Visita Domiciliaria"
// @Success 200 {object} models.Response{error=boolean,data=models.VisitaDomiciliaria,code=integer,type=string,msg=string} "Visita Domiciliaria actualizada exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Failure 404 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Visita Domiciliaria no encontrada"
// @Router /v1/visita-residente/{id} [PUT]
func (h *handlerVisitaResidente) UpdateVisitaResidente(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := models.VisitaDomiciliaria{}
	id := c.Params("id")

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
	visitaResidente, code, err := service.VisitaResidente().UpdateVisitaResidente(
		id,
		fmt.Sprintf("%d", req.AlumnoID),
		req.Estado,
		func() string {
			if req.Comentario != nil {
				return *req.Comentario
			} else {
				return ""
			}
		}(),
		func() string {
			if req.ImagenURL != nil {
				return *req.ImagenURL
			} else {
				return ""
			}
		}(),
	)

	if err != nil {
		logger.Error.Printf("Error updating visita residente: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = visitaResidente
	res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
	return c.Status(fiber.StatusOK).JSON(res)
}

// DeleteVisitaResidente godoc
// @Summary Eliminar una Visita Domiciliaria
// @Description Método que permite eliminar una Visita Domiciliaria por su ID
// @Tags Visita Residente
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "ID de la Visita Domiciliaria"
// @Success 200 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Visita Domiciliaria eliminada exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Failure 404 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Visita Domiciliaria no encontrada"
// @Router /v1/visita-residente/{id} [DELETE]
func (h *handlerVisitaResidente) DeleteVisitaResidente(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	id := c.Params("id")

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
	code, err := service.VisitaResidente().DeleteVisitaResidente(id)

	if err != nil {
		logger.Error.Printf("Error deleting visita residente: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = nil
	res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
	return c.Status(fiber.StatusOK).JSON(res)
}

// GetAlumnosPendientesVisita godoc
// @Summary Obtener alumnos pendientes de visita
// @Description Método que permite obtener alumnos que necesitan visita domiciliaria por convocatoria
// @Tags Visita Residente
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param convocatoria_id query string false "ID de la convocatoria"
// @Success 200 {object} models.Response{error=boolean,data=[]models.AlumnoPendienteVisita,code=integer,type=string,msg=string} "Alumnos pendientes obtenidos exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/visita-residente/alumnos-pendientes [GET]
func (h *handlerVisitaResidente) GetAlumnosPendientesVisita(c *fiber.Ctx) error {
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

	convocatoriaID := c.Query("convocatoria_id")
	var convocatoriaIDPtr *string
	if convocatoriaID != "" {
		convocatoriaIDPtr = &convocatoriaID
	}

	service := general_visit.NewGeneralVisitService(h.db, user, h.txID)
	alumnos, err := service.VisitaResidente().GetAlumnosPendientesVisita(convocatoriaIDPtr)

	if err != nil {
		logger.Error.Printf("Error getting alumnos pendientes visita: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(3)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = alumnos
	res.Code, res.Type, res.Msg = h.msg.GetByCode(29)
	return c.Status(fiber.StatusOK).JSON(res)
}

// GetEstadisticasVisitas godoc
// @Summary Obtener estadísticas de visitas
// @Description Método que permite obtener estadísticas generales de las visitas domiciliarias
// @Tags Visita Residente
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Response{error=boolean,data=models.EstadisticasVisita,code=integer,type=string,msg=string} "Estadísticas obtenidas exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/visita-residente/estadisticas [GET]
func (h *handlerVisitaResidente) GetEstadisticasVisitas(c *fiber.Ctx) error {
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
	// Crear filtros si se especifica convocatoria_id
	var filtros *visita_residente.FiltrosVisita
	convocatoriaIDStr := c.Query("convocatoria_id")
	if convocatoriaIDStr != "" {
		convocatoriaID, err := strconv.ParseUint(convocatoriaIDStr, 10, 64)
		if err != nil {
			logger.Error.Printf("Invalid convocatoria_id parameter: %v", err)
			res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
			return c.Status(fiber.StatusBadRequest).JSON(res)
		}
		filtros = &visita_residente.FiltrosVisita{
			ConvocatoriaID: &convocatoriaID,
		}
	} else {
		filtros = &visita_residente.FiltrosVisita{}
	}

	service := general_visit.NewGeneralVisitService(h.db, user, h.txID)
	estadisticas, err := service.VisitaResidente().GetEstadisticasVisitas(filtros)

	if err != nil {
		logger.Error.Printf("Error getting estadisticas visitas: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(3)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = estadisticas
	res.Code, res.Type, res.Msg = h.msg.GetByCode(29)
	return c.Status(fiber.StatusOK).JSON(res)
}

// GetEstadisticasPorFacultad godoc
// @Summary Obtener estadísticas de visitas por facultad
// @Description Método que permite obtener estadísticas de visitas domiciliarias agrupadas por facultad
// @Tags Visita Residente
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Response{error=boolean,data=[]models.EstadisticasPorFacultad,code=integer,type=string,msg=string} "Estadísticas por facultad obtenidas exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/visita-residente/estadisticas/facultad [GET]
func (h *handlerVisitaResidente) GetEstadisticasPorEscuelaProfesional(c *fiber.Ctx) error {
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
	estadisticas, err := service.VisitaResidente().GetEstadisticasPorEscuelaProfesional(nil)

	if err != nil {
		logger.Error.Printf("Error getting estadisticas por escuela profesional: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(3)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = estadisticas
	res.Code, res.Type, res.Msg = h.msg.GetByCode(29)
	return c.Status(fiber.StatusOK).JSON(res)
}

// GetEstadisticasPorLugarProcedencia godoc
// @Summary Obtener estadísticas de visitas por lugar de procedencia
// @Description Método que permite obtener estadísticas de visitas domiciliarias agrupadas por lugar de procedencia
// @Tags Visita Residente
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.Response{error=boolean,data=[]models.EstadisticasPorLugarProcedencia,code=integer,type=string,msg=string} "Estadísticas por lugar de procedencia obtenidas exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/visita-residente/estadisticas/lugar-procedencia [GET]
func (h *handlerVisitaResidente) GetEstadisticasPorLugarProcedencia(c *fiber.Ctx) error {
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
	estadisticas, err := service.VisitaResidente().GetEstadisticasPorLugarProcedencia(nil)

	if err != nil {
		logger.Error.Printf("Error getting estadisticas por lugar procedencia: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(3)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = estadisticas
	res.Code, res.Type, res.Msg = h.msg.GetByCode(29)
	return c.Status(fiber.StatusOK).JSON(res)
}

// ExisteVisitaAlumno godoc
// @Summary Verificar si un alumno ya tiene visita domiciliaria
// @Description Método que permite verificar si un alumno ya tiene una visita domiciliaria registrada
// @Tags Visita Residente
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param alumno_id query string true "ID del alumno"
// @Success 200 {object} models.Response{error=boolean,data=map[string]interface{},code=integer,type=string,msg=string} "Resultado de la verificación"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/visita-residente/existe-alumno [GET]
func (h *handlerVisitaResidente) ExisteVisitaAlumno(c *fiber.Ctx) error {
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

	alumnoID := c.Query("alumno_id")
	if alumnoID == "" {
		logger.Error.Printf("Missing alumno_id parameter")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	service := general_visit.NewGeneralVisitService(h.db, user, h.txID)
	existe, err := service.VisitaResidente().ExistsByAlumnoID(alumnoID)

	if err != nil {
		logger.Error.Printf("Error checking if alumno has visit: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(3)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = map[string]interface{}{
		"alumno_id":    alumnoID,
		"tiene_visita": existe,
	}
	res.Code, res.Type, res.Msg = h.msg.GetByCode(29)
	return c.Status(fiber.StatusOK).JSON(res)
}

// GetAlumnosPendientesPorDepartamento godoc
// @Summary Obtener alumnos pendientes de visita por departamento
// @Description Método que permite obtener alumnos pendientes de visita filtrados por convocatoria y departamento
// @Tags Visita Residente
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param convocatoria_id query string true "ID de la convocatoria"
// @Param departamento query string true "Nombre del departamento"
// @Success 200 {object} models.Response{error=boolean,data=[]models.AlumnoPendienteVisitaPorDepartamento,code=integer,type=string,msg=string} "Alumnos pendientes por departamento obtenidos exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/visita-residente/alumnos-pendientes/departamento [GET]
func (h *handlerVisitaResidente) GetAlumnosPendientesPorDepartamento(c *fiber.Ctx) error {
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

	// Obtener parámetros requeridos
	convocatoriaIDStr := c.Query("convocatoria_id")
	departamento := c.Query("departamento")

	if convocatoriaIDStr == "" {
		logger.Error.Printf("Missing convocatoria_id parameter")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	if departamento == "" {
		logger.Error.Printf("Missing departamento parameter")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	convocatoriaID, err := strconv.ParseUint(convocatoriaIDStr, 10, 64)
	if err != nil {
		logger.Error.Printf("Invalid convocatoria_id parameter: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	service := general_visit.NewGeneralVisitService(h.db, user, h.txID)
	alumnos, err := service.VisitaResidente().GetAlumnosPendientesPorDepartamento(convocatoriaID, departamento)

	if err != nil {
		logger.Error.Printf("Error getting alumnos pendientes por departamento: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(3)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = alumnos
	res.Code, res.Type, res.Msg = h.msg.GetByCode(29)
	return c.Status(fiber.StatusOK).JSON(res)
}

// GetTodosAlumnosPorConvocatoria godoc
// @Summary Obtener todos los alumnos de una convocatoria con información completa
// @Description Método que permite obtener todos los alumnos de una convocatoria específica con toda su información para descarga/estadística
// @Tags Visita Residente
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param convocatoria_id query integer true "ID de la convocatoria"
// @Success 200 {object} models.Response{error=boolean,data=[]visita_residente.AlumnoCompleto,code=integer,type=string,msg=string} "Lista de alumnos obtenida exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Router /v1/visita-residente/alumnos-completos [GET]
func (h *handlerVisitaResidente) GetTodosAlumnosPorConvocatoria(c *fiber.Ctx) error {
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

	// Obtener parámetro requerido
	convocatoriaIDStr := c.Query("convocatoria_id")

	if convocatoriaIDStr == "" {
		logger.Error.Printf("Missing convocatoria_id parameter")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	convocatoriaID, err := strconv.ParseUint(convocatoriaIDStr, 10, 64)
	if err != nil {
		logger.Error.Printf("Invalid convocatoria_id parameter: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	service := general_visit.NewGeneralVisitService(h.db, user, h.txID)
	alumnos, err := service.VisitaResidente().GetTodosAlumnosPorConvocatoria(convocatoriaID)

	if err != nil {
		logger.Error.Printf("Error getting todos los alumnos por convocatoria: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(3)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	res.Error = false
	res.Data = alumnos
	res.Code, res.Type, res.Msg = h.msg.GetByCode(29)
	return c.Status(fiber.StatusOK).JSON(res)
}
