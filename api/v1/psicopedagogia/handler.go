package psicopedagogia

import (
	"dbu-api/internal/logger"
	psicopedagogia "dbu-api/pkg/psicopedagogia"
	"dbu-api/pkg/psicopedagogia/citas"
	"dbu-api/pkg/psicopedagogia/diagnosticos"
	"dbu-api/pkg/psicopedagogia/encuestas"
	"dbu-api/pkg/psicopedagogia/participantes"
	"dbu-api/pkg/psicopedagogia/preguntas"
	"dbu-api/pkg/psicopedagogia/respuestas"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/jmoiron/sqlx"
)

type handlerpsicopedagogia struct {
	db   *sqlx.DB
	txID string
}

// --------------------------------------------------
// Handlers para estudiante
// --------------------------------------------------
func (h *handlerpsicopedagogia) GetEstudianteByDni(c *fiber.Ctx) error {
	res := ResponseData{Error: true}
	dni := c.Params("dni")
	tipoParticipante := c.Query("tipoParticipante", "")
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	data, _, err := srv.SrvEstudiantes.GetEstudianteByDni(dni, tipoParticipante)
	if err != nil {
		logger.Error.Printf("No se pudo recuperar el estudiante: %v", err)
		res.Code, res.Type, res.Msg = 23, "error", "could not retrieve estudiante"
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "success", ""
	if data == nil {
		res.Msg = "El alumno con DNI: " + dni + " no fue encontrado"
		return c.Status(http.StatusOK).JSON(res)
	}
	res.Detalle = data
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerpsicopedagogia) GetEstudianteByNameDni(c *fiber.Ctx) error {
	res := ResponseData{Error: true}
	dni := c.Params("dni")
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	data, _, err := srv.SrvEstudiantes.GetNameEstudianteByDni(dni)
	if err != nil {
		logger.Error.Printf("No se pudo recuperar el estudiante: %v", err)
		res.Code, res.Type, res.Msg = 23, "error", "could not retrieve estudiante"
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "success", ""
	if data == nil {
		res.Msg = "El alumno con DNI: " + dni + " no fue encontrado"
		return c.Status(http.StatusOK).JSON(res)
	}
	res.Detalle = data
	return c.Status(http.StatusOK).JSON(res)
}

// --------------------------------------------------
// Handlers para Participantes
// --------------------------------------------------
func (h *handlerpsicopedagogia) GetAll(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", 10)
	participantes, totalPages, err := srv.SrvParticipantes.GetAll(page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, "Error obteniendo participantes", nil, fiber.StatusInternalServerError, "error"})
	}
	return c.JSON(ResponseData{false, "Participantes obtenidos",
		fiber.Map{
			"items":      participantes,
			"totalPages": totalPages,
		}, fiber.StatusOK, "success"})
}

func (h *handlerpsicopedagogia) GetByID(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	id, _ := c.ParamsInt("id")
	participante, err := srv.SrvParticipantes.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, "Error obteniendo participante", nil, fiber.StatusInternalServerError, "error"})
	}
	return c.JSON(ResponseData{false, "Participante obtenido", participante, fiber.StatusOK, "success"})
}

