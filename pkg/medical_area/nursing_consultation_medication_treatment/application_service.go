package nursing_consultation_medication_treatment

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerMedicationTreatment interface {
	CreateMedicationTreatment(id, consulta_enfermeria_id, nombre_generico_medicamento, via_administracion, hora_aplicacion, responsable_atencion, observaciones string, areaSolicitante, especialistaSolicitante *string) (*MedicationTreatment, int, error)
	UpdateMedicationTreatment(id, consulta_enfermeria_id, nombre_generico_medicamento, via_administracion, hora_aplicacion, responsable_atencion, observaciones string, areaSolicitante, especialistaSolicitante *string) (*MedicationTreatment, int, error)
	DeleteMedicationTreatment(id string) (int, error)
	DeleteMedicationTreatmentByIDConsultation(id string) (int, error)
	GetMedicationTreatmentByIDConsultation(id string) ([]*MedicationTreatment, int, error)
	GetMedicationTreatmentByID(id string) (*MedicationTreatment, int, error)
	GetAllMedicationTreatment() ([]*MedicationTreatment, error)
}

type service struct {
	repository ServicesMedicationTreatmentRepository
	user       *models.User
	txID       string
}

func NewMedicationTreatmentService(repository ServicesMedicationTreatmentRepository, user *models.User, TxID string) PortsServerMedicationTreatment {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateMedicationTreatment(id, consulta_enfermeria_id, nombre_generico_medicamento, via_administracion, hora_aplicacion, responsable_atencion, observaciones string, areaSolicitante, especialistaSolicitante *string) (*MedicationTreatment, int, error) {

	m := NewMedicationTreatment(id, consulta_enfermeria_id, nombre_generico_medicamento, via_administracion, hora_aplicacion, responsable_atencion, observaciones, areaSolicitante, especialistaSolicitante)
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

func (s *service) UpdateMedicationTreatment(id, consulta_enfermeria_id, nombre_generico_medicamento, via_administracion, hora_aplicacion, responsable_atencion, observaciones string, areaSolicitante, especialistaSolicitante *string) (*MedicationTreatment, int, error) {
	m := NewMedicationTreatment(id, consulta_enfermeria_id, nombre_generico_medicamento, via_administracion, hora_aplicacion, responsable_atencion, observaciones, areaSolicitante, especialistaSolicitante)
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

func (s *service) DeleteMedicationTreatment(id string) (int, error) {
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

func (s *service) DeleteMedicationTreatmentByIDConsultation(id string) (int, error) {
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

func (s *service) GetMedicationTreatmentByIDConsultation(id string) ([]*MedicationTreatment, int, error) {
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

func (s *service) GetMedicationTreatmentByID(id string) (*MedicationTreatment, int, error) {
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

func (s *service) GetAllMedicationTreatment() ([]*MedicationTreatment, error) {
	return s.repository.getAll()
}
