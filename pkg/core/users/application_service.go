package users

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
)

type PortsServerUsers interface {
	GetUsersByID(id int64) (*User, int, error)
	GetUsersByUsername(username string) (*User, int, error)
}

type service struct {
	repository ServicesUsersRepository
	user       *models.User
	txID       string
}

func NewUsersService(repository ServicesUsersRepository, user *models.User, TxID string) PortsServerUsers {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) GetUsersByID(id int64) (*User, int, error) {
	m, err := s.repository.getByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetUsersByUsername(username string) (*User, int, error) {
	m, err := s.repository.getByUsername(username)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}
