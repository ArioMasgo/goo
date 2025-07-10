package residence_robot

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type ResidenceRobot struct {
	ID               string    `db:"id" json:"id" valid:"uuid,required"`
	ResidenceID      string    `db:"residence_id" json:"residence_id" valid:"uuid,required"`
	PromptTokens     int       `db:"prompt_tokens" json:"prompt_tokens" valid:"required"`
	CompletionTokens int       `db:"completion_tokens" json:"completion_tokens" valid:"required"`
	TotalTokens      int       `db:"total_tokens" json:"total_tokens" valid:"required"`
	CreatedAt        time.Time `db:"created_at" json:"created_at" valid:"-"`
}

func NewResidenceRobot(id, residenceID string, promptTokens, completionTokens, totalTokens int) *ResidenceRobot {
	now := time.Now()
	return &ResidenceRobot{
		ID:               id,
		ResidenceID:      residenceID,
		PromptTokens:     promptTokens,
		CompletionTokens: completionTokens,
		TotalTokens:      totalTokens,
		CreatedAt:        now,
	}
}

func (m *ResidenceRobot) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
