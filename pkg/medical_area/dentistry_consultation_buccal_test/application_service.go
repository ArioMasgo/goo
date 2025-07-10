package dentistry_consultation_buccal_test

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerBuccalTest interface {
	CreateBuccalTest(id, consulta_odontologia_id, odontograma_img, cpod, observacion, ihos, comentarios string) (*BuccalTest, int, error)
	UpdateBuccalTest(id, consulta_odontologia_id, odontograma_img, cpod, observacion, ihos, comentarios string) (*BuccalTest, int, error)
	DeleteBuccalTest(id string) (int, error)
	DeleteBuccalTestByIDConsultation(id string) (int, error)
	GetBuccalTestByID(id string) (*BuccalTest, int, error)
	GetBuccalTestByIDConsultation(id string) (*BuccalTest, int, error)
	GetAllBuccalTest() ([]*BuccalTest, error)
}

type service struct {
	repository ServicesBuccalTestRepository
	user       *models.User
	txID       string
}

func NewBuccalTestService(repository ServicesBuccalTestRepository, user *models.User, TxID string) PortsServerBuccalTest {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateBuccalTest(id, consulta_odontologia_id, odontograma_img, cpod, observacion, ihos, comentarios string) (*BuccalTest, int, error) {

	m := NewBuccalTest(id, consulta_odontologia_id, odontograma_img, cpod, observacion, ihos, comentarios)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create buccal test :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateBuccalTest(id, consulta_odontologia_id, odontograma_img, cpod, observacion, ihos, comentarios string) (*BuccalTest, int, error) {
	m := NewBuccalTest(id, consulta_odontologia_id, odontograma_img, cpod, observacion, ihos, comentarios)
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
		logger.Error.Println(s.txID, " - couldn't update buccal test :", err)
		return nil, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteBuccalTest(id string) (int, error) {
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

func (s *service) DeleteBuccalTestByIDConsultation(id string) (int, error) {
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

func (s *service) GetBuccalTestByID(id string) (*BuccalTest, int, error) {
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

func (s *service) GetBuccalTestByIDConsultation(id string) (*BuccalTest, int, error) {
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

func (s *service) GetAllBuccalTest() ([]*BuccalTest, error) {
	return s.repository.getAll()
}
