package alumnos

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

func newConvocatoriasSqlServerRepository(db *sqlx.DB, user *models.User, txID string) *sqlserver {
	return &sqlserver{
		DB:   db,
		user: user,
		TxID: txID,
	}
}

func (s *sqlserver) getStudentsAcceptedBySubmission(submissionID int64) ([]*Alumno, error) {
	var ms []*Alumno
	const sqlGetAll = `SELECT alm.id, alm.codigo_estudiante, alm.DNI, alm.nombres, alm.apellido_paterno, alm.apellido_materno, alm.sexo, alm.facultad, alm.escuela_profesional, alm.ultimo_semestre, alm.modalidad_ingreso,
lugar_procedencia, alm.lugar_nacimiento, alm.edad, alm.correo_institucional, alm.direccion, alm.fecha_nacimiento, alm.correo_personal, alm.celular_estudiante, alm.celular_padre, alm.estado_matricula, alm.
creditos_matriculados, alm.num_semestres_cursados, alm.pps, alm.ppa, alm.tca, alm.convocatoria_id, alm.created_at, alm.updated_at FROM solicitudes sol
JOIN servicio_solicitado ssol ON (sol.id = ssol.solicitud_id)
JOIN alumnos alm ON (sol.alumno_id = alm.id)
WHERE sol.convocatoria_id = ? AND ssol.estado = 'aprobado' AND alm.num_semestres_cursados > 1 AND NOT EXISTS (
   SELECT 1 
   FROM asignacion_cuartos ac 
   WHERE ac.alumno_id = alm.id
);`

	err := s.DB.Select(&ms, sqlGetAll, submissionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) getStudentsAcceptedBySubmissionNewbie(submissionID int64) ([]*Alumno, error) {
	var ms []*Alumno
	const sqlGetAll = `SELECT alm.id, alm.codigo_estudiante, alm.DNI, alm.nombres, alm.apellido_paterno, alm.apellido_materno, alm.sexo, alm.facultad, alm.escuela_profesional, alm.ultimo_semestre, alm.modalidad_ingreso,
lugar_procedencia, alm.lugar_nacimiento, alm.edad, alm.correo_institucional, alm.direccion, alm.fecha_nacimiento, alm.correo_personal, alm.celular_estudiante, alm.celular_padre, alm.estado_matricula, alm.
creditos_matriculados, alm.num_semestres_cursados, alm.pps, alm.ppa, alm.tca, alm.convocatoria_id, alm.created_at, alm.updated_at FROM solicitudes sol
JOIN servicio_solicitado ssol ON (sol.id = ssol.solicitud_id)
JOIN alumnos alm ON (sol.alumno_id = alm.id)
WHERE sol.convocatoria_id = ? AND ssol.estado = 'aprobado' AND alm.num_semestres_cursados <= 1 AND NOT EXISTS (
   SELECT 1 
   FROM asignacion_cuartos ac 
   WHERE ac.alumno_id = alm.id
);`

	err := s.DB.Select(&ms, sqlGetAll, submissionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return ms, err
	}
	return ms, nil
}

func (s *sqlserver) getStudentAcceptedBySubmission(submissionID, studentID int64) (*Alumno, error) {
	const sqlGetByID = `SELECT alm.id, alm.codigo_estudiante, alm.DNI, alm.nombres, alm.apellido_paterno, alm.apellido_materno, alm.sexo, alm.facultad, alm.escuela_profesional, alm.ultimo_semestre, alm.modalidad_ingreso,
lugar_procedencia, alm.lugar_nacimiento, alm.edad, alm.correo_institucional, alm.direccion, alm.fecha_nacimiento, alm.correo_personal, alm.celular_estudiante, alm.celular_padre, alm.estado_matricula, alm.
creditos_matriculados, alm.num_semestres_cursados, alm.pps, alm.ppa, alm.tca, alm.convocatoria_id, alm.created_at, alm.updated_at FROM solicitudes sol
JOIN servicio_solicitado ssol ON (sol.id = ssol.solicitud_id)
JOIN alumnos alm ON (sol.alumno_id = alm.id)
WHERE sol.convocatoria_id = ? AND alm.id = ? AND ssol.estado = 'aprobado' limit 1;`
	mdl := Alumno{}
	err := s.DB.Get(&mdl, sqlGetByID, submissionID, studentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return &mdl, err
	}
	return &mdl, nil
}

// getStudentsByResidenceANDBySubmission: ¡Refactorizado para seguridad y eficiencia!
func (s *sqlserver) getStudentsByResidenceANDBySubmission(residenceID string, submissionID int64, page, limit int, filter string) ([]*StudentInformation, error) {
	var ms []*StudentInformation

	// Construcción dinámica de la cláusula WHERE y los argumentos
	var whereClauses []string
	var args []interface{}

	// Condiciones WHERE obligatorias
	whereClauses = append(whereClauses, "r.id = ?")
	args = append(args, residenceID)

	whereClauses = append(whereClauses, "ac.convocatoria_id = ?")
	args = append(args, submissionID)

	whereClauses = append(whereClauses, "ac.estado = 'activo'")

	// Condición de filtro opcional
	if filter != "" {
		// Usamos placeholders para el filtro para prevenir inyección SQL
		whereClauses = append(whereClauses, "(CONCAT(a.nombres, ' ', a.apellido_paterno, ' ', a.apellido_materno) LIKE ? OR a.codigo_estudiante = ? OR a.DNI LIKE ?)")
		args = append(args, "%"+filter+"%", filter, "%"+filter+"%") // Se enlaza el mismo valor de filtro múltiples veces
	}

	// SQL base de la consulta
	sqlGetAll := `
        SELECT
            a.id,
            a.DNI,
            CONCAT(a.nombres, ' ', a.apellido_paterno, ' ', a.apellido_materno) as full_name,
            a.codigo_estudiante as code,
            a.escuela_profesional as professional_school,
            a.facultad as faculty,
            COALESCE(CAST(c.numero AS CHAR), '') as room,
            r.id as residence,
            ac.fecha_asignacion as admission_date
        FROM residencias r
        JOIN cuartos c ON c.residencia_id = r.id
        JOIN asignacion_cuartos ac ON ac.cuarto_id = c.id
        JOIN alumnos a ON ac.alumno_id = a.id`

	// Agrega la cláusula WHERE
	if len(whereClauses) > 0 {
		sqlGetAll += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	// Agrega ORDER BY
	sqlGetAll += "\nORDER BY c.piso, c.numero"

	// Agrega LIMIT y OFFSET para la paginación
	if page != 0 && limit != 0 {
		sqlGetAll += "\nLIMIT ? OFFSET ?"
		args = append(args, limit, (page-1)*limit)
	}

	err := s.DB.Select(&ms, sqlGetAll, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("error executing query in getStudentsByResidenceANDBySubmission: %w", err)
	}
	return ms, nil
}

// getTotalStudentsByResidenceANDBySubmission: Nueva función para el conteo de registros.
func (s *sqlserver) getTotalStudentsByResidenceANDBySubmission(residenceID string, submissionID int64, filter string) (int, error) {
	var total int

	var whereClauses []string
	var args []interface{}

	// Condiciones WHERE obligatorias
	whereClauses = append(whereClauses, "r.id = ?")
	args = append(args, residenceID)

	whereClauses = append(whereClauses, "ac.convocatoria_id = ?")
	args = append(args, submissionID)

	whereClauses = append(whereClauses, "ac.estado = 'activo'")

	// Condición de filtro opcional
	if filter != "" {
		whereClauses = append(whereClauses, "(CONCAT(a.nombres, ' ', a.apellido_paterno, ' ', a.apellido_materno) LIKE ? OR a.codigo_estudiante = ? OR a.DNI LIKE ?)")
		args = append(args, "%"+filter+"%", filter, "%"+filter+"%")
	}

	// SQL base para el conteo
	sqlCount := `
        SELECT
            COUNT(a.id)
        FROM residencias r
        JOIN cuartos c ON c.residencia_id = r.id
        JOIN asignacion_cuartos ac ON ac.cuarto_id = c.id
        JOIN alumnos a ON ac.alumno_id = a.id`

	// Agrega la cláusula WHERE
	if len(whereClauses) > 0 {
		sqlCount += " WHERE " + strings.Join(whereClauses, " AND ")
	}

	err := s.DB.Get(&total, sqlCount, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil // Si no hay filas, el conteo es 0.
		}
		return 0, fmt.Errorf("error executing count query in getTotalStudentsByResidenceANDBySubmission: %w", err)
	}
	return total, nil
}

func (s *sqlserver) getStudentsBySubmission(submissionID int, page, limit int, gender string, statusService string, departmentRequirementID int) ([]*StudentInformationSubmission, error) {
	var ms []*StudentInformationSubmission
	sqlGetAll := `SELECT
    a.id,
    a.DNI,
    CONCAT(a.nombres, ' ', a.apellido_paterno, ' ', a.apellido_materno) as full_name,
    COALESCE(
        dep.name,
        CASE 
            WHEN ds.opcion_seleccion IS NOT NULL THEN ds.opcion_seleccion
            ELSE 'SIN DEPARTAMENTO'
        END
    ) as department,
    a.sexo as sex,
    a.codigo_estudiante as code,
    a.escuela_profesional as professional_school,
    a.facultad as faculty,
    COALESCE(
       CASE 
          WHEN ac.estado = 'activo' THEN CAST(c.numero AS CHAR)
          ELSE ''
       END, 
       ''
    ) as room,
    COALESCE(
       CASE 
          WHEN ac.estado = 'activo' THEN r.id
          ELSE ''
       END, 
       ''
    ) as residence,
    COALESCE(ac.fecha_asignacion, srv_sol.updated_at) as admission_date
FROM solicitudes sol
    JOIN servicio_solicitado srv_sol ON srv_sol.solicitud_id = sol.id
    JOIN alumnos a ON sol.alumno_id = a.id
    LEFT JOIN detalle_solicitudes ds ON ds.solicitud_id = sol.id AND ds.requisito_id = ?
    LEFT JOIN departaments dep ON dep.id = ds.opcion_seleccion
    LEFT JOIN asignacion_cuartos ac ON ac.alumno_id = a.id
    LEFT JOIN cuartos c ON c.id = ac.cuarto_id
    LEFT JOIN residencias r ON r.id = c.residencia_id
WHERE sol.convocatoria_id = ? AND srv_sol.servicio_id = 2`

	if len(statusService) > 1 {
		sqlGetAll += fmt.Sprintf(" AND srv_sol.estado = '%s'", statusService)
	}

	if len(gender) > 1 {
		sqlGetAll += fmt.Sprintf(" AND a.sexo = '%s'", gender[:1])
	}

	sqlGetAll += fmt.Sprintf("\n\tORDER BY c.piso, c.numero")

	if page != 0 || limit != 0 {
		sqlGetAll += fmt.Sprintf("\n\tLIMIT %d OFFSET %d", limit, (page-1)*limit)
	}

	err := s.DB.Select(&ms, sqlGetAll, departmentRequirementID, submissionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return ms, nil
}

func (s *sqlserver) getTotalStudentsBySubmission(submissionID int, gender string, statusService string, departmentRequirementID int) (int, error) {
	var total int

	sqlCount := `SELECT COUNT(a.id)
    FROM solicitudes sol
       JOIN servicio_solicitado srv_sol ON srv_sol.solicitud_id = sol.id
       JOIN alumnos a ON sol.alumno_id = a.id
       LEFT JOIN asignacion_cuartos ac ON ac.alumno_id = a.id
       LEFT JOIN cuartos c ON c.id = ac.cuarto_id
       LEFT JOIN residencias r ON r.id = c.residencia_id
       -- Si usas LEFT JOIN para department, también agrégalo aquí
       LEFT JOIN detalle_solicitudes ds ON ds.solicitud_id = sol.id AND ds.requisito_id = ? 
    WHERE sol.convocatoria_id = ? AND srv_sol.servicio_id = 2`

	args := []interface{}{departmentRequirementID, submissionID} // Parámetros para la consulta principal

	if len(statusService) > 1 {
		sqlCount += " AND srv_sol.estado = ?"
		args = append(args, statusService)
	}

	if len(gender) > 1 {
		sqlCount += " AND a.sexo = ?"
		args = append(args, gender[:1])
	}

	err := s.DB.Get(&total, sqlCount, args...)
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (s *sqlserver) getStudentsBySubmissionExcel(submissionID int) ([]*models.StudentExcel, error) {
	var ms []*models.StudentExcel
	const sqlGetAll = `SELECT
		a.id,
		CONCAT(a.nombres, ' ', a.apellido_paterno, ' ', a.apellido_materno) as full_name,
		a.codigo_estudiante as code,
		a.sexo as sex,
		COALESCE(a.lugar_procedencia, '') as home,
		a.direccion as address,
		a.num_semestres_cursados as number_semesters_completed,
		a.pps as pps,
		a.escuela_profesional as professional_school,
		a.facultad as faculty,
		COALESCE(
			CASE 
				WHEN ac.estado = 'activo' THEN CAST(c.numero AS CHAR)
				ELSE ''
			END, 
			''
		) as room,
		COALESCE(
			CASE 
				WHEN ac.estado = 'activo' THEN r.id
				ELSE ''
			END, 
			''
		) as residence,
		COALESCE(ac.fecha_asignacion, srv_sol.updated_at) as admission_date,
		COALESCE(ac.estado, '') status,
		COALESCE(ac.updated_at, '') as update_date
	FROM solicitudes sol
		JOIN servicio_solicitado srv_sol ON srv_sol.solicitud_id = sol.id
		JOIN alumnos a ON sol.alumno_id = a.id
		LEFT JOIN asignacion_cuartos ac ON ac.alumno_id = a.id
		LEFT JOIN cuartos c ON c.id = ac.cuarto_id
		LEFT JOIN residencias r ON r.id = c.residencia_id
	WHERE sol.convocatoria_id = ? AND srv_sol.servicio_id = 2
	ORDER BY c.piso, c.numero`

	err := s.DB.Select(&ms, sqlGetAll, submissionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return ms, nil
}

func (s *sqlserver) getStudentsByRooms(rooms []string, submissionID int64) ([]*StudentInformation, error) {
	var ms []*StudentInformation

	sqlGetAll := fmt.Sprintf(`
        SELECT 
            a.id,
            a.DNI,
            CONCAT(a.nombres, ' ', a.apellido_paterno, ' ', a.apellido_materno) as full_name,
            a.codigo_estudiante as code,
            a.escuela_profesional as professional_school,
            a.facultad as faculty,
            COALESCE(CAST(c.numero AS CHAR), '') as room,
            '' as residence,
            ac.fecha_asignacion as admission_date
        FROM cuartos c
        JOIN asignacion_cuartos ac ON ac.cuarto_id = c.id
        JOIN alumnos a ON ac.alumno_id = a.id
        WHERE ac.convocatoria_id = ? AND ac.estado = 'activo' AND ac.cuarto_id IN ('%s')
        ORDER BY c.piso, c.numero`, strings.Join(rooms, "','"))

	err := s.DB.Select(&ms, sqlGetAll, submissionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return ms, nil
}
