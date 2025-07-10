package participantes

type Participante struct {
	ID                 int     `json:"id" db:"id_participante"`
	Tipo               string  `json:"tipo_participante" db:"tipo_participante"`
	Nombre             string  `json:"nombre" db:"nombre"`
	Apellido           string  `json:"apellido" db:"apellido"`
	DNI                string  `json:"dni,omitempty" db:"dni"`
	Estado             string  `json:"estado" db:"estado"`
	CreatedAt          string  `json:"created_at" db:"created_at"`
	UpdatedAt          string  `json:"updated_at" db:"updated_at"`
	ColegioProcedencia string  `json:"colegio_procedencia,omitempty" db:"colegio_procedencia"`
	AnioIngreso        int     `json:"anio_ingreso,omitempty" db:"anio_ingreso"`
	Escuela            string  `json:"escuela,omitempty" db:"escuela"`
	CodigoEstudiante   string  `json:"codigo_estudiante,omitempty" db:"codigo_estudiante"`
	FechaNacimiento    string  `json:"fecha_nacimiento,omitempty" db:"fecha_nacimiento"`
	Edad               int     `json:"edad,omitempty" db:"edad"`
	LugarNacimiento    string  `json:"lugar_nacimiento,omitempty" db:"lugar_nacimiento"`
	ModalidadIngreso   string  `json:"modalidad_ingreso,omitempty" db:"modalidad_ingreso"`
	NumeroAtencion     int     `json:"numero_atencion,omitempty" db:"numero_atencion"`
	Sexo               string  `json:"sexo,omitempty" db:"sexo"`
	NumTelefono        string  `json:"num_telefono,omitempty" db:"num_telefono"`
	EstadoEvaluacion   string  `json:"estado_evaluacion,omitempty" db:"estado_evaluacion"`
	Diagnostico        string  `json:"diagnostico,omitempty" db:"diagnostico"`
	ConQuienesVive     string  `json:"con_quienes_vive_actualmente,omitempty" db:"con_quienes_vive_actualmente"`
	SemestreCursa      string  `json:"semestre_cursa,omitempty" db:"semestre_cursa"`
	Direccion          string  `json:"direccion,omitempty" db:"direccion"`
	Profesion          string  `json:"profesion" db:"profesion"`
	EstadoCivil        *string `json:"estado_civil" db:"estado_civil"`
	LaboraEnUnas       bool    `json:"labora_en_unas" db:"labora_en_unas"`
	GradoInstruccion   *string `json:"grado_instruccion" db:"grado_instruccion"`
}

type Answer struct {
	ParticipanteID int    `db:"id_participante" json:"studentId"`
	EncuestaID     int    `json:"id_encuesta" db:"id_encuesta"`
	QuestionID     int    `db:"id_pregunta" json:"questionID"`
	Answer         string `db:"respuesta" json:"answer"`
}
