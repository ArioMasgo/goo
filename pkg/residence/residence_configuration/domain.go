package residence_configuration

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type ResidenceConfiguration struct {
	ID                      string     `db:"id" json:"id" valid:"uuid,required"`
	PercentageFcea          float64    `db:"porcentaje_fcea" json:"percentage_fcea" valid:"required"`
	PercentageEngineering   float64    `db:"porcentaje_ingenieria" json:"percentage_engineering" valid:"required"`
	MinimumGradeFcea        float64    `db:"nota_minima_fcea" json:"minimum_grade_fcea" valid:"required"`
	MinimumGradeEngineering float64    `db:"nota_minima_ingenieria" json:"minimum_grade_engineering" valid:"required"`
	ResidenceID             string     `db:"residencia_id" json:"residence_id" valid:"uuid,required"`
	IsNewbie                bool       `db:"es_cachimbo" json:"is_newbie"`
	CreatedBy               int64      `db:"created_by" json:"created_by" valid:"required"`
	CreatedAt               *time.Time `db:"created_at" json:"created_at"`
	UpdatedBy               int64      `db:"updated_by" json:"updated_by" valid:"required"`
	UpdatedAt               *time.Time `db:"updated_at" json:"updated_at"`
}

func NewResidenceConfiguration(id string, percentageFcea float64, percentageEngineering float64, minimumGradeFcea float64, minimumGradeEngineering float64, residenceID string, isNewbie bool, createdBy int64) *ResidenceConfiguration {
	now := time.Now()
	return &ResidenceConfiguration{
		ID:                      id,
		PercentageFcea:          percentageFcea,
		PercentageEngineering:   percentageEngineering,
		MinimumGradeFcea:        minimumGradeFcea,
		MinimumGradeEngineering: minimumGradeEngineering,
		ResidenceID:             residenceID,
		IsNewbie:                isNewbie,
		CreatedAt:               &now,
		UpdatedAt:               &now,
		CreatedBy:               createdBy,
		UpdatedBy:               createdBy,
	}
}

func (m *ResidenceConfiguration) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
