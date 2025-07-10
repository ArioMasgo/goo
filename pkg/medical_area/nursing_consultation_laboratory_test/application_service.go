package nursing_consultation_laboratory_test

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerLaboratoryTest interface {
	CreateLaboratoryTest(id, consulta_enfermeria_id, serologia, bk, hemograma, examen_orina, colesterol, glucosa, comentarios string) (*LaboratoryTest, int, error)
	UpdateLaboratoryTest(id, consulta_enfermeria_id, serologia, bk, hemograma, examen_orina, colesterol, glucosa, comentarios string) (*LaboratoryTest, int, error)
	DeleteLaboratoryTest(id string) (int, error)
	DeleteLaboratoryTestByIDConsultation(id string) (int, error)
	GetLaboratoryTestByID(id string) (*LaboratoryTest, int, error)
	GetAllLaboratoryTest() ([]*LaboratoryTest, error)
}

type service struct {
	repository ServicesLaboratoryTestRepository
	user       *models.User
	txID       string
}

func NewLaboratoryTestService(repository ServicesLaboratoryTestRepository, user *models.User, TxID string) PortsServerLaboratoryTest {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateLaboratoryTest(id, consulta_enfermeria_id, serologia, bk, hemograma, examen_orina, colesterol, glucosa, comentarios string) (*LaboratoryTest, int, error) {

	m := NewLaboratoryTest(id, consulta_enfermeria_id, serologia, bk, hemograma, examen_orina, colesterol, glucosa, comentarios)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create laboratory test :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateLaboratoryTest(id, consulta_enfermeria_id, serologia, bk, hemograma, examen_orina, colesterol, glucosa, comentarios string) (*LaboratoryTest, int, error) {
	m := NewLaboratoryTest(id, consulta_enfermeria_id, serologia, bk, hemograma, examen_orina, colesterol, glucosa, comentarios)
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

func (s *service) DeleteLaboratoryTest(id string) (int, error) {
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

func (s *service) DeleteLaboratoryTestByIDConsultation(id string) (int, error) {
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

func (s *service) GetLaboratoryTestByID(id string) (*LaboratoryTest, int, error) {
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

func (s *service) GetAllLaboratoryTest() ([]*LaboratoryTest, error) {
	return s.repository.getAll()
}
