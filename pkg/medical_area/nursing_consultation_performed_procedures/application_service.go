package nursing_consultation_performed_procedures

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerPerformedProcedures interface {
	CreatePerformedProcedures(id, consulta_enfermeria_id, procedimiento, numero_recibo, comentarios, costo, fecha_pago string, areaSolicitante, especialistaSolicitante *string) (*PerformedProcedures, int, error)
	UpdatePerformedProcedures(id, consulta_enfermeria_id, procedimiento, numero_recibo, comentarios, costo, fecha_pago string, areaSolicitante, especialistaSolicitante *string) (*PerformedProcedures, int, error)
	DeletePerformedProcedures(id string) (int, error)
	DeletePerformedProceduresByIDConsultation(id string) (int, error)
	GetPerformedProceduresByID(id string) (*PerformedProcedures, int, error)
	GetAllPerformedProcedures() ([]*PerformedProcedures, error)
	GetPerformedProceduresByIDConsultation(id string) ([]*PerformedProcedures, int, error)
	GetPerformedProceduresExcel(fecha_inicio, fecha_fin string) ([]*models.PerformedProceduresExcel, error)
}

type service struct {
	repository ServicesPerformedProceduresRepository
	user       *models.User
	txID       string
}

func NewPerformedProceduresService(repository ServicesPerformedProceduresRepository, user *models.User, TxID string) PortsServerPerformedProcedures {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreatePerformedProcedures(id, consulta_enfermeria_id, procedimiento, numero_recibo, comentarios, costo, fecha_pago string, areaSolicitante, especialistaSolicitante *string) (*PerformedProcedures, int, error) {

	m := NewPerformedProcedures(id, consulta_enfermeria_id, procedimiento, numero_recibo, comentarios, costo, fecha_pago, areaSolicitante, especialistaSolicitante)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create create performed procedures :", err)
		return m, 3, err
	}

	return m, 29, nil
}

func (s *service) UpdatePerformedProcedures(id, consulta_enfermeria_id, procedimiento, numero_recibo, comentarios, costo, fecha_pago string, areaSolicitante, especialistaSolicitante *string) (*PerformedProcedures, int, error) {
	m := NewPerformedProcedures(id, consulta_enfermeria_id, procedimiento, numero_recibo, comentarios, costo, fecha_pago, areaSolicitante, especialistaSolicitante)
	valid, err := m.valid()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't meet validations:", err)
		return nil, 15, err
	}
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return nil, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update performed procedures :", err)
		return nil, 18, err
	}
	return m, 29, nil
}

func (s *service) DeletePerformedProcedures(id string) (int, error) {
	if err := uuid.Validate(id); err != nil {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return 15, fmt.Errorf("id is required")
	}

	if err := s.repository.delete(id); err != nil {
		if err.Error() == "rows affected error" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update row:", err)
		return 20, err
	}
	return 28, nil
}

func (s *service) DeletePerformedProceduresByIDConsultation(id string) (int, error) {
	if err := uuid.Validate(id); err != nil {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return 15, fmt.Errorf("id is required")
	}

	if err := s.repository.deleteByIDConsultation(id); err != nil {
		if err.Error() == "rows affected error" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update row:", err)
		return 20, err
	}
	return 28, nil
}

func (s *service) GetPerformedProceduresByID(id string) (*PerformedProcedures, int, error) {
	if err := uuid.Validate(id); err != nil {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return nil, 15, fmt.Errorf("id is required")
	}
	m, err := s.repository.getByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllPerformedProcedures() ([]*PerformedProcedures, error) {
	return s.repository.getAll()
}

func (s *service) GetPerformedProceduresByIDConsultation(id string) ([]*PerformedProcedures, int, error) {
	if err := uuid.Validate(id); err != nil {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return nil, 15, fmt.Errorf("id is required")
	}
	m, err := s.repository.GetByIDConsultation(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByIDPatient row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetPerformedProceduresExcel(fecha_inicio, fecha_fin string) ([]*models.PerformedProceduresExcel, error) {
	m, err := s.repository.getAllByDateExcel(fecha_inicio, fecha_fin)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get performed procedures excel:", err)
		return nil, err
	}
	return m, nil
}
