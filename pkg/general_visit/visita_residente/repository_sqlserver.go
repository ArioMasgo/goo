package visita_residente

import (
	"database/sql"
	"dbu-api/internal/logger"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

type sqlServerRepository struct {
	db   *sqlx.DB
	txID string
}

func newVisitaResidenteSqlServerRepository(db *sqlx.DB, txID string) ServiceVisitaResidenteRepository {
	return &sqlServerRepository{
		db:   db,
		txID: txID,
	}
}

func (s *sqlServerRepository) create(visita *VisitaDomiciliaria) error {
	query := `
		INSERT INTO Visitas_Domiciliarias (alumno_id, estado, comentario, imagen_url, id_usuario)
		VALUES (?, ?, ?, ?, ?)
	`
		result, err := s.db.Exec(query, visita.AlumnoID, visita.Estado, visita.Comentario, visita.ImagenURL, visita.IDUsuario)
	if err != nil {
		logger.Error.Printf("%s - Error al crear visita domiciliaria: %v", s.txID, err)
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error al obtener ID generado: %w", err)
	}

	visita.ID = uint64(id)
	
	return nil
}

func (s *sqlServerRepository) getByID(id uint64) (*VisitaDomiciliaria, error) {
    query := `
        SELECT 
            vd.id, vd.alumno_id, vd.estado, vd.comentario, vd.imagen_url, 
            vd.id_usuario, vd.fecha_creacion, vd.fecha_actualizacion,
            CONCAT(a.nombres, ' ', a.apellido_paterno, ' ', a.apellido_materno) as alumno_nombre,
            a.codigo_estudiante as alumno_codigo,
            a.escuela_profesional,
            a.lugar_procedencia,
            u.full_name as usuario_nombre
        FROM Visitas_Domiciliarias vd
        LEFT JOIN alumnos a ON vd.alumno_id = a.id
        LEFT JOIN users u ON vd.id_usuario = u.id
        WHERE vd.id = ?
    `
    
    visita := &VisitaDomiciliaria{}
    err := s.db.Get(visita, query, id)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("record not found")
        }
        logger.Error.Printf("%s - Error al obtener visita por ID: %v", s.txID, err)
        return nil, err
    }
    
    return visita, nil
}

func (s *sqlServerRepository) update(visita *VisitaDomiciliaria) error {
	query := `
		UPDATE Visitas_Domiciliarias 
		SET alumno_id = ?, estado = ?, comentario = ?, imagen_url = ?, 
			id_usuario = ?, fecha_actualizacion = NOW()
		WHERE id = ?
	`
	
	result, err := s.db.Exec(query, visita.AlumnoID, visita.Estado, visita.Comentario, 
		visita.ImagenURL, visita.IDUsuario, visita.ID)
	if err != nil {
		logger.Error.Printf("%s - Error al actualizar visita: %v", s.txID, err)
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return fmt.Errorf("rows affected error")
	}
	
	return nil
}

func (s *sqlServerRepository) delete(id uint64) error {
	query := "DELETE FROM Visitas_Domiciliarias WHERE id = ?"
	
	result, err := s.db.Exec(query, id)
	if err != nil {
		logger.Error.Printf("%s - Error al eliminar visita: %v", s.txID, err)
		return err
	}
	
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	
	if rowsAffected == 0 {
		return fmt.Errorf("rows affected error")
	}
	
	return nil
}

func (s *sqlServerRepository) getAll() ([]*VisitaDomiciliaria, error) {
    query := `
        SELECT 
            vd.id, vd.alumno_id, vd.estado, vd.comentario, vd.imagen_url, 
            vd.id_usuario, vd.fecha_creacion, vd.fecha_actualizacion,
            CONCAT(a.nombres, ' ', a.apellido_paterno, ' ', a.apellido_materno) as alumno_nombre,
            a.codigo_estudiante as alumno_codigo,
            a.escuela_profesional,
            a.lugar_procedencia,
            u.full_name as usuario_nombre
        FROM Visitas_Domiciliarias vd
        LEFT JOIN alumnos a ON vd.alumno_id = a.id
        LEFT JOIN users u ON vd.id_usuario = u.id
        ORDER BY vd.fecha_creacion DESC
    `
    
    visitas := []*VisitaDomiciliaria{}
    err := s.db.Select(&visitas, query)
    if err != nil {
        logger.Error.Printf("%s - Error al obtener todas las visitas: %v", s.txID, err)
        return nil, err
    }
    
    return visitas, nil
}

