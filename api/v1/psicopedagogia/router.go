package psicopedagogia

import (
	"dbu-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterPsicopedagogia(app *fiber.App, db *sqlx.DB, txID string) {
	handler := handlerpsicopedagogia{db: db, txID: txID}
	v1 := app.Group("/v1/psicopedagogia")
	// 🔹 Rutas para Estudiante
	estudiantesGroup := v1.Group("/estudiantes")
	estudiantesGroup.Get("/:dni", handler.GetEstudianteByDni)
	estudiantesGroup.Get("/basic/:dni", handler.GetEstudianteByNameDni)

	// 🔹 Rutas para Participantes
	participantesGroup := v1.Group("/participantes")
	participantesGroup.Use(middleware.JWTProtected())
	participantesGroup.Get("/", handler.GetAll)
	participantesGroup.Get("/:id", handler.GetByID)
	participantesGroup.Post("/", handler.Create)
	participantesGroup.Put("/:id", handler.Update)
	participantesGroup.Delete("/:id", handler.Delete)
	participantesGroup.Post("/search-result", handler.SearchParticipants)

	// 🔹 Rutas para Preguntas
	preguntasGroup := v1.Group("/preguntas")
	preguntasGroup.Get("/", handler.GetAllPreguntas)
	preguntasGroup.Use(middleware.JWTProtected())
	preguntasGroup.Get("/:id", handler.GetPreguntaByID)
	preguntasGroup.Post("/", handler.CreatePregunta)
	preguntasGroup.Put("/:id", handler.UpdatePregunta)
	preguntasGroup.Delete("/:id", handler.DeletePregunta)

	// 🔹 Rutas para Encuestas
	encuestasGroup := v1.Group("/encuestas")
	encuestasGroup.Get("/active-srq", handler.HasActiveSRQ)
	encuestasGroup.Use(middleware.JWTProtected())
	encuestasGroup.Use(middleware.JWTProtected())
	encuestasGroup.Get("/", handler.GetAllEncuestas)
	encuestasGroup.Get("/:id", handler.GetEncuestaByID)
	encuestasGroup.Post("/", handler.CreateEncuesta)
	encuestasGroup.Put("/:id", handler.UpdateEncuesta)
	encuestasGroup.Delete("/:id", handler.DeleteEncuesta)

	// 🔹 Rutas para Respuestas
	respuestasGroup := v1.Group("/respuestas")
	respuestasGroup.Use(middleware.JWTProtected())
	respuestasGroup.Get("/", handler.GetAllRespuestas)
	respuestasGroup.Get("/:id", handler.GetRespuestaByID)
	respuestasGroup.Post("/", handler.CreateRespuesta)
	respuestasGroup.Put("/:id", handler.UpdateRespuesta)
	respuestasGroup.Delete("/:id", handler.DeleteRespuesta)
	respuestasGroup.Get("/participante/:idParticipante", handler.GetResponsesPerParticipant)
	respuestasGroup.Post("/participante", handler.GetAllByParticipanteIdAndNumeroAtencion)

	// 🔹 Rutas para Historial de Encuestas
	historialGroup := v1.Group("/historial")
	historialGroup.Post("/has-student-responded", handler.HasStudentResponded)
	historialGroup.Post("/", handler.CreateHistorialWithAnswers)
	historialGroup.Post("/save", handler.SaveHistory)
	historialGroup.Get("/key-url-exists", handler.KeyUrlExists)
	historialGroup.Use(middleware.JWTProtected())
	historialGroup.Post("/filtered", handler.GetHistorialFiltered)
	historialGroup.Get("/latest", handler.GetLatestHistorial)

	// 🔹 Rutas para PDFs
	pdfsGroup := v1.Group("/pdfs")
	pdfsGroup.Use(middleware.JWTProtected())
	pdfsGroup.Get("/srq/:id", handler.GeneratePDF_SRQ_Handler)

	// 🔹 Rutas para Diagnostico
	diagnosticoGroup := v1.Group("/diagnostico")
	diagnosticoGroup.Use(middleware.JWTProtected())
	diagnosticoGroup.Get("/", handler.GetDiagnosticoAll)
	diagnosticoGroup.Post("/", handler.DiagnosticoCreate)

	// 🔹 Rutas para Citas
	citasGroup := v1.Group("/citas")
	citasGroup.Post("/", handler.CreateCita)
	citasGroup.Use(middleware.JWTProtected())
	citasGroup.Get("/", handler.GetAllCitas)
	citasGroup.Get("/:id", handler.GetCitaByID)
	citasGroup.Put("/update", handler.UpdateCita)
	citasGroup.Delete("/:id", handler.DeleteCita)

}
