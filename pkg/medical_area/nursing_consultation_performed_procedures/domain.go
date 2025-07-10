package nursing_consultation_performed_procedures

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type PerformedProcedures struct {
	ID                      string     `json:"id" db:"id"`
	IDConsultaEnfermeria    string     `json:"consulta_enfermeria_id" db:"consulta_enfermeria_id"`
	Procedimiento           string     `json:"procedimiento" db:"procedimiento" valid:"-"`
	NumeroRecibo            string     `json:"numero_recibo" db:"numero_recibo"`
	Comentarios             string     `json:"comentarios" db:"comentarios"`
	Costo                   string     `json:"costo" db:"costo"`
	FechaPago               string     `json:"fecha_pago" db:"fecha_pago"`
	AreaSolicitante         *string    `json:"area_solicitante" db:"area_solicitante"`
	EspecialistaSolicitante *string    `json:"especialista_solicitante" db:"especialista_solicitante"`
	IsDeleted               bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted             *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt               *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator             *string    `json:"user_creator" db:"user_creator"`
	CreatedAt               *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt               *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
}

func NewPerformedProcedures(id string, consulta_enfermeria_id string, procedimiento string, numero_recibo string, comentarios string, costo string, fecha_pago string, areaSolicitante, especialistaSolicitante *string) *PerformedProcedures {
	now := time.Now()
	return &PerformedProcedures{
		ID:                      id,
		IDConsultaEnfermeria:    consulta_enfermeria_id,
		Procedimiento:           procedimiento,
		NumeroRecibo:            numero_recibo,
		Comentarios:             comentarios,
		Costo:                   costo,
		FechaPago:               fecha_pago,
		AreaSolicitante:         areaSolicitante,
		EspecialistaSolicitante: especialistaSolicitante,
		IsDeleted:               false,
		CreatedAt:               &now,
		UpdatedAt:               &now,
	}
}

func (m *PerformedProcedures) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
