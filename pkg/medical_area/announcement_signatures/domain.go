package announcement_signatures

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type AnnouncementSignatures struct {
	ID               string     `json:"id" db:"id" valid:"uuid,required"`
	ConvocatoriaID   int64      `json:"convocatoria_id" db:"convocatoria_id" valid:"required"`
	PacienteID       string     `json:"paciente_id" db:"paciente_id" valid:"uuid,required"`
	FirmaEnfermeria  string     `json:"firma_enfermeria" db:"firma_enfermeria"`
	FirmaMedicina    string     `json:"firma_medicina" db:"firma_medicina"`
	FirmaOdontologia string     `json:"firma_odontologia" db:"firma_odontologia"`
	FirmaPsicologia  string     `json:"firma_psicologia" db:"firma_psicologia"`
	IsDeleted        bool       `json:"is_deleted" db:"is_deleted"`
	UserDeleted      *string    `json:"user_deleted" db:"user_deleted"`
	DeletedAt        *time.Time `json:"deleted_at" db:"deleted_at"`
	UserCreator      string     `json:"user_creator" db:"user_creator"`
	CreatedAt        *time.Time `json:"created_at" db:"created_at" valid:"required"`
	UpdatedAt        *time.Time `json:"updated_at" db:"updated_at" valid:"required"`
}

type Announcement struct {
	ID          int    `json:"id" db:"id"`
	UserID      *int   `json:"user_id" db:"user_id"`
	FechaInicio string `json:"fecha_inicio" db:"fecha_inicio"`
	FechaFin    string `json:"fecha_fin" db:"fecha_fin"`
	Nombre      string `json:"nombre" db:"nombre"`
	Activo      bool   `json:"activo" db:"activo"`
}

func NewAnnouncementSignatures(id string, convocatoria_id int64, paciente_id, firma_enfermeria, firma_medicina, firma_odontologia, firma_psicologia string) *AnnouncementSignatures {
	now := time.Now()
	return &AnnouncementSignatures{
		ID:               id,
		ConvocatoriaID:   convocatoria_id,
		PacienteID:       paciente_id,
		FirmaEnfermeria:  firma_enfermeria,
		FirmaMedicina:    firma_medicina,
		FirmaOdontologia: firma_odontologia,
		FirmaPsicologia:  firma_psicologia,
		IsDeleted:        false,
		CreatedAt:        &now,
		UpdatedAt:        &now,
	}
}

func (m *AnnouncementSignatures) valid() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
