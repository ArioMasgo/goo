package asignaciones_cuartos

import (
	"database/sql"
	"dbu-api/internal/models"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

// sqlServer estructura de conexión a la BD de mssql
type sqlserver struct {
	DB   *sqlx.DB
	user *models.User
	TxID string
}

func newAsignacionesCuartosSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

// Create registra en la BD
func (s *sqlserver) create(m *RoomAssignment) error {
	const sqlInsert = `INSERT INTO asignacion_cuartos (id, alumno_id, cuarto_id, convocatoria_id, 
                       fecha_asignacion, estado, observaciones, created_at, updated_at)
                       VALUES (:id, :alumno_id, :cuarto_id, :convocatoria_id,
                       :fecha_asignacion, :estado, :observaciones, :created_at, :updated_at)`
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
func (s *sqlserver) update(m *RoomAssignment) error {
	const sqlUpdate = `UPDATE asignacion_cuartos SET 
                      alumno_id = :alumno_id, 
                      cuarto_id = :cuarto_id,
                      convocatoria_id = :convocatoria_id,
                      fecha_asignacion = :fecha_asignacion,
                      estado = :estado,
                      observaciones = :observaciones,
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
func (s *sqlserver) delete(m *RoomAssignment) error {
	const sqlDelete = `UPDATE asignacion_cuartos SET estado = :estado, observaciones = :observaciones
    WHERE alumno_id = :alumno_id AND cuarto_id = :cuarto_id 
    AND convocatoria_id = :convocatoria_id`
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
func (s *sqlserver) getByID(id string) (*RoomAssignment, error) {
	const sqlGetByID = `SELECT id, alumno_id, cuarto_id, convocatoria_id, 
                       fecha_asignacion, estado, observaciones, created_at, updated_at 
                       FROM asignacion_cuartos WHERE id = ?`
	mdl := RoomAssignment{}
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
func (s *sqlserver) getAll() ([]*RoomAssignment, error) {
	var ms []*RoomAssignment
	const sqlGetAll = `SELECT id, alumno_id, cuarto_id, convocatoria_id, 
                      fecha_asignacion, estado, observaciones, created_at, updated_at 
                      FROM asignacion_cuartos`

	err := s.DB.Select(&ms, sqlGetAll)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) getRoomAssignmentByRoomIDSubmissionID(roomID string, submissionID int64) ([]*RoomAssignment, error) {
	var ms []*RoomAssignment
	const sqlGetByID = `SELECT id, alumno_id, cuarto_id, convocatoria_id, 
                       fecha_asignacion, estado, observaciones, created_at, updated_at 
                       FROM asignacion_cuartos WHERE cuarto_id = ? AND convocatoria_id = ?`
	err := s.DB.Select(&ms, sqlGetByID, roomID, submissionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) multiAssign(assignments []string) error {
	sqlInsert := fmt.Sprintf(`
        INSERT INTO asignacion_cuartos (
            id,
            alumno_id,
            cuarto_id,
            convocatoria_id,
            fecha_asignacion,
            estado,
            observaciones,
            created_at,
            updated_at
        ) VALUES %s;
    `, strings.Join(assignments, ", "))

	rs, err := s.DB.Exec(sqlInsert)
	if err != nil {
		return err
	}
	if i, _ := rs.RowsAffected(); i == 0 {
		return fmt.Errorf("rows affected error")
	}

	return nil
}

func (s *sqlserver) getAllRoomAssignmentsByStudentIDANDSubmissionID(studentID, callID int64) ([]*RoomAssignment, error) {
	var ms []*RoomAssignment
	const sqlGetAll = `SELECT id, alumno_id, cuarto_id, convocatoria_id, 
                      fecha_asignacion, estado, observaciones, created_at, updated_at 
                      FROM asignacion_cuartos WHERE alumno_id = ? AND convocatoria_id = ?`

	err := s.DB.Select(&ms, sqlGetAll, studentID, callID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}
