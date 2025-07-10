package core

import (
	"dbu-api/internal/models"
	"dbu-api/pkg/core/level_user_premissions"
	"dbu-api/pkg/core/premissions"
	"dbu-api/pkg/core/users"
	"github.com/jmoiron/sqlx"
)

type ServerCore struct {
	SrvUsers                users.PortsServerUsers
	SrvLevelUserPermissions level_user_permissions.PortsServerLevelUserPermissions
	SrvPermissions          permissions.PortsServerPermission
}

func NewServerCore(db *sqlx.DB, usr *models.User, txID string) *ServerCore {
	return &ServerCore{
		SrvUsers:                users.NewUsersService(users.FactoryStorage(db, usr, txID), usr, txID),
		SrvLevelUserPermissions: level_user_permissions.NewLevelUserPermissionsService(level_user_permissions.FactoryStorage(db, usr, txID), usr, txID),
		SrvPermissions:          permissions.NewPermissionService(permissions.FactoryStorage(db, usr, txID), usr, txID),
	}
}
