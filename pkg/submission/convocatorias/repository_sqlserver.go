package convocatorias

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"dbu-api/internal/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// sqlServer estructura de conexión a la BD de mssql
type sqlserver struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newConvocatoriasSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

func (s *sqlserver) create(m *Convocatorias) error {
	date := time.Now()
	m.UpdatedAt = &date
	m.CreatedAt = &date

	const sqlInsert = `INSERT INTO convocatorias 
        (fecha_inicio, fecha_fin, nombre, user_id, credito_minimo, nota_minima, 
        created_at, updated_at) 
        VALUES 
        (:fecha_inicio, :fecha_fin, :nombre, :user_id, :credito_minimo, :nota_minima, 
        :created_at, :updated_at)`

	rs, err := s.DB.NamedExec(sqlInsert, &m)
	if err != nil {
		return err
	}

	id, err := rs.LastInsertId()
	if err != nil {
		return err
	}

	if id == 0 {
		return fmt.Errorf("rows affected error")
	}

	m.ID = id // Cambiado a int64 para coincidir con bigint unsigned
	return nil
}

// Update actualiza un registro en la BD
func (s *sqlserver) update(m *Convocatorias) error {
	date := time.Now()
	m.UpdatedAt = &date

	const sqlUpdate = `UPDATE convocatorias 
        SET fecha_inicio = :fecha_inicio, 
            fecha_fin = :fecha_fin, 
            nombre = :nombre, 
            user_id = :user_id, 
            credito_minimo = :credito_minimo, 
            nota_minima = :nota_minima, 
            updated_at = :updated_at 
        WHERE id = :id`

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
func (s *sqlserver) delete(id int64) error {
	const sqlDelete = `DELETE FROM convocatorias WHERE id = :id`
	m := Convocatorias{ID: id}
	rs, err := s.DB.NamedExec(sqlDelete, &m)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}
	return nil
}

// GetByID consulta un registro por su ID
func (s *sqlserver) getByID(id int64) (*Convocatorias, error) {
	const sqlGetByID = `SELECT id, fecha_inicio, fecha_fin, nombre, user_id, 
        credito_minimo, nota_minima, created_at, updated_at 
        FROM convocatorias WHERE id = ?`

	mdl := Convocatorias{}
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
func (s *sqlserver) getAll() ([]*Convocatorias, error) {
	var ms []*Convocatorias
	const sqlGetAll = `SELECT id, fecha_inicio, fecha_fin, nombre, user_id, 
        credito_minimo, nota_minima, created_at, updated_at 
        FROM convocatorias`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) getAllByService(id int64) ([]*Convocatorias, error) {
	var ms []*Convocatorias
	const sqlGetAll = `SELECT c.id, c.fecha_inicio, c.fecha_fin, c.nombre, c.user_id, c.credito_minimo, c.nota_minima, c.created_at, c.updated_at 
FROM convocatorias c
LEFT JOIN convocatoria_servicio cs on (cs.convocatoria_id = c.id)
WHERE cs.servicio_id = ?`

	err := s.DB.Select(&ms, sqlGetAll, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) getActive() (*Convocatorias, error) {
	const sqlGetByID = `SELECT 
            id, 
            fecha_inicio, 
            fecha_fin, 
            nombre,
            user_id,
            created_at,
            updated_at,
            credito_minimo,
            nota_minima
        FROM convocatorias
        WHERE fecha_inicio <= CURRENT_TIMESTAMP()
        AND fecha_fin >= CURRENT_TIMESTAMP() LIMIT 1;`

	mdl := Convocatorias{}
	err := s.DB.Get(&mdl, sqlGetByID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}
