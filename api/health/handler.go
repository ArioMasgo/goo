package health

import (
	"dbu-api/internal/models"
	"dbu-api/pkg/orchestrator/response_messages"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type handlerHealth struct {
	db   *sqlx.DB
	txID string
	msg  response_messages.Message
}

// Health godoc
// @Summary Verificar estado del servicio
// @Description Endpoint para verificar el estado y conectividad del servicio
// @Tags Health
// @Accept json
// @Produce json
// @Success 201 {object} models.Response "Sistema conectado"
// @Failure 500 {object} models.Response "Error interno del servidor"
// @Router /health [GET]
func (h *handlerHealth) Health(c *fiber.Ctx) error {
	res := models.Response{Error: false, Data: "Sistema conectado"}

	res.Code, res.Type, res.Msg = h.msg.GetByCode(210)
	return c.Status(fiber.StatusOK).JSON(res)
}
