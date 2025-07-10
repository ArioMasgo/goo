package historialencuesta

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

type HistorialRepository struct {
	DB *sqlx.DB
}

func NewHistorialRepository(db *sqlx.DB) *HistorialRepository {
	return &HistorialRepository{DB: db}
}

func (r *HistorialRepository) Create(historial *HistorialEncuesta) (int, error) {
	var query string
	var args map[string]interface{}

	if historial.IDEncuesta == 0 && historial.DiagnosticoID != 0 {
		query = `
			INSERT INTO historial 
			(id_participante, fecha_respuesta, diagnostico, notes, 
			num_telefono, con_quienes_vive_actualmente, estado_evaluacion, semestre_cursa, 
			direccion, quien_financia_carrera, motivo_consulta, situacion_actual, 
			otros_procedimientos, created_date, es_srq, diagnostico_id, key_url) 
			VALUES (:id_participante, :fecha_respuesta, :diagnostico, :notes, 
					:num_telefono, :con_quienes_vive_actualmente, :estado_evaluacion, :semestre_cursa, 
					:direccion, :quien_financia_carrera, :motivo_consulta, :situacion_actual, 
					:otros_procedimientos, :created_date, :es_srq, :diagnostico_id, :key_url)`

		args = map[string]interface{}{
			"id_participante":              historial.IDParticipante,
			"fecha_respuesta":              historial.FechaRespuesta,
			"diagnostico":                  historial.Diagnostico,
			"notes":                        historial.Notes,
			"num_telefono":                 historial.NumTelefono,
			"con_quienes_vive_actualmente": historial.ConQuienesVive,
			"estado_evaluacion":            historial.EstadoEvaluacion,
			"semestre_cursa":               historial.SemestreCursa,
			"direccion":                    historial.Direccion,
			"quien_financia_carrera":       historial.QuienFinanciaCarrera,
			"motivo_consulta":              historial.MotivoConsulta,
			"situacion_actual":             historial.SituacionActual,
			"otros_procedimientos":         historial.OtrosProcedimientos,
			"created_date":                 historial.CreatedDate,
			"es_srq":                       historial.EsSRQ,
			"diagnostico_id":               historial.DiagnosticoID,
			"key_url":                      historial.KeyUrl,
		}
	} else if historial.DiagnosticoID == 0 && historial.IDEncuesta != 0 {
		query = `
			INSERT INTO historial 
			(id_participante, id_encuesta, fecha_respuesta, diagnostico, notes, 
			num_telefono, con_quienes_vive_actualmente, estado_evaluacion, semestre_cursa, 
			direccion, quien_financia_carrera, motivo_consulta, situacion_actual, 
			otros_procedimientos, created_date, es_srq, key_url) 
			VALUES (:id_participante, :id_encuesta, :fecha_respuesta, :diagnostico, :notes, 
					:num_telefono, :con_quienes_vive_actualmente, :estado_evaluacion, :semestre_cursa, 
					:direccion, :quien_financia_carrera, :motivo_consulta, :situacion_actual, 
					:otros_procedimientos, :created_date, :es_srq, :key_url)`

		args = map[string]interface{}{
			"id_participante":              historial.IDParticipante,
			"id_encuesta":                  historial.IDEncuesta,
			"fecha_respuesta":              historial.FechaRespuesta,
			"diagnostico":                  historial.Diagnostico,
			"notes":                        historial.Notes,
			"num_telefono":                 historial.NumTelefono,
			"con_quienes_vive_actualmente": historial.ConQuienesVive,
			"estado_evaluacion":            historial.EstadoEvaluacion,
			"semestre_cursa":               historial.SemestreCursa,
			"direccion":                    historial.Direccion,
			"quien_financia_carrera":       historial.QuienFinanciaCarrera,
			"motivo_consulta":              historial.MotivoConsulta,
			"situacion_actual":             historial.SituacionActual,
			"otros_procedimientos":         historial.OtrosProcedimientos,
			"created_date":                 historial.CreatedDate,
			"es_srq":                       historial.EsSRQ,
			"key_url":                      historial.KeyUrl,
		}
	} else {
		query = `
			INSERT INTO historial 
			(id_participante, fecha_respuesta, diagnostico, notes, 
			num_telefono, con_quienes_vive_actualmente, estado_evaluacion, semestre_cursa, 
			direccion, quien_financia_carrera, motivo_consulta, situacion_actual, 
			otros_procedimientos, created_date, es_srq, key_url) 
			VALUES (:id_participante, :fecha_respuesta, :diagnostico, :notes, 
					:num_telefono, :con_quienes_vive_actualmente, :estado_evaluacion, :semestre_cursa, 
					:direccion, :quien_financia_carrera, :motivo_consulta, :situacion_actual, 
					:otros_procedimientos, :created_date, :es_srq, :key_url)`

		args = map[string]interface{}{
			"id_participante":              historial.IDParticipante,
			"fecha_respuesta":              historial.FechaRespuesta,
			"diagnostico":                  historial.Diagnostico,
			"notes":                        historial.Notes,
			"num_telefono":                 historial.NumTelefono,
			"con_quienes_vive_actualmente": historial.ConQuienesVive,
			"estado_evaluacion":            historial.EstadoEvaluacion,
			"semestre_cursa":               historial.SemestreCursa,
			"direccion":                    historial.Direccion,
			"quien_financia_carrera":       historial.QuienFinanciaCarrera,
			"motivo_consulta":              historial.MotivoConsulta,
			"situacion_actual":             historial.SituacionActual,
			"otros_procedimientos":         historial.OtrosProcedimientos,
			"created_date":                 historial.CreatedDate,
			"es_srq":                       historial.EsSRQ,
			"key_url":                      historial.KeyUrl,
		}
	}

	stmt, err := r.DB.PrepareNamed(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(args)
	if err != nil {
		return 0, err
	}

	// Obtener el ID insertado
	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastID), nil
}

