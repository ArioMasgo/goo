package nursing_consultation_accompanying_data

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

func newAccompanyingDataSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s sqlserver) create(m *AccompanyingData) error {
	const sqlInsertAcompanante = `INSERT INTO enfermeria_datos_acompanante (id, consulta_enfermeria_id, dni, nombres_apellidos, edad) VALUES (:id, :consulta_enfermeria_id, :dni, :nombres_apellidos, :edad)`
	rs, err := s.DB.NamedExec(sqlInsertAcompanante, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) update(m *AccompanyingData) error {
	const sqlUpdate = `UPDATE enfermeria_datos_acompanante SET dni = :dni, nombres_apellidos = :nombres_apellidos, edad = :edad, user_creator = :user_creator, updated_at = :updated_at WHERE id = :id`
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
	const psqlDelete = `DELETE FROM enfermeria_datos_acompanante WHERE id = :id`
	m := AccompanyingData{ID: id}
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
	const psqlDelete = `DELETE FROM enfermeria_datos_acompanante WHERE consulta_enfermeria_id = :consulta_enfermeria_id`
	m := AccompanyingData{IDConsultaEnfermeria: consulta_enfermeria_id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) getByID(id string) (*AccompanyingData, error) {
	const sqlGetByID = `SELECT id, consulta_enfermeria_id, dni, nombres_apellidos, edad, created_at, updated_at FROM enfermeria_datos_acompanante WHERE consulta_enfermeria_id = ?`
	mdl := AccompanyingData{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s sqlserver) getAll() ([]*AccompanyingData, error) {
	var ms []*AccompanyingData
	const sqlGetAll = `SELECT id, consulta_enfermeria_id, dni, nombres_apellidos, edad, created_at, updated_at FROM consulta_enfermeria`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}