func (h *handlerpsicopedagogia) Create(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	var p RequestParticipantes
	if err := c.BodyParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{
			true, "Datos inválidos", nil, fiber.StatusBadRequest, "error",
		})
	}
	if err := srv.SrvParticipantes.Create(&participantes.Participante{
		Tipo:               p.Tipo,
		Nombre:             p.Nombre,
		Apellido:           p.Apellido,
		DNI:                p.DNI,
		Estado:             p.Estado,
		CreatedAt:          p.CreatedAt,
		UpdatedAt:          p.UpdatedAt,
		ColegioProcedencia: p.ColegioProcedencia,
		AnioIngreso:        p.AnioIngreso,
		Escuela:            p.Escuela,
		CodigoEstudiante:   p.CodigoEstudiante,
		FechaNacimiento:    p.FechaNacimiento,
		Edad:               p.Edad,
		LugarNacimiento:    p.LugarNacimiento,
		ModalidadIngreso:   p.ModalidadIngreso,
		NumeroAtencion:     p.NumeroAtencion,
		Sexo:               p.Sexo,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{
			true, err.Error(), nil, fiber.StatusInternalServerError, "error",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(ResponseData{
		false, "Participante creado", p, fiber.StatusCreated, "success",
	})
}

func (h *handlerpsicopedagogia) Update(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	id, _ := c.ParamsInt("id")
	var p RequestParticipantes

	if err := c.BodyParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{true, "Datos inválidos", nil, fiber.StatusBadRequest, "error"})
	}

	if err := srv.SrvParticipantes.Update(id, &participantes.Participante{
		Tipo:               p.Tipo,
		Nombre:             p.Nombre,
		Apellido:           p.Apellido,
		DNI:                p.DNI,
		Estado:             p.Estado,
		CreatedAt:          p.CreatedAt,
		UpdatedAt:          p.UpdatedAt,
		ColegioProcedencia: p.ColegioProcedencia,
		AnioIngreso:        p.AnioIngreso,
		Escuela:            p.Escuela,
		CodigoEstudiante:   p.CodigoEstudiante,
		FechaNacimiento:    p.FechaNacimiento,
		Edad:               p.Edad,
		LugarNacimiento:    p.LugarNacimiento,
		ModalidadIngreso:   p.ModalidadIngreso,
		NumeroAtencion:     p.NumeroAtencion,
		Sexo:               p.Sexo,
		NumTelefono:        p.NumTelefono,
		EstadoEvaluacion:   p.EstadoEvaluacion,
		ConQuienesVive:     p.ConQuienesVive,
		SemestreCursa:      p.SemestreCursa,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, err.Error(), nil, fiber.StatusInternalServerError, "error"})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseData{false, "Participante actualizado", nil, fiber.StatusOK, "success"})
}

func (h *handlerpsicopedagogia) Delete(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	id, _ := c.ParamsInt("id")
	if err := srv.SrvParticipantes.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, err.Error(), nil, fiber.StatusInternalServerError, "error"})
	}
	return c.Status(fiber.StatusNoContent).JSON(ResponseData{false, "Participante eliminado", nil, fiber.StatusNoContent, "success"})
}

// --------------------------------------------------
// Handlers para Preguntas
// --------------------------------------------------
func (h *handlerpsicopedagogia) GetAllPreguntas(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", 10)
	preguntas, err := srv.SrvPreguntas.GetAll(page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, "Error obteniendo preguntas", nil, fiber.StatusInternalServerError, "error"})
	}
	return c.JSON(ResponseData{false, "Preguntas obtenidas", preguntas, fiber.StatusOK, "success"})
}

func (h *handlerpsicopedagogia) GetPreguntaByID(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	id, _ := c.ParamsInt("id")
	pregunta, err := srv.SrvPreguntas.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, "Error obteniendo pregunta", nil, fiber.StatusInternalServerError, "error"})
	}
	return c.JSON(ResponseData{false, "Pregunta obtenida", pregunta, fiber.StatusOK, "success"})
}

func (h *handlerpsicopedagogia) CreatePregunta(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	var p RequestPreguntas

	if err := c.BodyParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{true, "Datos inválidos", nil, fiber.StatusBadRequest, "error"})
	}

	if _, err := p.Valid(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{true, err.Error(), nil, fiber.StatusBadRequest, "error"})
	}

	if err := srv.SrvPreguntas.Create(&preguntas.Pregunta{
		TextoPregunta: p.TextoPregunta,
		IsMandatory:   p.IsMandatory,
		Order:         p.Order,
		Type:          p.Type,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, err.Error(), nil, fiber.StatusInternalServerError, "error"})
	}

	return c.Status(fiber.StatusCreated).JSON(ResponseData{false, "Pregunta creada", p, fiber.StatusCreated, "success"})
}

func (h *handlerpsicopedagogia) UpdatePregunta(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	id, _ := c.ParamsInt("id")
	var p RequestPreguntas

	if err := c.BodyParser(&p); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{true, "Datos inválidos", nil, fiber.StatusBadRequest, "error"})
	}

	if err := srv.SrvPreguntas.Update(id, &preguntas.Pregunta{
		TextoPregunta: p.TextoPregunta,
		IsMandatory:   p.IsMandatory,
		Order:         p.Order,
		Type:          p.Type,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, err.Error(), nil, fiber.StatusInternalServerError, "error"})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseData{false, "Pregunta actualizada", nil, fiber.StatusOK, "success"})
}

