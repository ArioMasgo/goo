package dentistry_consultation_odontogram_review

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

func newOdontogramReviewSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s sqlserver) create(m *OdontogramReview) error {
	const sqlInsertOdontogram = `INSERT INTO odontologia_revision_odontograma (id, consulta_odontologia_id, caries, erupcionado, perdido, costo, fecha_pago, cpod, diagnostico, mes, comentarios) VALUES (:id, :consulta_odontologia_id, :caries, :erupcionado, :perdido, :costo, :fecha_pago, :cpod, :diagnostico, :mes, :comentarios)`
	rs, err := s.DB.NamedExec(sqlInsertOdontogram, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) update(m *OdontogramReview) error {
	const sqlUpdate = `UPDATE odontologia_revision_odontograma SET consulta_odontologia_id = :consulta_odontologia_id, caries = :caries, erupcionado = :erupcionado, perdido = :perdido, costo = :costo, fecha_pago = :fecha_pago, cpod = :cpod, diagnostico = :diagnostico, mes = :mes, comentarios = :comentarios WHERE id = :id`
	_, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	return nil
}

func (s sqlserver) delete(id string) error {
	const psqlDelete = `DELETE FROM odontologia_revision_odontograma WHERE id = :id`
	m := OdontogramReview{ID: id}
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
	const psqlDelete = `DELETE FROM odontologia_revision_odontograma WHERE consulta_odontologia_id = :id`
	m := OdontogramReview{IDConsultaOdontologia: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) getByID(id string) (*OdontogramReview, error) {
	const sqlGetByID = `SELECT id, consulta_odontologia_id, caries, erupcionado, perdido, costo, fecha_pago, cpod, diagnostico, mes, comentarios FROM odontologia_revision_odontograma WHERE id = ?`
	mdl := OdontogramReview{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s sqlserver) getByIDConsultation(id string) (*OdontogramReview, error) {
	const sqlGetByID = `SELECT id, consulta_odontologia_id, caries, erupcionado, perdido, costo, fecha_pago, cpod, diagnostico, mes, comentarios FROM odontologia_revision_odontograma WHERE consulta_odontologia_id = ?`
	mdl := OdontogramReview{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s sqlserver) getAll() ([]*OdontogramReview, error) {
	var ms []*OdontogramReview
	const sqlGetAll = `SELECT id, consulta_odontologia_id, caries, erupcionado, perdido, costo, fecha_pago, cpod, diagnostico, mes, comentarios FROM odontologia_revision_odontograma`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
