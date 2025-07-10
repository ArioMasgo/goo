package third_parties

import (
	"dbu-api/internal/middleware"
	"dbu-api/pkg/orchestrator/response_messages"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterThirdParty(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerThirdParty{db: db, txID: txID, msg: response_messages.NewMsg(db)}
	v1 := app.Group("/v1")
	submissions := v1.Group("/third_parties")
	submissions.Use(middleware.JWTProtected())
	submissions.Get("/:id/alumnos-aceptados", h.StudentDebts)
}
