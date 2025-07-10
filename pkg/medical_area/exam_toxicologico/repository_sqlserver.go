package exam_toxicologico

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type sqlServer struct {
	db   *sqlx.DB
	txID string
}

func newRegistroToxicologicoSqlServerRepository(db *sqlx.DB, txID string) ServicesRegistroToxicologicoRepository {
	return &sqlServer{
		db:   db,
		txID: txID,
	}
}

func (s *sqlServer) create(m *RegistroToxicologico) error {
	query := `INSERT INTO registro_toxicologico (alumno_id, convocatoria_id, estado, comentario, id_usuario, fecha_creacion, fecha_actualizacion) 
              VALUES (?, ?, ?, ?, ?, ?, ?)`
	
	_, err := s.db.Exec(query, m.AlumnoID, m.ConvocatoriaID, m.Estado, m.Comentario, m.IDUsuario, m.FechaCreacion, m.FechaActualizacion)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlServer) update(m *RegistroToxicologico) error {
	query := `UPDATE registro_toxicologico SET estado = ?, comentario = ?, id_usuario = ?, fecha_actualizacion = ? WHERE id = ?`
	
	result, err := s.db.Exec(query, m.Estado, m.Comentario, m.IDUsuario, time.Now(), m.ID)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}
	
	return nil
}

func (s *sqlServer) delete(id int64) error {
	query := `DELETE FROM registro_toxicologico WHERE id = ?`
	
	result, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}
	
	return nil
}

func (s *sqlServer) getByID(id int64) (*RegistroToxicologico, error) {
	query := `SELECT id, alumno_id, convocatoria_id, estado, comentario, id_usuario, fecha_creacion, fecha_actualizacion 
              FROM registro_toxicologico WHERE id = ?`
	
	var registro RegistroToxicologico
	err := s.db.Get(&registro, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	
	return &registro, nil
}

func (s *sqlServer) getByAlumnoAndConvocatoria(alumnoID int64, convocatoriaID int64) (*RegistroToxicologico, error) {
	query := `SELECT id, alumno_id, convocatoria_id, estado, comentario, id_usuario, fecha_creacion, fecha_actualizacion 
              FROM registro_toxicologico WHERE alumno_id = ? AND convocatoria_id = ?`
	
	var registro RegistroToxicologico
	err := s.db.Get(&registro, query, alumnoID, convocatoriaID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	
	return &registro, nil
}

func (s *sqlServer) getAll() ([]*RegistroToxicologico, error) {
	query := `SELECT id, alumno_id, convocatoria_id, estado, comentario, id_usuario, fecha_creacion, fecha_actualizacion 
              FROM registro_toxicologico ORDER BY fecha_creacion DESC`
	
	var registros []*RegistroToxicologico
	err := s.db.Select(&registros, query)
	if err != nil {
		return nil, err
	}
	
	return registros, nil
}

func (s *sqlServer) getEstadosByConvocatoria(convocatoriaID int64) ([]*EstadoToxicologicoConvocatoria, error) {
	query := `SELECT 
                a.id as alumno_id,
                a.codigo_estudiante,
                a.nombres,
                a.apellido_paterno,
                a.apellido_materno,
                a.escuela_profesional,
                COALESCE(MAX(rt.estado), 'pendiente') as estado,
                MAX(rt.comentario) as comentario,
                MAX(rt.fecha_actualizacion) as fecha_examen,
                MAX(u.full_name) as usuario_nombre
              FROM solicitudes s
              INNER JOIN servicio_solicitado ss ON s.id = ss.solicitud_id
              INNER JOIN alumnos a ON s.alumno_id = a.id
              LEFT JOIN registro_toxicologico rt ON rt.alumno_id = a.id AND rt.convocatoria_id = s.convocatoria_id
              LEFT JOIN users u ON rt.id_usuario = u.id
              WHERE s.convocatoria_id = ? AND ss.estado = 'aprobado'
              GROUP BY a.id, a.codigo_estudiante, a.nombres, a.apellido_paterno, a.apellido_materno, a.escuela_profesional
              ORDER BY a.codigo_estudiante ASC`
	
	var estados []*EstadoToxicologicoConvocatoria
	err := s.db.Select(&estados, query, convocatoriaID)
	if err != nil {
		return nil, err
	}
	
	return estados, nil
}

func (s *sqlServer) existsByAlumnoAndConvocatoria(alumnoID int64, convocatoriaID int64) (bool, error) {
	query := `SELECT COUNT(*) FROM registro_toxicologico WHERE alumno_id = ? AND convocatoria_id = ?`
	
	var count int
	err := s.db.Get(&count, query, alumnoID, convocatoriaID)
	if err != nil {
		return false, err
	}
	
	return count > 0, nil
}