func (h *handlerpsicopedagogia) DeletePregunta(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	id, _ := c.ParamsInt("id")

	if err := srv.SrvPreguntas.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, err.Error(), nil, fiber.StatusInternalServerError, "error"})
	}

	return c.Status(fiber.StatusNoContent).JSON(ResponseData{false, "Pregunta eliminada", nil, fiber.StatusNoContent, "success"})
}

// --------------------------------------------------
// Handlers para Encuestas
// --------------------------------------------------
func (h *handlerpsicopedagogia) GetAllEncuestas(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", 10)
	encuestas, err := srv.SrvEncuestas.GetAll(page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, "Error obteniendo encuestas", nil, fiber.StatusInternalServerError, "error"})
	}
	return c.JSON(ResponseData{false, "Encuestas obtenidas", encuestas, fiber.StatusOK, "success"})
}

func (h *handlerpsicopedagogia) GetEncuestaByID(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	id, _ := c.ParamsInt("id")
	encuesta, err := srv.SrvEncuestas.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, "Error obteniendo encuesta", nil, fiber.StatusInternalServerError, "error"})
	}
	return c.JSON(ResponseData{false, "Encuesta obtenida", encuesta, fiber.StatusOK, "success"})
}

func (h *handlerpsicopedagogia) CreateEncuesta(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	var e RequestEncuestas

	if err := c.BodyParser(&e); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{true, "Datos inválidos", nil, fiber.StatusBadRequest, "error"})
	}

	isValid, err := e.Valid()
	if !isValid {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{true, err.Error(), nil, fiber.StatusBadRequest, "error"})
	}
	encuestaId, err := srv.SrvEncuestas.Create(&encuestas.Encuesta{
		Nombre:      e.Nombre,
		Descripcion: e.Descripcion,
		Estado:      e.Estado,
		FechaInicio: e.FechaInicio,
		FechaFin:    e.FechaFin,
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, err.Error(), nil, fiber.StatusInternalServerError, "error"})
	}

	return c.Status(fiber.StatusCreated).JSON(ResponseData{false, "Encuesta creada", encuestaId, fiber.StatusCreated, "success"})
}

func (h *handlerpsicopedagogia) UpdateEncuesta(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	id, _ := c.ParamsInt("id")
	var e RequestEncuestas

	if err := c.BodyParser(&e); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{true, "Datos inválidos", nil, fiber.StatusBadRequest, "error"})
	}

	if err := srv.SrvEncuestas.Update(id, &encuestas.Encuesta{
		Nombre:      e.Nombre,
		Descripcion: e.Descripcion,
		Estado:      e.Estado,
		FechaInicio: e.FechaInicio,
		FechaFin:    e.FechaFin,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, err.Error(), nil, fiber.StatusInternalServerError, "error"})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseData{false, "Encuesta actualizada", nil, fiber.StatusOK, "success"})
}

func (h *handlerpsicopedagogia) DeleteEncuesta(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	id, _ := c.ParamsInt("id")

	if err := srv.SrvEncuestas.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, err.Error(), nil, fiber.StatusInternalServerError, "error"})
	}

	return c.Status(fiber.StatusNoContent).JSON(ResponseData{false, "Encuesta eliminada", nil, fiber.StatusNoContent, "success"})
}

func (h *handlerpsicopedagogia) HasActiveSRQ(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)

	idEncuesta, fechaFin, hasActive, err := srv.SrvEncuestas.HasActiveSRQ()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{
			Error:   true,
			Msg:     "Error obteniendo encuesta",
			Detalle: nil,
			Code:    fiber.StatusInternalServerError,
			Type:    "error",
		})
	}

	return c.JSON(ResponseData{
		Error: false,
		Msg:   "Consulta realizada con éxito",
		Detalle: map[string]interface{}{
			"id_encuesta": idEncuesta,
			"fecha_fin":   fechaFin,
			"activa":      hasActive,
		},
		Code: fiber.StatusOK,
		Type: "success",
	})
}

