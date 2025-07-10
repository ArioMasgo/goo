package consultation_medical_area

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerConsultationMedicalArea interface {
	CreateConsultationMedicalArea(id, paciente_id, fecha_consulta, area_medica string) (*ConsultationMedicalArea, int, error)
	UpdateConsultationMedicalArea(id, paciente_id, fecha_consulta, area_medica string) (*ConsultationMedicalArea, int, error)
	DeleteConsultationMedicalArea(id string) (int, error)
	GetConsultationMedicalAreaByID(id string) (*ConsultationMedicalArea, int, error)
	GetConsultationMedicalAreasByPatientID(id string) ([]*ConsultationMedicalArea, int, error)
	GetConsultationMedicalAreaByPatientDNI(dni string) ([]*ConsultationMedicalArea, int, error)
	GetAllConsultationMedicalArea() ([]*ConsultationMedicalArea, error)
	GetAllConsultationMedicalAreaNursingByDateExcel(area_medica, fecha_inicio, fecha_fin string) ([]*models.ConsultationPatientsMedicalAreaExcel, error)
	GetAllConsultationMedicalAreaDentistryByDateExcel(area_medica, fecha_inicio, fecha_fin string) ([]*models.ConsultationPatientsMedicalAreaExcel, error)
	GetAllConsultationMedicalAreaMedicalByDateExcel(area_medica, fecha_inicio, fecha_fin string) ([]*models.ConsultationPatientsMedicalAreaExcel, error)
}

type service struct {
	repository ServicesConsultationMedicalAreaRepository
	user       *models.User
	txID       string
}

func NewConsultationMedicalAreaService(repository ServicesConsultationMedicalAreaRepository, user *models.User, TxID string) PortsServerConsultationMedicalArea {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateConsultationMedicalArea(id, paciente_id, fecha_consulta, area_medica string) (*ConsultationMedicalArea, int, error) {

	m := NewConsultationMedicalArea(id, paciente_id, fecha_consulta, area_medica)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create consultation :", err)
		return m, 3, err
	}

	return m, 29, nil
}

func (s *service) UpdateConsultationMedicalArea(id, paciente_id, fecha_consulta, area_medica string) (*ConsultationMedicalArea, int, error) {
	m := NewConsultationMedicalArea(id, paciente_id, fecha_consulta, area_medica)
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
		logger.Error.Println(s.txID, " - couldn't update consultation :", err)
		return nil, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteConsultationMedicalArea(id string) (int, error) {
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

func (s *service) GetConsultationMedicalAreaByID(id string) (*ConsultationMedicalArea, int, error) {
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

func (s *service) GetAllConsultationMedicalArea() ([]*ConsultationMedicalArea, error) {
	return s.repository.getAll()
}

func (s *service) GetAllConsultationMedicalAreaNursingByDateExcel(area_medica, fecha_inicio, fecha_fin string) ([]*models.ConsultationPatientsMedicalAreaExcel, error) {
	return s.repository.getAllByNursingDateExcel(area_medica, fecha_inicio, fecha_fin)
}

func (s *service) GetAllConsultationMedicalAreaDentistryByDateExcel(area_medica, fecha_inicio, fecha_fin string) ([]*models.ConsultationPatientsMedicalAreaExcel, error) {
	return s.repository.getAllByDentistryDateExcel(area_medica, fecha_inicio, fecha_fin)
}

func (s *service) GetAllConsultationMedicalAreaMedicalByDateExcel(area_medica, fecha_inicio, fecha_fin string) ([]*models.ConsultationPatientsMedicalAreaExcel, error) {
	return s.repository.getAllByMedicalDateExcel(area_medica, fecha_inicio, fecha_fin)
}

func (s *service) GetConsultationMedicalAreasByPatientID(id string) ([]*ConsultationMedicalArea, int, error) {
	if err := uuid.Validate(id); err != nil {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return nil, 15, fmt.Errorf("id is required")
	}
	m, err := s.repository.GetAllByPatientID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByIDPatient row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetConsultationMedicalAreaByPatientDNI(dni string) ([]*ConsultationMedicalArea, int, error) {
	m, err := s.repository.getByPatientDNI(dni)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}
