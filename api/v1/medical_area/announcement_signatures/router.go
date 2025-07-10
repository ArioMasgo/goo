package announcement_signatures

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterMedicalArea(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerMedicalArea{db: db, txID: txID}
	v1 := app.Group("/v1")
	v1.Get("/area_medica/firmas/paciente/dni/:dni", h.GetAnnouncementSignaturesByDNI)
}
