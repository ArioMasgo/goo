package patients

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"regexp"
)

type PortsServerPatients interface {
	CreatePatients(id string, codigo_sga string, dni string, nombres string, apellidos string, sexo string, edad string, estado_civil string, grupo_sanguineo string, fecha_nacimiento string, lugar_nacimiento string, procedencia string, escuela_profesional string, ocupacion string, correo_electronico string, numero_celular string, direccion string, tipo_persona string, factor_rh string, alergias string, ram bool) (*Patients, int, error)
	UpdatePatients(patient *Patients) (int, error)
	DeletePatients(id string) (int, error)
	GetPatientsByID(id string) (*Patients, int, error)
	GetPatientsByDNI(dni string) (*Patients, int, error)
	GetAllPatients() ([]*Patients, error)
	CountPatients(dni, names, surnames string) (int64, error)
	SearchPaginationPatients(dni, names, surnames string, limit, offset int64) ([]*Patients, error)
}

type service struct {
	repository ServicesPatientsRepository
	user       *models.User
	txID       string
}

func NewPatientsService(repository ServicesPatientsRepository, user *models.User, TxID string) PortsServerPatients {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreatePatients(id string, codigo_sga string, dni string, nombres string, apellidos string, sexo string, edad string, estado_civil string, grupo_sanguineo string, fecha_nacimiento string, lugar_nacimiento string, procedencia string, escuela_profesional string, ocupacion string, correo_electronico string, numero_celular string, direccion string, tipo_persona string, factor_rh string, alergias string, ram bool) (*Patients, int, error) {

	exists, err := s.repository.existsByDNI(dni)
	if err != nil {
		logger.Error.Println(s.txID, " - error checking if patient exists by DNI:", err)
		return nil, 3, err
	}
	if exists {
		logger.Error.Println(s.txID, " - patient with DNI already exists")
		return nil, 16, fmt.Errorf("patient with DNI already exists")
	}

	m := NewPatients(id, codigo_sga, dni, nombres, apellidos, sexo, edad, estado_civil, grupo_sanguineo, fecha_nacimiento, lugar_nacimiento, procedencia, escuela_profesional, ocupacion, correo_electronico, numero_celular, direccion, tipo_persona, factor_rh, alergias, ram)
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

func (s *service) UpdatePatients(patient *Patients) (int, error) {
	valid, err := patient.valid()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't meet validations:", err)
		return 15, err
	}
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return 15, err
	}
	if err := s.repository.update(patient); err != nil {
		logger.Error.Println(s.txID, " - couldn't update patient :", err)
		return 18, err
	}
	return 29, nil
}

func (s *service) DeletePatients(id string) (int, error) {
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

func (s *service) GetPatientsByID(id string) (*Patients, int, error) {
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

func (s *service) GetPatientsByDNI(dni string) (*Patients, int, error) {
	matched, err := regexp.MatchString(`^\d{8}$`, dni)
	if err != nil || !matched {
		logger.Error.Println(s.txID, " - don't meet validations:", errors.New("dni must be a string of 8 digits"))
		return nil, 15, errors.New("dni must be a string of 8 digits")
	}

	m, err := s.repository.getByDNI(dni)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByDNI row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllPatients() ([]*Patients, error) {
	return s.repository.getAll()
}

func (s *service) CountPatients(dni, names, surnames string) (int64, error) {
	return s.repository.countPaginationPatients(dni, names, surnames)
}

func (s *service) SearchPaginationPatients(dni, names, surnames string, limit, offset int64) ([]*Patients, error) {
	return s.repository.searchPaginationPatients(dni, names, surnames, limit, offset)
}
