package nursing_consultation_sexuality_test

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerSexualityTest interface {
	CreateSexualityTest(id, consulta_enfermeria_id, actividad_sexual, planificacion_familiar, comentarios string) (*SexualityTest, int, error)
	UpdateSexualityTest(id, consulta_enfermeria_id, actividad_sexual, planificacion_familiar, comentarios string) (*SexualityTest, int, error)
	DeleteSexualityTest(id string) (int, error)
	DeleteSexualityTestByIDConsultation(id string) (int, error)
	GetSexualityTestByID(id string) (*SexualityTest, int, error)
	GetAllSexualityTest() ([]*SexualityTest, error)
}

type service struct {
	repository ServicesSexualityTestRepository
	user       *models.User
	txID       string
}

func NewSexualityTestService(repository ServicesSexualityTestRepository, user *models.User, TxID string) PortsServerSexualityTest {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateSexualityTest(id, consulta_enfermeria_id, actividad_sexual, planificacion_familiar, comentarios string) (*SexualityTest, int, error) {

	m := NewSexualityTest(id, consulta_enfermeria_id, actividad_sexual, planificacion_familiar, comentarios)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create sexuality test :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateSexualityTest(id, consulta_enfermeria_id, actividad_sexual, planificacion_familiar, comentarios string) (*SexualityTest, int, error) {
	m := NewSexualityTest(id, consulta_enfermeria_id, actividad_sexual, planificacion_familiar, comentarios)
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
		logger.Error.Println(s.txID, " - couldn't update sexuality test :", err)
		return nil, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteSexualityTest(id string) (int, error) {
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

func (s *service) DeleteSexualityTestByIDConsultation(id string) (int, error) {
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

func (s *service) GetSexualityTestByID(id string) (*SexualityTest, int, error) {
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

func (s *service) GetAllSexualityTest() ([]*SexualityTest, error) {
	return s.repository.getAll()
}