// --------------------------------------------------
// Handlers para Respuestas
// --------------------------------------------------
func (h *handlerpsicopedagogia) GetAllRespuestas(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", 10)
	respuestas, err := srv.SrvRespuestas.GetAll(page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, "Error obteniendo respuestas", nil, fiber.StatusInternalServerError, "error"})
	}
	return c.JSON(ResponseData{false, "Respuestas obtenidas", respuestas, fiber.StatusOK, "success"})
}

func (h *handlerpsicopedagogia) GetRespuestaByID(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	id, _ := c.ParamsInt("id")
	respuesta, err := srv.SrvRespuestas.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, "Error obteniendo respuesta", nil, fiber.StatusInternalServerError, "error"})
	}
	return c.JSON(ResponseData{false, "Respuesta obtenida", respuesta, fiber.StatusOK, "success"})
}

func (h *handlerpsicopedagogia) GetResponsesPerParticipant(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	idParticipante, err := c.ParamsInt("idParticipante")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{true, "IDParticipante inválido", nil, fiber.StatusBadRequest, "error"})
	}

	respuestas, err := srv.SrvRespuestas.GetResponsesPerParticipant(idParticipante)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, "Error obteniendo respuestas", nil, fiber.StatusInternalServerError, "error"})
	}

	return c.JSON(ResponseData{false, "Respuestas obtenidas", respuestas, fiber.StatusOK, "success"})
}

func (h *handlerpsicopedagogia) GetAllByParticipanteIdAndNumeroAtencion(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	if err := c.BodyParser(&RequestRespuesta); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{true, "Cuerpo de la solicitud inválido", nil, fiber.StatusBadRequest, "error"})
	}
	if RequestRespuesta.IDParticipante == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{true, "IDParticipante es requerido", nil, fiber.StatusBadRequest, "error"})
	}
	if RequestRespuesta.NumeroAtencion == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{true, "NumeroAtencion es requerido", nil, fiber.StatusBadRequest, "error"})
	}
	respuestas, err := srv.SrvRespuestas.GetAllByParticipanteIdAndNumeroAtencion(RequestRespuesta.IDParticipante, RequestRespuesta.NumeroAtencion, RequestRespuesta.Page, RequestRespuesta.PageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, "Error obteniendo respuestas", nil, fiber.StatusInternalServerError, "error"})
	}
	return c.JSON(ResponseData{false, "Respuestas obtenidas", respuestas, fiber.StatusOK, "success"})
}

func (h *handlerpsicopedagogia) CreateRespuesta(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	var r RequestRespuestas

	if err := c.BodyParser(&r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{true, "Datos inválidos", nil, fiber.StatusBadRequest, "error"})
	}

	if _, err := r.Valid(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{true, err.Error(), nil, fiber.StatusBadRequest, "error"})
	}

	if err := srv.SrvRespuestas.Create(&respuestas.Respuesta{
		IDParticipante: r.IDParticipante,
		IDEncuesta:     r.IDEncuesta,
		IDPregunta:     r.IDPregunta,
		Respuesta:      r.Respuesta,
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, err.Error(), nil, fiber.StatusInternalServerError, "error"})
	}

	return c.Status(fiber.StatusCreated).JSON(ResponseData{false, "Respuesta creada", r, fiber.StatusCreated, "success"})
}

func (h *handlerpsicopedagogia) UpdateRespuesta(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	id, _ := c.ParamsInt("id")
	var r RequestRespuestas

	if err := c.BodyParser(&r); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{true, "Datos inválidos", nil, fiber.StatusBadRequest, "error"})
	}

	if err := srv.SrvRespuestas.Update(id, &respuestas.Respuesta{
		IDParticipante: r.IDParticipante,
		IDEncuesta:     r.IDEncuesta,
		IDPregunta:     r.IDPregunta,
		Respuesta:      r.Respuesta,
		CreatedAt:      r.CreatedAt,
		UpdatedAt:      r.UpdatedAt,
	}); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, err.Error(), nil, fiber.StatusInternalServerError, "error"})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseData{false, "Respuesta actualizada", nil, fiber.StatusOK, "success"})
}

