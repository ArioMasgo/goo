package dentistry_consultation

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
	medicalArea.Post("/consulta_odontologia", h.CreateDentistryConsultation)
	medicalArea.Put("/consulta_odontologia", h.UpdateDentistryConsultation)
	medicalArea.Get("/consultas_odontologia", h.GetAllDentistryConsultation)
	medicalArea.Get("/consulta_odontologia/:id", h.GetDentistryConsultationByID)
	medicalArea.Delete("/consulta_odontologia/:id", h.DeleteDentistryConsultation)

	medicalArea.Get("/consultas_odontologia/paciente/:id", h.GetDentistryConsultationByIDPatient)
	medicalArea.Get("/consultas_odontologia/paciente/dni/:id", h.GetDentistryConsultationByDNIPatient)
}
