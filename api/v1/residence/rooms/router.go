package rooms

import (
	"dbu-api/internal/middleware"
	"dbu-api/pkg/orchestrator/response_messages"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterRooms(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerRooms{db: db, txID: txID, msg: response_messages.NewMsg(db)}
	v1 := app.Group("/v1")
	residence := v1.Group("/residencias")
	residence.Use(middleware.JWTProtected())

	residence.Put("/cuartos", h.UpdateRooms)
	residence.Get("/:id/cuartos", h.GetAllRoomsByResidence)
	residence.Delete("/cuartos/:id/eliminar-asignacion", h.DeleteAssignmentRoom)
	residence.Post("/cuartos/:id/asignar", h.AssignmentRoom)
}
