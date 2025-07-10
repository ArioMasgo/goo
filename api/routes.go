package api

import (
	"dbu-api/api/health"
	authentication "dbu-api/api/v1/auth"
	"dbu-api/api/v1/automation/room_assignment"
	"dbu-api/api/v1/medical_area/payments"
	"dbu-api/api/v1/medical_area/reports"

	"dbu-api/api/v1/general_visit/visita_general"
	"dbu-api/api/v1/general_visit/visita_residente"
	"dbu-api/api/v1/medical_area/announcement_signatures"
	"dbu-api/api/v1/medical_area/dentistry_consultation"
	"dbu-api/api/v1/medical_area/exam_toxicologico"
	"dbu-api/api/v1/medical_area/medical_consultation"
	"dbu-api/api/v1/medical_area/nursing_consultation"
	"dbu-api/api/v1/medical_area/nursing_consultation_assignment"
	"dbu-api/api/v1/medical_area/patients"
	"dbu-api/api/v1/psicopedagogia"
	"dbu-api/api/v1/residence/residences"
	"dbu-api/api/v1/residence/rooms"
	"dbu-api/api/v1/submission/submissions"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Se cargan los loggerHttp, y los allowedOrigins (registrosHttp) (permisos de origen)
func routes(loggerHttp bool, allowedOrigins string, db *sqlx.DB) *fiber.App {
	app := fiber.New()

	prometheus := fiberprometheus.New("DBU API REST")
	prometheus.RegisterAt(app, "/metrics")

	app.Get("/doc/*", swagger.New(swagger.Config{
		URL:         "/doc/doc.json",
		DeepLinking: false,
	}))

	app.Use(recover.New())
	app.Use(prometheus.Middleware)
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowHeaders: "Origin, X-Requested-With, Content-Type, Accept, Authorization, signature, submission_id",
		AllowMethods: "GET,POST,PUT,PATCH,DELETE",
	}))

	if loggerHttp {
		app.Use(logger.New())
	}

	TxID := uuid.New().String()

	loadRoutes(app, TxID, db)

	return app
}

// Aqui se cargan las direcciones o las ubicaciones de las funciones Handler
func loadRoutes(app *fiber.App, TxID string, db *sqlx.DB) {
	submissions.RouterSubmissions(app, db, TxID)
	residences.RouterResidences(app, db, TxID)
	rooms.RouterRooms(app, db, TxID)
	patients.RouterMedicalArea(app, db, TxID)
	nursing_consultation.RouterMedicalArea(app, db, TxID)
	dentistry_consultation.RouterMedicalArea(app, db, TxID)
	medical_consultation.RouterMedicalArea(app, db, TxID)
	authentication.RouterAuthentication(app, db, TxID)
	room_assignment.RouterRoomAssignment(app, db, TxID)
	health.RouterHealth(app, db, TxID)
	psicopedagogia.RouterPsicopedagogia(app, db, TxID)
	visita_general.RouterVisitaGeneral(app, db, TxID)
	visita_residente.RouterVisitaResidente(app, db, TxID)
	nursing_consultation_assignment.RouterMedicalArea(app, db, TxID)
	announcement_signatures.RouterMedicalArea(app, db, TxID)
	reports.RouterMedicalArea(app, db, TxID)
	exam_toxicologico.RouterMedicalArea(app, db, TxID)
	payments.Router(app, db, TxID)
}

