package nursing_consultation_medication_treatment

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type MedicationTreatment struct {
	ID                        string     `json:"id" db:"id"`
	IDConsultaEnfermeria      string     `json:"consulta_enfermeria_id" db:"consulta_enfermeria_id"`
	NombreGenericoMedicamento string     `json:"nombre_generico_medicamento" db:"nombre_generico_medicamento"`
	ViaAdministracion         string     `json:"via_administracion" db:"via_administracion"`
	HoraAplicacion            string     `json:"hora_aplicacion" db:"hora_aplicacion"`
	ResponsableAtencion       string     `json:"responsable_atencion" db:"responsable_atencion"`
	Observaciones             string     `json:"observaciones" db:"observaciones"`
	AreaSolicitante           *string    `json:"area_solicitante" db:"area_solicitante"`
	EspecialistaSolicitante   *string    `json:"especialista_solicitante" db:"especialista_solicitante"`
	IsDeleted                 bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted               *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt                 *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator               *string    `json:"user_creator" db:"user_creator"`
	CreatedAt                 *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt                 *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
}

func NewMedicationTreatment(id, consulta_enfermeria_id, nombre_generico_medicamento, via_administracion, hora_aplicacion, responsable_atencion, observaciones string, areaSolicitante, especialistaSolicitante *string) *MedicationTreatment {
	now := time.Now()
	return &MedicationTreatment{
		ID:                        id,
		IDConsultaEnfermeria:      consulta_enfermeria_id,
		NombreGenericoMedicamento: nombre_generico_medicamento,
		ViaAdministracion:         via_administracion,
		HoraAplicacion:            hora_aplicacion,
		ResponsableAtencion:       responsable_atencion,
		Observaciones:             observaciones,
		AreaSolicitante:           areaSolicitante,
		EspecialistaSolicitante:   especialistaSolicitante,
		IsDeleted:                 false,
		CreatedAt:                 &now,
		UpdatedAt:                 &now,
	}
}

func (m *MedicationTreatment) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
