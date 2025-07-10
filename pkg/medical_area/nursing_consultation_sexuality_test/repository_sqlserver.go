package nursing_consultation_sexuality_test

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

func newSexualityTestSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s sqlserver) create(m *SexualityTest) error {
	const sqlInsertSexualidad = `INSERT INTO enfermeria_examen_sexualidad (id, consulta_enfermeria_id, actividad_sexual, planificacion_familiar, comentarios) VALUES (:id, :consulta_enfermeria_id, :actividad_sexual, :planificacion_familiar, :comentarios)`
	rs, err := s.DB.NamedExec(sqlInsertSexualidad, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) update(m *SexualityTest) error {
	const sqlUpdate = `UPDATE enfermeria_examen_sexualidad SET actividad_sexual = :actividad_sexual, planificacion_familiar = :planificacion_familiar, comentarios = :comentarios, user_creator = :user_creator, updated_at = :updated_at WHERE id = :id`
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
	const psqlDelete = `DELETE FROM enfermeria_examen_sexualidad WHERE id = :id`
	m := SexualityTest{ID: id}
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
	const psqlDelete = `DELETE FROM enfermeria_examen_sexualidad WHERE consulta_enfermeria_id = :consulta_enfermeria_id`
	m := SexualityTest{IDConsultaEnfermeria: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) getByID(id string) (*SexualityTest, error) {
	const sqlGetByID = `SELECT id, consulta_enfermeria_id, actividad_sexual, planificacion_familiar, comentarios, created_at, updated_at
                        FROM enfermeria_examen_sexualidad WHERE consulta_enfermeria_id = ?`

	mdl := SexualityTest{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &mdl, nil
}

func (s sqlserver) getAll() ([]*SexualityTest, error) {
	var ms []*SexualityTest
	const sqlGetAll = `SELECT id, consulta_enfermeria_id, actividad_sexual, planificacion_familiar, comentarios, created_at, updated_at FROM enfermeria_examen_sexualidad`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}
