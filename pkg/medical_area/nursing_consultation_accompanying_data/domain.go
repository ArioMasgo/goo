package nursing_consultation_accompanying_data

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type AccompanyingData struct {
	ID                   string     `json:"id" db:"id"`
	CreatedAt            *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt            *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
	IsDeleted            bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted          *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt            *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator          *string    `json:"user_creator" db:"user_creator"`
	IDConsultaEnfermeria string     `json:"consulta_enfermeria_id" db:"consulta_enfermeria_id"`
	DNI                  string     `json:"dni" db:"dni" valid:"-"`
	NombresApellidos     string     `json:"nombres_apellidos" db:"nombres_apellidos" valid:"-"`
	Edad                 string     `json:"edad" db:"edad" valid:"-"`
}

func NewAccompanyingData(id string, consulta_enfermeria_id string, dni string, nombres_apellidos string, edad string) *AccompanyingData {
	now := time.Now()
	return &AccompanyingData{
		ID:                   id,
		IDConsultaEnfermeria: consulta_enfermeria_id,
		DNI:                  dni,
		NombresApellidos:     nombres_apellidos,
		Edad:                 edad,
		IsDeleted:            false,
		CreatedAt:            &now,
		UpdatedAt:            &now,
	}
}

func (m *AccompanyingData) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
