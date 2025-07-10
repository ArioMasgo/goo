package dentistry_consultation_buccal_procedure

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

func newBuccalProcedureSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s *sqlserver) existsByReceipt(recibo string) (bool, error) {
	var exists bool = false
	var count int
	const sqlExistsByReceipt = `SELECT COUNT(1) FROM odontologia_procedimiento WHERE recibo = ?`
	err := s.DB.QueryRow(sqlExistsByReceipt, recibo).Scan(&count)
	if err != nil {
		return false, err
	}
	if count > 0 {
		exists = true
	}
	return exists, nil
}

func (s sqlserver) create(m *BuccalProcedure) error {
	const sqlInsertBucal = `INSERT INTO odontologia_procedimiento (id, consulta_odontologia_id, tipo_procedimiento, recibo, costo, fecha_pago, pieza_dental, comentarios) VALUES (:id, :consulta_odontologia_id, :tipo_procedimiento, :recibo, :costo, :fecha_pago, :pieza_dental, :comentarios)`
	rs, err := s.DB.NamedExec(sqlInsertBucal, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) update(m *BuccalProcedure) error {
	const sqlUpdate = `UPDATE odontologia_procedimiento SET tipo_procedimiento = :tipo_procedimiento, recibo = :recibo, costo = :costo, fecha_pago = :fecha_pago, pieza_dental = :pieza_dental, comentarios = :comentarios, updated_at = :updated_at WHERE id = :id`
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
	const psqlDelete = `DELETE FROM odontologia_procedimiento WHERE id = :id`
	m := BuccalProcedure{ID: id}
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
	const psqlDelete = `DELETE FROM odontologia_procedimiento WHERE consulta_odontologia_id = :consulta_odontologia_id`
	m := BuccalProcedure{IDConsultaOdontologia: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) getByID(id string) (*BuccalProcedure, error) {
	const sqlGetByID = `SELECT id, consulta_odontologia_id, tipo_procedimiento, recibo, costo, fecha_pago, pieza_dental, comentarios, created_at, updated_at FROM odontologia_procedimiento WHERE id = ?`
	mdl := BuccalProcedure{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s sqlserver) getByIDConsultation(id string) (*BuccalProcedure, error) {
	const sqlGetByID = `SELECT id, consulta_odontologia_id, tipo_procedimiento, recibo, costo, fecha_pago, pieza_dental, comentarios, created_at, updated_at FROM odontologia_procedimiento WHERE consulta_odontologia_id = ?`
	mdl := BuccalProcedure{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s sqlserver) getAll() ([]*BuccalProcedure, error) {
	var ms []*BuccalProcedure
	const sqlGetAll = `SELECT id, consulta_odontologia_id, tipo_procedimiento, recibo, costo, fecha_pago, pieza_dental, comentarios, created_at, updated_at FROM odontologia_procedimiento`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s sqlserver) getBuccalProceduresExcel(fecha_inicio, fecha_fin string) ([]*models.PerformedProceduresExcel, error) {
	const sqlGetAll = `
		SELECT 
			cam.id as id, 
			cam.fecha_consulta as fecha_consulta, 
			p.tipo_persona as tipo_persona,
			p.escuela_profesional as escuela_profesional,
			(CASE 
				WHEN oc.consulta_odontologia_id IS NOT NULL THEN 'CONSULTA DENTAL'
				WHEN oe.consulta_odontologia_id IS NOT NULL THEN 'DIAGNÓSTICO EXAMEN BUCAL'
				ELSE (CASE 
						WHEN op.tipo_procedimiento in ('profilaxis dental','destartraje dental') THEN 'PERIODONCIA PROFILAXIS'
						WHEN op.tipo_procedimiento in ('resina simple','resina compuesta') THEN 'OPERATORIA RESINA'
						WHEN op.tipo_procedimiento in ('exodoncia simple','exodoncia compleja','curetaje alveolar','apertura cameral','endodoncia anterior','endodoncia posterior','cementado de corona') THEN 'CIRUGÍA EXODONCIA'
						WHEN op.tipo_procedimiento in ('aplicación de flúor gel') THEN 'PREVENCIÓN FLUORIZACIÓN'
						WHEN op.tipo_procedimiento in ('prevención colutorio') THEN 'PREVENCIÓN: COLUTORIO CH2+CPC 0.12%'
						WHEN op.tipo_procedimiento in ('prevención IHO - TCA cepillado') THEN 'PREVENCIÓN: IHO - TCA CEPILLADO'
						WHEN op.tipo_procedimiento in ('radiografía') THEN 'RADIOGRAFÍA DENTAL'
						ELSE ''
					END)
			END) AS tipo_procedimiento,
			p.sexo as sexo
		FROM consultas_areas_medicas cam
		JOIN pacientes p ON p.id = cam.paciente_id
		LEFT JOIN odontologia_procedimiento op ON op.consulta_odontologia_id = cam.id
		LEFT JOIN odontologia_consulta oc ON oc.consulta_odontologia_id = cam.id
		LEFT JOIN odontologia_examen oe ON oe.consulta_odontologia_id = cam.id
		WHERE (oc.consulta_odontologia_id IS NOT NULL 
		OR oe.consulta_odontologia_id IS NOT NULL 
		OR op.consulta_odontologia_id IS NOT NULL)
		AND cam.created_at >= ?
		AND cam.created_at <= ?
		`
	var ms []*models.PerformedProceduresExcel
	err := s.DB.Select(&ms, sqlGetAll, fecha_inicio, fecha_fin)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
