package historialencuesta

import (
	"dbu-api/internal/logger"
	"dbu-api/pkg/psicopedagogia/participantes"
	"dbu-api/pkg/psicopedagogia/respuestas"
	"errors"
	"fmt"
	"strings"
	"time"
)

type PortsHistorialService interface {
	GuardarEncuestaWithAnswers(historial *HistorialEncuesta, participante *participantes.Participante, respuestas []respuestas.Respuesta) (string, error)
	GetHistorialFiltered(dni, apellido, fechaInicio, fechaFin string) ([]Historial, error)
	HasStudentResponded(dni string, encuestaID int) (bool, error)
	GetLatestHistorial(limit, offset int) ([]Historial, error)
	SaveHistory(historial *Historial, participante *participantes.Participante) error
	KeyUrlExists(keyUrl string) (bool, error)
}

type service struct {
	repository       ServicesHistorialEncuestaRepository
	HistorialRepo    *HistorialRepository
	RespuestaRepo    *respuestas.RespuestaRepository
	ParticipanteRepo *participantes.ParticipanteRepository
	txID             string
}

func NewHistorialService(
	repository ServicesHistorialEncuestaRepository,
	historialRepo *HistorialRepository,
	respuestaRepo *respuestas.RespuestaRepository,
	participanteRepo *participantes.ParticipanteRepository,
	txID string,
) PortsHistorialService {
	return &service{
		repository:       repository,
		HistorialRepo:    historialRepo,
		RespuestaRepo:    respuestaRepo,
		ParticipanteRepo: participanteRepo,
		txID:             txID,
	}
}

func (s *service) GuardarEncuestaWithAnswers(historial *HistorialEncuesta, participante *participantes.Participante, respuestas []respuestas.Respuesta) (string, error) {
	tx, err := s.HistorialRepo.DB.Beginx()
	if err != nil {
		return "", err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	estadoDiagnostico, estadoEvaluacion, participanteID, numAtencion, err := s.ProcesarParticipante(participante, respuestas)
	if err != nil {
		logger.Error.Println(s.txID, " - format date:", err)
		return "", err
	}
	historial.IDParticipante = participanteID
	historial.FechaRespuesta = time.Now()
	historial.CreatedDate = timestamp
	historial.Diagnostico = estadoDiagnostico
	historial.EstadoEvaluacion = estadoEvaluacion
	historialID, err := s.repository.Create(historial)
	if err != nil {
		logger.Error.Println(s.txID, " - save history:", err)
		return "", errors.New("error al guardar historial")
	}

	for i := range respuestas {
		respuestas[i].IDParticipante = historial.IDParticipante
		respuestas[i].IDEncuesta = historial.IDEncuesta
		respuestas[i].IdHistorial = historialID
		respuestas[i].NumeroAtencion = numAtencion
		respuestas[i].CreatedAt = timestamp
		respuestas[i].UpdatedAt = timestamp
	}

	if err := s.RespuestaRepo.CreateBatch(respuestas); err != nil {
		logger.Error.Println(s.txID, " - save answers:", err)
		return "", errors.New("error al guardar respuestas")
	}

	return estadoEvaluacion, tx.Commit()
}

func (s *service) ProcesarParticipante(participante *participantes.Participante, respuestas []respuestas.Respuesta) (string, string, int, int, error) {
	existingParticipante, err := s.ParticipanteRepo.GetByDNI(participante.DNI)
	if err != nil {
		logger.Error.Println(s.txID, " - search participante:", err)
		return "", "", 0, 0, errors.New("error al buscar participante")
	}

	diagnosticoData := ProcessSurveyAnswers(respuestas)
	estadoDiagnostico := diagnosticoData["estado_diagnostico"].(string)
	estadoEvaluacion := "Pendiente"
	if estadoDiagnostico == "Estable" {
		estadoEvaluacion = "Aprobado"
	}

	if existingParticipante == nil {
		//Crear nuevo participante
		participante.Estado = "Activo"
		participante.NumeroAtencion++
		participante.Diagnostico = estadoDiagnostico
		participante.EstadoEvaluacion = estadoEvaluacion
		numeroDeAtencion := participante.NumeroAtencion
		participanteID, err := s.ParticipanteRepo.Create(participante)
		if err != nil {
			return "", "", 0, 0, errors.New("error al crear participante")
		}
		return estadoDiagnostico, estadoEvaluacion, participanteID, numeroDeAtencion, nil
	}

	//Actualizar participante existente
	fechaNacimiento, err := time.Parse(time.RFC3339, existingParticipante.FechaNacimiento)
	if err != nil {
		fechaNacimiento, err = time.Parse("2006-01-02", existingParticipante.FechaNacimiento)
		if err != nil {
			return "", "", 0, 0, errors.New("error parsing time")
		}
	}

	hoy := time.Now()
	edad := hoy.Year() - fechaNacimiento.Year()
	if hoy.Before(time.Date(hoy.Year(), fechaNacimiento.Month(), fechaNacimiento.Day(), 0, 0, 0, 0, hoy.Location())) {
		edad--
	}
	existingParticipante.Edad = edad
	existingParticipante.NumeroAtencion++
	existingParticipante.Diagnostico = estadoDiagnostico
	existingParticipante.EstadoEvaluacion = estadoEvaluacion
	numeroDeAtencion := existingParticipante.NumeroAtencion
	if err := s.ParticipanteRepo.Update(existingParticipante.ID, existingParticipante); err != nil {
		return "", "", 0, 0, fmt.Errorf("error al actualizar participante: %w", err)
	}

	return estadoDiagnostico, estadoEvaluacion, existingParticipante.ID, numeroDeAtencion, nil
}

func ProcessSurveyAnswers(respuestas []respuestas.Respuesta) map[string]interface{} {
	var score float64
	for _, respuesta := range respuestas {
		if strings.ToUpper(respuesta.Respuesta) == "SI" {
			score += 1.0
		}
	}
	estadoDiagnostico := "Estable"
	if score >= 5 && score <= 6 {
		estadoDiagnostico = "Moderado"
	} else if score >= 7 {
		estadoDiagnostico = "Grave"
	}
	return map[string]interface{}{
		"estado_diagnostico": estadoDiagnostico,
		"score":              score,
	}
}

func (s *service) GetHistorialFiltered(dni, apellido, fechaInicio, fechaFin string) ([]Historial, error) {
	if dni == "" && apellido == "" && fechaInicio == "" && fechaFin == "" {
		return nil, fmt.Errorf("se requiere al menos un criterio de búsqueda")
	}

	historial, err := s.repository.GetHistorialFiltered(dni, apellido, fechaInicio, fechaFin)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo historial: %w", err)
	}

	if len(historial) == 0 {
		return nil, fmt.Errorf("no se encontraron registros para los filtros proporcionados")
	}

	return historial, nil
}

