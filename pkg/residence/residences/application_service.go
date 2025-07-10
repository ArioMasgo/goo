package residences

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type PortsServerResidence interface {
	CreateResidence(id, name, description, gender, address, status string) (*Residence, int, error)
	UpdateResidence(id, name, description, gender, address, status string) (*Residence, int, error)
	DeleteResidence(id string) (int, error)
	GetResidenceByID(id string) (*Residence, int, error)
	GetAllResidence() ([]*Residence, error)
}

type service struct {
	repository ServiceResidenceRepository
	user       *models.User
	txID       string
}

func NewResidenceService(repository ServiceResidenceRepository, user *models.User, TxID string) PortsServerResidence {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateResidence(id, name, description, gender, address, status string) (*Residence, int, error) {
	m := NewResidence(id, strings.ToUpper(name), gender, description, address, status, s.user.ID)
	valid, err := m.valid()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't meet validations:", err)
		return nil, 15, err
	}
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:")
		return nil, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Residence :", err)
		return nil, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateResidence(id, name, description, gender, address, status string) (*Residence, int, error) {
	m := NewResidence(id, strings.ToUpper(name), gender, description, address, status, s.user.ID)
	valid, err := m.valid()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't meet validations:", err)
		return nil, 15, err
	}
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:")
		return nil, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Residence :", err)
		return nil, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteResidence(id string) (int, error) {
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

func (s *service) GetResidenceByID(id string) (*Residence, int, error) {
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

func (s *service) GetAllResidence() ([]*Residence, error) {
	return s.repository.getAll()
}
