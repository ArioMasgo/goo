package asignaciones_cuartos

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type RoomAssignment struct {
	ID             string     `json:"id" db:"id" valid:"-"`
	StudentID      int64      `json:"alumno_id" db:"alumno_id" valid:"required"`
	RoomID         string     `json:"cuarto_id" db:"cuarto_id" valid:"required"`
	CallID         int64      `json:"convocatoria_id" db:"convocatoria_id" valid:"required"`
	AssignmentDate time.Time  `json:"fecha_asignacion" db:"fecha_asignacion" valid:"required"`
	Status         string     `json:"estado" db:"estado" valid:"in(activo|desocupado|suspendido|cancelado),required"`
	Observations   *string    `json:"observaciones" db:"observaciones"`
	CreatedAt      *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at" db:"updated_at"`
}

func NewRoomAssignment(id string, studentID int64, roomID string, callID int64, assignmentDate time.Time, status, observation string) *RoomAssignment {
	now := time.Now()
	return &RoomAssignment{
		ID:             id,
		StudentID:      studentID,
		RoomID:         roomID,
		CallID:         callID,
		AssignmentDate: assignmentDate,
		Status:         status,
		Observations:   &observation,
		CreatedAt:      &now,
		UpdatedAt:      &now,
	}
}

func (m *RoomAssignment) Valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
