package visita_general

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"

	"github.com/google/uuid"
)

type PortsServerVisitaGeneral interface {
	CreateVisitaGeneral(id, tipoUsuario string, codigoEstudiante, dni *string, nombreCompleto string, genero *string, edad *int, escuela, area *string, motivoAtencion, descripcionMotivo string, urlImagen, departamento, provincia, distrito *string, lugarAtencion string) (*VisitaGeneral, int, error)
	GetAllVisitaGeneral() ([]*VisitaGeneral, error)
	GetVisitaGeneralByID(id string) (*VisitaGeneral, int, error)
	UpdateVisitaGeneral(id, tipoUsuario string, codigoEstudiante, dni *string, nombreCompleto string, genero *string, edad *int, escuela, area *string, motivoAtencion, descripcionMotivo string, urlImagen, departamento, provincia, distrito *string, lugarAtencion string) (*VisitaGeneral, int, error)
	DeleteVisitaGeneral(id string) (int, error)
	GetAllDepartments() ([]*Departamento, error)
	GetProvincesByDepartment(departmentID string) ([]*Provincia, error)
	GetDistrictsByProvince(provinceID string) ([]*Distrito, error)
	GetLocationHierarchy() (*LocationResponse, error)
}

type service struct {
	repository ServiceVisitaGeneralRepository
	user       *models.User
	txID       string
}

func NewVisitaGeneralService(repository ServiceVisitaGeneralRepository, user *models.User, TxID string) PortsServerVisitaGeneral {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateVisitaGeneral(id, tipoUsuario string, codigoEstudiante, dni *string, nombreCompleto string, genero *string, edad *int, escuela, area *string, motivoAtencion, descripcionMotivo string, urlImagen, departamento, provincia, distrito *string, lugarAtencion string) (*VisitaGeneral, int, error) {
	if id == "" {
		id = uuid.New().String()
	}

	userID := &s.user.ID
	m := NewVisitaGeneral(id, tipoUsuario, codigoEstudiante, dni, nombreCompleto, genero, edad, escuela, area, motivoAtencion, descripcionMotivo, urlImagen, departamento, provincia, distrito, lugarAtencion, userID)
	valid, err := m.valid()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't meet validations:", err)
		return nil, 15, err
	}
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:")
		return nil, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create VisitaGeneral :", err)
		return nil, 3, err
	}
	return m, 210, nil
}

func (s *service) GetAllVisitaGeneral() ([]*VisitaGeneral, error) {
	ms, err := s.repository.getAll()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get all VisitaGeneral:", err)
		return nil, fmt.Errorf("couldn't get all VisitaGeneral: %v", err)
	}
	return ms, nil
}

func (s *service) GetVisitaGeneralByID(id string) (*VisitaGeneral, int, error) {
	if id == "" {
		logger.Error.Println(s.txID, " - ID must not be empty")
		return nil, 2, fmt.Errorf("ID must not be empty")
	}
	
	m, err := s.repository.getByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			logger.Error.Println(s.txID, " - record not found")
			return nil, 108, fmt.Errorf("record not found")
		}
		logger.Error.Println(s.txID, " - couldn't get VisitaGeneral by ID:", err)
		return nil, 3, err
	}
	return m, 210, nil
}

func (s *service) UpdateVisitaGeneral(id, tipoUsuario string, codigoEstudiante, dni *string, nombreCompleto string, genero *string, edad *int, escuela, area *string, motivoAtencion, descripcionMotivo string, urlImagen, departamento, provincia, distrito *string, lugarAtencion string) (*VisitaGeneral, int, error) {
	userID := &s.user.ID
	m := NewVisitaGeneral(id, tipoUsuario, codigoEstudiante, dni, nombreCompleto, genero, edad, escuela, area, motivoAtencion, descripcionMotivo, urlImagen, departamento, provincia, distrito, lugarAtencion, userID)
	valid, err := m.valid()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't meet validations:", err)
		return nil, 15, err
	}
	if !valid {
		logger.Error.Println(s.txID, " - don't meet validations:")
		return nil, 15, err
	}
	if err := s.repository.update(m); err != nil {
		if err.Error() == "rows affected error" {
			return nil, 108, fmt.Errorf("record not found")
		}
		logger.Error.Println(s.txID, " - couldn't update VisitaGeneral :", err)
		return nil, 3, err
	}
	return m, 210, nil
}

func (s *service) DeleteVisitaGeneral(id string) (int, error) {
	if err := s.repository.delete(id); err != nil {
		if err.Error() == "rows affected error" {
			return 108, fmt.Errorf("record not found")
		}
		logger.Error.Println(s.txID, " - couldn't delete VisitaGeneral :", err)
		return 3, err
	}
	return 210, nil
}

func (s *service) GetAllDepartments() ([]*Departamento, error) {
	departments, err := s.repository.getAllDepartments()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get departments:", err)
		return nil, err
	}
	return departments, nil
}

func (s *service) GetProvincesByDepartment(departmentID string) ([]*Provincia, error) {
	if departmentID == "" {
		return nil, fmt.Errorf("department ID is required")
	}
	
	provinces, err := s.repository.getProvincesByDepartment(departmentID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get provinces for department:", departmentID, err)
		return nil, err
	}
	return provinces, nil
}

func (s *service) GetDistrictsByProvince(provinceID string) ([]*Distrito, error) {
	if provinceID == "" {
		return nil, fmt.Errorf("province ID is required")
	}
	
	districts, err := s.repository.getDistrictsByProvince(provinceID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get districts for province:", provinceID, err)
		return nil, err
	}
	return districts, nil
}

func (s *service) GetLocationHierarchy() (*LocationResponse, error) {
	hierarchy, err := s.repository.getLocationHierarchy()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get location hierarchy:", err)
		return nil, err
	}
	return hierarchy, nil
}
