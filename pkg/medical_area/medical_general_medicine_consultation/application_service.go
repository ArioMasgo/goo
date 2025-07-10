package medical_general_medicine_consultation

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerGeneralMedicineConsultation interface {
	CreateGeneralMedicineConsultation(id, consulta_id, fecha_hora, anamnesis, examen_clinico, indicaciones string) (*GeneralMedicineConsultation, int, error)
	UpdateGeneralMedicineConsultation(id, consulta_id, fecha_hora, anamnesis, examen_clinico, indicaciones string) (*GeneralMedicineConsultation, int, error)
	DeleteGeneralMedicineConsultation(id string) (int, error)
	GetGeneralMedicineConsultationByID(id string) (*GeneralMedicineConsultation, int, error)
	GetGeneralMedicineConsultationByIDPatient(id string) ([]*GeneralMedicineConsultation, int, error)
	GetAllGeneralMedicineConsultation() ([]*GeneralMedicineConsultation, error)
	GetGeneralMedicineConsultationExcel(fecha_inicio, fecha_fin string) ([]*models.ConsultationIntegralAttentionExcel, error)
}

type service struct {
	repository ServicesGeneralMedicineConsultationRepository
	user       *models.User
	txID       string
}

func NewGeneralMedicineConsultationService(repository ServicesGeneralMedicineConsultationRepository, user *models.User, TxID string) PortsServerGeneralMedicineConsultation {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateGeneralMedicineConsultation(id, consulta_id, fecha_hora, anamnesis, examen_clinico, indicaciones string) (*GeneralMedicineConsultation, int, error) {

	m := NewGeneralMedicineConsultation(id, consulta_id, fecha_hora, anamnesis, examen_clinico, indicaciones)
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

func (s *service) UpdateGeneralMedicineConsultation(id, consulta_id, fecha_hora, anamnesis, examen_clinico, indicaciones string) (*GeneralMedicineConsultation, int, error) {
	m := NewGeneralMedicineConsultation(id, consulta_id, fecha_hora, anamnesis, examen_clinico, indicaciones)
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

func (s *service) DeleteGeneralMedicineConsultation(id string) (int, error) {
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

func (s *service) GetGeneralMedicineConsultationByID(id string) (*GeneralMedicineConsultation, int, error) {
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

func (s *service) GetGeneralMedicineConsultationByIDPatient(id string) ([]*GeneralMedicineConsultation, int, error) {
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

func (s *service) GetAllGeneralMedicineConsultation() ([]*GeneralMedicineConsultation, error) {
	return s.repository.getAll()
}

func (s *service) GetGeneralMedicineConsultationExcel(fecha_inicio, fecha_fin string) ([]*models.ConsultationIntegralAttentionExcel, error) {
	m, err := s.repository.getAllByDateExcel(fecha_inicio, fecha_fin)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get consultation integran attention excel:", err)
		return nil, err
	}
	return m, nil
}
