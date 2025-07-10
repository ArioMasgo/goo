package nursing_consultation_vaccine

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

func newVaccineSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s sqlserver) create(m *Vaccine) error {
	const sqlInsertVacuna = `INSERT INTO enfermeria_vacuna (id, consulta_enfermeria_id, tipo_vacuna, fecha_dosis, comentarios) VALUES (:id, :consulta_enfermeria_id, :tipo_vacuna, :fecha_dosis, :comentarios)`
	rs, err := s.DB.NamedExec(sqlInsertVacuna, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) update(m *Vaccine) error {
	const sqlUpdate = `UPDATE enfermeria_vacuna SET tipo_vacuna = :tipo_vacuna, fecha_dosis = :fecha_dosis, comentarios = :comentarios, user_creator = :user_creator, updated_at = :updated_at WHERE id = :id`
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
	const psqlDelete = `DELETE FROM enfermeria_vacuna WHERE id = :id`
	m := Vaccine{ID: id}
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
	const psqlDeleteByIDConsultation = `DELETE FROM enfermeria_vacuna WHERE consulta_enfermeria_id = :consulta_enfermeria_id`
	m := Vaccine{IDConsultaEnfermeria: consulta_enfermeria_id}
	rs, err := s.DB.NamedExec(psqlDeleteByIDConsultation, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) getByID(id string) ([]*Vaccine, error) {
	var ms []*Vaccine
	const sqlGetByID = `SELECT id, consulta_enfermeria_id, tipo_vacuna, fecha_dosis, comentarios, created_at, updated_at FROM enfermeria_vacuna WHERE consulta_enfermeria_id = ?`

	err := s.DB.Select(&ms, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}

func (s sqlserver) getByIDPatient(id string) ([]*Vaccine, error) {
	var ms []*Vaccine
	const sqlGetByID = `SELECT ev.id, ev.consulta_enfermeria_id, ev.tipo_vacuna, ev.fecha_dosis, ev.created_at, ev.updated_at FROM enfermeria_vacuna ev JOIN consultas_areas_medicas cam on cam.id = ev.consulta_enfermeria_id WHERE cam.paciente_id = ?`

	err := s.DB.Select(&ms, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}

func (s sqlserver) getAll() ([]*Vaccine, error) {
	var ms []*Vaccine
	const sqlGetAll = `SELECT id, consulta_enfermeria_id, tipo_vacuna, fecha_dosis, comentarios, created_at, updated_at FROM enfermeria_vacuna`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}

func (s sqlserver) getAllTypesVaccines() ([]*TypesVaccines, error) {
	var ms []*TypesVaccines
	const sqlGetAll = `SELECT id, nombre, estado, duracion_meses, created_at, updated_at FROM tipos_vacunas`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}

func (s sqlserver) getAllVaccinesByPatientDni(dni string) ([]*Vaccine, error) {
	var ms []*Vaccine
	const query = `SELECT ev.* FROM enfermeria_vacuna ev 
	JOIN consultas_areas_medicas cam ON ev.consulta_enfermeria_id = cam.id 
	JOIN pacientes p ON cam.paciente_id = p.id 
	WHERE p.dni = ?`

	err := s.DB.Select(&ms, query, dni)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}
