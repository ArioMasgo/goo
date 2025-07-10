package exam_toxicologico

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/middleware"
	"dbu-api/internal/models"

	//"dbu-api/pkg/core"
	"dbu-api/pkg/medical_area/exam_toxicologico"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type handlerExamToxicologico struct {
	db   *sqlx.DB
	txID string
}

// CreateRegistroToxicologico godoc
// @Summary Crear un registro toxicológico
// @Description Método que permite crear un registro toxicológico para un alumno
// @tags Examen Toxicológico
// @Accept json
// @Produce json
// @Param models.RegistroToxicologicoRequest body models.RegistroToxicologicoRequest true "Datos para crear registro toxicológico"
// @Success 201 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Router /api/v1/area_medica/examen_toxicologico [POST]
func (h *handlerExamToxicologico) CreateRegistroToxicologico(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := models.RegistroToxicologicoRequest{}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	if err := c.BodyParser(&req); err != nil {
		logger.Error.Printf("couldn't parse body request, error: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", "invalid request body"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	isValid, err := req.Valid()
	if err != nil {
		logger.Error.Printf("couldn't validate body request, error: %v", err)
		res.Code, res.Type, res.Msg = 15, "error", "validation failed"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if !isValid {
		logger.Error.Printf("invalid request body")
		res.Code, res.Type, res.Msg = 15, "error", "validation failed"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	// Usar el ID del usuario autenticado si no se proporciona
	userID := user.ID
	if req.IDUsuario != nil {
		userID = *req.IDUsuario
	}

	//srv := core.NewServerCore(h.db, user, h.txID)
	repository := exam_toxicologico.FactoryStorage(h.db, h.txID)
	service := exam_toxicologico.NewRegistroToxicologicoService(repository, user, h.txID)

	registro, code, err := service.CreateRegistroToxicologico(
		req.AlumnoID,
		req.ConvocatoriaID,
		req.Estado,
		req.Comentario,
		&userID,
	)

	if err != nil {
		logger.Error.Printf("couldn't create registro toxicológico, error: %v", err)
		res.Code, res.Type, res.Msg = code, "error", err.Error()
		if code == 16 {
			return c.Status(http.StatusConflict).JSON(res)
		}
		return c.Status(http.StatusInternalServerError).JSON(res)
	}

	res.Data = registro
	res.Code, res.Type, res.Msg, res.Error = code, "success", "Registro toxicológico creado exitosamente", false
	return c.Status(http.StatusCreated).JSON(res)
}

// UpdateRegistroToxicologico godoc
// @Summary Actualizar un registro toxicológico
// @Description Método que permite actualizar el estado de un registro toxicológico
// @tags Examen Toxicológico
// @Accept json
// @Produce json
// @Param id path int true "ID del registro toxicológico"
// @Param models.RegistroToxicologicoUpdate body models.RegistroToxicologicoUpdate true "Datos para actualizar registro toxicológico"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /api/v1/area_medica/examen_toxicologico/{id} [PUT]
func (h *handlerExamToxicologico) UpdateRegistroToxicologico(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := models.RegistroToxicologicoUpdate{}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		logger.Error.Printf("couldn't parse id parameter, error: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", "invalid id parameter"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if err := c.BodyParser(&req); err != nil {
		logger.Error.Printf("couldn't parse body request, error: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", "invalid request body"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	isValid, err := req.Valid()
	if err != nil {
		logger.Error.Printf("couldn't validate body request, error: %v", err)
		res.Code, res.Type, res.Msg = 15, "error", "validation failed"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if !isValid {
		logger.Error.Printf("invalid request body")
		res.Code, res.Type, res.Msg = 15, "error", "validation failed"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	// Usar el ID del usuario autenticado si no se proporciona
	userID := user.ID
	if req.IDUsuario != nil {
		userID = *req.IDUsuario
	}

	repository := exam_toxicologico.FactoryStorage(h.db, h.txID)
	service := exam_toxicologico.NewRegistroToxicologicoService(repository, user, h.txID)

	registro, code, err := service.UpdateRegistroToxicologico(
		id,
		req.Estado,
		req.Comentario,
		&userID,
	)

	if err != nil {
		logger.Error.Printf("couldn't update registro toxicológico, error: %v", err)
		res.Code, res.Type, res.Msg = code, "error", err.Error()
		if code == 22 {
			return c.Status(http.StatusNotFound).JSON(res)
		}
		return c.Status(http.StatusInternalServerError).JSON(res)
	}

	res.Data = registro
	res.Code, res.Type, res.Msg, res.Error = code, "success", "Registro toxicológico actualizado exitosamente", false
	return c.Status(http.StatusOK).JSON(res)
}

// GetRegistroToxicologicoByID godoc
// @Summary Obtener registro toxicológico por ID
// @Description Método que permite obtener un registro toxicológico por su ID
// @tags Examen Toxicológico
// @Accept json
// @Produce json
// @Param id path int true "ID del registro toxicológico"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /api/v1/area_medica/examen_toxicologico/{id} [GET]
func (h *handlerExamToxicologico) GetRegistroToxicologicoByID(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		logger.Error.Printf("couldn't parse id parameter, error: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", "invalid id parameter"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	repository := exam_toxicologico.FactoryStorage(h.db, h.txID)
	service := exam_toxicologico.NewRegistroToxicologicoService(repository, user, h.txID)

	registro, code, err := service.GetRegistroToxicologicoByID(id)

	if err != nil {
		logger.Error.Printf("couldn't get registro toxicológico, error: %v", err)
		res.Code, res.Type, res.Msg = code, "error", err.Error()
		if code == 22 {
			return c.Status(http.StatusNotFound).JSON(res)
		}
		return c.Status(http.StatusInternalServerError).JSON(res)
	}

	res.Data = registro
	res.Code, res.Type, res.Msg, res.Error = code, "success", "Registro toxicológico obtenido exitosamente", false
	return c.Status(http.StatusOK).JSON(res)
}

// GetEstadosByConvocatoria godoc
// @Summary Obtener estados toxicológicos por convocatoria
// @Description Método que permite obtener todos los estados toxicológicos de una convocatoria
// @tags Examen Toxicológico
// @Accept json
// @Produce json
// @Param convocatoria_id path int true "ID de la convocatoria"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Router /api/v1/area_medica/examen_toxicologico/convocatoria/{convocatoria_id} [GET]
func (h *handlerExamToxicologico) GetEstadosByConvocatoria(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	convocatoriaID, err := strconv.ParseInt(c.Params("convocatoria_id"), 10, 64)
	if err != nil {
		logger.Error.Printf("couldn't parse convocatoria_id parameter, error: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", "invalid convocatoria_id parameter"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	repository := exam_toxicologico.FactoryStorage(h.db, h.txID)
	service := exam_toxicologico.NewRegistroToxicologicoService(repository, user, h.txID)

	estados, err := service.GetEstadosByConvocatoria(convocatoriaID)

	if err != nil {
		logger.Error.Printf("couldn't get estados by convocatoria, error: %v", err)
		res.Code, res.Type, res.Msg = 3, "error", err.Error()
		return c.Status(http.StatusInternalServerError).JSON(res)
	}

	res.Data = estados
	res.Code, res.Type, res.Msg, res.Error = 29, "success", "Estados toxicológicos obtenidos exitosamente", false
	return c.Status(http.StatusOK).JSON(res)
}

// DeleteRegistroToxicologico godoc
// @Summary Eliminar registro toxicológico
// @Description Método que permite eliminar un registro toxicológico
// @tags Examen Toxicológico
// @Accept json
// @Produce json
// @Param id path int true "ID del registro toxicológico"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Failure 404 {object} models.Response
// @Router /api/v1/area_medica/examen_toxicologico/{id} [DELETE]
func (h *handlerExamToxicologico) DeleteRegistroToxicologico(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		logger.Error.Printf("couldn't parse id parameter, error: %v", err)
		res.Code, res.Type, res.Msg = 1, "error", "invalid id parameter"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	repository := exam_toxicologico.FactoryStorage(h.db, h.txID)
	service := exam_toxicologico.NewRegistroToxicologicoService(repository, user, h.txID)

	code, err := service.DeleteRegistroToxicologico(id)

	if err != nil {
		logger.Error.Printf("couldn't delete registro toxicológico, error: %v", err)
		res.Code, res.Type, res.Msg = code, "error", err.Error()
		if code == 22 {
			return c.Status(http.StatusNotFound).JSON(res)
		}
		return c.Status(http.StatusInternalServerError).JSON(res)
	}

	res.Code, res.Type, res.Msg, res.Error = code, "success", "Registro toxicológico eliminado exitosamente", false
	return c.Status(http.StatusOK).JSON(res)
}

// GetAllRegistrosToxicologicos godoc
// @Summary Obtener todos los registros toxicológicos
// @Description Método que permite obtener todos los registros toxicológicos
// @tags Examen Toxicológico
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 401 {object} models.Response
// @Router /api/v1/area_medica/examen_toxicologico [GET]
func (h *handlerExamToxicologico) GetAllRegistrosToxicologicos(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	repository := exam_toxicologico.FactoryStorage(h.db, h.txID)
	service := exam_toxicologico.NewRegistroToxicologicoService(repository, user, h.txID)

	registros, err := service.GetAllRegistrosToxicologicos()

	if err != nil {
		logger.Error.Printf("couldn't get all registros toxicológicos, error: %v", err)
		res.Code, res.Type, res.Msg = 3, "error", err.Error()
		return c.Status(http.StatusInternalServerError).JSON(res)
	}

	res.Data = registros
	res.Code, res.Type, res.Msg, res.Error = 29, "success", "Registros toxicológicos obtenidos exitosamente", false
	return c.Status(http.StatusOK).JSON(res)
}
