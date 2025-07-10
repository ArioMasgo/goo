package nursing_consultation_visual_test

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerVisualTest interface {
	CreateVisualTest(id, consulta_enfermeria_id, ojo_derecho, ojo_izquierdo, presion_ocular, comentarios string) (*VisualTest, int, error)
	UpdateVisualTest(id, consulta_enfermeria_id, ojo_derecho, ojo_izquierdo, presion_ocular, comentarios string) (*VisualTest, int, error)
	DeleteVisualTest(id string) (int, error)
	DeleteVisualTestByIDConsultation(id string) (int, error)
	GetVisualTestByID(id string) (*VisualTest, int, error)
	GetAllVisualTest() ([]*VisualTest, error)
}

type service struct {
	repository ServicesVisualTestRepository
	user       *models.User
	txID       string
}

func NewVisualTestService(repository ServicesVisualTestRepository, user *models.User, TxID string) PortsServerVisualTest {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateVisualTest(id, consulta_enfermeria_id, ojo_derecho, ojo_izquierdo, presion_ocular, comentarios string) (*VisualTest, int, error) {

	m := NewVisualTest(id, consulta_enfermeria_id, ojo_derecho, ojo_izquierdo, presion_ocular, comentarios)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create visual test :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateVisualTest(id, consulta_enfermeria_id, ojo_derecho, ojo_izquierdo, presion_ocular, comentarios string) (*VisualTest, int, error) {
	m := NewVisualTest(id, consulta_enfermeria_id, ojo_derecho, ojo_izquierdo, presion_ocular, comentarios)
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

func (s *service) DeleteVisualTest(id string) (int, error) {
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

func (s *service) DeleteVisualTestByIDConsultation(id string) (int, error) {
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

func (s *service) GetVisualTestByID(id string) (*VisualTest, int, error) {
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

func (s *service) GetAllVisualTest() ([]*VisualTest, error) {
	return s.repository.getAll()
}