func (s *sqlServerRepository) getAllWithFilters(filtros *FiltrosVisita) ([]*VisitaDomiciliaria, error) {
	query := `
		SELECT 
			vd.id, vd.alumno_id, vd.estado, vd.comentario, vd.imagen_url, 
			vd.id_usuario, vd.fecha_creacion, vd.fecha_actualizacion,
			CONCAT(a.nombres, ' ', a.apellido_paterno, ' ', a.apellido_materno) as alumno_nombre,
			a.codigo_estudiante as alumno_codigo,
			a.escuela_profesional,
			a.lugar_procedencia,
			u.full_name as usuario_nombre
		FROM Visitas_Domiciliarias vd
		LEFT JOIN alumnos a ON vd.alumno_id = a.id
		LEFT JOIN users u ON vd.id_usuario = u.id
	`
	
	conditions, args := s.buildWhereClause(filtros)
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}
	
	query += " ORDER BY vd.fecha_creacion DESC"
	
	if filtros != nil && filtros.Limit > 0 {
		query += fmt.Sprintf(" LIMIT %d OFFSET %d", filtros.Limit, filtros.Offset)
	}
	
	visitas := []*VisitaDomiciliaria{}
	err := s.db.Select(&visitas, query, args...)
	if err != nil {
		logger.Error.Printf("%s - Error al obtener visitas con filtros: %v", s.txID, err)
		return nil, err
	}
	
	return visitas, nil
}

func (s *sqlServerRepository) getByAlumnoID(alumnoID uint64) (*VisitaDomiciliaria, error) {
	query := `
		SELECT 
			vd.id, vd.alumno_id, vd.estado, vd.comentario, vd.imagen_url, 
			vd.id_usuario, vd.fecha_creacion, vd.fecha_actualizacion,
			CONCAT(a.nombres, ' ', a.apellido_paterno, ' ', a.apellido_materno) as alumno_nombre,
			a.codigo_estudiante as alumno_codigo,
			a.escuela_profesional,
			a.lugar_procedencia,
			u.full_name as usuario_nombre
		FROM Visitas_Domiciliarias vd
		LEFT JOIN alumnos a ON vd.alumno_id = a.id
		LEFT JOIN users u ON vd.id_usuario = u.id
		WHERE vd.alumno_id = ?
	`
	
	visita := &VisitaDomiciliaria{}
	err := s.db.Get(visita, query, alumnoID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("record not found")
		}
		logger.Error.Printf("%s - Error al obtener visita por alumno: %v", s.txID, err)
		return nil, err
	}
	
	return visita, nil
}

func (s *sqlServerRepository) existsByAlumnoID(alumnoID uint64) (bool, error) {
	query := "SELECT COUNT(*) FROM Visitas_Domiciliarias WHERE alumno_id = ?"
	
	var count int
	err := s.db.Get(&count, query, alumnoID)
	if err != nil {
		logger.Error.Printf("%s - Error al verificar existencia de visita: %v", s.txID, err)
		return false, err
	}
	
	return count > 0, nil
}

func (s *sqlServerRepository) getAlumnosPendientesVisitaPorConvocatoria(convocatoriaID uint64) ([]*AlumnoPendienteVisita, error) {
	query := `
		SELECT DISTINCT  
			a.id AS alumno_id,
			a.codigo_estudiante AS codigo,
			CONCAT(a.nombres, ' ', a.apellido_paterno, ' ', a.apellido_materno) AS nombre,
			a.DNI AS dni,
			a.celular_estudiante AS celular,
			a.direccion,
			a.escuela_profesional,
			a.lugar_procedencia,
			s.id AS solicitud_id,
			c.id AS convocatoria_id,
			c.nombre AS convocatoria_nombre
		FROM alumnos a
		INNER JOIN solicitudes s ON a.id = s.alumno_id
		INNER JOIN convocatorias c ON s.convocatoria_id = c.id
		INNER JOIN servicio_solicitado ss ON s.id = ss.solicitud_id
		INNER JOIN servicios srv ON ss.servicio_id = srv.id
		LEFT JOIN Visitas_Domiciliarias vd ON a.id = vd.alumno_id
		WHERE 
			c.id = ?
			AND ss.estado = 'aprobado'
			AND ss.servicio_id = 2
			AND vd.id IS NULL
		ORDER BY a.nombres, a.apellido_paterno;

	`
	
	alumnos := []*AlumnoPendienteVisita{}
	err := s.db.Select(&alumnos, query, convocatoriaID)
	if err != nil {
		logger.Error.Printf("%s - Error al obtener alumnos pendientes por convocatoria: %v", s.txID, err)
		return nil, err
	}
	
	return alumnos, nil
}

