package nursing_consultation_preferential_test

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

func newPreferentialTestSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s sqlserver) create(m *PreferentialTest) error {
	const sqlInsertPreferencial = `INSERT INTO enfermeria_examen_preferencial (id, consulta_enfermeria_id, aparato_respiratorio, aparato_cardiovascular, aparato_digestivo, aparato_genitourinario, papanicolau, examen_mama, comentarios) VALUES (:id, :consulta_enfermeria_id, :aparato_respiratorio, :aparato_cardiovascular, :aparato_digestivo, :aparato_genitourinario, :papanicolau, :examen_mama, :comentarios)`
	rs, err := s.DB.NamedExec(sqlInsertPreferencial, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) update(m *PreferentialTest) error {
	const sqlUpdate = `UPDATE enfermeria_examen_preferencial SET aparato_respiratorio = :aparato_respiratorio, aparato_cardiovascular = :aparato_cardiovascular, aparato_digestivo = :aparato_digestivo, aparato_genitourinario = :aparato_genitourinario, papanicolau = :papanicolau, examen_mama = :examen_mama, comentarios = :comentarios, user_creator = :user_creator, updated_at = :updated_at WHERE id = :id`
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
	const psqlDelete = `DELETE FROM enfermeria_examen_preferencial WHERE id = :id`
	m := PreferentialTest{ID: id}
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
	const psqlDelete = `DELETE FROM enfermeria_examen_preferencial WHERE consulta_enfermeria_id = :consulta_enfermeria_id`
	m := PreferentialTest{IDConsultaEnfermeria: consulta_enfermeria_id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) getByID(id string) (*PreferentialTest, error) {
	const sqlGetByID = `SELECT id, consulta_enfermeria_id, aparato_respiratorio, aparato_cardiovascular, aparato_digestivo, aparato_genitourinario, papanicolau, examen_mama, comentarios, created_at, updated_at
                        FROM enfermeria_examen_preferencial WHERE consulta_enfermeria_id = ?`

	mdl := PreferentialTest{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &mdl, nil
}

func (s sqlserver) getAll() ([]*PreferentialTest, error) {
	var ms []*PreferentialTest
	const sqlGetAll = `SELECT id, consulta_enfermeria_id, aparato_respiratorio, aparato_cardiovascular, aparato_digestivo, aparato_genitourinario, papanicolau, examen_mama, comentarios, created_at, updated_at FROM enfermeria_examen_preferencial`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}
