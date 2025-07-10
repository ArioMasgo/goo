package visita_residente

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type VisitaDomiciliaria struct {
	ID                 uint64    `json:"id" db:"id"`
	AlumnoID           uint64    `json:"alumno_id" db:"alumno_id" valid:"required"`
	Estado             string    `json:"estado" db:"estado" valid:"in(pendiente|verificado|observado),required"`
	Comentario         *string   `json:"comentario" db:"comentario"`
	ImagenURL          *string   `json:"imagen_url" db:"imagen_url"`
	IDUsuario          *uint64   `json:"id_usuario" db:"id_usuario"`
	FechaCreacion      time.Time `json:"fecha_creacion" db:"fecha_creacion"`
	FechaActualizacion time.Time `json:"fecha_actualizacion" db:"fecha_actualizacion"`
	AlumnoNombre       *string   `json:"alumno_nombre,omitempty" db:"alumno_nombre"`
	AlumnoCodigo       *string   `json:"alumno_codigo,omitempty" db:"alumno_codigo"`
	EscuelaProfesional *string   `json:"escuela_profesional,omitempty" db:"escuela_profesional"`
	LugarProcedencia   *string   `json:"lugar_procedencia,omitempty" db:"lugar_procedencia"`
	UsuarioNombre      *string   `json:"usuario_nombre,omitempty" db:"usuario_nombre"`
}

func NewVisitaDomiciliaria(alumnoID uint64, estado string, comentario, imagenURL *string, idUsuario *uint64) *VisitaDomiciliaria {
	now := time.Now()
	return &VisitaDomiciliaria{
		AlumnoID:  alumnoID,
		Estado:    estado,
		Comentario: comentario,
		ImagenURL: imagenURL,
		IDUsuario: idUsuario,
		FechaCreacion: now,
		FechaActualizacion: now,
	}
}

func (m *VisitaDomiciliaria) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}

type AlumnoPendienteVisita struct {
	AlumnoID           uint64 `json:"alumno_id" db:"alumno_id"`
	Codigo             string `json:"codigo" db:"codigo"`
	Nombre             string `json:"nombre" db:"nombre"`
	DNI                string `json:"dni" db:"dni"`
	Celular            string `json:"celular" db:"celular"`
	Direccion          string `json:"direccion" db:"direccion"`
	EscuelaProfesional string  `json:"escuela_profesional" db:"escuela_profesional"` 
    LugarProcedencia   string `json:"lugar_procedencia,omitempty" db:"lugar_procedencia"` 
	SolicitudID        uint64 `json:"solicitud_id" db:"solicitud_id"`
	ConvocatoriaID     uint64 `json:"convocatoria_id" db:"convocatoria_id"`
	ConvocatoriaNombre string `json:"convocatoria_nombre" db:"convocatoria_nombre"`
}

type EstadisticasVisita struct {
	TotalVisitas     int64 `json:"total_visitas" db:"total_visitas"`
	Pendientes       int64 `json:"pendientes" db:"pendientes"`
	Verificadas      int64 `json:"verificadas" db:"verificadas"`
	Observadas       int64 `json:"observadas" db:"observadas"`
	VisitasDelMes    int64 `json:"visitas_del_mes" db:"visitas_del_mes"`
	AlumnosSinVisita int64 `json:"alumnos_sin_visita" db:"alumnos_sin_visita"`
}

type EstadisticasPorEscuelaProfesional struct {
	EscuelaProfesional string `json:"escuela_profesional" db:"escuela_profesional"`
	TotalVisitados    	 int64  `json:"total_visitados" db:"total_visitados"`
}

type EstadisticasPorLugarProcedencia struct {
	Departamento   string `json:"departamento" db:"departamento"`
	TotalVisitados int64  `json:"total_visitados" db:"total_visitados"`
}

type AlumnoPendienteVisitaPorDepartamento struct {
	AlumnoID           uint64 `json:"alumno_id" db:"alumno_id"`
	Codigo             string `json:"codigo" db:"codigo"`
	Nombre             string `json:"nombre" db:"nombre"`
	EscuelaProfesional string `json:"escuela_profesional" db:"escuela_profesional"`
	Departamento       string `json:"departamento" db:"departamento"`
	Provincia          string `json:"provincia" db:"provincia"`
	Distrito           string `json:"distrito" db:"distrito"`
	Direccion          string `json:"direccion" db:"direccion"`
	Celular            string `json:"celular" db:"celular"`
	CelularPadre       string `json:"celular_padre" db:"celular_padre"`
	ConvocatoriaNombre string `json:"convocatoria_nombre" db:"convocatoria_nombre"`
}


type FiltrosVisita struct {
	Estado         *string    `json:"estado,omitempty"`
	AlumnoID       *uint64    `json:"alumno_id,omitempty"`
	IDUsuario      *uint64    `json:"id_usuario,omitempty"`
	FechaInicio    *time.Time `json:"fecha_inicio,omitempty"`
	FechaFin       *time.Time `json:"fecha_fin,omitempty"`
	ConvocatoriaID *uint64    `json:"convocatoria_id,omitempty"`
	Limit          int        `json:"limit,omitempty"`
	Offset         int        `json:"offset,omitempty"`
}

type AlumnoCompleto struct {
	AlumnoID              uint64  `json:"alumno_id" db:"alumno_id"`
	Codigo                string  `json:"codigo" db:"codigo"`
	Nombre                string  `json:"nombre" db:"nombre"`
	DNI                   string  `json:"dni" db:"dni"`
	Celular               string  `json:"celular" db:"celular"`
	CelularPadre          string  `json:"celular_padre" db:"celular_padre"`
	Direccion             string  `json:"direccion" db:"direccion"`
	LugarProcedencia      string  `json:"lugar_procedencia" db:"lugar_procedencia"`
	EscuelaProfesional    string  `json:"escuela_profesional" db:"escuela_profesional"`
	Departamento          string  `json:"departamento" db:"departamento"`
	Provincia             string  `json:"provincia" db:"provincia"`
	Distrito              string  `json:"distrito" db:"distrito"`
	SolicitudID           uint64  `json:"solicitud_id" db:"solicitud_id"`
	ConvocatoriaID        uint64  `json:"convocatoria_id" db:"convocatoria_id"`
	ConvocatoriaNombre    string  `json:"convocatoria_nombre" db:"convocatoria_nombre"`
	EstadoVisita          string  `json:"estado_visita" db:"estado_visita"`
}