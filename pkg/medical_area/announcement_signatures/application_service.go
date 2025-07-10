package announcement_signatures

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerAnnouncementSignatures interface {
	CreateAnnouncementSignatures(id string, convocatoria_id int64, paciente_id, firma_id, area string) (*AnnouncementSignatures, int, error)
	UpdateAnnouncementSignatures(id string, convocatoria_id int64, paciente_id, firma_id, area string) (*AnnouncementSignatures, int, error)
	DeleteAnnouncementSignatures(id string) (int, error)
	GetAnnouncementSignaturesByID(id string) (*AnnouncementSignatures, int, error)
	GetAnnouncementSignaturesByIDPatient(id string) (*AnnouncementSignatures, int, error)
	GetAnnouncementSignaturesByDNIPatient(dni string) (*AnnouncementSignatures, int, error)
	GetAllAnnouncementSignatures() ([]*AnnouncementSignatures, error)
	GetAnnouncement() (*Announcement, error)
}

type service struct {
	repository ServicesAnnouncementSignaturesRepository
	user       *models.User
	txID       string
}

func NewAnnouncementSignaturesService(repository ServicesAnnouncementSignaturesRepository, user *models.User, TxID string) PortsServerAnnouncementSignatures {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateAnnouncementSignatures(id string, convocatoria_id int64, paciente_id, firma_id, area string) (*AnnouncementSignatures, int, error) {
	firma_enfermeria, firma_medicina, firma_odontologia, firma_psicologia := "", "", "", ""

	if area == "enfermería" {
		firma_enfermeria = firma_id
	}

	if area == "odontología" {
		firma_odontologia = firma_id
	}

	if area == "medicina" {
		firma_medicina = firma_id
	}

	if area == "psicología" {
		firma_psicologia = firma_id
	}

	m := NewAnnouncementSignatures(id, convocatoria_id, paciente_id, firma_enfermeria, firma_medicina, firma_odontologia, firma_psicologia)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create announcement signatures :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateAnnouncementSignatures(id string, convocatoria_id int64, paciente_id, firma_id, area string) (*AnnouncementSignatures, int, error) {
	firma_enfermeria, firma_medicina, firma_odontologia, firma_psicologia := "", "", "", ""

	announcementSignatures, code, err := s.GetAnnouncementSignaturesByIDPatient(paciente_id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't meet validations:", err)
		return nil, code, err
	}

	if area == "enfermería" {
		firma_enfermeria = firma_id
		firma_odontologia = announcementSignatures.FirmaOdontologia
		firma_medicina = announcementSignatures.FirmaMedicina
		firma_psicologia = announcementSignatures.FirmaPsicologia
	}

	if area == "odontología" {
		firma_odontologia = firma_id
		firma_enfermeria = announcementSignatures.FirmaEnfermeria
		firma_medicina = announcementSignatures.FirmaMedicina
		firma_psicologia = announcementSignatures.FirmaPsicologia
	}

	if area == "medicina" {
		firma_medicina = firma_id
		firma_enfermeria = announcementSignatures.FirmaEnfermeria
		firma_odontologia = announcementSignatures.FirmaOdontologia
		firma_psicologia = announcementSignatures.FirmaPsicologia
	}

	if area == "psicología" {
		firma_psicologia = firma_id
		firma_enfermeria = announcementSignatures.FirmaEnfermeria
		firma_odontologia = announcementSignatures.FirmaOdontologia
		firma_medicina = announcementSignatures.FirmaMedicina
	}

	m := NewAnnouncementSignatures(id, convocatoria_id, paciente_id, firma_enfermeria, firma_medicina, firma_odontologia, firma_psicologia)
	valid, err := m.valid()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't meet validations:", err)
		return nil, 15, err
	}
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return nil, 15, err
	}
	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update announcement signatures :", err)
		return nil, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteAnnouncementSignatures(id string) (int, error) {
	if err := uuid.Validate(id); err != nil {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return 15, fmt.Errorf("id is required")
	}

	if err := s.repository.delete(id); err != nil {
		if err.Error() == "rows affected error" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update row:", err)
		return 20, err
	}
	return 28, nil
}

func (s *service) GetAnnouncementSignaturesByID(id string) (*AnnouncementSignatures, int, error) {
	if err := uuid.Validate(id); err != nil {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return nil, 15, fmt.Errorf("id is required")
	}
	m, err := s.repository.getByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAnnouncementSignaturesByIDPatient(id string) (*AnnouncementSignatures, int, error) {
	if err := uuid.Validate(id); err != nil {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return nil, 15, fmt.Errorf("id is required")
	}
	m, err := s.repository.getByIDPatient(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAnnouncementSignaturesByDNIPatient(id string) (*AnnouncementSignatures, int, error) {
	m, err := s.repository.getByDNIPatient(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllAnnouncementSignatures() ([]*AnnouncementSignatures, error) {
	return s.repository.getAll()
}

func (s *service) GetAnnouncement() (*Announcement, error) {
	return s.repository.GetAnnouncement()
}
