package models

import "github.com/asaskevich/govalidator"

type ConsultationAssignment struct {
	IDConsulta   string `json:"consulta_id"`
	AreaAsignada string `json:"area_asignada"`
}

func (m *ConsultationAssignment) ValidConsultationAssignment() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
