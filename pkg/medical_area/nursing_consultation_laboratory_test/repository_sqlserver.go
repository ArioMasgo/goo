package nursing_consultation_laboratory_test

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

func newLaboratoryTestSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s sqlserver) create(m *LaboratoryTest) error {
	const sqlInsertLaboratorio = `INSERT INTO enfermeria_examen_laboratorio (id, consulta_enfermeria_id, serologia, bk, hemograma, examen_orina, colesterol, glucosa, comentarios) VALUES (:id, :consulta_enfermeria_id, :serologia, :bk, :hemograma, :examen_orina, :colesterol, :glucosa, :comentarios)`
	rs, err := s.DB.NamedExec(sqlInsertLaboratorio, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) update(m *LaboratoryTest) error {
	const sqlUpdate = `UPDATE enfermeria_examen_laboratorio SET serologia = :serologia, bk = :bk, hemograma = :hemograma, examen_orina = :examen_orina, colesterol = :colesterol, glucosa = :glucosa, comentarios = :comentarios, user_creator = :user_creator, updated_at = :updated_at WHERE id = :id`
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
	const psqlDelete = `DELETE FROM enfermeria_examen_laboratorio WHERE id = :id`
	m := LaboratoryTest{ID: id}
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
	const psqlDelete = `DELETE FROM enfermeria_examen_laboratorio WHERE consulta_enfermeria_id = :consulta_enfermeria_id`
	m := LaboratoryTest{IDConsultaEnfermeria: consulta_enfermeria_id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) getByID(id string) (*LaboratoryTest, error) {
	const sqlGetByID = `SELECT id, consulta_enfermeria_id, serologia, bk, hemograma, examen_orina, colesterol, glucosa, comentarios, created_at, updated_at FROM enfermeria_examen_laboratorio WHERE consulta_enfermeria_id = ?`
	mdl := LaboratoryTest{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s sqlserver) getAll() ([]*LaboratoryTest, error) {
	var ms []*LaboratoryTest
	const sqlGetAll = `SELECT id, consulta_enfermeria_id, serologia, bk, hemograma, examen_orina, colesterol, glucosa, comentarios, created_at, updated_at FROM enfermeria_examen_laboratorio`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}