func (h *handlerpsicopedagogia) DeleteRespuesta(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	id, _ := c.ParamsInt("id")

	if err := srv.SrvRespuestas.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, err.Error(), nil, fiber.StatusInternalServerError, "error"})
	}

	return c.Status(fiber.StatusNoContent).JSON(ResponseData{false, "Respuesta eliminada", nil, fiber.StatusNoContent, "success"})
}

// --------------------------------------------------
// 📌 **Crear historial y respuestas asociadas**
// --------------------------------------------------

func (h *handlerpsicopedagogia) GetLatestHistorial(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("pageSize", 10)
	participantes, err := srv.SrvHistorialEncuesta.GetLatestHistorial(page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{true, "Error obteniendo participantes - historial", nil, fiber.StatusInternalServerError, "error"})
	}
	return c.JSON(ResponseData{false, "Participantes obtenidos", participantes, fiber.StatusOK, "success"})
}

func (h *handlerpsicopedagogia) GetHistorialFiltered(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)

	// Define the request structure
	var request struct {
		DNI         string `json:"dni"`
		Apellido    string `json:"apellido"`
		FechaInicio string `json:"fecha_inicio"`
		FechaFin    string `json:"fecha_fin"`
	}

	// Parse request body
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{
			Error:   true,
			Msg:     "Datos inválidos",
			Detalle: nil,
			Code:    fiber.StatusBadRequest,
			Type:    "error",
		})
	}
	if request.DNI == "" && request.Apellido == "" && request.FechaInicio == "" && request.FechaFin == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{
			Error:   true,
			Msg:     "Debe proporcionar al menos un criterio de búsqueda",
			Detalle: nil,
			Code:    fiber.StatusBadRequest,
			Type:    "error",
		})
	}
	historial, err := srv.SrvHistorialEncuesta.GetHistorialFiltered(request.DNI, request.Apellido, request.FechaInicio, request.FechaFin)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{
			Error:   true,
			Msg:     "Error obteniendo historial",
			Detalle: err.Error(),
			Code:    fiber.StatusInternalServerError,
			Type:    "error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseData{
		Error:   false,
		Msg:     "Historial obtenido correctamente",
		Detalle: historial,
		Code:    fiber.StatusOK,
		Type:    "success",
	})
}

func (h *handlerpsicopedagogia) CreateHistorialWithAnswers(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	var request GuardarEncuestaRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{Error: true, Msg: "Datos inválidos", Detalle: nil, Code: fiber.StatusBadRequest, Type: "error"})
	}

	participante := ConvertRequestToParticipante(&request.Participante)
	respuestas := ConvertRequestToRespuestas(request.Respuestas, request.EncuestaID)
	historial := ConvertRequestToHistorialEncuesta(&request.Participante, request.EncuestaID)
	estadoEvaluacion, err := srv.SrvHistorialEncuesta.GuardarEncuestaWithAnswers(historial, participante, respuestas)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{Error: true, Msg: err.Error(), Detalle: nil, Code: fiber.StatusInternalServerError, Type: "error"})
	}

	return c.Status(fiber.StatusCreated).JSON(ResponseData{
		Error:   false,
		Msg:     "Historial y respuestas creadas correctamente",
		Detalle: estadoEvaluacion,
		Code:    fiber.StatusCreated,
		Type:    "success",
	})
}

func (h *handlerpsicopedagogia) HasStudentResponded(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	var request struct {
		DNI        string `json:"dni"`
		EncuestaID int    `json:"id_encuesta"`
	}
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{
			Error:   true,
			Msg:     "Solicitud inválida. Verifica el formato de los datos.",
			Detalle: nil,
			Code:    fiber.StatusBadRequest,
			Type:    "error",
		})
	}
	if request.DNI == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{
			Error:   true,
			Msg:     "El DNI del estudiante es obligatorio",
			Detalle: nil,
			Code:    fiber.StatusBadRequest,
			Type:    "error",
		})
	}
	if request.EncuestaID <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{
			Error:   true,
			Msg:     "El ID de la encuesta es obligatorio y debe ser un número válido",
			Detalle: nil,
			Code:    fiber.StatusBadRequest,
			Type:    "error",
		})
	}
	hasResponded, err := srv.SrvHistorialEncuesta.HasStudentResponded(request.DNI, request.EncuestaID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{
			Error:   true,
			Msg:     "Error verificando la respuesta del estudiante",
			Detalle: err.Error(),
			Code:    fiber.StatusInternalServerError,
			Type:    "error",
		})
	}

	return c.JSON(ResponseData{
		Error: false,
		Msg:   "Consulta realizada con éxito",
		Detalle: map[string]interface{}{
			"dni":         request.DNI,
			"id_encuesta": request.EncuestaID,
			"respondida":  hasResponded,
		},
		Code: fiber.StatusOK,
		Type: "success",
	})
}

