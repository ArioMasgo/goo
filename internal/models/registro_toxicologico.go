package models

import (
	"time"

	"github.com/asaskevich/govalidator"
)

// RegistroToxicologico Model struct para registro_toxicologico
type RegistroToxicologico struct {
	ID                 int64     `db:"id" json:"id"`
	AlumnoID           int64     `db:"alumno_id" json:"alumno_id" valid:"required"`
	ConvocatoriaID     int64     `db:"convocatoria_id" json:"convocatoria_id" valid:"required"`
	Estado             string    `db:"estado" json:"estado" valid:"required,in(verificado|observado|pendiente)"`
	Comentario         *string   `db:"comentario" json:"comentario"`
	IDUsuario          *int64    `db:"id_usuario" json:"id_usuario"`
	FechaCreacion      time.Time `db:"fecha_creacion" json:"fecha_creacion"`
	FechaActualizacion time.Time `db:"fecha_actualizacion" json:"fecha_actualizacion"`
}

// RegistroToxicologicoRequest struct para requests de creación/actualización
type RegistroToxicologicoRequest struct {
	AlumnoID       int64   `json:"alumno_id" valid:"required"`
	ConvocatoriaID int64   `json:"convocatoria_id" valid:"required"`
	Estado         string  `json:"estado" valid:"required,in(verificado|observado|pendiente)"`
	Comentario     *string `json:"comentario"`
	IDUsuario      *int64  `json:"id_usuario"`
}

// RegistroToxicologicoUpdate struct para requests de actualización (solo estado y comentario)
type RegistroToxicologicoUpdate struct {
	Estado     string  `json:"estado" valid:"required,in(verificado|observado|pendiente)"`
	Comentario *string `json:"comentario"`
	IDUsuario  *int64  `json:"id_usuario"`
}

// RegistroToxicologicoResponse struct para respuestas con información adicional
type RegistroToxicologicoResponse struct {
	ID                 int64     `json:"id"`
	AlumnoID           int64     `json:"alumno_id"`
	CodigoEstudiante   string    `json:"codigo_estudiante"`
	ConvocatoriaID     int64     `json:"convocatoria_id"`
	Estado             string    `json:"estado"`
	Comentario         *string   `json:"comentario"`
	IDUsuario          *int64    `json:"id_usuario"`
	UsuarioNombre      *string   `json:"usuario_nombre"`
	FechaCreacion      time.Time `json:"fecha_creacion"`
	FechaActualizacion time.Time `json:"fecha_actualizacion"`
}

// EstadoToxicologicoConvocatoria struct para obtener estados por convocatoria
type EstadoToxicologicoConvocatoria struct {
	AlumnoID         int64      `json:"alumno_id" db:"alumno_id"`
	CodigoEstudiante string     `json:"codigo_estudiante" db:"codigo_estudiante"`
	Estado           string     `json:"estado" db:"estado"`
	Comentario       *string    `json:"comentario" db:"comentario"`
	FechaExamen      *time.Time `json:"fecha_examen" db:"fecha_examen"`
	UsuarioNombre    *string    `json:"usuario_nombre" db:"usuario_nombre"`
}

// Valid valida el modelo RegistroToxicologico
func (m *RegistroToxicologico) Valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Valid valida el request de creación
func (m *RegistroToxicologicoRequest) Valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Valid valida el request de actualización
func (m *RegistroToxicologicoUpdate) Valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
