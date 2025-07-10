package convocatorias

import (
	"fmt"
	"time"

	"dbu-api/internal/logger"
	"dbu-api/internal/models"
)

type PortsServerConvocatorias interface {
	CreateSubmissions(fechaInicio, fechaFin time.Time, nombre string, userId int64, creditoMinimo, notaMinima int) (*Convocatorias, int, error)
	UpdateSubmissions(id int64, fechaInicio, fechaFin time.Time, nombre string, userId int64, creditoMinimo, notaMinima int) (*Convocatorias, int, error)
	DeleteSubmissions(id int64) (int, error)
	GetSubmissionsByID(id int64) (*Convocatorias, int, error)
	GetAllSubmissions() ([]*Convocatorias, error)
	GetAllSubmissionsByService(id int64) ([]*Convocatorias, error)
	GetActiveSubmissions() (*Convocatorias, int, error)
}

type service struct {
	repository ServicesConvocatoriasRepository
	user       *models.User
	txID       string
}

func NewConvocatoriasService(repository ServicesConvocatoriasRepository, user *models.User, TxID string) PortsServerConvocatorias {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateSubmissions(fechaInicio, fechaFin time.Time, nombre string, userId int64, creditoMinimo, notaMinima int) (*Convocatorias, int, error) {
	credMin := creditoMinimo
	notaMin := notaMinima
	m := NewCreateSubmissions(&fechaInicio, &fechaFin, nombre, userId, &credMin, &notaMin)

	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create Convocatorias :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateSubmissions(id int64, fechaInicio, fechaFin time.Time, nombre string, userId int64, creditoMinimo, notaMinima int) (*Convocatorias, int, error) {
	credMin := creditoMinimo
	notaMin := notaMinima
	m := NewSubmissions(id, &fechaInicio, &fechaFin, nombre, userId, &credMin, &notaMin)

	if id == 0 {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return m, 15, fmt.Errorf("id is required")
	}

	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update Convocatorias :", err)
		return m, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteSubmissions(id int64) (int, error) {
	if id == 0 {
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

func (s *service) GetSubmissionsByID(id int64) (*Convocatorias, int, error) {
	if id == 0 {
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

func (s *service) GetAllSubmissions() ([]*Convocatorias, error) {
	return s.repository.getAll()
}

func (s *service) GetAllSubmissionsByService(id int64) ([]*Convocatorias, error) {
	return s.repository.getAllByService(id)
}

func (s *service) GetActiveSubmissions() (*Convocatorias, int, error) {
	m, err := s.repository.getActive()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}
