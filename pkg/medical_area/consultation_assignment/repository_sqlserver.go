package consultation_assignment

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type sqlserver struct {
	DB   *sqlx.DB
	TxID string
}

func newConsultationAssignmentSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s sqlserver) create(m *ConsultationAssignment) error {
	const sqlInsertConsulta = `INSERT INTO area_medica_asignada (id, consulta_id, area_asignada, is_deleted, user_deleted, deleted_at, user_creator, created_at, updated_at) VALUES (UUID(), :consulta_id, :area_asignada, :is_deleted, :user_deleted, :deleted_at, :user_creator, :created_at, :updated_at)`

	rs, err := s.DB.NamedExec(sqlInsertConsulta, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) update(m *ConsultationAssignment) error {
	const sqlUpdate = `UPDATE area_medica_asignada SET area_asignada = :area_asignada, user_creator = :user_creator, updated_at = :updated_at WHERE id = :id`
	_, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	return nil
}

func (s sqlserver) updateByIDConsultation(m *ConsultationAssignment) error {
	const sqlUpdate = `UPDATE area_medica_asignada SET area_asignada = :area_asignada, user_creator = :user_creator, updated_at = :updated_at WHERE consulta_id = :consulta_id`
	_, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	return nil
}

func (s sqlserver) delete(id string) error {
	const psqlDelete = `DELETE FROM area_medica_asignada WHERE id = :id`
	m := ConsultationAssignment{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) deleteByIDConsultation(id string) error {
	const psqlDelete = `DELETE FROM area_medica_asignada WHERE consulta_id = :id`
	m := ConsultationAssignment{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) getByID(id string) (*ConsultationAssignment, error) {
	const sqlGetByID = `SELECT id, consulta_id, area_asignada, is_deleted, user_deleted, deleted_at, user_creator, created_at, updated_at FROM area_medica_asignada WHERE id = ?`

	mdl := ConsultationAssignment{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &mdl, nil
}

func (s sqlserver) getByIDConsultation(id string) (*ConsultationAssignment, error) {
	const sqlGetByIDConsultation = `SELECT id, consulta_id, area_asignada, is_deleted, user_deleted, deleted_at, user_creator, created_at, updated_at FROM area_medica_asignada WHERE consulta_id = ? order by created_at desc limit 1`

	mdl := ConsultationAssignment{}
	err := s.DB.Get(&mdl, sqlGetByIDConsultation, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &mdl, nil
}

func (s sqlserver) getAll() ([]*ConsultationAssignment, error) {
	var ms []*ConsultationAssignment
	const sqlGetAll = `SELECT id, consulta_id, area_asignada, is_deleted, user_deleted, deleted_at, user_creator, created_at, updated_at FROM area_medica_asignada`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}
