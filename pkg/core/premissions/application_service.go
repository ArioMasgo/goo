package permissions

import (
	"dbu-api/internal/models"
)

type PortsServerPermission interface {
	GetAllByIDs(ids string) ([]*Permission, error)
}

type service struct {
	repository ServicesPermissionRepository
	user       *models.User
	txID       string
}

func NewPermissionService(repository ServicesPermissionRepository, user *models.User, TxID string) PortsServerPermission {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) GetAllByIDs(ids string) ([]*Permission, error) {
	return s.repository.getAllByIDs(ids)
}
