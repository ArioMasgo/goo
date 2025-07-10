package models

import (
	"github.com/asaskevich/govalidator"
	"time"
)

// User  Model struct User
type User struct {
	ID              int64      `db:"id" json:"id"`
	Username        string     `db:"username" json:"username"`
	FullName        string     `db:"full_name" json:"full_name,omitempty"`
	Email           string     `db:"email" json:"email" valid:"required"`
	EmailVerifiedAt *time.Time `db:"email_verified_at" json:"email_verified_at" valid:"required"`
	Password        string     `db:"password" json:"password,omitempty"`
	IPAddress       string     `db:"ip_address" json:"ip_address" valid:"required"`
	IDLevelUser     int64      `db:"id_level_user" json:"id_level_user" valid:"required"`
	LastUser        *int64     `db:"last_user" json:"last_user"`
	RememberToken   *string    `db:"remember_token" json:"remember_token"`
	StatusID        int64      `db:"status_id" json:"status_id" valid:"required"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at" valid:"required"`
	UpdatedAt       time.Time  `db:"updated_at" json:"updated_at" valid:"required"`
	IsDeleted       bool       `db:"is_deleted" json:"is_deleted"`
	UserDeleter     *int64     `db:"user_deleter" json:"user_deleter"`
	DeletedAt       *time.Time `db:"deleted_at" json:"deleted_at"`
}

func (m *User) Valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
