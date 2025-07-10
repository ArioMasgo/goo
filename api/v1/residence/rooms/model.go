package rooms

import (
	"github.com/asaskevich/govalidator"
)

type RoomRequest struct {
	ID       string `json:"id" valid:"uuid,required"`
	Capacity int    `json:"capacity" valid:"numeric,required"`
	Status   string `json:"status" valid:"in(mantenimiento|deshabilitado|habilitado),required"`
}

type AssignmentRoomRequest struct {
	StudentID    int64 `json:"student_id" valid:"required"`
	SubmissionID int64 `json:"submission_id" valid:"required"`
}

type DeleteAssignmentRoomRequest struct {
	StudentID    int64  `json:"student_id" valid:"required"`
	SubmissionID int64  `json:"submission_id" valid:"required"`
	Status       string `json:"status" valid:"in(activo|desocupado|suspendido|cancelado),required"`
	Observation  string `json:"observation" valid:"required"`
}

func (m *RoomRequest) Valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *AssignmentRoomRequest) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *DeleteAssignmentRoomRequest) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
