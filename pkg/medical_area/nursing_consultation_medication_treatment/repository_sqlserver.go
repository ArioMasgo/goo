package nursing_consultation_medication_treatment

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

func newMedicationTreatmentSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s sqlserver) create(m *MedicationTreatment) error {
	const sqlInsertConsulta = `INSERT INTO enfermeria_tratamiento_medicamentoso (id, consulta_enfermeria_id, nombre_generico_medicamento, via_administracion, hora_aplicacion, responsable_atencion, observaciones, area_solicitante, especialista_solicitante, is_deleted, user_creator, created_at, updated_at) VALUES (:id, :consulta_enfermeria_id, :nombre_generico_medicamento, :via_administracion, :hora_aplicacion, :responsable_atencion, :observaciones, :area_solicitante, :especialista_solicitante, :is_deleted, :user_creator, :created_at, :updated_at)`

	rs, err := s.DB.NamedExec(sqlInsertConsulta, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) update(m *MedicationTreatment) error {
	const sqlUpdate = `UPDATE enfermeria_tratamiento_medicamentoso SET nombre_generico_medicamento = :nombre_generico_medicamento, via_administracion = :via_administracion, hora_aplicacion = :hora_aplicacion, responsable_atencion = :responsable_atencion, observaciones = :observaciones, area_solicitante = :area_solicitante, especialista_solicitante = :especialista_solicitante, updated_at = :updated_at WHERE id = :id`
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
	const psqlDelete = `DELETE FROM enfermeria_tratamiento_medicamentoso WHERE id = :id`
	m := MedicationTreatment{ID: id}
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
	const psqlDelete = `DELETE FROM enfermeria_tratamiento_medicamentoso WHERE consulta_enfermeria_id = :consulta_enfermeria_id`
	m := MedicationTreatment{IDConsultaEnfermeria: consulta_enfermeria_id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) getByID(id string) (*MedicationTreatment, error) {
	const sqlGetByID = `SELECT id, consulta_enfermeria_id, nombre_generico_medicamento, via_administracion, hora_aplicacion, responsable_atencion, observaciones, area_solicitante, especialista_solicitante,is_deleted, user_deleted, deleted_at, user_creator, created_at, updated_at FROM enfermeria_tratamiento_medicamentoso WHERE id = ?`

	mdl := MedicationTreatment{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &mdl, nil
}

func (s sqlserver) getByIDConsultation(id string) ([]*MedicationTreatment, error) {
	var ms []*MedicationTreatment
	const sqlGetByIDConsultation = `SELECT id, consulta_enfermeria_id, nombre_generico_medicamento, via_administracion, hora_aplicacion, responsable_atencion, observaciones, area_solicitante, especialista_solicitante,is_deleted, user_deleted, deleted_at, user_creator, created_at, updated_at FROM enfermeria_tratamiento_medicamentoso WHERE consulta_enfermeria_id = ?`
	err := s.DB.Select(&ms, sqlGetByIDConsultation, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}

func (s sqlserver) getAll() ([]*MedicationTreatment, error) {
	var ms []*MedicationTreatment
	const sqlGetAll = `SELECT id,  consulta_enfermeria_id, nombre_generico_medicamento, via_administracion, hora_aplicacion, responsable_atencion, observaciones, area_solicitante, especialista_solicitante, is_deleted, user_deleted, deleted_at, user_creator, created_at, updated_at FROM enfermeria_tratamiento_medicamentoso`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}
