package payments

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/middleware"
	"dbu-api/internal/models"
	"dbu-api/pkg/medical_area"
	"dbu-api/pkg/orchestrator/response_messages"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handler struct {
	db   *sqlx.DB
	txID string
	msg  response_messages.Message
}

func (h *handler) SearchPayment(c *fiber.Ctx) error {
	req := RequestPayment{}
	res := models.Response{Error: true}
	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(53)
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	if err := c.BodyParser(&req); err != nil {
		logger.Error.Printf("couldn't parse body request, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	isValid, err := req.valid()
	if err != nil {
		logger.Error.Printf("couldn't validate body request, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	if !isValid {
		logger.Error.Println("couldn't validate body request")
		res.Code, res.Type, res.Msg = h.msg.GetByCode(1)
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := medical_area.NewServerMedicalArea(h.db, user, h.txID)
	paymentConcept, code, err := srv.SrvPaymentsConcept.Search(req.Dni, "odontologia", req.TipoServicio, req.NombreServicio, req.Recibo)
	if err != nil {
		logger.Error.Printf("couldn't create dentistry consultation, error: %v", err)
		res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
		return c.Status(http.StatusInternalServerError).JSON(res)
	}

	res.Data = paymentConcept
	res.Error = false
	res.Code, res.Type, res.Msg = h.msg.GetByCode(code)
	return c.Status(http.StatusOK).JSON(res)
}
