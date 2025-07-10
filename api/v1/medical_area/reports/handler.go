package reports

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/middleware"
	"dbu-api/internal/models"
	"dbu-api/pkg/orchestrator/low_code_medical_area"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type handlerMedicalArea struct {
	db   *sqlx.DB
	txID string
}

// GetReportNursingConsultation godoc
// @Summary Obtener reporte de consultas de enfermería
// @Description Método que permite obtener el reporte de consultas de enfermería
// @Tags Reportes
// @Accept json
// @Produce json
// @Param area_medica query string true "Área médica (obligatoria)"
// @Param fecha_inicio query string false "Fecha de inicio (opcional) - formato: YYYY-MM-DD"
// @Param fecha_fin query string false "Fecha de fin (opcional) - formato: YYYY-MM-DD"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/reporte/consultas [get]
func (h *handlerMedicalArea) GetReportConsultation(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	areaMedica := c.Query("area_medica")
	if areaMedica == "" {
		logger.Error.Println("missing id query param")
		res.Code, res.Type, res.Msg = 1, "", "Medical Area is required"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	validAreas := map[string]bool{
		"enfermería":  true,
		"odontología": true,
		"medicina":    true,
	}

	if !validAreas[areaMedica] {
		logger.Error.Printf("invalid area_medica value: %s", areaMedica)
		res.Code, res.Type, res.Msg = 2, "", "Invalid medical area. Valid values: enfermería, odontología, medicina"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	fechaInicio := c.Query("fecha_inicio")
	fechaFin := c.Query("fecha_fin")

	srv := low_code_medical_area.NewReportsMedicalArea(h.db, user, h.txID)

	data, code, err := srv.GetReportMedicalConsultationByMedicalAreaLowCode(areaMedica, fechaInicio, fechaFin)
	if err != nil {
		logger.Error.Printf("couldn't get report by date, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", ""
	return c.Status(http.StatusOK).JSON(res)
}

// GetReportNursingConsultation godoc
// @Summary Obtener reporte de consultas de enfermería
// @Description Método que permite obtener el reporte de consultas de enfermería
// @Tags Reportes
// @Accept json
// @Produce json
// @Param area_medica query string true "Área médica (obligatoria)"
// @Param fecha_inicio query string false "Fecha de inicio (opcional) - formato: YYYY-MM-DD"
// @Param fecha_fin query string false "Fecha de fin (opcional) - formato: YYYY-MM-DD"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/reporte/consultas [get]
func (h *handlerMedicalArea) GetReportNursingFrame(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	numberFrame := c.Query("numero_cuadro")
	if numberFrame == "" {
		logger.Error.Println("missing id query param")
		res.Code, res.Type, res.Msg = 1, "", "Number frame is required"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	validFrame := map[string]bool{
		"1": true,
		"2": true,
		"3": true,
		"4": true,
		"5": true,
		"6": true,
		"7": true,
		"8": true,
		"9": true,
	}

	if !validFrame[numberFrame] {
		logger.Error.Printf("invalid area_medica value: %s", numberFrame)
		res.Code, res.Type, res.Msg = 2, "", "Invalid number frame, Valid values: 1,2,3,4,5,6,7,8,9"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	months := c.Query("meses")
	year := c.Query("anio")
	yearStart := c.Query("anio_inicial")
	yearEnd := c.Query("anio_final")

	srv := low_code_medical_area.NewReportsMedicalArea(h.db, user, h.txID)

	data, code, err := srv.GetReportNursingFrameLowCode(numberFrame, months, year, yearStart, yearEnd)
	if err != nil {
		logger.Error.Printf("couldn't get report, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", ""
	return c.Status(http.StatusOK).JSON(res)
}

// GetReportNursingConsultation godoc
// @Summary Obtener reporte de consultas de enfermería
// @Description Método que permite obtener el reporte de consultas de enfermería
// @Tags Reportes
// @Accept json
// @Produce json
// @Param area_medica query string true "Área médica (obligatoria)"
// @Param fecha_inicio query string false "Fecha de inicio (opcional) - formato: YYYY-MM-DD"
// @Param fecha_fin query string false "Fecha de fin (opcional) - formato: YYYY-MM-DD"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/reporte/consultas [get]
func (h *handlerMedicalArea) GetReportDentistryFrame(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	numeroCuadro := c.Query("numero_cuadro")
	if numeroCuadro == "" {
		logger.Error.Println("missing id query param")
		res.Code, res.Type, res.Msg = 1, "", "Number frame is required"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	validFrame := map[string]bool{
		"1": true,
		"2": true,
	}

	if !validFrame[numeroCuadro] {
		logger.Error.Printf("invalid area_medica value: %s", numeroCuadro)
		res.Code, res.Type, res.Msg = 2, "", "Invalid number frame, Valid values: 1,2"
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	meses := c.Query("meses")
	anio := c.Query("anio")

	srv := low_code_medical_area.NewReportsMedicalArea(h.db, user, h.txID)

	data, code, err := srv.GetReportDentistryFrameLowCode(numeroCuadro, meses, anio)
	if err != nil {
		logger.Error.Printf("couldn't get report by date, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", ""
	return c.Status(http.StatusOK).JSON(res)
}

// GetReportNursingConsultation godoc
// @Summary Obtener reporte de consultas de enfermería
// @Description Método que permite obtener el reporte de consultas de enfermería
// @Tags Reportes
// @Accept json
// @Produce json
// @Param area_medica query string true "Área médica (obligatoria)"
// @Param fecha_inicio query string false "Fecha de inicio (opcional) - formato: YYYY-MM-DD"
// @Param fecha_fin query string false "Fecha de fin (opcional) - formato: YYYY-MM-DD"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Failure 202 {object} models.Response
// @Router /api/v1/area_medica/reporte/consultas [get]
func (h *handlerMedicalArea) GetReportMedicalFrame(c *fiber.Ctx) error {
	res := models.Response{Error: true}

	bearer := c.Get("Authorization")

	user, err := middleware.GetUser(bearer, h.db)
	if err != nil {
		logger.Error.Printf("Unauthenticated user, error: %v", err)
		res.Code, res.Type, res.Msg = 9, "error", "unauthenticated"
		return c.Status(http.StatusUnauthorized).JSON(res)
	}

	numeroCuadro := c.Query("numero_cuadro")
	if numeroCuadro == "" {
		logger.Error.Println("missing id query param")
		res.Code, res.Type, res.Msg = 1, "", "Number frame is required"
		return c.Status(http.StatusBadRequest).JSON(res)
	}

	validFrame := map[string]bool{
		"1":  true,
		"3":  true,
		"8":  true,
		"12": true,
	}

	if !validFrame[numeroCuadro] {
		logger.Error.Printf("invalid area_medica value: %s", numeroCuadro)
		res.Code, res.Type, res.Msg = 2, "", "Invalid number frame, Valid values: 1,3,8,12"
		return c.Status(http.StatusBadRequest).JSON(res)
	}
	meses := c.Query("meses")
	anio := c.Query("anio")

	srv := low_code_medical_area.NewReportsMedicalArea(h.db, user, h.txID)

	data, code, err := srv.GetReportMedicalFrameLowCode(numeroCuadro, meses, anio)
	if err != nil {
		logger.Error.Printf("couldn't get report by date, error: %v", err)
		res.Code, res.Type, res.Msg = code, "", ""
		return c.Status(http.StatusAccepted).JSON(res)
	}

	res.Data = data
	res.Error = false
	res.Code, res.Type, res.Msg = 29, "", ""
	return c.Status(http.StatusOK).JSON(res)
}
