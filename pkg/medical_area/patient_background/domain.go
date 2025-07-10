package patient_background

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type PatientBackground struct {
	ID          string     `json:"id" db:"id" valid:"uuid,required"`
	IDPaciente  string     `json:"paciente_id" db:"paciente_id"`
	Nombre      string     `json:"nombre_antecedente" db:"nombre_antecedente" valid:"-"`
	Estado      string     `json:"estado_antecedente" db:"estado_antecedente" valid:"-"`
	IsDeleted   bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt   *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator string     `json:"user_creator" db:"user_creator"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

func NewPatientBackground(id string, paciente_id string, nombre_antecedente string, estado_antecedente string) *PatientBackground {
	return &PatientBackground{
		ID:         id,
		IDPaciente: paciente_id,
		Nombre:     nombre_antecedente,
		Estado:     estado_antecedente,
		IsDeleted:  false,
	}
}

func (m *PatientBackground) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
