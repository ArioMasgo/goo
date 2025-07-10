package level_user_permissions

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type LevelUserPermission struct {
	ID           int64     `db:"id" json:"id"`
	PermissionID int64     `db:"permission_id" json:"permission_id" valid:"required"`
	LevelUserID  int64     `db:"level_user_id" json:"level_user_id" valid:"required"`
	CreatedAt    time.Time `db:"created_at" json:"created_at" valid:"required"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at" valid:"required"`
}

func NewLevelUserPermission(
	id int64,
	permissionId int64,
	levelUserId int64,
) *LevelUserPermission {
	now := time.Now()
	return &LevelUserPermission{
		ID:           id,
		LevelUserID:  levelUserId,
		PermissionID: permissionId,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func (m *LevelUserPermission) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
