package authentication

import (
	"dbu-api/pkg/orchestrator/response_messages"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterAuthentication(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerAuthentication{db: db, txID: txID, msg: response_messages.NewMsg(db)}
	v1 := app.Group("/v1")
	v1.Post("/login", h.Login)
	v1.Get("/refresh-token", h.Login)

}
