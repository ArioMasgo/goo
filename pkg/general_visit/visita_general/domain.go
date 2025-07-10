package visita_general

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type VisitaGeneral struct {
	ID                string     `db:"id" json:"id" valid:"uuid,required"`
	TipoUsuario       string     `db:"tipo_usuario" json:"tipo_usuario" valid:"in(alumno|docente|administrativo),required"`
	CodigoEstudiante  *string    `db:"codigo_estudiante" json:"codigo_estudiante"`
	DNI               *string    `db:"dni" json:"dni"`
	NombreCompleto    string     `db:"nombre_completo" json:"nombre_completo" valid:"required"`
	Genero            *string    `db:"genero" json:"genero" valid:"in(M|F)"`
	Edad              *int       `db:"edad" json:"edad"`
	Escuela           *string    `db:"escuela" json:"escuela"`
	Area              *string    `db:"area" json:"area"`
	MotivoAtencion    string     `db:"motivo_atencion" json:"motivo_atencion" valid:"required"`
	DescripcionMotivo string     `db:"descripcion_motivo" json:"descripcion_motivo" valid:"required"`
	URLImagen         *string    `db:"url_imagen" json:"url_imagen"`
	Departamento      *string    `db:"departamento" json:"departamento"`
	Provincia         *string    `db:"provincia" json:"provincia"`
	Distrito          *string    `db:"distrito" json:"distrito"`
	LugarAtencion     string     `db:"lugar_atencion" json:"lugar_atencion" valid:"required"`
	CreatedBy         *int64     `db:"created_by" json:"created_by"`
	CreatedAt         *time.Time `db:"created_at" json:"created_at"`
	UpdatedBy         *int64     `db:"updated_by" json:"updated_by"`
	UpdatedAt         *time.Time `db:"updated_at" json:"updated_at"`
}

func NewVisitaGeneral(id string, tipoUsuario string, codigoEstudiante *string, dni *string, nombreCompleto string,
	genero *string, edad *int, escuela *string, area *string, motivoAtencion string,
	descripcionMotivo string, urlImagen *string, departamento *string, provincia *string, distrito *string,
	lugarAtencion string, createdBy *int64) *VisitaGeneral {
	now := time.Now()
	return &VisitaGeneral{
		ID:                id,
		TipoUsuario:       tipoUsuario,
		CodigoEstudiante:  codigoEstudiante,
		DNI:               dni,
		NombreCompleto:    nombreCompleto,
		Genero:            genero,
		Edad:              edad,
		Escuela:           escuela,
		Area:              area,
		MotivoAtencion:    motivoAtencion,
		DescripcionMotivo: descripcionMotivo,
		URLImagen:         urlImagen,
		Departamento:      departamento,
		Provincia:         provincia,
		Distrito:          distrito,
		LugarAtencion:     lugarAtencion,
		CreatedAt:         &now,
		UpdatedAt:         &now,
		CreatedBy:         createdBy,
		UpdatedBy:         createdBy,
	}
}

func (m *VisitaGeneral) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}

type Departamento struct {
	ID        string     `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at"`
}

type Provincia struct {
	ID           string     `db:"id" json:"id"`
	Name         string     `db:"name" json:"name"`
	DepartmentID string     `db:"departament_id" json:"department_id"`
	CreatedAt    *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at"`
}

type Distrito struct {
	ID         string     `db:"id" json:"id"`
	Name       string     `db:"name" json:"name"`
	ProvinceID string     `db:"province_id" json:"province_id"`
	CreatedAt  *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at" json:"updated_at"`
}

type LocationResponse struct {
	Departamentos []DepartamentoWithProvincias `json:"departamentos"`
}

type DepartamentoWithProvincias struct {
	ID         string                   `json:"id"`
	Name       string                   `json:"name"`
	Provincias []ProvinciaWithDistritos `json:"provincias"`
}

type ProvinciaWithDistritos struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Distritos []Distrito `json:"distritos"`
}
