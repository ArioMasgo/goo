package room_assignment

import (
	"dbu-api/internal/middleware"
	"dbu-api/pkg/orchestrator/response_messages"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterRoomAssignment(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerRoomAssignment{db: db, txID: txID, msg: response_messages.NewMsg(db)}
	v1 := app.Group("/v1")
	submissions := v1.Group("/automatizacion")
	submissions.Use(middleware.JWTProtected())
	submissions.Post("/asignacion-cuartos", h.ExecuteRoomAssignment)
}
