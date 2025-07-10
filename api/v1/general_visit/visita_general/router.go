package visita_general

import (
	"dbu-api/internal/middleware"
	"dbu-api/pkg/orchestrator/response_messages"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func RouterVisitaGeneral(app *fiber.App, db *sqlx.DB, txID string) {
	h := handlerVisitaGeneral{db: db, txID: txID, msg: response_messages.NewMsg(db)}
	v1 := app.Group("/v1")
	visitaGeneral := v1.Group("/visita-general")
	visitaGeneral.Use(middleware.JWTProtected())
	
	visitaGeneral.Get("/departamentos", h.GetAllDepartments)
	visitaGeneral.Get("/departamentos/:departmentId/provincias", h.GetProvincesByDepartment)
	visitaGeneral.Get("/provincias/:provinceId/distritos", h.GetDistrictsByProvince)
	visitaGeneral.Get("/ubicaciones", h.GetLocationHierarchy)
	
	visitaGeneral.Post("/", h.CreateVisitaGeneral)
	visitaGeneral.Get("/", h.GetAllVisitaGeneral)
	visitaGeneral.Put("/", h.UpdateVisitaGeneral)
	
	visitaGeneral.Get("/:id", h.GetVisitaGeneralByID)
	visitaGeneral.Delete("/:id", h.DeleteVisitaGeneral)
}