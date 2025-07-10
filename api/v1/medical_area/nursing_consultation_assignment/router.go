package nursing_consultation_assignment

import (
	"dbu-api/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterMedicalArea(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerMedicalArea{db: db, txID: txID}
	v1 := app.Group("/v1")
	medicalArea := v1.Group("/area_medica")
	medicalArea.Use(middleware.JWTProtected())
	medicalArea.Post("/consulta_enfermeria_asignacion", h.CreateConsultationAssignment)
	medicalArea.Put("/consulta_enfermeria_asignacion", h.UpdateConsultationAssignment)
	medicalArea.Get("/consulta_enfermeria_asignacion", h.GetAllConsultationAssignment)
	medicalArea.Get("/consulta_enfermeria_asignacion/:id", h.GetConsultationAssignmentByID)
	medicalArea.Delete("/consulta_enfermeria_asignacion/:id", h.DeleteConsultationAssignment)
}
