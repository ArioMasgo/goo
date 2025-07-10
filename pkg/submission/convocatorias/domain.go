package convocatorias

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Convocatorias struct {
	ID            int64      `json:"id" db:"id" valid:"-"`
	FechaInicio   *time.Time `json:"fecha_inicio" db:"fecha_inicio" valid:"optional"`
	FechaFin      *time.Time `json:"fecha_fin" db:"fecha_fin" valid:"optional"`
	Nombre        string     `json:"nombre" db:"nombre" valid:"required"`
	UserId        int64      `json:"user_id" db:"user_id" valid:"required"`
	CreditoMinimo *int       `json:"credito_minimo" db:"credito_minimo" valid:"optional"`
	NotaMinima    *int       `json:"nota_minima" db:"nota_minima" valid:"optional"`
	CreatedAt     *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at" db:"updated_at"`
}

func NewSubmissions(id int64, fechaInicio *time.Time, fechaFin *time.Time, nombre string, userId int64, creditoMinimo *int, notaMinima *int) *Convocatorias {
	return &Convocatorias{
		ID:            id,
		FechaInicio:   fechaInicio,
		FechaFin:      fechaFin,
		Nombre:        nombre,
		UserId:        userId,
		CreditoMinimo: creditoMinimo,
		NotaMinima:    notaMinima,
	}
}

func NewCreateSubmissions(fechaInicio *time.Time, fechaFin *time.Time, nombre string, userId int64, creditoMinimo *int, notaMinima *int) *Convocatorias {
	return &Convocatorias{
		FechaInicio:   fechaInicio,
		FechaFin:      fechaFin,
		Nombre:        nombre,
		UserId:        userId,
		CreditoMinimo: creditoMinimo,
		NotaMinima:    notaMinima,
	}
}

func (m *Convocatorias) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
