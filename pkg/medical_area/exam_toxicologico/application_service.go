package exam_toxicologico

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"errors"
	"fmt"
)

type PortsServerRegistroToxicologico interface {
	CreateRegistroToxicologico(alumnoID int64, convocatoriaID int64, estado string, comentario *string, idUsuario *int64) (*RegistroToxicologico, int, error)
	UpdateRegistroToxicologico(id int64, estado string, comentario *string, idUsuario *int64) (*RegistroToxicologico, int, error)
	DeleteRegistroToxicologico(id int64) (int, error)
	GetRegistroToxicologicoByID(id int64) (*RegistroToxicologico, int, error)
	GetRegistroToxicologicoByAlumnoAndConvocatoria(alumnoID int64, convocatoriaID int64) (*RegistroToxicologico, int, error)
	GetAllRegistrosToxicologicos() ([]*RegistroToxicologico, error)
	GetEstadosByConvocatoria(convocatoriaID int64) ([]*EstadoToxicologicoConvocatoria, error)
}

type service struct {
	repository ServicesRegistroToxicologicoRepository
	user       *models.User
	txID       string
}

func NewRegistroToxicologicoService(repository ServicesRegistroToxicologicoRepository, user *models.User, TxID string) PortsServerRegistroToxicologico {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateRegistroToxicologico(alumnoID int64, convocatoriaID int64, estado string, comentario *string, idUsuario *int64) (*RegistroToxicologico, int, error) {

	exists, err := s.repository.existsByAlumnoAndConvocatoria(alumnoID, convocatoriaID)
	if err != nil {
		logger.Error.Println(s.txID, " - error checking if registro exists:", err)
		return nil, 3, err
	}
	if exists {
		logger.Error.Println(s.txID, " - registro toxicológico already exists for this alumno and convocatoria")
		return nil, 16, fmt.Errorf("registro toxicológico already exists for this alumno and convocatoria")
	}

	m := NewRegistroToxicologico(alumnoID, convocatoriaID, estado, comentario, idUsuario)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}

	if err := s.repository.create(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't create registro toxicológico:", err)
		return m, 3, err
	}

	return m, 29, nil
}

func (s *service) UpdateRegistroToxicologico(id int64, estado string, comentario *string, idUsuario *int64) (*RegistroToxicologico, int, error) {
	existing, err := s.repository.getByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - error getting registro:", err)
		return nil, 3, err
	}
	if existing == nil {
		logger.Error.Println(s.txID, " - registro toxicológico not found")
		return nil, 22, errors.New("registro toxicológico not found")
		}	
	m := UpdateRegistroToxicologico(id, estado, comentario, idUsuario)

	if valid, err := m.ValidForUpdate(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	m.AlumnoID = existing.AlumnoID
	m.ConvocatoriaID = existing.ConvocatoriaID
	m.FechaCreacion = existing.FechaCreacion

	if err := s.repository.update(m); err != nil {
		logger.Error.Println(s.txID, " - couldn't update registro toxicológico:", err)
		return m, 18, err
	}

	updated, err := s.repository.getByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - error getting updated registro:", err)
		return m, 3, err
	}

	return updated, 29, nil
}

func (s *service) DeleteRegistroToxicologico(id int64) (int, error) {
	existing, err := s.repository.getByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - error getting registro:", err)
		return 3, err
	}
	if existing == nil {
		logger.Error.Println(s.txID, " - registro toxicológico not found")
		return 22, errors.New("registro toxicológico not found")
	}

	if err := s.repository.delete(id); err != nil {
		logger.Error.Println(s.txID, " - couldn't delete registro toxicológico:", err)
		return 1, err
	}

	return 28, nil
}

func (s *service) GetRegistroToxicologicoByID(id int64) (*RegistroToxicologico, int, error) {
	m, err := s.repository.getByID(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get registro toxicológico by ID:", err)
		return nil, 22, err
	}
	if m == nil {
		logger.Error.Println(s.txID, " - registro toxicológico not found")
		return nil, 22, errors.New("registro toxicológico not found")
	}

	return m, 29, nil
}

func (s *service) GetRegistroToxicologicoByAlumnoAndConvocatoria(alumnoID int64, convocatoriaID int64) (*RegistroToxicologico, int, error) {
	m, err := s.repository.getByAlumnoAndConvocatoria(alumnoID, convocatoriaID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get registro toxicológico by alumno and convocatoria:", err)
		return nil, 22, err
	}
	if m == nil {
		return nil, 22, nil
	}

	return m, 29, nil
}

func (s *service) GetAllRegistrosToxicologicos() ([]*RegistroToxicologico, error) {
	registros, err := s.repository.getAll()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get all registros toxicológicos:", err)
		return nil, err
	}

	return registros, nil
}

func (s *service) GetEstadosByConvocatoria(convocatoriaID int64) ([]*EstadoToxicologicoConvocatoria, error) {
	estados, err := s.repository.getEstadosByConvocatoria(convocatoriaID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get estados by convocatoria:", err)
		return nil, err
	}

	return estados, nil
}
