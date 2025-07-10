package nursing_consultation_preferential_test

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type PreferentialTest struct {
	ID                    string     `json:"id" db:"id"`
	CreatedAt             *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt             *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
	IsDeleted             bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted           *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt             *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator           *string    `json:"user_creator" db:"user_creator"`
	IDConsultaEnfermeria  string     `json:"consulta_enfermeria_id" db:"consulta_enfermeria_id"`
	AparatoRespiratorio   string     `json:"aparato_respiratorio" db:"aparato_respiratorio" valid:"-"`
	AparatoCardiovascular string     `json:"aparato_cardiovascular" db:"aparato_cardiovascular" valid:"-"`
	AparatoDigestivo      string     `json:"aparato_digestivo" db:"aparato_digestivo" valid:"-"`
	AparatoGenitourinario string     `json:"aparato_genitourinario" db:"aparato_genitourinario" valid:"-"`
	Papanicolau           string     `json:"papanicolau" db:"papanicolau" valid:"-"`
	ExamenMama            string     `json:"examen_mama" db:"examen_mama" valid:"-"`
	Comentarios           string     `json:"comentarios" db:"comentarios" valid:"-"`
}

func NewPreferentialTest(id string, consulta_enfermeria_id string, aparato_respiratorio string, aparato_cardiovascular string, aparato_digestivo string, aparato_genitourinario string, papanicolau string, examen_mama string, comentarios string) *PreferentialTest {
	now := time.Now()
	return &PreferentialTest{
		ID:                    id,
		IDConsultaEnfermeria:  consulta_enfermeria_id,
		AparatoRespiratorio:   aparato_respiratorio,
		AparatoCardiovascular: aparato_cardiovascular,
		AparatoDigestivo:      aparato_digestivo,
		AparatoGenitourinario: aparato_genitourinario,
		Papanicolau:           papanicolau,
		ExamenMama:            examen_mama,
		Comentarios:           comentarios,
		IsDeleted:             false,
		CreatedAt:             &now,
		UpdatedAt:             &now,
	}
}

func (m *PreferentialTest) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
