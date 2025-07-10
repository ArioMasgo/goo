package dentistry_consultation

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerDentistryConsultation interface {
	CreateDentistryConsultation(id, paciente_id, fecha_consulta string) (*DentistryConsultation, int, error)
	UpdateDentistryConsultation(id, paciente_id, fecha_consulta string) (*DentistryConsultation, int, error)
	DeleteDentistryConsultation(id string) (int, error)
	GetDentistryConsultationByID(id string) (*DentistryConsultation, int, error)
	GetDentistryConsultationByIDPatient(id string) ([]*DentistryConsultation, int, error)
	GetAllDentistryConsultation() ([]*DentistryConsultation, error)
}

type service struct {
	repository ServicesDentistryConsultationRepository
	user       *models.User
	txID       string
}

func NewDentistryConsultationService(repository ServicesDentistryConsultationRepository, user *models.User, TxID string) PortsServerDentistryConsultation {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateDentistryConsultation(id, paciente_id, fecha_consulta string) (*DentistryConsultation, int, error) {

	m := NewDentistryConsultation(id, paciente_id, fecha_consulta)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create dentistry consultation :", err)
		return m, 3, err
	}

	return m, 29, nil
}

func (s *service) UpdateDentistryConsultation(id, paciente_id, fecha_consulta string) (*DentistryConsultation, int, error) {
	m := NewDentistryConsultation(id, paciente_id, fecha_consulta)
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
		logger.Error.Println(s.txID, " - couldn't update dentistry consultation :", err)
		return nil, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteDentistryConsultation(id string) (int, error) {
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

func (s *service) GetDentistryConsultationByID(id string) (*DentistryConsultation, int, error) {
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

func (s *service) GetDentistryConsultationByIDPatient(id string) ([]*DentistryConsultation, int, error) {
	if err := uuid.Validate(id); err != nil {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return nil, 15, fmt.Errorf("id is required")
	}
	m, err := s.repository.getByIDPatient(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllDentistryConsultation() ([]*DentistryConsultation, error) {
	return s.repository.getAll()
}