func (r *HistorialRepository) GetHistorialFiltered(dni, apellido, fechaInicio, fechaFin string) ([]Historial, error) {
	var historiales []Historial
	var filters []string
	var args []interface{}

	query := `
		SELECT 
			h.id_historial,
			h.id_participante,
			h.id_encuesta,
			p.nombre,
			p.apellido,
			h.fecha_respuesta,
			h.es_srq,
			h.fecha_respuesta,
			h.diagnostico,
			h.num_telefono,
			h.con_quienes_vive_actualmente,
			h.estado_evaluacion,
			h.semestre_cursa,
			h.direccion,
			h.quien_financia_carrera,
			h.motivo_consulta,
			h.situacion_actual,
			h.otros_procedimientos,
			h.notas_atencion,
			h.instrumentos_utilizados,
			h.resultados_obtenidos
		FROM historial h
		JOIN participantes p ON h.id_participante = p.id_participante
	`

	if dni != "" {
		filters = append(filters, "p.dni = ?")
		args = append(args, dni)
	}
	if apellido != "" {
		filters = append(filters, "p.apellido LIKE ?")
		args = append(args, "%"+apellido+"%")
	}
	if fechaInicio != "" && fechaFin != "" {
		filters = append(filters, "h.fecha_respuesta BETWEEN ? AND ?")
		args = append(args, fechaInicio, fechaFin)
	} else if fechaInicio != "" {
		filters = append(filters, "h.fecha_respuesta >= ?")
		args = append(args, fechaInicio)
	} else if fechaFin != "" {
		filters = append(filters, "h.fecha_respuesta <= ?")
		args = append(args, fechaFin)
	}

	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}

	err := r.DB.Select(&historiales, query, args...)
	return historiales, err
}

func (r *HistorialRepository) HasStudentResponded(dni string, encuestaID int) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM historial h
			  INNER JOIN participantes p ON h.id_participante = p.id_participante
			  WHERE p.dni = ? AND h.id_encuesta = ?`

	err := r.DB.Get(&count, query, dni, encuestaID)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *HistorialRepository) GetLatestHistorial(limit, offset int) ([]Historial, error) {
	var historiales []Historial

	query := `
		SELECT h.*,
			   e.dni, 
			   e.numero_atencions, 
			   e.tipo_paciente, 
			   e.nombre, 
			   e.apellido, 
			   e.escuela
		FROM historial h
		INNER JOIN (
			SELECT id_participante, MAX(id_historial) AS max_id
			FROM historial
			GROUP BY id_participante
		) latest ON h.id_participante = latest.id_participante 
		AND h.id_historial = latest.max_id
		INNER JOIN estudiantes e ON h.id_participante = e.id_participante
		LIMIT ? OFFSET ?
	`

	err := r.DB.Select(&historiales, query, limit, offset)
	return historiales, err
}

func (r *HistorialRepository) KeyUrlExists(keyUrl string) (bool, error) {
	var exists bool
	query := `SELECT COUNT(*) > 0 FROM historial WHERE key_url = ?`

	err := r.DB.Get(&exists, query, keyUrl)
	if err != nil {
		return false, err
	}

	return exists, nil
}
