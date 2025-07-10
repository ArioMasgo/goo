package submissions

import (
	"dbu-api/internal/middleware"
	"dbu-api/pkg/orchestrator/response_messages"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterSubmissions(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerResidencias{db: db, txID: txID, msg: response_messages.NewMsg(db)}
	v1 := app.Group("/v1")
	submissions := v1.Group("/convocatorias")
	submissions.Use(middleware.JWTProtected())
	submissions.Get("/:id/alumnos-aceptados", h.StudentsBySubmissions)
	submissions.Get("/:id/reporte-residencias", h.StudentReportsBySubmissions)
}
