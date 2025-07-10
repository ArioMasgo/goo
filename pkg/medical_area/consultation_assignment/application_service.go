package consultation_assignment

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerConsultationAssignment interface {
	CreateConsultationAssignment(id, consulta_id, area_asignada string) (*ConsultationAssignment, int, error)
	UpdateConsultationAssignment(id, consulta_id, area_asignada string) (*ConsultationAssignment, int, error)
	UpdateConsultationAssignmentByIDConsultation(id, consulta_id, area_asignada string) (*ConsultationAssignment, int, error)
	DeleteConsultationAssignment(id string) (int, error)
	DeleteConsultationAssignmentByIDConsultation(id string) (int, error)
	GetConsultationAssignmentByID(id string) (*ConsultationAssignment, int, error)
	GetConsultationAssignmentByIDConsultation(id string) (*ConsultationAssignment, int, error)
	GetAllConsultationAssignment() ([]*ConsultationAssignment, error)
	GetConsultationAssignmentByArea(area string) (*ConsultationAssignment, int, error)
}

type service struct {
	repository ServicesConsultationAssignmentRepository
	user       *models.User
	txID       string
}

func NewConsultationAssignmentService(repository ServicesConsultationAssignmentRepository, user *models.User, TxID string) PortsServerConsultationAssignment {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateConsultationAssignment(id, consulta_id, area_asignada string) (*ConsultationAssignment, int, error) {

	m := NewConsultationAssignment(id, consulta_id, area_asignada)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create patient :", err)
		return m, 3, err
	}

	return m, 29, nil
}

func (s *service) UpdateConsultationAssignment(id, consulta_id, area_asignada string) (*ConsultationAssignment, int, error) {
	m := NewConsultationAssignment(id, consulta_id, area_asignada)
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
		logger.Error.Println(s.txID, " - couldn't update nursing consultation :", err)
		return nil, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteConsultationAssignment(id string) (int, error) {
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

func (s *service) GetConsultationAssignmentByID(id string) (*ConsultationAssignment, int, error) {
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

func (s *service) UpdateConsultationAssignmentByIDConsultation(id, consulta_id, area_asignada string) (*ConsultationAssignment, int, error) {
	m := NewConsultationAssignment(id, consulta_id, area_asignada)
	valid, err := m.valid()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't meet validations:", err)
		return nil, 15, err
	}
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return nil, 15, err
	}
	if err := s.repository.updateByIDConsultation(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update nursing consultation :", err)
		return nil, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteConsultationAssignmentByIDConsultation(id string) (int, error) {
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

func (s *service) GetConsultationAssignmentByIDConsultation(id string) (*ConsultationAssignment, int, error) {
	if err := uuid.Validate(id); err != nil {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return nil, 15, fmt.Errorf("id is required")
	}
	m, err := s.repository.getByIDConsultation(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllConsultationAssignment() ([]*ConsultationAssignment, error) {
	return s.repository.getAll()
}

func (s *service) GetConsultationAssignmentByArea(area string) (*ConsultationAssignment, int, error) {

	m, err := s.repository.getByIDConsultation(area)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}