func (s *sqlServerRepository) getTodosAlumnosPendientesVisita() ([]*AlumnoPendienteVisita, error) {
	query := `
		SELECT DISTINCT 
			a.id AS alumno_id,
			a.codigo_estudiante AS codigo,
			CONCAT(a.nombres, ' ', a.apellido_paterno, ' ', a.apellido_materno) AS nombre,
			a.DNI AS dni,
			a.celular_estudiante AS celular,
			a.direccion,
			a.escuela_profesional,
			a.lugar_procedencia,
			s.id AS solicitud_id,
			c.id AS convocatoria_id,
			c.nombre AS convocatoria_nombre
		FROM alumnos a
		INNER JOIN solicitudes s ON a.id = s.alumno_id
		INNER JOIN convocatorias c ON s.convocatoria_id = c.id
		INNER JOIN servicio_solicitado ss ON s.id = ss.solicitud_id
		INNER JOIN servicios srv ON ss.servicio_id = srv.id
		LEFT JOIN Visitas_Domiciliarias vd ON a.id = vd.alumno_id
		WHERE ss.estado = 'aprobado'
		AND ss.servicio_id = 2
		AND vd.id IS NULL
		ORDER BY c.nombre, a.nombres, a.apellido_paterno
	`
	
	alumnos := []*AlumnoPendienteVisita{}
	err := s.db.Select(&alumnos, query)
	if err != nil {
		logger.Error.Printf("%s - Error al obtener todos los alumnos pendientes: %v", s.txID, err)
		return nil, err
	}
	
	return alumnos, nil
}

func (s *sqlServerRepository) getEstadisticas(convocatoriaID *uint64) (*EstadisticasVisita, error) {
	query := `
		SELECT 
			COALESCE(COUNT(*), 0) as total_visitas,
			COALESCE(SUM(CASE WHEN vd.estado = 'pendiente' THEN 1 ELSE 0 END), 0) as pendientes,
			COALESCE(SUM(CASE WHEN vd.estado = 'verificado' THEN 1 ELSE 0 END), 0) as verificadas,
			COALESCE(SUM(CASE WHEN vd.estado = 'observado' THEN 1 ELSE 0 END), 0) as observadas,
			COALESCE(SUM(CASE WHEN YEAR(vd.fecha_creacion) = YEAR(NOW()) 
				AND MONTH(vd.fecha_creacion) = MONTH(NOW()) THEN 1 ELSE 0 END), 0) as visitas_del_mes
		FROM Visitas_Domiciliarias vd
	`
	
	args := []interface{}{}
	
	if convocatoriaID != nil {
		query += `
			INNER JOIN alumnos a ON vd.alumno_id = a.id
			INNER JOIN solicitudes s ON a.id = s.alumno_id
			WHERE s.convocatoria_id = ?
		`
		args = append(args, *convocatoriaID)
	}
	
	stats := &EstadisticasVisita{}
	err := s.db.Get(stats, query, args...)
	if err != nil {
		logger.Error.Printf("%s - Error al obtener estadísticas: %v", s.txID, err)
		return nil, err
	}
	queryAlumnosSinVisita := `
		SELECT COALESCE(COUNT(DISTINCT a.id), 0)
		FROM alumnos a
		INNER JOIN solicitudes s ON a.id = s.alumno_id
		INNER JOIN servicio_solicitado ss ON s.id = ss.solicitud_id
		INNER JOIN servicios srv ON ss.servicio_id = srv.id
		LEFT JOIN Visitas_Domiciliarias vd ON a.id = vd.alumno_id
		WHERE 
			srv.id = 2
			AND ss.estado = 'aprobado'
			AND vd.id IS NULL
	`
	
	argsAlumnosSinVisita := []interface{}{}
	
	if convocatoriaID != nil {
		queryAlumnosSinVisita += " AND s.convocatoria_id = ?"
		argsAlumnosSinVisita = append(argsAlumnosSinVisita, *convocatoriaID)
	}
	
	err = s.db.Get(&stats.AlumnosSinVisita, queryAlumnosSinVisita, argsAlumnosSinVisita...)
	if err != nil {
		logger.Error.Printf("%s - Error al obtener alumnos sin visita: %v", s.txID, err)
		return nil, err
	}
	
	return stats, nil
}

