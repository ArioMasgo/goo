package models

import (
	"time"
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
	AlumnoNombre  *string `json:"alumno_nombre,omitempty" db:"alumno_nombre"`
	AlumnoCodigo  *string `json:"alumno_codigo,omitempty" db:"alumno_codigo"`
	UsuarioNombre *string `json:"usuario_nombre,omitempty" db:"usuario_nombre"`
}

type AlumnoPendienteVisita struct {
	AlumnoID           uint64 `json:"alumno_id" db:"alumno_id"`
	Codigo             string `json:"codigo" db:"codigo"`
	Nombre             string `json:"nombre" db:"nombre"`
	DNI                string `json:"dni" db:"dni"`
	Celular            string `json:"celular" db:"celular"`
	Direccion          string `json:"direccion" db:"direccion"`
	ConvocatoriaID     uint64 `json:"convocatoria_id" db:"convocatoria_id"`
	ConvocatoriaNombre string `json:"convocatoria_nombre" db:"convocatoria_nombre"`
}
