package patients

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
	medicalArea.Post("/paciente", h.CreatePatients)
	medicalArea.Put("/paciente", h.UpdatePatients)
	medicalArea.Get("/pacientes", h.GetAllPatients)
	medicalArea.Post("/pacientes/get", h.GetPatients)
	medicalArea.Get("/paciente/:id", h.GetPatientsByID)
	medicalArea.Delete("/paciente/:id", h.DeletePatients)
	//medicalArea.Post("/", h.CreateAttention)

	medicalArea.Get("/paciente/dni/:dni", h.GetPatientsByDNI)
}
