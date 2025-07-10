package users

import (
	"dbu-api/internal/models"
	"time"

	"github.com/asaskevich/govalidator"
)

type User models.User

func NewUser(
	id int64,
	username string,
	fullName string,
	email string,
	password string,
	ipAddress string,
	idLevelUser int64,
	statusID int64,
) *User {
	now := time.Now()

	return &User{
		ID:              id,
		Username:        username,
		FullName:        fullName,
		Email:           email,
		EmailVerifiedAt: nil,
		Password:        password,
		IPAddress:       ipAddress,
		IDLevelUser:     idLevelUser,
		StatusID:        statusID,
		CreatedAt:       now,
		UpdatedAt:       now,
		IsDeleted:       false,
	}
}

func MapperUser(user *User) *models.User {
	if user == nil {
		return nil
	}

	return &models.User{
		ID:              user.ID,
		Username:        user.Username,
		FullName:        user.FullName,
		Email:           user.Email,
		EmailVerifiedAt: user.EmailVerifiedAt,
		Password:        user.Password,
		IPAddress:       user.IPAddress,
		IDLevelUser:     user.IDLevelUser,
		StatusID:        user.StatusID,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		IsDeleted:       user.IsDeleted,
		DeletedAt:       user.DeletedAt,
	}
}

func (m *User) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
