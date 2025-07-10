package patient_background

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerPatientBackground interface {
	CreatePatientBackground(id, paciente_id, nombre_antecedente, estado_antecedente string) (int, error)
	CreateAntecedents(patientID string, antecedents []*models.RequestAntecedents) (int, error)
	UpdatePatientBackground(id, paciente_id, nombre_antecedente, estado_antecedente string) (int, error)
	DeletePatientBackground(id string) (int, error)
	DeletePatientBackgroundByIDPatient(id string) (int, error)
	GetPatientBackgroundByID(id string) (*PatientBackground, int, error)
	GetPatientBackgroundByIDPatient(id string) ([]*PatientBackground, int, error)
	GetAllPatientBackground() ([]*PatientBackground, error)
}

type service struct {
	repository ServicesPatientBackgroundRepository
	user       *models.User
	txID       string
}

func NewPatientBackgroundService(repository ServicesPatientBackgroundRepository, user *models.User, TxID string) PortsServerPatientBackground {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreatePatientBackground(id, paciente_id, nombre_antecedente, estado_antecedente string) (int, error) {

	m := NewPatientBackground(id, paciente_id, nombre_antecedente, estado_antecedente)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create patient :", err)
		return 3, err
	}
	return 29, nil
}

func (s *service) CreateAntecedents(patientID string, antecedents []*models.RequestAntecedents) (int, error) {
	for _, antecedent := range antecedents {
		m := NewPatientBackground(antecedent.ID, patientID, antecedent.Nombre, antecedent.Estado)
		if valid, err := m.valid(); !valid {
			logger.Error.Println(s.txID, " - don't meet validations:", err)
			return 15, err
		}
		if err := s.repository.create(m); err != nil {
			if err.Error() == "rows affected error" {
				return 108, nil
			}
			logger.Error.Println(s.txID, " - couldn't create antecedent :", err)
			return 3, err
		}
	}
	return 29, nil
}

func (s *service) UpdatePatientBackground(id, paciente_id, nombre_antecedente, estado_antecedente string) (int, error) {
	m := NewPatientBackground(id, paciente_id, nombre_antecedente, estado_antecedente)
	valid, err := m.valid()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't meet validations:", err)
		return 15, err
	}
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update patient :", err)
		return 18, err
	}
	return 29, nil
}

func (s *service) DeletePatientBackground(id string) (int, error) {
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

func (s *service) DeletePatientBackgroundByIDPatient(id string) (int, error) {
	if err := uuid.Validate(id); err != nil {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return 15, fmt.Errorf("id is required")
	}

	if err := s.repository.deleteByIDPatient(id); err != nil {
		if err.Error() == "rows affected error" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update row:", err)
		return 20, err
	}
	return 28, nil
}

func (s *service) GetPatientBackgroundByID(id string) (*PatientBackground, int, error) {
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

func (s *service) GetPatientBackgroundByIDPatient(id string) ([]*PatientBackground, int, error) {
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

func (s *service) GetAllPatientBackground() ([]*PatientBackground, error) {
	return s.repository.getAll()
}