func (h *handlerpsicopedagogia) SaveHistory(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	var request GuardarHistorialRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{Error: true, Msg: "Datos inválidos", Detalle: nil, Code: fiber.StatusBadRequest, Type: "error"})
	}
	participante := ConvertRequestToParticipante(&request.Participante)
	historial := ConvertRequestToHistorial(&request.Participante)
	err := srv.SrvHistorialEncuesta.SaveHistory(historial, participante)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{Error: true, Msg: err.Error(), Detalle: nil, Code: fiber.StatusInternalServerError, Type: "error"})
	}
	return c.Status(fiber.StatusCreated).JSON(ResponseData{
		Error:   false,
		Msg:     "Historial y respuestas creadas correctamente",
		Detalle: true,
		Code:    fiber.StatusCreated,
		Type:    "success",
	})
}

func (h *handlerpsicopedagogia) KeyUrlExists(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	keyUrl := c.Query("key_url")
	if keyUrl == "" {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{
			Error:   true,
			Msg:     "El parámetro 'key_url' es requerido",
			Detalle: nil,
			Code:    fiber.StatusBadRequest,
			Type:    "error",
		})
	}
	exists, err := srv.SrvHistorialEncuesta.KeyUrlExists(keyUrl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{
			Error:   true,
			Msg:     err.Error(),
			Detalle: nil,
			Code:    fiber.StatusInternalServerError,
			Type:    "error",
		})
	}
	return c.Status(fiber.StatusOK).JSON(ResponseData{
		Error:   false,
		Msg:     "Consulta realizada correctamente",
		Detalle: exists,
		Code:    fiber.StatusOK,
		Type:    "success",
	})
}

// --------------------------------------------------
// 📌 **Metodo para busqueda general
// --------------------------------------------------
func (h *handlerpsicopedagogia) SearchParticipants(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	res := ResponseData{Error: true}
	var filters map[string]string
	if err := c.BodyParser(&filters); err != nil {
		logger.Error.Printf("couldn't parse filters, error: %v", err)
		res.Code, res.Type, res.Msg = 22, "error", "could not parse request body"
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	criteria := make(map[string]interface{})
	for key, value := range filters {
		criteria[key] = value
	}
	data, code, err := srv.SrvParticipantes.SearchParticipants(criteria)
	if err != nil {
		logger.Error.Printf("couldn't search participants, error: %v", err)
		res.Code, res.Type, res.Msg = 23, "error", "could not retrieve results"
		return c.Status(http.StatusAccepted).JSON(res)
	}
	res.Detalle = data
	res.Error = false
	res.Code, res.Type, res.Msg = code, "success", ""
	return c.Status(http.StatusOK).JSON(res)
}

// --------------------------------------------------
// 📌 **Metodo PDFs
// --------------------------------------------------
func (h *handlerpsicopedagogia) GeneratePDF_SRQ_Handler(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	participantID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid participant ID",
		})
	}

	pdfBytes, err := srv.SrvPDFs.GeneratePDF_SRQ(participantID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": "Failed to generate PDF",
		})
	}

	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "attachment; filename=SRQ_Report.pdf")
	return c.Send(pdfBytes)
}

// --------------------------------------------------
// 📌 **Metodo Diagnosticos
// --------------------------------------------------

func (h *handlerpsicopedagogia) GetDiagnosticoByID(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "invalid id parameter",
		})
	}
	diagnostico, err := srv.SrvDiagnostico.GetByID(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(diagnostico)
}

