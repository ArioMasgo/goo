package dentistry_consultation_odontogram_review

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type RequestDentistryConsultation struct {
	ConsultaOdontologia DentistryConsultation `json:"consulta_enfermeria"`
	ExamenBucal         *BuccalTest           `json:"revision_rutina"`
	RevisionOdontograma *OdontogramReview     `json:"datos_acompanante"`
	IsDeleted           bool                  `json:"is_deleted" db:"is_deleted"`
	UserDeleted         *string               `json:"user_deleted" db:"user_deleted"`
	DeletedAt           *time.Time            `json:"deleted_at" db:"deleted_at"`
	UserCreator         string                `json:"user_creator" db:"user_creator"`
	CreatedAt           *time.Time            `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt           *time.Time            `json:"updated_at" db:"updated_at" valid:"required"`
}

type DentistryConsultation struct {
	ID            string     `json:"id" db:"id"`
	IDPaciente    string     `json:"paciente_id" db:"paciente_id"`
	FechaConsulta string     `json:"fecha_consulta" db:"fecha_consulta"`
	IsDeleted     bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted   *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt     *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator   string     `json:"user_creator" db:"user_creator"`
	CreatedAt     *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt     *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
}

type BuccalTest struct {
	ID                      string     `json:"id" db:"id"`
	IDConsultaOdontologia   string     `json:"consulta_odontologia_id" db:"consulta_odontologia_id"`
	CapacidadMasticatoria   string     `json:"capacidad_masticatoria" db:"capacidad_masticatoria"`
	Encias                  string     `json:"encias" db:"encias"`
	CariesDentales          string     `json:"caries_dentales" db:"caries_dentales"`
	EdentulismoParcialTotal string     `json:"edentulismo_parcial_total" db:"edentulismo_parcial_total"`
	PortadorProtesisDental  string     `json:"portador_protesis_dental" db:"portador_protesis_dental"`
	EstadoHigieneBucal      string     `json:"estado_higiene_bucal" db:"estado_higiene_bucal"`
	UrgenciaTratamiento     string     `json:"urgencia_tratamiento" db:"urgencia_tratamiento"`
	Fluorizacion            string     `json:"fluorizacion" db:"fluorizacion"`
	Destartraje             string     `json:"destartraje" db:"destartraje"`
	Comentarios             string     `json:"comentarios" db:"comentarios"`
	IsDeleted               bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted             *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt               *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator             string     `json:"user_creator" db:"user_creator"`
	CreatedAt               *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt               *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
}

type OdontogramReview struct {
	ID                    string     `json:"id" db:"id"`
	IDConsultaOdontologia string     `json:"consulta_odontologia_id" db:"consulta_odontologia_id"`
	Caries                string     `json:"caries" db:"caries"`
	Erupcionado           string     `json:"erupcionado" db:"erupcionado"`
	Perdido               string     `json:"perdido" db:"perdido"`
	Costo                 string     `json:"costo" db:"costo"`
	FechaPago             string     `json:"fecha_pago" db:"fecha_pago"`
	Cpod                  string     `json:"cpod" db:"cpod"`
	Diagnostico           string     `json:"diagnostico" db:"diagnostico"`
	Mes                   string     `json:"mes" db:"mes"`
	Comentarios           string     `json:"comentarios" db:"comentarios"`
	IsDeleted             bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted           *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt             *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator           string     `json:"user_creator" db:"user_creator"`
	CreatedAt             *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt             *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
}

func NewOdontogramReview(id string, consulta_odontologia_id string, caries string, erupcionado string, perdido string, costo string, fecha_pago string, cpod string, diagnostico string, mes string, comentarios string) *OdontogramReview {
	now := time.Now()
	return &OdontogramReview{
		ID:                    id,
		IDConsultaOdontologia: consulta_odontologia_id,
		Caries:                caries,
		Erupcionado:           erupcionado,
		Perdido:               perdido,
		Costo:                 costo,
		FechaPago:             fecha_pago,
		Cpod:                  cpod,
		Diagnostico:           diagnostico,
		Mes:                   mes,
		Comentarios:           comentarios,
		IsDeleted:             false,
		CreatedAt:             &now,
		UpdatedAt:             &now,
	}
}

func (m *DentistryConsultation) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *BuccalTest) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *OdontogramReview) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
