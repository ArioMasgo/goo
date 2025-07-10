package nursing_consultation

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
	medicalArea.Post("/consulta_enfermeria", h.CreateNursingConsultation)
	medicalArea.Put("/consulta_enfermeria", h.UpdateNursingConsultation)
	medicalArea.Get("/consultas_enfermeria", h.GetAllNursingConsultation)
	medicalArea.Get("/consulta_enfermeria/:id", h.GetNursingConsultationByID)
	medicalArea.Delete("/consulta_enfermeria/:id", h.DeleteNursingConsultation)

	medicalArea.Get("/consultas_enfermeria/paciente/:id", h.GetNursingConsultationByIDPatient)
	medicalArea.Get("/consultas_enfermeria/paciente/dni/:id", h.GetNursingConsultationByDNIPatient)

	medicalArea.Get("/consultas_enfermeria/vacunas", h.GetAllTypesVaccines)
	medicalArea.Get("/consultas_enfermeria/vacunas_requeridas/paciente/:id", h.GetTypesVaccineRequired)
	medicalArea.Get("/consultas_enfermeria/vacunas/paciente/:dni", h.GetAllVaccinesByPatientDni)
}
