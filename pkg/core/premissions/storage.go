package permissions

import (
	"dbu-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type ServicesPermissionRepository interface {
	getAllByIDs(ids string) ([]*Permission, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesPermissionRepository {
	return newPermissionMysqlServerRepository(db, user, txID)
}