func (h *handlerpsicopedagogia) GetDiagnosticoAll(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	res := ResponseData{Error: true}
	diagnosticos, err := srv.SrvDiagnostico.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	res.Detalle = diagnosticos
	res.Error = false
	res.Code, res.Type, res.Msg = 200, "success", ""
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerpsicopedagogia) DiagnosticoCreate(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	res := ResponseData{Error: true}
	var diagnostico diagnosticos.Diagnostico
	if err := c.BodyParser(&diagnostico); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "invalid request body",
		})
	}
	diagnosticoId, err := srv.SrvDiagnostico.Create(&diagnostico)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}
	res.Detalle = diagnosticoId
	res.Error = false
	res.Code, res.Type, res.Msg = 200, "success", ""
	return c.Status(http.StatusOK).JSON(res)
}

func (h *handlerpsicopedagogia) DiagnosticoUpdate(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	var diagnostico diagnosticos.Diagnostico
	if err := c.BodyParser(&diagnostico); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "invalid request body",
		})
	}

	if err := srv.SrvDiagnostico.Update(&diagnostico); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(diagnostico)
}

func (h *handlerpsicopedagogia) DiagnosticoDelete(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "invalid id parameter",
		})
	}

	if err := srv.SrvDiagnostico.Delete(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Diagnostico deleted successfully",
	})
}

// --------------------------------------------------
// Handlers para Citas
// --------------------------------------------------
func (h *handlerpsicopedagogia) GetAllCitas(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	// page := c.QueryInt("page", 1)
	// pageSize := c.QueryInt("pageSize", 10)

	citas, err := srv.SrvCitas.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{
			true, "Error obteniendo citas", nil, fiber.StatusInternalServerError, "error",
		})
	}
	return c.JSON(ResponseData{false, "Citas obtenidas", citas, fiber.StatusOK, "success"})
}

func (h *handlerpsicopedagogia) GetCitaByID(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	id, _ := c.ParamsInt("id")

	cita, err := srv.SrvCitas.GetByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{
			true, "Error obteniendo la cita", nil, fiber.StatusInternalServerError, "error",
		})
	}
	return c.JSON(ResponseData{false, "Cita obtenida", cita, fiber.StatusOK, "success"})
}

func (h *handlerpsicopedagogia) CreateCita(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	var cita RequestCita

	if err := c.BodyParser(&cita); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{
			true, "Datos inválidos", nil, fiber.StatusBadRequest, "error",
		})
	}

	idCita, err := srv.SrvCitas.Create(&citas.Cita{
		DNI:         cita.DNI,
		Nombre:      cita.Nombre,
		Apellido:    cita.Apellido,
		Facultad:    cita.Facultad,
		FechaInicio: cita.FechaInicio,
		FechaFin:    cita.FechaInicio.Add(time.Hour),
		Estado:      "Pendiente",
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{
			true, err.Error(), nil, fiber.StatusInternalServerError, "error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(ResponseData{
		false, "Cita creada", map[string]int64{"id": idCita}, fiber.StatusCreated, "success",
	})
}

func (h *handlerpsicopedagogia) UpdateCita(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	var cita RequestCita

	if err := c.BodyParser(&cita); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseData{
			true, "Datos inválidos", nil, fiber.StatusBadRequest, "error",
		})
	}

	err := srv.SrvCitas.Update(&citas.Cita{
		DNI:         cita.DNI,
		Nombre:      cita.Nombre,
		Apellido:    cita.Apellido,
		Facultad:    cita.Facultad,
		FechaInicio: cita.FechaInicio,
		FechaFin:    cita.FechaFin,
		Estado:      cita.Estado,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{
			true, err.Error(), nil, fiber.StatusInternalServerError, "error",
		})
	}

	return c.Status(fiber.StatusOK).JSON(ResponseData{
		false, "Cita actualizada", nil, fiber.StatusOK, "success",
	})
}

func (h *handlerpsicopedagogia) DeleteCita(c *fiber.Ctx) error {
	srv := psicopedagogia.NewServerpsicopedagogia(h.db, nil, h.txID)
	id, _ := c.ParamsInt("id")

	if err := srv.SrvCitas.Delete(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseData{
			true, err.Error(), nil, fiber.StatusInternalServerError, "error",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(ResponseData{
		false, "Cita eliminada", nil, fiber.StatusNoContent, "success",
	})
}
