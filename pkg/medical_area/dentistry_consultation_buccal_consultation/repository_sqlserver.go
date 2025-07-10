package dentistry_consultation_buccal_consultation

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

func newBuccalConsultationSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s sqlserver) create(m *BuccalConsultation) error {
	const sqlInsertBucal = `INSERT INTO odontologia_consulta (id, consulta_odontologia_id, relato, diagnostico, examen_auxiliar, examen_clinico, tratamiento, indicaciones, comentarios) VALUES (:id, :consulta_odontologia_id, :relato, :diagnostico, :examen_auxiliar, :examen_clinico, :tratamiento, :indicaciones, :comentarios)`
	rs, err := s.DB.NamedExec(sqlInsertBucal, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) update(m *BuccalConsultation) error {
	const sqlUpdate = `UPDATE odontologia_consulta SET  relato = :relato, diagnostico = :diagnostico, examen_auxiliar = :examen_auxiliar, examen_clinico = :examen_clinico, tratamiento = :tratamiento, indicaciones = :indicaciones, comentarios = :comentarios, updated_at = :updated_at WHERE id = :id`
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
	const psqlDelete = `DELETE FROM odontologia_consulta WHERE id = :id`
	m := BuccalConsultation{ID: id}
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
	const psqlDelete = `DELETE FROM odontologia_consulta WHERE consulta_odontologia_id = :consulta_odontologia_id`
	m := BuccalConsultation{IDConsultaOdontologia: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) getByID(id string) (*BuccalConsultation, error) {
	const sqlGetByID = `SELECT id, consulta_odontologia_id, relato, diagnostico, examen_auxiliar, examen_clinico, tratamiento, indicaciones, comentarios, created_at, updated_at FROM odontologia_consulta WHERE id = ?`
	mdl := BuccalConsultation{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s sqlserver) getByIDConsultation(id string) (*BuccalConsultation, error) {
	const sqlGetByID = `SELECT id, consulta_odontologia_id, relato, diagnostico, examen_auxiliar, examen_clinico, tratamiento, indicaciones, comentarios, created_at, updated_at FROM odontologia_consulta WHERE consulta_odontologia_id = ?`
	mdl := BuccalConsultation{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s sqlserver) getAll() ([]*BuccalConsultation, error) {
	var ms []*BuccalConsultation
	const sqlGetAll = `SELECT id, consulta_odontologia_id, relato, diagnostico, examen_auxiliar, examen_clinico, tratamiento, indicaciones, comentarios, created_at, updated_at FROM odontologia_consulta`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
