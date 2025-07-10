package nursing_consultation_routine_review

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type RoutineReview struct {
	ID                       string     `json:"id" db:"id"`
	CreatedAt                *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt                *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
	IsDeleted                bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted              *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt                *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator              *string    `json:"user_creator" db:"user_creator"`
	IDConsultaEnfermeria     string     `json:"consulta_enfermeria_id" db:"consulta_enfermeria_id"`
	FiebreUltimoQuinceDias   string     `json:"fiebre_ultimo_quince_dias" db:"fiebre_ultimo_quince_dias" valid:"-"`
	TosMasQuinceDias         string     `json:"tos_mas_quince_dias" db:"tos_mas_quince_dias" valid:"-"`
	SecrecionLesionGenitales string     `json:"secrecion_lesion_genitales" db:"secrecion_lesion_genitales" valid:"-"`
	FechaUltimaRegla         string     `json:"fecha_ultima_regla" db:"fecha_ultima_regla" valid:"-"`
	Comentarios              string     `json:"comentarios" db:"comentarios" valid:"-"`
}

func NewRoutineReview(id string, consulta_enfermeria_id string, fiebre_ultimo_quince_dias string, tos_mas_quince_dias string, secrecion_lesion_genitales string, fecha_ultima_regla string, comentarios string) *RoutineReview {
	now := time.Now()
	return &RoutineReview{
		ID:                       id,
		IDConsultaEnfermeria:     consulta_enfermeria_id,
		FiebreUltimoQuinceDias:   fiebre_ultimo_quince_dias,
		TosMasQuinceDias:         tos_mas_quince_dias,
		SecrecionLesionGenitales: secrecion_lesion_genitales,
		FechaUltimaRegla:         fecha_ultima_regla,
		Comentarios:              comentarios,
		IsDeleted:                false,
		CreatedAt:                &now,
		UpdatedAt:                &now,
	}
}

func (m *RoutineReview) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
