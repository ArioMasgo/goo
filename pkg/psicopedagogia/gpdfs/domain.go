package gpdfs

type DataPDFRespuesta struct {
	TextoPregunta string `json:"texto_pregunta" db:"texto_pregunta"`
	Respuesta     string `json:"respuesta" db:"respuesta"`
}

type ParticipantePDF struct {
	ID                        int     `json:"id" db:"id_participante"`
	Tipo                      string  `json:"tipo_participante" db:"tipo_participante"`
	Nombre                    string  `json:"nombre" db:"nombre"`
	Apellido                  string  `json:"apellido" db:"apellido"`
	Correo                    string  `json:"correo" db:"correo"`
	DNI                       *string `json:"dni,omitempty" db:"dni"`
	Estado                    string  `json:"estado" db:"estado"`
	Score                     float64 `json:"score" db:"score"`
	Diagnostico               string  `json:"diagnostico" db:"diagnostico"`
	Notes                     string  `json:"notes" db:"notes"`
	CreatedAt                 string  `json:"created_at" db:"created_at"`
	UpdatedAt                 string  `json:"updated_at" db:"updated_at"`
	NumTelefono               string  `json:"num_telefono" db:"num_telefono"`
	ColegioProcedencia        string  `json:"colegio_procedencia" db:"colegio_procedencia"`
	AnioIngreso               int     `json:"anio_ingreso" db:"anio_ingreso"`
	ConQuienesViveActualmente string  `json:"con_quienes_vive_actualmente" db:"con_quienes_vive_actualmente"`
	EstadoEvaluacion          string  `json:"estado_evaluacion" db:"estado_evaluacion"`
	SemestreCursa             string  `json:"semestre_cursa" db:"semestre_cursa"`
	Escuela                   string  `json:"escuela" db:"escuela"`
	CodigoEstudiante          string  `json:"codigo_estudiante" db:"codigo_estudiante"`
	FechaNacimiento           string  `json:"fecha_nacimiento" db:"fecha_nacimiento"`
	Edad                      int     `json:"edad" db:"edad"`
	LugarNacimiento           string  `json:"lugar_nacimiento" db:"lugar_nacimiento"`
	ModalidadIngreso          string  `json:"modalidad_ingreso" db:"modalidad_ingreso"`
	NumeroAtencion            int     `json:"numero_atencion" db:"numero_atencion"`
	Sexo                      string  `json:"sexo,omitempty" db:"sexo"`
	Direccion                 string  `json:"direccion,omitempty" db:"direccion"`
	Profesion                 string  `json:"profesion" db:"profesion"`
	EstadoCivil               *string `json:"estado_civil" db:"estado_civil"`
	LaboraEnUnas              bool    `json:"labora_en_unas" db:"labora_en_unas"`
	GradoInstruccion          *string `json:"grado_instruccion" db:"grado_instruccion"`
}
