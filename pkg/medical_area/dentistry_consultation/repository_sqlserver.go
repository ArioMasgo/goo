package dentistry_consultation

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

func newDentistryConsultationSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s sqlserver) create(m *DentistryConsultation) error {
	const sqlInsertConsulta = `INSERT INTO consulta_odontologia (id, paciente_id, fecha_consulta, user_creator, created_at, updated_at) VALUES (:id, :paciente_id, :fecha_consulta, :user_creator, :created_at, :updated_at)`

	rs, err := s.DB.NamedExec(sqlInsertConsulta, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) update(m *DentistryConsultation) error {
	const sqlUpdate = `UPDATE consulta_odontologia SET paciente_id = :paciente_id, fecha_consulta = :fecha_consulta, user_creator = :user_creator, created_at = :created_at, updated_at = :updated_at WHERE id = :id`
	_, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	return nil
}

func (s sqlserver) delete(id string) error {
	const psqlDelete = `DELETE FROM consulta_odontologia WHERE id = :id`
	m := DentistryConsultation{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) getByID(id string) (*DentistryConsultation, error) {
	const sqlGetByID = `SELECT id, paciente_id, fecha_consulta, user_creator, created_at, updated_at FROM consulta_odontologia WHERE id = ?`
	mdl := DentistryConsultation{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s sqlserver) getAll() ([]*DentistryConsultation, error) {
	var ms []*DentistryConsultation
	const sqlGetAll = `SELECT id, paciente_id, fecha_consulta, user_creator, created_at, updated_at FROM consulta_odontologia`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s sqlserver) getByIDPatient(id string) ([]*DentistryConsultation, error) {
	var ms []*DentistryConsultation
	const sqlGetAll = `SELECT id, paciente_id, fecha_consulta, user_creator, created_at, updated_at FROM consulta_odontologia WHERE paciente_id = ?`

	err := s.DB.Select(&ms, sqlGetAll, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
