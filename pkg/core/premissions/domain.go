package permissions

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Permission struct {
	ID          int64     `db:"id" json:"id"`
	Description string    `db:"description" json:"description" valid:"required"`
	Method      string    `db:"method" json:"method" valid:"required"`
	Path        string    `db:"path" json:"path" valid:"required"`
	CreatedAt   time.Time `db:"created_at" json:"created_at" valid:"required"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at" valid:"required"`
}

func NewPermission(
	id int64,
	description string,
	method string,
	path string,
) *Permission {
	now := time.Now()
	return &Permission{
		ID:          id,
		Description: description,
		Method:      method,
		Path:        path,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (m *Permission) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