func (s *service) HasStudentResponded(dni string, encuestaID int) (bool, error) {
	return s.repository.HasStudentResponded(dni, encuestaID)
}

func (s *service) GetLatestHistorial(limit, offset int) ([]Historial, error) {
	return s.repository.GetLatestHistorial(limit, offset)
}

func (s *service) SaveHistory(historial *Historial, participante *participantes.Participante) error {
	existingParticipante, err := s.ParticipanteRepo.GetByDNI(historial.DNI)
	if err != nil {
		return fmt.Errorf("error al buscar participante: %w", err)
	}
	var idPac int
	if existingParticipante == nil {
		participante.NumeroAtencion++
		idPac, err = s.ParticipanteRepo.Create(participante)
		if err != nil {
			return fmt.Errorf("error al crear participante: %w", err)
		}
	} else {
		existingParticipante.Direccion = participante.Direccion
		existingParticipante.NumTelefono = participante.NumTelefono
		existingParticipante.Edad = participante.Edad
		existingParticipante.ConQuienesVive = participante.ConQuienesVive
		existingParticipante.SemestreCursa = participante.SemestreCursa
		existingParticipante.ColegioProcedencia = participante.ColegioProcedencia
		existingParticipante.NumeroAtencion++
		err = s.ParticipanteRepo.Update(existingParticipante.ID, existingParticipante)
		if err != nil {
			return fmt.Errorf("error al actualizar participante: %w", err)
		}
		idPac = existingParticipante.ID
	}
	historialEncuesta := &HistorialEncuesta{
		IDParticipante:       idPac,
		FechaRespuesta:       time.Now(),
		Diagnostico:          "No Aplica",
		EstadoEvaluacion:     "No Aplica",
		EsSRQ:                historial.EsSRQ,
		NumTelefono:          historial.NumTelefono,
		ConQuienesVive:       historial.ConQuienesVive,
		SemestreCursa:        historial.SemestreCursa,
		Direccion:            historial.Direccion,
		QuienFinanciaCarrera: historial.QuienFinanciaCarrera,
		MotivoConsulta:       historial.MotivoConsulta,
		SituacionActual:      historial.SituacionActual,
		OtrosProcedimientos:  historial.OtrosProcedimientos,
		DiagnosticoID:        historial.DiagnosticoID,
		KeyUrl:               historial.KeyUrl,
	}

	_, err = s.repository.Create(historialEncuesta)
	if err != nil {
		logger.Error.Println(s.txID, " - save saveHistory:", err)
		return fmt.Errorf("error al guardar historial: %w", err)
	}

	return nil
}

func (s *service) KeyUrlExists(keyUrl string) (bool, error) {
	return s.repository.KeyUrlExists(keyUrl)
}
