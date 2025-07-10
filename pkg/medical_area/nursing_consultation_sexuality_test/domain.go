package nursing_consultation_sexuality_test

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type SexualityTest struct {
	ID                    string     `json:"id" db:"id"`
	CreatedAt             *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt             *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
	IsDeleted             bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted           *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt             *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator           *string    `json:"user_creator" db:"user_creator"`
	IDConsultaEnfermeria  string     `json:"consulta_enfermeria_id" db:"consulta_enfermeria_id"`
	ActividadSexual       string     `json:"actividad_sexual" db:"actividad_sexual" valid:"-"`
	PlanificacionFamiliar string     `json:"planificacion_familiar" db:"planificacion_familiar" valid:"-"`
	Comentarios           string     `json:"comentarios" db:"comentarios" valid:"-"`
}

func NewSexualityTest(id string, consulta_enfermeria_id string, actividad_sexual string, planificacion_familiar string, comentarios string) *SexualityTest {
	now := time.Now()
	return &SexualityTest{
		ID:                    id,
		IDConsultaEnfermeria:  consulta_enfermeria_id,
		ActividadSexual:       actividad_sexual,
		PlanificacionFamiliar: planificacion_familiar,
		Comentarios:           comentarios,
		IsDeleted:             false,
		CreatedAt:             &now,
		UpdatedAt:             &now,
	}
}

func (m *SexualityTest) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
