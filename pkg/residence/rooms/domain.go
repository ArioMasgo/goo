package rooms

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Room struct {
	ID          string     `db:"id" json:"id" valid:"uuid,required"`
	Number      int        `db:"numero" json:"number" valid:"required"`
	Capacity    int        `db:"capacidad" json:"capacity" valid:"required"`
	Status      string     `db:"estado" json:"status" valid:"in(mantenimiento|deshabilitado|habilitado),required"`
	Floor       int        `db:"piso" json:"floor" valid:"required"`
	ResidenceID string     `db:"residencia_id" json:"residence_id" valid:"required"`
	CreatedBy   int64      `db:"created_by" json:"created_by" valid:"required"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	UpdatedBy   int64      `db:"updated_by" json:"updated_by" valid:"required"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
}

func NewRoom(id string, number int, residenceId string, capacity int, status string, floor int, createdBy int64) *Room {
	now := time.Now()
	return &Room{
		ID:          id,
		Number:      number,
		ResidenceID: residenceId,
		Capacity:    capacity,
		Status:      status,
		Floor:       floor,
		CreatedBy:   createdBy,
		CreatedAt:   &now,
		UpdatedBy:   createdBy,
		UpdatedAt:   &now,
	}
}

func (m *Room) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
