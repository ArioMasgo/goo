package reports

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
	medicalArea.Get("/reporte/consultas", h.GetReportConsultation)
	medicalArea.Get("/reporte/enfermeria", h.GetReportNursingFrame)
	medicalArea.Get("/reporte/odontologia", h.GetReportDentistryFrame)
	medicalArea.Get("/reporte/medicina", h.GetReportMedicalFrame)
}
