package residence_robot

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerResidenceRobot interface {
	CreateResidenceRobot(residenceID string, promptTokens, completionTokens, totalTokens int) (*ResidenceRobot, int, error)
}

type service struct {
	repository ServiceResidenceRobotRepository // Repositorio para ResidenceRobot
	user       *models.User                    // Usuario asociado
	txID       string                          // ID de transacción
}

func NewResidenceRobotService(repository ServiceResidenceRobotRepository, user *models.User, txID string) PortsServerResidenceRobot {
	return &service{repository: repository, user: user, txID: txID}
}

func (s *service) CreateResidenceRobot(residenceID string, promptTokens, completionTokens, totalTokens int) (*ResidenceRobot, int, error) {
	m := NewResidenceRobot(uuid.New().String(), residenceID, promptTokens, completionTokens, totalTokens)

	valid, err := m.valid()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't meet validations:", err)
		return nil, 15, err
	}
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:")
		return nil, 15, fmt.Errorf("validation failed")
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create ResidenceRobot:", err)
		return nil, 3, err
	}
	return m, 29, nil
}
