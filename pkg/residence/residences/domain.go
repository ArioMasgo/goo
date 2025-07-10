package residences

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Residence struct {
	ID          string     `db:"id" json:"id" valid:"uuid,required"`
	Name        string     `db:"nombre" json:"name" valid:"required"`
	Gender      string     `db:"genero" json:"gender" valid:"in(femenino|masculino),required"`
	Description string     `db:"description" json:"description" valid:"required"`
	Address     string     `db:"direccion" json:"address" valid:"required"`
	Status      string     `db:"estado" json:"status" valid:"in(mantenimiento|deshabilitado|habilitado),required"`
	CreatedBy   int64      `db:"created_by" json:"created_by" valid:"required"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	UpdatedBy   int64      `db:"updated_by" json:"updated_by" valid:"required"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
}

func NewResidence(id string, name string, gender string, description string, address string, status string, createdBy int64) *Residence {
	now := time.Now()
	return &Residence{
		ID:          id,
		Name:        name,
		Gender:      gender,
		Description: description,
		Address:     address,
		Status:      status,
		CreatedAt:   &now,
		UpdatedAt:   &now,
		CreatedBy:   createdBy,
		UpdatedBy:   createdBy,
	}
}

func (m *Residence) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
