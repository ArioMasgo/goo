package patients

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Patients struct {
	ID                 string     `json:"id" db:"id" valid:"uuid,required"`
	CodigoSGA          string     `json:"codigo_sga" db:"codigo_sga" valid:"-"`
	DNI                string     `json:"dni" db:"dni" valid:"required"`
	Nombres            string     `json:"nombres" db:"nombres" valid:"required"`
	Apellidos          string     `json:"apellidos" db:"apellidos" valid:"required"`
	Sexo               string     `json:"sexo" db:"sexo" valid:"-"`
	Edad               string     `json:"edad" db:"edad" valid:"-"`
	EstadoCivil        string     `json:"estado_civil" db:"estado_civil" valid:"-"`
	GrupoSanguineo     string     `json:"grupo_sanguineo" db:"grupo_sanguineo" valid:"-"`
	FechaNacimiento    string     `json:"fecha_nacimiento" db:"fecha_nacimiento" valid:"-"`
	LugarNacimiento    string     `json:"lugar_nacimiento" db:"lugar_nacimiento" valid:"-"`
	Procedencia        string     `json:"procedencia" db:"procedencia" valid:"-"`
	EscuelaProfesional string     `json:"escuela_profesional" db:"escuela_profesional" valid:"-"`
	Ocupacion          string     `json:"ocupacion" db:"ocupacion" valid:"-"`
	CorreoElectronico  string     `json:"correo_electronico" db:"correo_electronico" valid:"-"`
	NumeroCelular      string     `json:"numero_celular" db:"numero_celular" valid:"required"`
	Direccion          string     `json:"direccion" db:"direccion" valid:"-"`
	TipoPersona        string     `json:"tipo_persona" db:"tipo_persona" valid:"required"`
	FactorRH           string     `json:"factor_rh" db:"factor_rh" valid:"-"`
	Alergias           string     `json:"alergias" db:"alergias" valid:"-"`
	RAM                bool       `json:"ram" db:"ram" valid:"-"`
	IsDeleted          bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted        *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt          *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator        string     `json:"user_creator" db:"user_creator"`
	CreatedAt          time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at" db:"updated_at"`
}

func NewPatients(id string, codigo_sga string, dni string, nombres string, apellidos string, sexo string, edad string, estado_civil string, grupo_sanguineo string, fecha_nacimiento string, lugar_nacimiento string, procedencia string, escuela_profesional string, ocupacion string, correo_electronico string, numero_celular string, direccion string, tipo_persona string, factor_rh string, alergias string, ram bool) *Patients {
	now := time.Now()
	return &Patients{
		ID:                 id,
		CodigoSGA:          codigo_sga,
		DNI:                dni,
		Nombres:            nombres,
		Apellidos:          apellidos,
		Sexo:               sexo,
		Edad:               edad,
		EstadoCivil:        estado_civil,
		GrupoSanguineo:     grupo_sanguineo,
		FechaNacimiento:    fecha_nacimiento,
		LugarNacimiento:    lugar_nacimiento,
		Procedencia:        procedencia,
		EscuelaProfesional: escuela_profesional,
		Ocupacion:          ocupacion,
		CorreoElectronico:  correo_electronico,
		NumeroCelular:      numero_celular,
		Direccion:          direccion,
		TipoPersona:        tipo_persona,
		FactorRH:           factor_rh,
		Alergias:           alergias,
		RAM:                ram,
		IsDeleted:          false,
		CreatedAt:          now,
		UpdatedAt:          now,
	}
}

func (m *Patients) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
