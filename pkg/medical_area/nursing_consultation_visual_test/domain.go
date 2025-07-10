package nursing_consultation_visual_test

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type VisualTest struct {
	ID                   string     `json:"id" db:"id"`
	CreatedAt            *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt            *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
	IsDeleted            bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted          *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt            *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator          *string    `json:"user_creator" db:"user_creator"`
	IDConsultaEnfermeria string     `json:"consulta_enfermeria_id" db:"consulta_enfermeria_id"`
	OjoDerecho           string     `json:"ojo_derecho" db:"ojo_derecho" valid:"-"`
	OjoIzquierdo         string     `json:"ojo_izquierdo" db:"ojo_izquierdo" valid:"-"`
	PresionOcular        string     `json:"presion_ocular" db:"presion_ocular" valid:"-"`
	Comentarios          string     `json:"comentarios" db:"comentarios" valid:"-"`
}

func NewVisualTest(id string, consulta_enfermeria_id string, ojo_derecho string, ojo_izquierdo string, presion_ocular string, comentarios string) *VisualTest {
	now := time.Now()
	return &VisualTest{
		ID:                   id,
		IDConsultaEnfermeria: consulta_enfermeria_id,
		OjoDerecho:           ojo_derecho,
		OjoIzquierdo:         ojo_izquierdo,
		PresionOcular:        presion_ocular,
		Comentarios:          comentarios,
		IsDeleted:            false,
		CreatedAt:            &now,
		UpdatedAt:            &now,
	}
}

func (m *VisualTest) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
