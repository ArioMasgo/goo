package payments

import (
	"dbu-api/internal/middleware"
	"dbu-api/pkg/orchestrator/response_messages"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Router(app *fiber.App, db *sqlx.DB, txID string) {
	h := handler{db: db, txID: txID, msg: response_messages.NewMsg(db)}
	v1 := app.Group("/v1")
	medicalArea := v1.Group("/area_medica")
	medicalArea.Use(middleware.JWTProtected())
	medicalArea.Post("/odontologia/pagos/procedimiento/buscar", h.SearchPayment)
}
