package visita_residente

import (
	"dbu-api/internal/middleware"
	"dbu-api/pkg/orchestrator/response_messages"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterVisitaResidente(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerVisitaResidente{db: db, txID: txID, msg: response_messages.NewMsg(db)}
	v1 := app.Group("/v1")
	visitaResidente := v1.Group("/visita-residente")
	visitaResidente.Use(middleware.JWTProtected())
	visitaResidente.Post("/", h.CreateVisitaResidente)
	visitaResidente.Get("/", h.GetAllVisitaResidente)
	visitaResidente.Get("/alumnos-pendientes", h.GetAlumnosPendientesVisita)
	visitaResidente.Get("/alumnos-pendientes/departamento", h.GetAlumnosPendientesPorDepartamento)
	visitaResidente.Get("/alumnos-completos", h.GetTodosAlumnosPorConvocatoria)
	visitaResidente.Get("/estadisticas", h.GetEstadisticasVisitas)
	visitaResidente.Get("/estadisticas/escuela-profesional", h.GetEstadisticasPorEscuelaProfesional)
	visitaResidente.Get("/estadisticas/lugar-procedencia", h.GetEstadisticasPorLugarProcedencia)
	visitaResidente.Get("/existe-alumno", h.ExisteVisitaAlumno)
	visitaResidente.Get("/:id", h.GetVisitaResidenteByID)
	visitaResidente.Put("/:id", h.UpdateVisitaResidente)
	visitaResidente.Delete("/:id", h.DeleteVisitaResidente)
}
