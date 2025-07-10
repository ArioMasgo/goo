package level_user_permissions

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
)

type PortsServerLevelUserPermissions interface {
	CreateLevelUserPermission(id int64, permissionId int64, levelUserId int64) (*LevelUserPermission, int, error)
	GetAllByLevelUser(levelUserID int64) ([]*LevelUserPermission, error)
}

type service struct {
	repository ServicesLevelUserPermissionsRepository
	user       *models.User
	txID       string
}

func NewLevelUserPermissionsService(repository ServicesLevelUserPermissionsRepository, user *models.User, TxID string) PortsServerLevelUserPermissions {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateLevelUserPermission(id int64, permissionId int64, levelUserId int64) (*LevelUserPermission, int, error) {
	m := NewLevelUserPermission(id, permissionId, levelUserId)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Residencias :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) GetAllByLevelUser(levelUserID int64) ([]*LevelUserPermission, error) {
	return s.repository.getAllByLevelUser(levelUserID)
}
