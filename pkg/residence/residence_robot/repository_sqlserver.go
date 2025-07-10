package residence_robot

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// sqlServer estructura de conexión a la BD de mssql
type sqlserver struct {
	DB   *sqlx.DB
	TxID string
}

func newResidenceRobotSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s *sqlserver) create(m *ResidenceRobot) error {
	const sqlInsert = `
        INSERT INTO residence_robot (id, residence_id, prompt_tokens, completion_tokens, total_tokens, created_at)
        VALUES (:id, :residence_id, :prompt_tokens, :completion_tokens, :total_tokens, :created_at)`
	rs, err := s.DB.NamedExec(sqlInsert, m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *sqlserver) update(m *ResidenceRobot) error {
	const sqlUpdate = `
        UPDATE residence_robot 
        SET prompt_tokens = :prompt_tokens, completion_tokens = :completion_tokens, total_tokens = :total_tokens
        WHERE id = :id`
	rs, err := s.DB.NamedExec(sqlUpdate, m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *sqlserver) getByID(id int64) (*ResidenceRobot, error) {
	const sqlGetByID = `
        SELECT id, residence_id, prompt_tokens, completion_tokens, total_tokens, created_at 
        FROM residence_robot 
        WHERE id = ?`
	mdl := ResidenceRobot{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *sqlserver) getAll() ([]*ResidenceRobot, error) {
	var ms []*ResidenceRobot
	const sqlGetAll = `
        SELECT id, residence_id, prompt_tokens, completion_tokens, total_tokens, created_at 
        FROM residence_robot`
	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return ms, nil
}

// GetByResidenceID consulta un registro por su ResidenceID
func (s *sqlserver) getByResidenceID(residenceID string) (*ResidenceRobot, error) {
	const sqlGetByResidenceID = `
        SELECT id, residence_id, prompt_tokens, completion_tokens, total_tokens, created_at 
        FROM residence_robot 
        WHERE residence_id = ?`
	mdl := ResidenceRobot{}
	err := s.DB.Get(&mdl, sqlGetByResidenceID, residenceID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &mdl, nil
}
