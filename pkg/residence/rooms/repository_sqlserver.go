package rooms

import (
	"database/sql"
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// sqlServer estructura de conexión a la BD de mssql
type sqlserver struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newRoomSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) create(m *Room) error {
	const sqlInsert = `INSERT INTO cuartos (id, numero, capacidad, estado, piso, residencia_id, created_by, created_at, updated_by, updated_at) 
  				 VALUES (:id, :numero, :capacidad, :estado, :piso, :residencia_id, :created_by, :created_at, :updated_by, :updated_at)`
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
func (s *sqlserver) update(m *Room) error {
	const sqlUpdate = `UPDATE cuartos SET numero = :numero, capacidad = :capacidad, estado = :estado, piso = :piso, 
  						residencia_id = :residencia_id, updated_by = :updated_by, updated_at = :updated_at 
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
func (s *sqlserver) delete(id string) error {

	// Physical delete
	const psqlDelete = `DELETE FROM cuartos WHERE id = :id`
	m := Room{ID: id}
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
func (s *sqlserver) getByID(id string) (*Room, error) {
	const sqlGetByID = `SELECT id, numero, capacidad, estado, piso, residencia_id, created_by, created_at, updated_by, updated_at 
  					FROM cuartos WHERE id = ?`
	mdl := Room{}
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
func (s *sqlserver) getAll() ([]*Room, error) {
	var ms []*Room
	const sqlGetAll = `SELECT id, numero, capacidad, estado, piso, residencia_id, created_by, created_at, updated_by, updated_at 
  					FROM cuartos`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) getAllRoomsByResidenceID(id string) ([]*Room, error) {
	var ms []*Room
	const sqlGetAll = `SELECT id, numero, capacidad, estado, piso, residencia_id, created_by, created_at, updated_by, updated_at 
  					FROM cuartos where residencia_id = ?`

	err := s.DB.Select(&ms, sqlGetAll, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) multiCreate(rooms []*Room) error {
	tx, err := s.DB.Beginx()
	if err != nil {
		logger.Error.Println(s.TxID, " - couldn't begin transaction", err)
		return err
	}
	defer tx.Rollback()

	query := `INSERT INTO cuartos (id, numero, capacidad, estado, piso, residencia_id, created_by, created_at, updated_by, updated_at) 
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	for _, room := range rooms {
		_, err := tx.Exec(query, room.ID, room.Number, room.Capacity, room.Status, room.Floor, room.ResidenceID, room.CreatedBy, room.CreatedAt, room.UpdatedBy, room.UpdatedAt)
		if err != nil {
			logger.Error.Println(s.TxID, " - couldn't insert room", err)
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		logger.Error.Println(s.TxID, " - couldn't commit transaction", err)
		return err
	}

	return nil
}

func (s *sqlserver) updateOnlyCharacteristics(m *Room) error {
	const sqlUpdate = `UPDATE cuartos SET capacidad = :capacidad, estado = :estado, updated_by = :updated_by, updated_at = :updated_at 
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

func (s *sqlserver) gtAllRoomsBySubmissionIDResidenceID(submissionID int64, residenceID string) ([]*Room, error) {
	var ms []*Room
	const sqlGetAll = `SELECT c.id, c.numero, c.capacidad, c.estado, c.piso, c.residencia_id, c.created_by, c.created_at, c.updated_by, c.updated_at FROM cuartos c
    where c.estado = 'habilitado' AND c.residencia_id = ? AND c.id NOT IN (
    SELECT ac.cuarto_id 
    FROM asignacion_cuartos ac 
    WHERE ac.convocatoria_id = ?
    AND ac.estado = 'activo') order by c.numero`

	err := s.DB.Select(&ms, sqlGetAll, residenceID, submissionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) getRoomsByResidence(residenceID string, page, limit int) ([]*Room, error) {
	var ms []*Room
	const sqlGetAll = `SELECT c.id, c.numero, c.capacidad, c.estado, c.piso, c.residencia_id, c.created_by, c.created_at, c.updated_by, c.updated_at 
  	FROM cuartos c
	WHERE c.residencia_id = ?
	ORDER BY c.piso, c.numero
	LIMIT ? OFFSET ?`

	offset := (page - 1) * limit
	err := s.DB.Select(&ms, sqlGetAll, residenceID, limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return ms, nil
}
