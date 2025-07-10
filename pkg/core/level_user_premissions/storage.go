package level_user_permissions

import (
	"dbu-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type ServicesLevelUserPermissionsRepository interface {
	create(permission *LevelUserPermission) error
	getAllByLevelUser(levelUserID int64) ([]*LevelUserPermission, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesLevelUserPermissionsRepository {
	return newLevelUserPermissionsMysqlServerRepository(db, user, txID)
}
