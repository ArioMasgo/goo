package users

import (
	"database/sql"
	"dbu-api/internal/models"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// sqlServer estructura de conexión a la BD de mssql
type mysqlserver struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newUserMysqlServerRepository(db *sqlx.DB, user *models.User, txID string) *mysqlserver {
	return &mysqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// GetByID consulta un registro por su ID
func (s *mysqlserver) getByID(id int64) (*User, error) {
	const sqlGetByID = `SELECT 
			id, username, full_name, email, email_verified_at, password, ip_address, id_level_user,
			last_user, remember_token, status_id, created_at, updated_at
		FROM users 
		WHERE id = ? LIMIT 1`
	var user User
	err := s.DB.Get(&user, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("system: user not found")
		}
		return nil, err
	}
	return &user, nil
}

// GetByID consulta un registro por su ID
func (s *mysqlserver) getByUsername(username string) (*User, error) {
	const sqlGetByID = `SELECT 
			id, username, full_name, email, email_verified_at, password, ip_address, id_level_user,
			last_user, remember_token, status_id, created_at, updated_at
		FROM users 
		WHERE username = ? LIMIT 1`
	var user User
	err := s.DB.Get(&user, sqlGetByID, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("system: user not found")
		}
		return nil, err
	}
	return &user, nil
}