func (s *sqlServerRepository) getEstadisticasPorEscuelaProfesional(convocatoriaID *uint64) ([]*EstadisticasPorEscuelaProfesional, error) {
	query := `
		SELECT 
			a.escuela_profesional,
			COUNT(*) AS total_visitados
		FROM Visitas_Domiciliarias vd
		INNER JOIN alumnos a ON vd.alumno_id = a.id
		GROUP BY a.escuela_profesional
		ORDER BY total_visitados DESC
	`
	estadisticas := []*EstadisticasPorEscuelaProfesional{}
	err := s.db.Select(&estadisticas, query)
	if err != nil {
		logger.Error.Printf("%s - Error al obtener estadísticas por escuela profesional: %v", s.txID, err)
		return nil, err
	}
	
	return estadisticas, nil
}

func (s *sqlServerRepository) getEstadisticasPorLugarProcedencia(convocatoriaID *uint64) ([]*EstadisticasPorLugarProcedencia, error) {
	query := `
		SELECT 
			(
				SELECT d.name
				FROM detalle_solicitudes ds
				JOIN requisitos r ON r.id = ds.requisito_id AND r.opciones = 'department'
				JOIN departaments d ON d.id = ds.opcion_seleccion
				WHERE ds.solicitud_id = s.id
				LIMIT 1
			) AS departamento,
			COUNT(*) AS total_visitados

		FROM Visitas_Domiciliarias vd
		INNER JOIN alumnos a ON vd.alumno_id = a.id
		INNER JOIN solicitudes s ON s.alumno_id = a.id

		GROUP BY departamento
		ORDER BY total_visitados DESC;

	`
	
	estadisticas := []*EstadisticasPorLugarProcedencia{}
	err := s.db.Select(&estadisticas, query)
	if err != nil {
		logger.Error.Printf("%s - Error al obtener estadísticas por lugar de procedencia: %v", s.txID, err)
		return nil, err
	}
	
	return estadisticas, nil
}

func (s *sqlServerRepository) getAlumnosPendientesPorDepartamento(convocatoriaID uint64, departamento string) ([]*AlumnoPendienteVisitaPorDepartamento, error) {
	query := `
		SELECT DISTINCT  
			a.id AS alumno_id,
			a.codigo_estudiante AS codigo,
			CONCAT(a.nombres, ' ', a.apellido_paterno, ' ', a.apellido_materno) AS nombre,
			a.escuela_profesional,

			-- Subconsultas para lugar de procedencia
			(SELECT d.name
			FROM detalle_solicitudes ds
			JOIN requisitos r ON r.id = ds.requisito_id AND r.opciones = 'department'
			JOIN departaments d ON d.id = ds.opcion_seleccion
			WHERE ds.solicitud_id = s.id
			LIMIT 1) AS departamento,

			(SELECT p.name
			FROM detalle_solicitudes ds
			JOIN requisitos r ON r.id = ds.requisito_id AND r.opciones = 'province'
			JOIN provinces p ON p.id = ds.opcion_seleccion
			WHERE ds.solicitud_id = s.id
			LIMIT 1) AS provincia,

			(SELECT di.name
			FROM detalle_solicitudes ds
			JOIN requisitos r ON r.id = ds.requisito_id AND r.opciones = 'district'
			JOIN districts di ON di.id = ds.opcion_seleccion
			WHERE ds.solicitud_id = s.id
			LIMIT 1) AS distrito,

			a.direccion,
			a.celular_estudiante AS celular,
			a.celular_padre AS celular_padre,
			c.nombre AS convocatoria_nombre

		FROM alumnos a
		INNER JOIN solicitudes s ON a.id = s.alumno_id
		INNER JOIN convocatorias c ON s.convocatoria_id = c.id
		INNER JOIN servicio_solicitado ss ON s.id = ss.solicitud_id
		INNER JOIN servicios srv ON ss.servicio_id = srv.id
		LEFT JOIN Visitas_Domiciliarias vd ON a.id = vd.alumno_id

		WHERE 
			c.id = ? -- id de convocatoria
			AND ss.estado = 'aprobado'
			AND ss.servicio_id = 2
			AND vd.id IS NULL

			-- Aquí filtramos por nombre de departamento obtenido desde la subconsulta
			AND LOWER((
				SELECT d.name
				FROM detalle_solicitudes ds
				JOIN requisitos r ON r.id = ds.requisito_id AND r.opciones = 'department'
				JOIN departaments d ON d.id = ds.opcion_seleccion
				WHERE ds.solicitud_id = s.id
				LIMIT 1
			)) = LOWER(?)

		ORDER BY provincia, distrito, a.nombres, a.apellido_paterno;

	`
	
	alumnos := []*AlumnoPendienteVisitaPorDepartamento{}
	err := s.db.Select(&alumnos, query, convocatoriaID, departamento)
	if err != nil {
		logger.Error.Printf("%s - Error al obtener alumnos pendientes por departamento: %v", s.txID, err)
		return nil, err
	}
	
	return alumnos, nil
}

