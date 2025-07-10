package medical_consultation

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

func newMedicalConsultationSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s sqlserver) create(m *MedicalConsultation) error {
	const sqlInsertConsulta = `INSERT INTO consultas_areas_medicas (id, paciente_id, fecha_consulta, area_medica, user_creator, created_at, updated_at) VALUES (:id, :paciente_id, :fecha_consulta, :area_medica, :user_creator, :created_at, :updated_at)`

	rs, err := s.DB.NamedExec(sqlInsertConsulta, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) update(m *MedicalConsultation) error {
	const sqlUpdate = `UPDATE consultas_areas_medicas SET paciente_id = :paciente_id, fecha_consulta = :fecha_consulta, area_medica = :area_medica, user_creator = :user_creator, created_at = :created_at, updated_at = :updated_at WHERE id = :id`
	_, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	return nil
}

func (s sqlserver) delete(id string) error {
	const psqlDelete = `DELETE FROM consultas_areas_medicas WHERE id = :id`
	m := MedicalConsultation{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) getByID(id string) (*MedicalConsultation, error) {
	const sqlGetByID = `SELECT id, paciente_id, fecha_consulta, area_medica FROM consultas_areas_medicas WHERE id = ?`
	mdl := MedicalConsultation{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s sqlserver) getByIDPatient(id string) ([]*MedicalConsultation, error) {
	var ms []*MedicalConsultation
	const sqlGetAll = `SELECT id, paciente_id, fecha_consulta, area_medica FROM consultas_areas_medicas WHERE paciente_id = ?`

	err := s.DB.Select(&ms, sqlGetAll, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s sqlserver) getAll() ([]*MedicalConsultation, error) {
	var ms []*MedicalConsultation
	const sqlGetAll = `SELECT id, paciente_id, fecha_consulta, area_medica FROM consultas_areas_medicas`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
