package medical_general_medicine_consultation

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type GeneralMedicineConsultation struct {
	ID            string     `json:"id" db:"id"`
	ConsultaID    string     `json:"consulta_id" db:"consulta_id"`
	FechaHora     string     `json:"fecha_hora" db:"fecha_hora"`
	Anamnesis     string     `json:"anamnesis" db:"anamnesis"`
	ExamenClinico string     `json:"examen_clinico" db:"examen_clinico"`
	Indicaciones  string     `json:"indicaciones" db:"indicaciones"`
	IsDeleted     bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted   *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt     *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator   string     `json:"user_creator" db:"user_creator"`
	CreatedAt     *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt     *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
}

func NewGeneralMedicineConsultation(id string, consulta_id string, fecha_hora string, anamnesis string, examen_clinico string, indicaciones string) *GeneralMedicineConsultation {
	now := time.Now()
	return &GeneralMedicineConsultation{
		ID:            id,
		ConsultaID:    consulta_id,
		FechaHora:     fecha_hora,
		Anamnesis:     anamnesis,
		ExamenClinico: examen_clinico,
		Indicaciones:  indicaciones,
		IsDeleted:     false,
		CreatedAt:     &now,
		UpdatedAt:     &now,
	}
}

func (m *GeneralMedicineConsultation) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
