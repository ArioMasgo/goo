package models

import "time"

type VisitaGeneral struct {
	ID                string  `json:"id"`
	TipoUsuario       string  `json:"tipo_usuario"`
	CodigoEstudiante  *string `json:"codigo_estudiante"`
	DNI               *string `json:"dni"`
	NombreCompleto    string  `json:"nombre_completo"`
	Genero            *string `json:"genero"`
	Edad              *int    `json:"edad"`
	Escuela           *string `json:"escuela"`
	Area              *string `json:"area"`
	MotivoAtencion    string  `json:"motivo_atencion"`
	DescripcionMotivo string  `json:"descripcion_motivo"`
	URLImagen         *string `json:"url_imagen"`
	Departamento      *string `json:"departamento"`
	Provincia         *string `json:"provincia"`
	Distrito          *string `json:"distrito"`
	LugarAtencion     string  `json:"lugar_atencion"`
}

// Modelos para dropdowns de ubicación
type Departamento struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Provincia struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	DepartmentID string     `json:"departament_id"`
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}

type Distrito struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	ProvinceID string     `json:"province_id"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
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
