package health

import (
	"dbu-api/pkg/orchestrator/response_messages"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterHealth(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerHealth{db: db, txID: txID, msg: response_messages.NewMsg(db)}
	v1 := app.Group("/health")
	v1.Get("/", h.Health)

}
