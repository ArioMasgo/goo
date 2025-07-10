package exam_toxicologico

import (
	"dbu-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterMedicalArea(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerExamToxicologico{db: db, txID: txID}
	v1 := app.Group("/v1")
	medicalArea := v1.Group("/area_medica")
	medicalArea.Use(middleware.JWTProtected())
	
	// Rutas CRUD para registros toxicológicos
	medicalArea.Post("/examen_toxicologico", h.CreateRegistroToxicologico)
	medicalArea.Put("/examen_toxicologico/:id", h.UpdateRegistroToxicologico)
	
	medicalArea.Get("/examen_toxicologico/:id", h.GetRegistroToxicologicoByID)
	medicalArea.Delete("/examen_toxicologico/:id", h.DeleteRegistroToxicologico)
	medicalArea.Get("/examen_toxicologico", h.GetAllRegistrosToxicologicos)
	
	// Ruta especial para obtener estados por convocatoria
	medicalArea.Get("/examen_toxicologico/convocatoria/:convocatoria_id", h.GetEstadosByConvocatoria)
}
