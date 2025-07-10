package level_user_permissions

import (
	"database/sql"
	"dbu-api/internal/models"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

// sqlServer estructura de conexión a la BD de mssql
type mysqlserver struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newLevelUserPermissionsMysqlServerRepository(db *sqlx.DB, user *models.User, txID string) *mysqlserver {
	return &mysqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// GetByID consulta un registro por su ID
func (s *mysqlserver) create(permission *LevelUserPermission) error {
	date := time.Now()
	permission.UpdatedAt = date
	permission.CreatedAt = date

	const sqlInsert = `INSERT INTO level_user_permissions (permission_id, level_user_id, created_at, updated_at) 
              VALUES (:permission_id, :level_user_id, :created_at, :updated_at)`
	rs, err := s.DB.NamedExec(sqlInsert, &permission)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s *mysqlserver) getAllByLevelUser(levelUserID int64) ([]*LevelUserPermission, error) {
	var ms []*LevelUserPermission
	const sqlGetAll = `SELECT id , permission_id, level_user_id, created_at, updated_at FROM level_user_permissions WHERE level_user_id = ?`

	err := s.DB.Select(&ms, sqlGetAll, levelUserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
