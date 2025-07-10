package permissions

import (
	"database/sql"
	"dbu-api/internal/models"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// sqlServer estructura de conexión a la BD de mssql
type mysqlserver struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newPermissionMysqlServerRepository(db *sqlx.DB, user *models.User, txID string) *mysqlserver {
	return &mysqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

func (s *mysqlserver) getAllByIDs(ids string) ([]*Permission, error) {
	var ms []*Permission
	var sqlGetAll = `SELECT id , description, method, path, created_at, updated_at FROM permissions WHERE id in (%s)`
	sqlGetAll = fmt.Sprintf(sqlGetAll, ids)

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
