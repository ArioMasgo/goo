package dentistry_consultation_buccal_consultation

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type BuccalConsultation struct {
	ID                    string     `json:"id" db:"id"`
	IDConsultaOdontologia string     `json:"consulta_odontologia_id" db:"consulta_odontologia_id"`
	Relato                string     `json:"relato" db:"relato" valid:"-"`
	Diagnostico           string     `json:"diagnostico" db:"diagnostico" valid:"-"`
	ExamenAuxiliar        string     `json:"examen_auxiliar" db:"examen_auxiliar" valid:"-"`
	ExamenClinico         *string    `json:"examen_clinico" db:"examen_clinico" valid:"-"`
	Tratamiento           string     `json:"tratamiento" db:"tratamiento" valid:"-"`
	Indicaciones          string     `json:"indicaciones" db:"indicaciones" valid:"-"`
	Comentarios           string     `json:"comentarios" db:"comentarios" valid:"-"`
	IsDeleted             bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted           *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt             *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator           string     `json:"user_creator" db:"user_creator"`
	CreatedAt             *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt             *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
}

func NewBuccalConsultation(id string, consulta_odontologia_id string, relato string, diagnostico string, examen_auxiliar string, examen_clinico *string, tratamiento string, indicaciones string, comentarios string) *BuccalConsultation {
	now := time.Now()
	return &BuccalConsultation{
		ID:                    id,
		IDConsultaOdontologia: consulta_odontologia_id,
		Relato:                relato,
		Diagnostico:           diagnostico,
		ExamenAuxiliar:        examen_auxiliar,
		ExamenClinico:         examen_clinico,
		Tratamiento:           tratamiento,
		Indicaciones:          indicaciones,
		Comentarios:           comentarios,
		IsDeleted:             false,
		CreatedAt:             &now,
		UpdatedAt:             &now,
	}
}

func (m *BuccalConsultation) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
