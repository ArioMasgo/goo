package users

import (
	"dbu-api/internal/models"
	"github.com/jmoiron/sqlx"
)

type ServicesUsersRepository interface {
	getByID(id int64) (*User, error)
	getByUsername(username string) (*User, error)
}

func FactoryStorage(db *sqlx.DB, user *models.User, txID string) ServicesUsersRepository {
	return newUserMysqlServerRepository(db, user, txID)
}
