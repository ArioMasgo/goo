package nursing_consultation_vaccine

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Vaccine struct {
	ID                   string     `json:"id" db:"id"`
	CreatedAt            *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt            *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
	IsDeleted            bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted          *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt            *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator          *string    `json:"user_creator" db:"user_creator"`
	IDConsultaEnfermeria string     `json:"consulta_enfermeria_id" db:"consulta_enfermeria_id"`
	TipoVacuna           string     `json:"tipo_vacuna" db:"tipo_vacuna" valid:"-"`
	FechaDosis           string     `json:"fecha_dosis" db:"fecha_dosis" valid:"-"`
	Comentarios          string     `json:"comentarios" db:"comentarios" valid:"-"`
}

type TypesVaccines struct {
	ID            string     `json:"id" db:"id"`
	Nombre        string     `json:"nombre" db:"nombre"`
	Estado        bool       `json:"estado" db:"estado"`
	DuracionMeses string     `json:"duracion_meses" db:"duracion_meses"`
	IsDeleted     bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted   *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt     *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator   *string    `json:"user_creator" db:"user_creator"`
	CreatedAt     *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt     *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
}

func NewVaccine(id string, consulta_enfermeria_id string, tipo_vacuna string, fecha_dosis string, comentarios string) *Vaccine {
	now := time.Now()
	return &Vaccine{
		ID:                   id,
		IDConsultaEnfermeria: consulta_enfermeria_id,
		TipoVacuna:           tipo_vacuna,
		FechaDosis:           fecha_dosis,
		Comentarios:          comentarios,
		IsDeleted:            false,
		CreatedAt:            &now,
		UpdatedAt:            &now,
	}
}

func (m *Vaccine) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
