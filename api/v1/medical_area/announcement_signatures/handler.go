package announcement_signatures

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"dbu-api/pkg/medical_area"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerMedicalArea struct {
	db   *sqlx.DB
	txID string
}

// GetPatientsByID godoc
// @Summary Obtiene una instancia de paciente por su id
// @Description Método que permite obtener una instancia del objeto paciente en la base de datos por su DNI
// @tags Pacientes
// @Accept json
// @Produce json
// @Param	id	path int true "Paciente ID"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/paciente/:id [GET]
func (h *handlerMedicalArea) GetAnnouncementSignaturesByDNI(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	dniStr := c.Params("dni")
	if dniStr == "" {
		logger.Error.Println("couldn't parse id request")
		res.Code, res.Type, res.Msg = 1, "", ""
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	srv := medical_area.NewServerMedicalArea(h.db, nil, h.txID)
	data, code, err := srv.SrvAnnouncementSignatures.GetAnnouncementSignaturesByDNIPatient(dniStr)
	if err != nil {
		logger.Error.Printf("couldn't get patient by DNI, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", ""
	return c.Status(http.StatusOK).JSON(res)
}
