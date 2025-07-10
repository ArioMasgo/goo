package rooms

import (
	"dbu-api/internal/authorization"
	"dbu-api/internal/logger"
	"dbu-api/internal/middleware"
	"dbu-api/internal/models"
	"dbu-api/pkg/orchestrator/low_code_residences"
	"dbu-api/pkg/orchestrator/response_messages"
	"dbu-api/pkg/residence"
	"dbu-api/pkg/submission"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerRooms struct {
	db   *sqlx.DB
	txID string
	msg  response_messages.Message
}

// UpdateRooms godoc
// @Summary Actualiza una instancia de cuartos
// @Description Método que permite Actualiza una instancia del objeto cuartos en la base de datos
// @tags Cuartos
// @Accept json
// @Produce json
// @Param RoomRequest body RoomRequest true "Datos para actualizar cuartos"
// @Security BearerAuth
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /v1/residencias/cuartos [PUT]
func (h *handlerRooms) UpdateRooms(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := RoomRequest{}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(51)
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	err = authorization.ValidPermissions(user, h.db, c)
	if err != nil {
		logger.Error.Printf("User does not have permission to call the api, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(10)
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	if err := c.BodyParser(&req); err != nil {
		logger.Error.Printf("couldn't parse body request, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	isValid, err := req.Valid()
	if err != nil {
		logger.Error.Printf("couldn't validate body request, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if !isValid {
		logger.Error.Println("couldn't validate body request")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := residence.NewServerResidence(h.db, user, h.txID)
	_, _, err = srv.SrvRoom.UpdateOnlyCharacteristicsRoom(req.ID, req.Capacity, req.Status)
	if err != nil {
		logger.Error.Printf("couldn't update Rooms, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(13)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = req
	res.Error = false
	res.Code, res.Type, res.Msg = h.msg.GetByCode(212)
	return c.Status(http.StatusOK).JSON(res)
}

// AssignmentRoom godoc
// @Summary Asignar estudiante a un cuarto
// @Description Asigna un estudiante a un cuarto específico durante una convocatoria activa
// @Tags Cuartos
// @Accept json
// @Produce json
// @Param id path string true "ID del cuarto"
// @Param Authorization header string true "Bearer token"
// @Param request body AssignmentRoomRequest true "Datos de asignación del estudiante"
// @Security BearerAuth
// @Success 200 {object} models.Response{error=boolean,data=models.AssignmentRoom,code=integer,type=string,msg=string} "Asignación exitosa"
// @Success 201 {object} models.Response{error=boolean,data=models.AssignmentRoom,code=integer,type=string,msg=string} "Asignación creada exitosamente"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud: puede incluir errores de validación, habitación sin capacidad (código 93), estudiante ya asignado (código 94), estudiante con otra asignación (código 95), o estudiante no aceptado (código 96)"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Failure 403 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Permisos insuficientes"
// @Failure 404 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Recurso no encontrado"
// @Failure 500 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error interno del servidor"
// @Router /v1/residencias/cuartos/{id}/asignar [POST]
func (h *handlerRooms) AssignmentRoom(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := AssignmentRoomRequest{}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(51)
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	err = authorization.ValidPermissions(user, h.db, c)
	if err != nil {
		logger.Error.Printf("User does not have permission to call the api, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(10)
		return c.Status(http.StatusForbidden).JSON(res)
	}

	idStr := c.Params("id")
	if idStr == "" {
		logger.Error.Println("couldn't parse id request")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if err := c.BodyParser(&req); err != nil {
		logger.Error.Printf("couldn't parse body request, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	isValid, err := req.valid()
	if err != nil {
		logger.Error.Printf("couldn't validate body request, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if !isValid {
		logger.Error.Println("couldn't validate body request")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := low_code_residences.NewResidence(h.db, user, h.txID)
	assignmentRoom, code, err := srv.AssignmentRoom(idStr, req.SubmissionID, req.StudentID)
	if err != nil {
		logger.Error.Printf("couldn't get Rooms, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(http.StatusInternalServerError).JSON(res)
	}

	res.Error = false
	res.Data = assignmentRoom
	res.Code, res.Type, res.Msg = h.msg.GetByCode(211)
	return c.Status(http.StatusOK).JSON(res)
}

// DeleteAssignmentRoom godoc
// @Summary Eliminar asignación de cuarto
// @Description Elimina la asignación de un estudiante a un cuarto específico
// @Tags Cuartos
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param Authorization header string true "Bearer token"
// @Param id path string true "ID del cuarto"
// @Param request body DeleteAssignmentRoomRequest true "Datos de la asignación a eliminar"
// @Success 200 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Eliminación exitosa"
// @Failure 400 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error en la solicitud"
// @Failure 401 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error de autenticación"
// @Failure 403 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Permisos insuficientes"
// @Failure 404 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Recurso no encontrado"
// @Failure 500 {object} models.Response{error=boolean,data=interface{},code=integer,type=string,msg=string} "Error interno del servidor"
// @Router /v1/residencias/cuartos/{id}/eliminar-asignacion [DELETE]
func (h *handlerRooms) DeleteAssignmentRoom(c *fiber.Ctx) error {
	res := models.Response{Error: true}
	req := DeleteAssignmentRoomRequest{}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(51)
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	err = authorization.ValidPermissions(user, h.db, c)
	if err != nil {
		logger.Error.Printf("User does not have permission to call the api, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(10)
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	if err := c.BodyParser(&req); err != nil {
		logger.Error.Printf("couldn't parse body request, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	isValid, err := req.valid()
	if err != nil {
		logger.Error.Printf("couldn't validate body request, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if !isValid {
		logger.Error.Println("couldn't validate body request")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	idStr := c.Params("id")
	if idStr == "" {
		logger.Error.Println("couldn't parse id request")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srvSubmission := submission.NewServerSubmission(h.db, user, h.txID)
	activeSubmission, _, err := srvSubmission.SrvConvocatorias.GetSubmissionsByID(req.SubmissionID)
	if err != nil {
		logger.Error.Println(h.txID, " - couldn't get active submission", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(15)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if activeSubmission == nil {
		logger.Error.Println(h.txID, " - no found active submission")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(4)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	srv := residence.NewServerResidence(h.db, user, h.txID)
	room, _, err := srv.SrvRoom.GetRoomByID(idStr)
	if err != nil {
		logger.Error.Printf("couldn't get Rooms, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(15)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	if room == nil {
		logger.Error.Printf("no found room")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(4)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	_, err = srv.SrvAssignmentRoom.DeleteRoomAssignment(req.StudentID, room.ID, activeSubmission.ID, req.Status, req.Observation)
	if err != nil {
		logger.Error.Printf("couldn't delete assignment room, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(14)
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Error = false
	res.Code, res.Type, res.Msg = h.msg.GetByCode(213)
	return c.Status(http.StatusOK).JSON(res)
}

// GetAllRoomsByResidence godoc
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
// @Router /v1/residencias/{id}/cuartos/ [GET]
func (h *handlerRooms) GetAllRoomsByResidence(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(51)
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

	submissionHeader := c.QueryInt("submission_id", 0)
	if submissionHeader == 0 {
		logger.Error.Printf("missing submission_id parameter")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(2)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	srv := low_code_residences.NewResidence(h.db, user, h.txID)
	students, _, _, err := srv.GetRoomsByResidenceLowCode(residenceID, int64(submissionHeader), page, limit)
	if err != nil {
		logger.Error.Printf("error getting students: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(16)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res.Data = students

	res.Error = false
	res.Code, res.Type, res.Msg = h.msg.GetByCode(215)
	return c.Status(fiber.StatusOK).JSON(res)
}
