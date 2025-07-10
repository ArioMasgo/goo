package dentistry_consultation_buccal_procedure

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type BuccalProcedure struct {
	ID                    string     `json:"id" db:"id"`
	IDConsultaOdontologia string     `json:"consulta_odontologia_id" db:"consulta_odontologia_id"`
	TipoProcedimiento     string     `json:"tipo_procedimiento" db:"tipo_procedimiento" valid:"required"`
	Recibo                string     `json:"recibo" db:"recibo"`
	Costo                 string     `json:"costo" db:"costo"`
	FechaPago             string     `json:"fecha_pago" db:"fecha_pago"`
	PiezaDental           string     `json:"pieza_dental" db:"pieza_dental"`
	Comentarios           string     `json:"comentarios" db:"comentarios"`
	IsDeleted             bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted           *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt             *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator           string     `json:"user_creator" db:"user_creator"`
	CreatedAt             *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt             *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
}

func NewBuccalProcedure(id string, consulta_odontologia_id string, tipo_procedimiento string, recibo string, costo string, fecha_pago string, pieza_dental string, comentarios string) *BuccalProcedure {
	now := time.Now()
	return &BuccalProcedure{
		ID:                    id,
		IDConsultaOdontologia: consulta_odontologia_id,
		TipoProcedimiento:     tipo_procedimiento,
		Recibo:                recibo,
		Costo:                 costo,
		FechaPago:             fecha_pago,
		PiezaDental:           pieza_dental,
		Comentarios:           comentarios,
		IsDeleted:             false,
		CreatedAt:             &now,
		UpdatedAt:             &now,
	}
}

func (m *BuccalProcedure) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
