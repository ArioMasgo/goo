package nursing_consultation_visual_test

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

func newVisualTestSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s sqlserver) create(m *VisualTest) error {
	const sqlInsertVisual = `INSERT INTO enfermeria_examen_visual (id, consulta_enfermeria_id, ojo_derecho, ojo_izquierdo, presion_ocular, comentarios) VALUES (:id, :consulta_enfermeria_id, :ojo_derecho, :ojo_izquierdo, :presion_ocular, :comentarios)`
	rs, err := s.DB.NamedExec(sqlInsertVisual, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) update(m *VisualTest) error {
	const sqlUpdate = `UPDATE enfermeria_examen_visual SET ojo_derecho = :ojo_derecho, ojo_izquierdo = :ojo_izquierdo, presion_ocular = :presion_ocular, comentarios = :comentarios, user_creator = :user_creator, updated_at = :updated_at WHERE id = :id`
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) delete(id string) error {
	const psqlDelete = `DELETE FROM enfermeria_examen_visual WHERE id = :id`
	m := VisualTest{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) deleteByIDConsultation(consulta_enfermeria_id string) error {
	const psqlDelete = `DELETE FROM enfermeria_examen_visual WHERE consulta_enfermeria_id = :consulta_enfermeria_id`
	m := VisualTest{IDConsultaEnfermeria: consulta_enfermeria_id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) getByID(id string) (*VisualTest, error) {
	const sqlGetByID = `SELECT id, consulta_enfermeria_id, ojo_derecho, ojo_izquierdo, presion_ocular, comentarios, created_at, updated_at
                        FROM enfermeria_examen_visual WHERE consulta_enfermeria_id = ?`

	mdl := VisualTest{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &mdl, nil
}

func (s sqlserver) getAll() ([]*VisualTest, error) {
	var ms []*VisualTest
	const sqlGetAll = `SELECT id, consulta_enfermeria_id, ojo_derecho, ojo_izquierdo, presion_ocular, comentarios, created_at, updated_at FROM enfermeria_examen_visual`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}
