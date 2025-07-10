package nursing_consultation_routine_review

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

func newRoutineReviewSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s sqlserver) create(m *RoutineReview) error {
	const sqlInsertRevision = `INSERT INTO enfermeria_revision_rutina (id, consulta_enfermeria_id,fiebre_ultimo_quince_dias, tos_mas_quince_dias, secrecion_lesion_genitales, fecha_ultima_regla, comentarios) VALUES (:id, :consulta_enfermeria_id, :fiebre_ultimo_quince_dias, :tos_mas_quince_dias, :secrecion_lesion_genitales, :fecha_ultima_regla, :comentarios)`
	rs, err := s.DB.NamedExec(sqlInsertRevision, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) update(m *RoutineReview) error {
	const sqlUpdate = `UPDATE enfermeria_revision_rutina SET fiebre_ultimo_quince_dias = :fiebre_ultimo_quince_dias, tos_mas_quince_dias = :tos_mas_quince_dias, secrecion_lesion_genitales = :secrecion_lesion_genitales, fecha_ultima_regla = :fecha_ultima_regla, comentarios = :comentarios, user_creator = :user_creator, updated_at = :updated_at WHERE id = :id`
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
	const psqlDelete = `DELETE FROM enfermeria_revision_rutina WHERE id = :id`
	m := RoutineReview{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) deleteByIDConsultation(enfermeria_revision_rutina string) error {
	const psqlDelete = `DELETE FROM enfermeria_revision_rutina WHERE consulta_enfermeria_id = :consulta_enfermeria_id`
	m := RoutineReview{IDConsultaEnfermeria: enfermeria_revision_rutina}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) getByID(id string) (*RoutineReview, error) {
	const sqlGetByID = `SELECT id, consulta_enfermeria_id,fiebre_ultimo_quince_dias, tos_mas_quince_dias, secrecion_lesion_genitales, fecha_ultima_regla, comentarios, created_at, updated_at
                        FROM enfermeria_revision_rutina WHERE consulta_enfermeria_id = ?`

	mdl := RoutineReview{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &mdl, nil
}

func (s sqlserver) getAll() ([]*RoutineReview, error) {
	var ms []*RoutineReview
	const sqlGetAll = `SELECT id, consulta_enfermeria_id,fiebre_ultimo_quince_dias, tos_mas_quince_dias, secrecion_lesion_genitales, fecha_ultima_regla, comentarios, created_at, updated_at FROM enfermeria_revision_rutina`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}
