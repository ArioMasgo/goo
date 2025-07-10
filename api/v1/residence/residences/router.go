package residences

import (
	"dbu-api/internal/middleware"
	"dbu-api/pkg/orchestrator/response_messages"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterResidences(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerResidences{db: db, txID: txID, msg: response_messages.NewMsg(db)}
	v1 := app.Group("/v1")
	residences := v1.Group("/residencias")
	residences.Use(middleware.JWTProtected())
	residences.Post("/", h.CreateResidences)
	residences.Put("/", h.UpdateResidences)
	residences.Get("/", h.GetAllResidences)
	residences.Delete("/:id", h.DeleteResidences)

	residences.Get("/:id/alumnos", h.GetAllStudentsByResidence)
	residences.Put("/:id/configuracion", h.UpdateConfigResidence)
}
