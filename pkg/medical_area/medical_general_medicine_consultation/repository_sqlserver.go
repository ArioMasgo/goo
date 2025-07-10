package medical_general_medicine_consultation

import (
	"database/sql"
	"dbu-api/internal/models"
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

func newGeneralMedicineConsultationSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s sqlserver) create(m *GeneralMedicineConsultation) error {
	const sqlInsertConsulta = `INSERT INTO medicina_consulta_general (id, consulta_id, fecha_hora, anamnesis, examen_clinico, indicaciones, is_deleted, user_creator, created_at, updated_at) VALUES (:id, :consulta_id,:fecha_hora, :anamnesis, :examen_clinico, :indicaciones, :is_deleted, :user_creator, :created_at, :updated_at)`

	rs, err := s.DB.NamedExec(sqlInsertConsulta, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) update(m *GeneralMedicineConsultation) error {
	const sqlUpdate = `UPDATE medicina_consulta_general SET fecha_hora = :fecha_hora, anamnesis = :anamnesis, examen_clinico = :examen_clinico, indicaciones = :indicaciones, updated_at = :updated_at WHERE id = :id`
	_, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	return nil
}

func (s sqlserver) delete(consulta_id string) error {
	const psqlDelete = `DELETE FROM medicina_consulta_general WHERE consulta_id = :consulta_id`
	m := GeneralMedicineConsultation{ConsultaID: consulta_id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) getByID(id string) (*GeneralMedicineConsultation, error) {
	const sqlGetByID = `SELECT id, fecha_hora, anamnesis, examen_clinico, indicaciones FROM medicina_consulta_general WHERE consulta_id = ?`
	mdl := GeneralMedicineConsultation{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s sqlserver) getByIDPatient(id string) ([]*GeneralMedicineConsultation, error) {
	var ms []*GeneralMedicineConsultation
	const sqlGetAll = `SELECT id, fecha_hora, anamnesis, examen_clinico, indicaciones FROM medicina_consulta_general WHERE id = ?`

	err := s.DB.Select(&ms, sqlGetAll, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s sqlserver) getAll() ([]*GeneralMedicineConsultation, error) {
	var ms []*GeneralMedicineConsultation
	const sqlGetAll = `SELECT id, fecha_hora, anamnesis, examen_clinico, indicaciones FROM medicina_consulta_general`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s sqlserver) getAllByDateExcel(fecha_inicio, fecha_fin string) ([]*models.ConsultationIntegralAttentionExcel, error) {
	var ms []*models.ConsultationIntegralAttentionExcel
	const sqlGetAll = `
		SELECT 
		cam.id as id, 
		cam.fecha_consulta as fecha_consulta, 
		p.tipo_persona as tipo_persona,
		p.escuela_profesional as escuela_profesional,
		p.sexo as sexo
	FROM consultas_areas_medicas cam
	JOIN pacientes p ON p.id = cam.paciente_id
	JOIN medicina_consulta_general mcg ON mcg.consulta_id = cam.id
	AND cam.created_at >= ?
	AND cam.created_at <= ?
	`

	err := s.DB.Select(&ms, sqlGetAll, fecha_inicio, fecha_fin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}
