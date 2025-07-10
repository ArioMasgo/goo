package consultation_assignment

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type ConsultationAssignment struct {
	ID           string     `json:"id" db:"id"`
	IDConsulta   string     `json:"consulta_id" db:"consulta_id"`
	AreaAsignada string     `json:"area_asignada" db:"area_asignada" valid:"-"`
	IsDeleted    bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted  *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt    *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator  *string    `json:"user_creator" db:"user_creator"`
	CreatedAt    *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt    *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
}

func NewConsultationAssignment(id string, consulta_id string, area_asignada string) *ConsultationAssignment {
	now := time.Now()
	return &ConsultationAssignment{
		ID:           id,
		IDConsulta:   consulta_id,
		AreaAsignada: area_asignada,
		IsDeleted:    false,
		CreatedAt:    &now,
		UpdatedAt:    &now,
	}
}

func (m *ConsultationAssignment) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
