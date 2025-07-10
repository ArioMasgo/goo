package nursing_consultation_preferential_test

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerPreferentialTest interface {
	CreatePreferentialTest(id, consulta_enfermeria_id, aparato_respiratorio, aparato_cardiovascular, aparato_digestivo, aparato_genitourinario, papanicolau, examen_mama, comentarios string) (*PreferentialTest, int, error)
	UpdatePreferentialTest(id, consulta_enfermeria_id, aparato_respiratorio, aparato_cardiovascular, aparato_digestivo, aparato_genitourinario, papanicolau, examen_mama, comentarios string) (*PreferentialTest, int, error)
	DeletePreferentialTest(id string) (int, error)
	DeletePreferentialTestByIDConsultation(id string) (int, error)
	GetPreferentialTestByID(id string) (*PreferentialTest, int, error)
	GetAllPreferentialTest() ([]*PreferentialTest, error)
}

type service struct {
	repository ServicesPreferentialTestRepository
	user       *models.User
	txID       string
}

func NewPreferentialTestService(repository ServicesPreferentialTestRepository, user *models.User, TxID string) PortsServerPreferentialTest {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreatePreferentialTest(id, consulta_enfermeria_id, aparato_respiratorio, aparato_cardiovascular, aparato_digestivo, aparato_genitourinario, papanicolau, examen_mama, comentarios string) (*PreferentialTest, int, error) {

	m := NewPreferentialTest(id, consulta_enfermeria_id, aparato_respiratorio, aparato_cardiovascular, aparato_digestivo, aparato_genitourinario, papanicolau, examen_mama, comentarios)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create preferential test :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdatePreferentialTest(id, consulta_enfermeria_id, aparato_respiratorio, aparato_cardiovascular, aparato_digestivo, aparato_genitourinario, papanicolau, examen_mama, comentarios string) (*PreferentialTest, int, error) {
	m := NewPreferentialTest(id, consulta_enfermeria_id, aparato_respiratorio, aparato_cardiovascular, aparato_digestivo, aparato_genitourinario, papanicolau, examen_mama, comentarios)
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
		logger.Error.Println(s.txID, " - couldn't update preferential test :", err)
		return nil, 18, err
	}
	return m, 29, nil
}

func (s *service) DeletePreferentialTest(id string) (int, error) {
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

func (s *service) DeletePreferentialTestByIDConsultation(id string) (int, error) {
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

func (s *service) GetPreferentialTestByID(id string) (*PreferentialTest, int, error) {
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

func (s *service) GetAllPreferentialTest() ([]*PreferentialTest, error) {
	return s.repository.getAll()
}
