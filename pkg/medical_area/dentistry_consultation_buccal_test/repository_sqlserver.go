package dentistry_consultation_buccal_test

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

func newBuccalTestSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s sqlserver) create(m *BuccalTest) error {
	const sqlInsertBucal = `INSERT INTO odontologia_examen (id, consulta_odontologia_id, odontograma_img, cpod, observacion, ihos, comentarios, created_at, updated_at) VALUES (:id, :consulta_odontologia_id, :odontograma_img, :cpod, :observacion, :ihos, :comentarios, :created_at, :updated_at)`
	rs, err := s.DB.NamedExec(sqlInsertBucal, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

func (s sqlserver) update(m *BuccalTest) error {
	const sqlUpdate = `UPDATE odontologia_examen SET odontograma_img = :odontograma_img, cpod = :cpod, observacion = :observacion, ihos = :ihos, comentarios = :comentarios, updated_at = :updated_at WHERE id = :id`
	_, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}

	return nil
}

func (s sqlserver) delete(id string) error {
	const psqlDelete = `DELETE FROM odontologia_examen WHERE id = :id`
	m := BuccalTest{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) deleteByIDConsultation(consulta_odontologia_id string) error {
	const psqlDelete = `DELETE FROM odontologia_examen WHERE consulta_odontologia_id = :consulta_odontologia_id`
	m := BuccalTest{IDConsultaOdontologia: consulta_odontologia_id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) getByID(id string) (*BuccalTest, error) {
	const sqlGetByID = `SELECT id, consulta_odontologia_id, odontograma_img, cpod, observacion, ihos, comentarios, created_at, updated_at FROM odontologia_examen WHERE id = ?`
	mdl := BuccalTest{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s sqlserver) getByIDConsultation(id string) (*BuccalTest, error) {
	const sqlGetByID = `SELECT id, consulta_odontologia_id, odontograma_img, cpod, observacion, ihos, comentarios, created_at, updated_at FROM odontologia_examen WHERE consulta_odontologia_id = ?`
	mdl := BuccalTest{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s sqlserver) getAll() ([]*BuccalTest, error) {
	var ms []*BuccalTest
	const sqlGetAll = `SELECT id, consulta_odontologia_id, odontograma_img, cpod, observacion, ihos, comentarios, created_at, updated_at FROM odontologia_examen`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
