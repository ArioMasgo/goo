package residence_configuration

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

func newResidenceConfigurationSqlServerRepository(db *sqlx.DB, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) create(m *ResidenceConfiguration) error {
	const sqlInsert = `INSERT INTO configuracion_residencias (id, porcentaje_fcea, porcentaje_ingenieria, nota_minima_fcea, nota_minima_ingenieria, residencia_id, es_cachimbo, created_by, created_at, updated_by, updated_at) 
	VALUES (:id, :porcentaje_fcea, :porcentaje_ingenieria, :nota_minima_fcea, :nota_minima_ingenieria, :residencia_id, :es_cachimbo, :created_by, :created_at, :updated_by, :updated_at)`
	rs, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

// Update actualiza un registro en la BD
func (s *sqlserver) update(m *ResidenceConfiguration) error {
	const sqlUpdate = `UPDATE configuracion_residencias SET porcentaje_fcea = :porcentaje_fcea, porcentaje_ingenieria = :porcentaje_ingenieria, nota_minima_fcea = :nota_minima_fcea, nota_minima_ingenieria = :nota_minima_ingenieria, es_cachimbo = :es_cachimbo, updated_by = :updated_by, updated_at = :updated_at WHERE id = :id`
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

// Delete elimina un registro de la BD
func (s *sqlserver) delete(id string) error {

	// Physical delete
	const psqlDelete = `DELETE FROM configuracion_residencias WHERE id = :id`
	m := ResidenceConfiguration{ID: id}
	rs, err := s.DB.NamedExec(psqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

// GetByID consulta un registro por su ID
func (s *sqlserver) getByID(id string) (*ResidenceConfiguration, error) {
	const sqlGetByID = `SELECT id, porcentaje_fcea, porcentaje_ingenieria, nota_minima_fcea, nota_minima_ingenieria, residencia_id, es_cachimbo, created_by, created_at, updated_by, updated_at FROM configuracion_residencias WHERE id = ?`
	mdl := ResidenceConfiguration{}
	err := s.DB.Get(&mdl, sqlGetByID, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// GetAll consulta todos los registros de la BD
func (s *sqlserver) getAll() ([]*ResidenceConfiguration, error) {
	var ms []*ResidenceConfiguration
	const sqlGetAll = `SELECT id, porcentaje_fcea, porcentaje_ingenieria, nota_minima_fcea, nota_minima_ingenieria, residencia_id, es_cachimbo, created_by, created_at, updated_by, updated_at FROM configuracion_residencias`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) getByResidenceID(residenciaID string) (*ResidenceConfiguration, error) {
	const sqlGetByID = `SELECT id, porcentaje_fcea, porcentaje_ingenieria, nota_minima_fcea, nota_minima_ingenieria, residencia_id, es_cachimbo, created_by, created_at, updated_by, updated_at FROM configuracion_residencias WHERE residencia_id = ?`
	mdl := ResidenceConfiguration{}
	err := s.DB.Get(&mdl, sqlGetByID, residenciaID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

func (s *sqlserver) updateByResidenceID(m *ResidenceConfiguration) error {
	const sqlUpdate = `UPDATE configuracion_residencias SET porcentaje_fcea = :porcentaje_fcea, porcentaje_ingenieria = :porcentaje_ingenieria, nota_minima_fcea = :nota_minima_fcea, nota_minima_ingenieria = :nota_minima_ingenieria, es_cachimbo = :es_cachimbo, updated_by = :updated_by, updated_at = :updated_at WHERE residencia_id = :residencia_id`
	rs, err := s.DB.NamedExec(sqlUpdate, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}
