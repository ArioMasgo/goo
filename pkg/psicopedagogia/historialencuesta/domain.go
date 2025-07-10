package historialencuesta

import "time"

type HistorialEncuesta struct {
	ID                   int       `json:"id_historial" db:"id_historial"`
	IDParticipante       int       `json:"id_participante" db:"id_participante"`
	IDEncuesta           int       `json:"id_encuesta" db:"id_encuesta"`
	FechaRespuesta       time.Time `json:"fecha_respuesta" db:"fecha_respuesta"`
	Calificacion         float64   `json:"calificacion,omitempty" db:"calificacion"`
	Diagnostico          string    `json:"diagnostico" db:"diagnostico"`
	Notes                string    `json:"notes,omitempty" db:"notes"`
	NumTelefono          string    `json:"num_telefono,omitempty" db:"num_telefono"`
	ConQuienesVive       string    `json:"con_quienes_vive_actualmente,omitempty" db:"con_quienes_vive_actualmente"`
	EstadoEvaluacion     string    `json:"estado_evaluacion" db:"estado_evaluacion"`
	SemestreCursa        string    `json:"semestre_cursa,omitempty" db:"semestre_cursa"`
	Direccion            string    `json:"direccion,omitempty" db:"direccion"`
	QuienFinanciaCarrera string    `json:"quien_financia_carrera,omitempty" db:"quien_financia_carrera"`
	MotivoConsulta       string    `json:"motivo_consulta,omitempty" db:"motivo_consulta"`
	SituacionActual      string    `json:"situacion_actual,omitempty" db:"situacion_actual"`
	OtrosProcedimientos  string    `json:"otros_procedimientos,omitempty" db:"otros_procedimientos"`
	DiagnosticoID        int       `json:"diagnostico_id,omitempty" db:"diagnostico_id"`
	CreatedDate          string    `json:"created_date" db:"created_date"`
	EsSRQ                bool      `json:"es_srq" db:"es_srq"`
	KeyUrl               string    `json:"key_url" db:"key_url"`
}

type Historial struct {
	IDHistorial            int       `json:"id_historial" db:"id_historial"`
	IDParticipante         int       `json:"id_participante" db:"id_participante"`
	IDEncuesta             *int      `json:"id_encuesta" db:"id_encuesta"`
	NombreEncuesta         string    `json:"nombre_encuesta" db:"nombre_encuesta"`
	FechaRespuesta         time.Time `json:"fecha_respuesta" db:"fecha_respuesta"`
	Diagnostico            string    `json:"diagnostico" db:"diagnostico"`
	Notes                  string    `json:"notes" db:"notes"`
	NumTelefono            string    `json:"num_telefono" db:"num_telefono"`
	ConQuienesVive         string    `json:"con_quienes_vive_actualmente" db:"con_quienes_vive_actualmente"`
	EstadoEvaluacion       string    `json:"estado_evaluacion" db:"estado_evaluacion"`
	SemestreCursa          string    `json:"semestre_cursa" db:"semestre_cursa"`
	Direccion              string    `json:"direccion" db:"direccion"`
	QuienFinanciaCarrera   string    `json:"quien_financia_carrera" db:"quien_financia_carrera"`
	MotivoConsulta         string    `json:"motivo_consulta" db:"motivo_consulta"`
	SituacionActual        string    `json:"situacion_actual" db:"situacion_actual"`
	OtrosProcedimientos    string    `json:"otros_procedimientos" db:"otros_procedimientos"`
	CreatedDate            time.Time `json:"created_date" db:"created_date"`
	EsSRQ                  bool      `json:"es_srq" db:"es_srq"`
	DiagnosticoID          int       `json:"diagnostico_id" db:"diagnostico_id"`
	DNI                    string    `json:"dni" db:"dni"`
	NumeroAtencion         int       `json:"numero_atencion" db:"numero_atencion"`
	TipoParticipante       string    `json:"tipo_participante" db:"tipo_participante"`
	NombreEstudiante       string    `json:"nombre_estudiante" db:"nombre"`
	ApellidoEstudiante     string    `json:"apellido_estudiante" db:"apellido"`
	Escuela                string    `json:"escuela" db:"escuela"`
	KeyUrl                 string    `json:"key_url" db:"key_url"`
	NotasAtencion          *string   `json:"notas_atencion" db:"notas_atencion"`
	InstrumentosUtilizados *string   `json:"instrumentos_utilizados" db:"instrumentos_utilizados"`
	ResultadosObtenidos    *string   `json:"resultados_obtenidos" db:"resultados_obtenidos"`
}
