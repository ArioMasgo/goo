package nursing_consultation_laboratory_test

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type LaboratoryTest struct {
	ID                   string     `json:"id" db:"id"`
	CreatedAt            *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt            *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
	IsDeleted            bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted          *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt            *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator          *string    `json:"user_creator" db:"user_creator"`
	IDConsultaEnfermeria string     `json:"consulta_enfermeria_id" db:"consulta_enfermeria_id"`
	Serologia            string     `json:"serologia" db:"serologia" valid:"-"`
	Bk                   string     `json:"bk" db:"bk" valid:"-"`
	Hemograma            string     `json:"hemograma" db:"hemograma" valid:"-"`
	ExamenOrina          string     `json:"examen_orina" db:"examen_orina" valid:"-"`
	Colesterol           string     `json:"colesterol" db:"colesterol" valid:"-"`
	Glucosa              string     `json:"glucosa" db:"glucosa" valid:"-"`
	Comentarios          string     `json:"comentarios" db:"comentarios" valid:"-"`
}

func NewLaboratoryTest(id string, consulta_enfermeria_id string, serologia string, bk string, hemograma string, examen_orina string, colesterol string, glucosa string, comentarios string) *LaboratoryTest {
	now := time.Now()
	return &LaboratoryTest{
		ID:                   id,
		IDConsultaEnfermeria: consulta_enfermeria_id,
		Serologia:            serologia,
		Bk:                   bk,
		Hemograma:            hemograma,
		ExamenOrina:          examen_orina,
		Colesterol:           colesterol,
		Glucosa:              glucosa,
		Comentarios:          comentarios,
		IsDeleted:            false,
		CreatedAt:            &now,
		UpdatedAt:            &now,
	}
}

func (m *LaboratoryTest) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