func (s *sqlServerRepository) getTodosAlumnosPorConvocatoria(convocatoriaID uint64) ([]*AlumnoCompleto, error) {
	query := `
		SELECT 
			a.id AS alumno_id,
			a.codigo_estudiante AS codigo,
			CONCAT(a.nombres, ' ', a.apellido_paterno, ' ', a.apellido_materno) AS nombre,
			a.DNI AS dni,
			a.celular_estudiante AS celular,
			a.celular_padre AS celular_padre,
			a.direccion,
			a.escuela_profesional,
			a.lugar_procedencia,

			-- Departamento desde detalle_solicitudes
			(SELECT d.name
			FROM detalle_solicitudes ds
			JOIN requisitos r ON r.id = ds.requisito_id AND r.opciones = 'department'
			JOIN departaments d ON d.id = ds.opcion_seleccion
			WHERE ds.solicitud_id = s.id
			LIMIT 1) AS departamento,

			-- Provincia desde detalle_solicitudes
			(SELECT p.name
			FROM detalle_solicitudes ds
			JOIN requisitos r ON r.id = ds.requisito_id AND r.opciones = 'province'
			JOIN provinces p ON p.id = ds.opcion_seleccion
			WHERE ds.solicitud_id = s.id
			LIMIT 1) AS provincia,

			-- Distrito desde detalle_solicitudes
			(SELECT di.name
			FROM detalle_solicitudes ds
			JOIN requisitos r ON r.id = ds.requisito_id AND r.opciones = 'district'
			JOIN districts di ON di.id = ds.opcion_seleccion
			WHERE ds.solicitud_id = s.id
			LIMIT 1) AS distrito,

			s.id AS solicitud_id,
			c.id AS convocatoria_id,
			c.nombre AS convocatoria_nombre,

			-- Estado de la visita (si existe)
			CASE 
				WHEN vd.id IS NOT NULL THEN vd.estado
				ELSE 'pendiente'
			END AS estado_visita

		FROM alumnos a
		INNER JOIN solicitudes s ON a.id = s.alumno_id
		INNER JOIN convocatorias c ON s.convocatoria_id = c.id
		INNER JOIN servicio_solicitado ss ON s.id = ss.solicitud_id
		INNER JOIN servicios srv ON ss.servicio_id = srv.id
		LEFT JOIN Visitas_Domiciliarias vd ON a.id = vd.alumno_id

		WHERE 
			c.id = ? -- ID de la convocatoria
			AND ss.estado = 'aprobado'
			AND srv.id = 2 -- Servicio: residencia

		ORDER BY estado_visita DESC, a.nombres, a.apellido_paterno;

	`
	
	alumnos := []*AlumnoCompleto{}
	err := s.db.Select(&alumnos, query, convocatoriaID)
	if err != nil {
		logger.Error.Printf("%s - Error al obtener todos los alumnos por convocatoria: %v", s.txID, err)
		return nil, err
	}
	
	return alumnos, nil
}

func (s *sqlServerRepository) buildWhereClause(filtros *FiltrosVisita) ([]string, []interface{}) {
	if filtros == nil {
		return nil, nil
	}
	
	conditions := []string{}
	args := []interface{}{}
	
	if filtros.Estado != nil {
		conditions = append(conditions, "vd.estado = ?")
		args = append(args, *filtros.Estado)
	}
	
	if filtros.AlumnoID != nil {
		conditions = append(conditions, "vd.alumno_id = ?")
		args = append(args, *filtros.AlumnoID)
	}
	
	if filtros.IDUsuario != nil {
		conditions = append(conditions, "vd.id_usuario = ?")
		args = append(args, *filtros.IDUsuario)
	}
	
	if filtros.FechaInicio != nil {
		conditions = append(conditions, "vd.fecha_creacion >= ?")
		args = append(args, *filtros.FechaInicio)
	}
	
	if filtros.FechaFin != nil {
		conditions = append(conditions, "vd.fecha_creacion <= ?")
		args = append(args, *filtros.FechaFin)
	}
		return conditions, args
}