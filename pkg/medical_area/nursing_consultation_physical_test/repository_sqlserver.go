package nursing_consultation_physical_test

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

func newPhysicalTestSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

func (s sqlserver) create(m *PhysicalTest) error {
	const sqlInsertFisico = `INSERT INTO enfermeria_examen_fisico (id, consulta_enfermeria_id, talla_peso, perimetro_cintura, indice_masa_corporal_img, presion_arterial, comentarios) VALUES (:id, :consulta_enfermeria_id, :talla_peso, :perimetro_cintura, :indice_masa_corporal_img, :presion_arterial, :comentarios)`
	rs, err := s.DB.NamedExec(sqlInsertFisico, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}
func (s sqlserver) update(m *PhysicalTest) error {
	const sqlUpdate = `UPDATE enfermeria_examen_fisico SET talla_peso = :talla_peso, perimetro_cintura = :perimetro_cintura, indice_masa_corporal_img = :indice_masa_corporal_img, presion_arterial = :presion_arterial, comentarios = :comentarios, user_creator = :user_creator, updated_at = :updated_at WHERE id = :id`
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
	const psqlDelete = `DELETE FROM enfermeria_examen_fisico WHERE id = :id`
	m := PhysicalTest{ID: id}
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
	const psqlDelete = `DELETE FROM enfermeria_examen_fisico WHERE consulta_enfermeria_id = :consulta_enfermeria_id`
	m := PhysicalTest{IDConsultaEnfermeria: consulta_enfermeria_id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s sqlserver) getByID(id string) (*PhysicalTest, error) {
	const sqlGetByID = `SELECT id, consulta_enfermeria_id, talla_peso, perimetro_cintura, indice_masa_corporal_img, presion_arterial, comentarios, created_at, updated_at
                        FROM enfermeria_examen_fisico WHERE consulta_enfermeria_id = ?`

	mdl := PhysicalTest{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &mdl, nil
}

func (s sqlserver) getAll() ([]*PhysicalTest, error) {
	var ms []*PhysicalTest
	const sqlGetAll = `SELECT id, consulta_enfermeria_id, talla_peso, perimetro_cintura, indice_masa_corporal_img, presion_arterial, comentarios, created_at, updated_at FROM enfermeria_examen_fisico`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}

	return ms, nil
}
