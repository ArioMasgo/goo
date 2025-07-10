package residence_configuration

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerResidenceConfiguration interface {
	CreateResidenceConfiguration(id string, percentageFcea float64, percentageEngineering float64, minimumGradeFcea float64, minimumGradeEngineering float64, residenceID string, isNewbie bool) (*ResidenceConfiguration, int, error)
	UpdateResidenceConfiguration(id string, percentageFcea float64, percentageEngineering float64, minimumGradeFcea float64, minimumGradeEngineering float64, residenceID string, isNewbie bool) (*ResidenceConfiguration, int, error)
	DeleteResidenceConfiguration(id string) (int, error)
	GetResidenceConfigurationByID(id string) (*ResidenceConfiguration, int, error)
	GetAllResidenceConfiguration() ([]*ResidenceConfiguration, error)
	GetResidenceConfigurationByResidenceID(residenceID string) (*ResidenceConfiguration, int, error)
	UpdateResidenceConfigurationByResidenceID(percentageFcea float64, percentageEngineering float64, minimumGradeFcea float64, minimumGradeEngineering float64, residenceID string, isNewbie bool) (*ResidenceConfiguration, int, error)
}

type service struct {
	repository ServiceResidenceConfigurationRepository
	user       *models.User
	txID       string
}

func NewResidenceConfigurationService(repository ServiceResidenceConfigurationRepository, user *models.User, TxID string) PortsServerResidenceConfiguration {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateResidenceConfiguration(id string, percentageFcea float64, percentageEngineering float64, minimumGradeFcea float64, minimumGradeEngineering float64, residenceID string, isNewbie bool) (*ResidenceConfiguration, int, error) {
	m := NewResidenceConfiguration(id, percentageFcea, percentageEngineering, minimumGradeFcea, minimumGradeEngineering, residenceID, isNewbie, s.user.ID)
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
		logger.Error.Println(s.txID, " - couldn't create ResidenceConfiguration :", err)
		return nil, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateResidenceConfiguration(id string, percentageFcea float64, percentageEngineering float64, minimumGradeFcea float64, minimumGradeEngineering float64, residenceID string, isNewbie bool) (*ResidenceConfiguration, int, error) {
	m := NewResidenceConfiguration(id, percentageFcea, percentageEngineering, minimumGradeFcea, minimumGradeEngineering, residenceID, isNewbie, s.user.ID)
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
		logger.Error.Println(s.txID, " - couldn't update ResidenceConfiguration :", err)
		return nil, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteResidenceConfiguration(id string) (int, error) {
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

func (s *service) GetResidenceConfigurationByID(id string) (*ResidenceConfiguration, int, error) {
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

func (s *service) GetAllResidenceConfiguration() ([]*ResidenceConfiguration, error) {
	return s.repository.getAll()
}

func (s *service) GetResidenceConfigurationByResidenceID(residenceID string) (*ResidenceConfiguration, int, error) {
	if err := uuid.Validate(residenceID); err != nil {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return nil, 15, fmt.Errorf("id is required")
	}
	m, err := s.repository.getByResidenceID(residenceID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) UpdateResidenceConfigurationByResidenceID(percentageFcea float64, percentageEngineering float64, minimumGradeFcea float64, minimumGradeEngineering float64, residenceID string, isNewbie bool) (*ResidenceConfiguration, int, error) {
	m := NewResidenceConfiguration(uuid.New().String(), percentageFcea, percentageEngineering, minimumGradeFcea, minimumGradeEngineering, residenceID, isNewbie, s.user.ID)
	valid, err := m.valid()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't meet validations:", err)
		return nil, 15, err
	}
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:")
		return nil, 15, err
	}
	if err := s.repository.updateByResidenceID(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update ResidenceConfiguration :", err)
		return nil, 18, err
	}
	return m, 29, nil
}
