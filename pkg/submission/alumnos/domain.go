package alumnos

import (
	"time"
)

type Alumno struct {
	ID                   int64      `json:"id" db:"id" valid:"-"`
	CodigoEstudiante     string     `json:"codigo_estudiante" db:"codigo_estudiante" valid:"required"`
	DNI                  string     `json:"DNI" db:"DNI" valid:"required"`
	Nombres              string     `json:"nombres" db:"nombres" valid:"required"`
	ApellidoPaterno      string     `json:"apellido_paterno" db:"apellido_paterno" valid:"required"`
	ApellidoMaterno      string     `json:"apellido_materno" db:"apellido_materno" valid:"required"`
	Sexo                 string     `json:"sexo" db:"sexo" valid:"required,in(masculino|femenino)"`
	Facultad             string     `json:"facultad" db:"facultad" valid:"required"`
	EscuelaProfesional   string     `json:"escuela_profesional" db:"escuela_profesional" valid:"required"`
	UltimoSemestre       string     `json:"ultimo_semestre" db:"ultimo_semestre" valid:"required"`
	ModalidadIngreso     string     `json:"modalidad_ingreso" db:"modalidad_ingreso" valid:"required"`
	LugarProcedencia     *string    `json:"lugar_procedencia" db:"lugar_procedencia" valid:"optional"`
	LugarNacimiento      *string    `json:"lugar_nacimiento" db:"lugar_nacimiento" valid:"optional"`
	Edad                 int32      `json:"edad" db:"edad" valid:"required"`
	CorreoInstitucional  string     `json:"correo_institucional" db:"correo_institucional" valid:"required,email"`
	Direccion            string     `json:"direccion" db:"direccion" valid:"required"`
	FechaNacimiento      time.Time  `json:"fecha_nacimiento" db:"fecha_nacimiento" valid:"required"`
	CorreoPersonal       string     `json:"correo_personal" db:"correo_personal" valid:"required,email"`
	CelularEstudiante    string     `json:"celular_estudiante" db:"celular_estudiante" valid:"required"`
	CelularPadre         string     `json:"celular_padre" db:"celular_padre" valid:"required"`
	EstadoMatricula      string     `json:"estado_matricula" db:"estado_matricula" valid:"required"`
	CreditosMatriculados string     `json:"creditos_matriculados" db:"creditos_matriculados" valid:"required"`
	NumSemestresCursados string     `json:"num_semestres_cursados" db:"num_semestres_cursados" valid:"required"`
	PPS                  string     `json:"pps" db:"pps" valid:"required"`
	PPA                  string     `json:"ppa" db:"ppa" valid:"required"`
	TCA                  string     `json:"tca" db:"tca" valid:"required"`
	ConvocatoriaID       int64      `json:"convocatoria_id" db:"convocatoria_id" valid:"required"`
	CreatedAt            *time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt            *time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type StudentInformation struct {
	ID                 int64  `json:"id" db:"id" valid:"uuid,required"`
	DNI                string `json:"DNI" db:"DNI"`
	FullName           string `json:"full_name" db:"full_name"`
	Code               string `json:"code" db:"code"`
	ProfessionalSchool string `json:"professional_school" db:"professional_school"`
	Faculty            string `json:"faculty" db:"faculty"`
	Room               string `json:"room" db:"room"`
	Residence          string `json:"residence,omitempty" db:"residence"`
	AdmissionDate      string `json:"admission_date" db:"admission_date"`
}

type StudentInformationSubmission struct {
	ID                 int64  `json:"id" db:"id" valid:"uuid,required"`
	DNI                string `json:"DNI" db:"DNI"`
	FullName           string `json:"full_name" db:"full_name"`
	Department         string `json:"department" db:"department"`
	Sex                string `json:"sex" db:"sex"`
	Code               string `json:"code" db:"code"`
	ProfessionalSchool string `json:"professional_school" db:"professional_school"`
	Faculty            string `json:"faculty" db:"faculty"`
	Room               string `json:"room" db:"room"`
	Residence          string `json:"residence,omitempty" db:"residence"`
	AdmissionDate      string `json:"admission_date" db:"admission_date"`
}
