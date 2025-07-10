package medical_consultation

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerMedicalConsultation interface {
	CreateMedicalConsultation(id, paciente_id, fecha_consulta, area_medica string) (*MedicalConsultation, int, error)
	UpdateMedicalConsultation(id, paciente_id, fecha_consulta, area_medica string) (*MedicalConsultation, int, error)
	DeleteMedicalConsultation(id string) (int, error)
	GetMedicalConsultationByID(id string) (*MedicalConsultation, int, error)
	GetMedicalConsultationByIDPatient(id string) ([]*MedicalConsultation, int, error)
	GetAllMedicalConsultation() ([]*MedicalConsultation, error)
}

type service struct {
	repository ServicesMedicalConsultationRepository
	user       *models.User
	txID       string
}

func NewMedicalConsultationService(repository ServicesMedicalConsultationRepository, user *models.User, TxID string) PortsServerMedicalConsultation {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateMedicalConsultation(id, paciente_id, fecha_consulta, area_medica string) (*MedicalConsultation, int, error) {

	m := NewMedicalConsultation(id, paciente_id, fecha_consulta, area_medica)
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

func (s *service) UpdateMedicalConsultation(id, paciente_id, fecha_consulta, area_medica string) (*MedicalConsultation, int, error) {
	m := NewMedicalConsultation(id, paciente_id, fecha_consulta, area_medica)
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
		logger.Error.Println(s.txID, " - couldn't update patient :", err)
		return nil, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteMedicalConsultation(id string) (int, error) {
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

func (s *service) GetMedicalConsultationByID(id string) (*MedicalConsultation, int, error) {
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

func (s *service) GetMedicalConsultationByIDPatient(id string) ([]*MedicalConsultation, int, error) {
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

func (s *service) GetAllMedicalConsultation() ([]*MedicalConsultation, error) {
	return s.repository.getAll()
}
