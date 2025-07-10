package exam_toxicologico

import (
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
)

type RegistroToxicologico struct {
	ID                 int64     `json:"id" db:"id"`
	AlumnoID           int64     `json:"alumno_id" db:"alumno_id" valid:"required"`
	ConvocatoriaID     int64     `json:"convocatoria_id" db:"convocatoria_id" valid:"required"`
	Estado             string    `json:"estado" db:"estado" valid:"required,in(verificado|observado|pendiente)"`
	Comentario         *string   `json:"comentario" db:"comentario"`
	IDUsuario          *int64    `json:"id_usuario" db:"id_usuario"`
	FechaCreacion      time.Time `json:"fecha_creacion" db:"fecha_creacion"`
	FechaActualizacion time.Time `json:"fecha_actualizacion" db:"fecha_actualizacion"`
}
type EstadoToxicologicoConvocatoria struct {
	AlumnoID          int64      `json:"alumno_id" db:"alumno_id"`
	CodigoEstudiante  string     `json:"codigo_estudiante" db:"codigo_estudiante"`
	Nombres           string     `json:"nombres" db:"nombres"`
	ApellidoPaterno   string     `json:"apellido_paterno" db:"apellido_paterno"`
	ApellidoMaterno   string     `json:"apellido_materno" db:"apellido_materno"`
	EscuelaProfesional string    `json:"escuela_profesional" db:"escuela_profesional"`
	Estado            string     `json:"estado" db:"estado"`
	Comentario        *string    `json:"comentario" db:"comentario"`
	FechaExamen       *time.Time `json:"fecha_examen" db:"fecha_examen"`
	UsuarioNombre     *string    `json:"usuario_nombre" db:"usuario_nombre"`
}

func NewRegistroToxicologico(alumnoID int64, convocatoriaID int64, estado string, comentario *string, idUsuario *int64) *RegistroToxicologico {
	return &RegistroToxicologico{
		AlumnoID:           alumnoID,
		ConvocatoriaID:     convocatoriaID,
		Estado:             estado,
		Comentario:         comentario,
		IDUsuario:          idUsuario,
		FechaCreacion:      time.Now(),
		FechaActualizacion: time.Now(),
	}
}

func UpdateRegistroToxicologico(id int64, estado string, comentario *string, idUsuario *int64) *RegistroToxicologico {
	return &RegistroToxicologico{
		ID:                 id,
		Estado:             estado,
		Comentario:         comentario,
		IDUsuario:          idUsuario,
		FechaActualizacion: time.Now(),
	}
}

func (r *RegistroToxicologico) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(r)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *RegistroToxicologico) ValidForUpdate() (bool, error) {
	if r.Estado == "" {
		return false, fmt.Errorf("estado is required")
	}
	validStates := map[string]bool{
		"verificado": true,
		"observado":  true,
		"pendiente":  true,
	}
	
	if !validStates[r.Estado] {
		return false, fmt.Errorf("estado must be one of: verificado, observado, pendiente")
	}
	
	return true, nil
}
