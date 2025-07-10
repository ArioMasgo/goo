package medical_consultation

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
	medicalArea.Post("/consulta_medicina", h.CreateMedicalConsultation)
	medicalArea.Put("/consulta_medicina", h.UpdateMedicalConsultation)
	medicalArea.Get("/consultas_medicina", h.GetAllMedicalConsultation)
	medicalArea.Get("/consulta_medicina/:id", h.GetMedicalConsultationByID)
	medicalArea.Delete("/consulta_medicina/:id", h.DeleteMedicalConsultation)

	medicalArea.Get("/consultas_medicina/paciente/:id", h.GetMedicalConsultationByIDPatient)
	medicalArea.Get("/consultas_medicina/paciente/dni/:id", h.GetMedicalConsultationByDNIPatient)
}
