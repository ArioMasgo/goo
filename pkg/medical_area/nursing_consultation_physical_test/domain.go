package nursing_consultation_physical_test

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type PhysicalTest struct {
	ID                    string     `json:"id" db:"id"`
	CreatedAt             *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt             *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
	IsDeleted             bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted           *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt             *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator           *string    `json:"user_creator" db:"user_creator"`
	IDConsultaEnfermeria  string     `json:"consulta_enfermeria_id" db:"consulta_enfermeria_id"`
	TallaPesos            string     `json:"talla_peso" db:"talla_peso" valid:"-"`
	PerimetroCintura      string     `json:"perimetro_cintura" db:"perimetro_cintura" valid:"-"`
	IndiceMasaCorporalImg string     `json:"indice_masa_corporal_img" db:"indice_masa_corporal_img" valid:"-"`
	PresionArterial       string     `json:"presion_arterial" db:"presion_arterial" valid:"-"`
	Comentarios           string     `json:"comentarios" db:"comentarios" valid:"-"`
}

func NewPhysicalTest(id string, consulta_enfermeria_id string, talla_peso string, perimetro_cintura string, indice_masa_corporal_img string, presion_arterial string, comentarios string) *PhysicalTest {
	now := time.Now()
	return &PhysicalTest{
		ID:                    id,
		IDConsultaEnfermeria:  consulta_enfermeria_id,
		TallaPesos:            talla_peso,
		PerimetroCintura:      perimetro_cintura,
		IndiceMasaCorporalImg: indice_masa_corporal_img,
		PresionArterial:       presion_arterial,
		Comentarios:           comentarios,
		IsDeleted:             false,
		CreatedAt:             &now,
		UpdatedAt:             &now,
	}
}

func (m *PhysicalTest) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
