package models

import (
	"github.com/asaskevich/govalidator"
)

type Patients struct {
	ID                 string                `json:"id" valid:"uuid,required"`
	CodigoSGA          string                `json:"codigo_sga" valid:"-"`
	DNI                string                `json:"dni" valid:"required"`
	Nombres            string                `json:"nombres" valid:"required"`
	Apellidos          string                `json:"apellidos" valid:"required"`
	Sexo               string                `json:"sexo" valid:"-"`
	Edad               string                `json:"edad" valid:"-"`
	EstadoCivil        string                `json:"estado_civil" valid:""`
	GrupoSanguineo     string                `json:"grupo_sanguineo" valid:"-"`
	FechaNacimiento    string                `json:"fecha_nacimiento" valid:"-"`
	LugarNacimiento    string                `json:"lugar_nacimiento" valid:"-"`
	Procedencia        string                `json:"procedencia" valid:"-"`
	EscuelaProfesional string                `json:"escuela_profesional" valid:"-"`
	Ocupacion          string                `json:"ocupacion" valid:"-"`
	CorreoElectronico  string                `json:"correo_electronico" valid:"-"`
	NumeroCelular      string                `json:"numero_celular" valid:"required"`
	Direccion          string                `json:"direccion" valid:"-"`
	TipoPersona        string                `json:"tipo_persona" valid:"required"`
	FactorRH           string                `json:"factor_rh" valid:"-"`
	Alergias           string                `json:"alergias" valid:"-"`
	RAM                bool                  `json:"ram" valid:"-"`
	Antecedentes       []*RequestAntecedents `json:"antecedentes" valid:"-"`
}

type ResponsePatientInfo struct {
	CodigoSGA   string `json:"codigo_sga" valid:"-"`
	DNI         string `json:"dni" valid:"required"`
	Nombres     string `json:"nombres" valid:"required"`
	Apellidos   string `json:"apellidos" valid:"required"`
	TipoPersona string `json:"tipo_persona" valid:"required"`
}

type RequestAntecedents struct {
	ID     string `json:"id" valid:"uuid"`
	Nombre string `json:"nombre_antecedente" valid:"-"`
	Estado string `json:"estado_antecedente" valid:"-"`
}

func (m *Patients) ValidPatients() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
