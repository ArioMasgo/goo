package medical_consultation

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type MedicalConsultation struct {
	ID            string     `json:"id" db:"id"`
	IDPaciente    string     `json:"paciente_id" db:"paciente_id"`
	FechaConsulta string     `json:"fecha_consulta" db:"fecha_consulta"`
	AreaMedica    string     `json:"area_medica" db:"area_medica"`
	IsDeleted     bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted   *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt     *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator   string     `json:"user_creator" db:"user_creator"`
	CreatedAt     *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt     *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
}

func NewMedicalConsultation(id string, paciente_id string, fecha_consulta string, area_medica string) *MedicalConsultation {
	now := time.Now()
	return &MedicalConsultation{
		ID:            id,
		IDPaciente:    paciente_id,
		FechaConsulta: fecha_consulta,
		AreaMedica:    area_medica,
		IsDeleted:     false,
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}
}

func (m *MedicalConsultation) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
