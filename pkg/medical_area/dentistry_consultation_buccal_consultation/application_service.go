package dentistry_consultation_buccal_consultation

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerBuccalConsultation interface {
	CreateBuccalConsultation(id, consulta_odontologia_id, relato, diagnostico, examen_auxiliar string, examen_clinico *string, tratamiento, indicaciones, comentarios string) (*BuccalConsultation, int, error)
	UpdateBuccalConsultation(id, consulta_odontologia_id, relato, diagnostico, examen_auxiliar string, examen_clinico *string, tratamiento, indicaciones, comentarios string) (*BuccalConsultation, int, error)
	DeleteBuccalConsultation(id string) (int, error)
	DeleteBuccalConsultationByIDConsultation(id string) (int, error)
	GetBuccalConsultationByID(id string) (*BuccalConsultation, int, error)
	GetBuccalConsultationByIDConsultation(id string) (*BuccalConsultation, int, error)
	GetAllBuccalConsultation() ([]*BuccalConsultation, error)
}

type service struct {
	repository ServicesBuccalConsultationRepository
	user       *models.User
	txID       string
}

func NewBuccalConsultationService(repository ServicesBuccalConsultationRepository, user *models.User, TxID string) PortsServerBuccalConsultation {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateBuccalConsultation(id, consulta_odontologia_id, relato, diagnostico, examen_auxiliar string, examen_clinico *string, tratamiento, indicaciones, comentarios string) (*BuccalConsultation, int, error) {

	m := NewBuccalConsultation(id, consulta_odontologia_id, relato, diagnostico, examen_auxiliar, examen_clinico, tratamiento, indicaciones, comentarios)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create buccal consultation :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateBuccalConsultation(id, consulta_odontologia_id, relato, diagnostico, examen_auxiliar string, examen_clinico *string, tratamiento, indicaciones, comentarios string) (*BuccalConsultation, int, error) {
	m := NewBuccalConsultation(id, consulta_odontologia_id, relato, diagnostico, examen_auxiliar, examen_clinico, tratamiento, indicaciones, comentarios)
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
		logger.Error.Println(s.txID, " - couldn't update buccal consultation :", err)
		return nil, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteBuccalConsultation(id string) (int, error) {
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

func (s *service) DeleteBuccalConsultationByIDConsultation(id string) (int, error) {
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

func (s *service) GetBuccalConsultationByID(id string) (*BuccalConsultation, int, error) {
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

func (s *service) GetBuccalConsultationByIDConsultation(id string) (*BuccalConsultation, int, error) {
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

func (s *service) GetAllBuccalConsultation() ([]*BuccalConsultation, error) {
	return s.repository.getAll()
}
