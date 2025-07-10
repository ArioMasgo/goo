package patient_background

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

func newPatientBackgroundSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s *sqlserver) create(m *PatientBackground) error {
	const sqlInsert = `INSERT INTO paciente_antecedentes (id, paciente_id, nombre_antecedente, estado_antecedente, is_deleted, user_creator, created_at, updated_at) VALUES (:id, :paciente_id, :nombre_antecedente, :estado_antecedente, :is_deleted, :user_creator, now(), now())`
	rs, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s *sqlserver) update(m *PatientBackground) error {
	const sqlUpdate = `UPDATE paciente_antecedentes SET nombre_antecedente = :nombre_antecedente, estado_antecedente = :estado_antecedente, updated_at = now() WHERE id = :id`
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s *sqlserver) delete(id string) error {
	const psqlDelete = `DELETE FROM paciente_antecedentes WHERE id = :id`
	m := PatientBackground{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s *sqlserver) deleteByIDPatient(id string) error {
	const psqlDelete = `DELETE FROM paciente_antecedentes WHERE paciente_id = :id`
	m := PatientBackground{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s *sqlserver) getByID(id string) (*PatientBackground, error) {
	const sqlGetByID = `SELECT id , paciente_id, nombre_antecedente, estado_antecedente, is_deleted, user_deleted, deleted_at, user_creator, created_at, updated_at FROM paciente_antecedentes WHERE id = ?`
	mdl := PatientBackground{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s *sqlserver) getAll() ([]*PatientBackground, error) {
	var ms []*PatientBackground
	const sqlGetAll = `SELECT id , paciente_id, nombre_antecedente, estado_antecedente, is_deleted, user_deleted, deleted_at, user_creator, created_at, updated_at FROM paciente_antecedentes`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) getByIDPatient(id string) ([]*PatientBackground, error) {
	var ms []*PatientBackground
	const sqlGetByIDPatient = `SELECT id , paciente_id, nombre_antecedente, estado_antecedente, is_deleted, user_deleted, deleted_at, user_creator, created_at, updated_at FROM paciente_antecedentes WHERE paciente_id = ?`

	err := s.DB.Select(&ms, sqlGetByIDPatient, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
