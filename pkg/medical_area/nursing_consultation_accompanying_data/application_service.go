package nursing_consultation_accompanying_data

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerAccompanyingData interface {
	CreateAccompanyingData(id, consulta_enfermeria_id, dni, nombres_apellidos, edad string) (*AccompanyingData, int, error)
	UpdateAccompanyingData(id, consulta_enfermeria_id, dni, nombres_apellidos, edad string) (*AccompanyingData, int, error)
	DeleteAccompanyingData(id string) (int, error)
	DeleteAccompanyingDataByIDConsultation(id string) (int, error)
	GetAccompanyingDataByID(id string) (*AccompanyingData, int, error)
	GetAllAccompanyingData() ([]*AccompanyingData, error)
}

type service struct {
	repository ServicesAccompanyingDataRepository
	user       *models.User
	txID       string
}

func NewAccompanyingDataService(repository ServicesAccompanyingDataRepository, user *models.User, TxID string) PortsServerAccompanyingData {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateAccompanyingData(id, consulta_enfermeria_id, dni, nombres_apellidos, edad string) (*AccompanyingData, int, error) {

	m := NewAccompanyingData(id, consulta_enfermeria_id, dni, nombres_apellidos, edad)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create patient :", err)
		return m, 3, err
	}

	return m, 29, nil
}

func (s *service) UpdateAccompanyingData(id, consulta_enfermeria_id, dni, nombres_apellidos, edad string) (*AccompanyingData, int, error) {
	m := NewAccompanyingData(id, consulta_enfermeria_id, dni, nombres_apellidos, edad)
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
		logger.Error.Println(s.txID, " - couldn't update nursing accompanying data :", err)
		return nil, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteAccompanyingData(id string) (int, error) {
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

func (s *service) DeleteAccompanyingDataByIDConsultation(id string) (int, error) {
	if err := uuid.Validate(id); err != nil {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return 15, fmt.Errorf("id is required")
	}

	if err := s.repository.deleteByIDConsultation(id); err != nil {
		if err.Error() == "rows affected error" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update row:", err)
		return 20, err
	}
	return 28, nil
}

func (s *service) GetAccompanyingDataByID(id string) (*AccompanyingData, int, error) {
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

func (s *service) GetAllAccompanyingData() ([]*AccompanyingData, error) {
	return s.repository.getAll()
}
