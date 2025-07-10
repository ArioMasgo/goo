package alumnos

type Estudiante struct {
	ID                   int    `json:"id" db:"id" valid:"-"`
	CodigoEstudiante     string `json:"codigo_estudiante" db:"codigo_estudiante"`
	DNI                  string `json:"dni" db:"DNI"`
	Nombres              string `json:"nombres" db:"nombres"`
	ApellidoPaterno      string `json:"apellido_paterno" db:"apellido_paterno"`
	ApellidoMaterno      string `json:"apellido_materno" db:"apellido_materno"`
	Sexo                 string `json:"sexo" db:"sexo"`
	Facultad             string `json:"facultad" db:"facultad"`
	EscuelaProfesional   string `json:"escuela_profesional" db:"escuela_profesional"`
	UltimoSemestre       string `json:"ultimo_semestre" db:"ultimo_semestre"`
	ModalidadIngreso     string `json:"modalidad_ingreso" db:"modalidad_ingreso"`
	LugarProcedencia     string `json:"lugar_procedencia" db:"lugar_procedencia"`
	LugarNacimiento      string `json:"lugar_nacimiento" db:"lugar_nacimiento"`
	Edad                 int    `json:"edad" db:"edad"`
	CorreoInstitucional  string `json:"correo_institucional" db:"correo_institucional"`
	Direccion            string `json:"direccion" db:"direccion"`
	FechaNacimiento      string `json:"fecha_nacimiento" db:"fecha_nacimiento"`
	CorreoPersonal       string `json:"correo_personal" db:"correo_personal"`
	CelularEstudiante    string `json:"celular_estudiante" db:"celular_estudiante"`
	CelularPadre         string `json:"celular_padre" db:"celular_padre"`
	EstadoMatricula      string `json:"estado_matricula" db:"estado_matricula"`
	CreditosMatriculados string `json:"creditos_matriculados" db:"creditos_matriculados"`
	NumSemestresCursados string `json:"num_semestres_cursados" db:"num_semestres_cursados"`
}

type BasicEstudiante struct {
	ID              int    `json:"id" db:"id" valid:"-"`
	Nombres         string `json:"nombres" db:"nombres"`
	ApellidoPaterno string `json:"apellido_paterno" db:"apellido_paterno"`
	ApellidoMaterno string `json:"apellido_materno" db:"apellido_materno"`
}
