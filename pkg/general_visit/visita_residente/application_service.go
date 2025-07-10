package visita_residente

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"strconv"
)

type PortsServerVisitaResidente interface {
	CreateVisitaResidente(alumnoID, estado, comentario, imagenURL string) (*VisitaDomiciliaria, int, error)
	GetAllVisitaResidente() ([]*VisitaDomiciliaria, error)
	GetVisitaResidenteByID(id string) (*VisitaDomiciliaria, int, error)
	UpdateVisitaResidente(id, alumnoID, estado, comentario, imagenURL string) (*VisitaDomiciliaria, int, error)
	DeleteVisitaResidente(id string) (int, error)
	GetAlumnosPendientesVisita(convocatoriaID *string) ([]*AlumnoPendienteVisita, error)
	GetEstadisticasVisitas(filtros *FiltrosVisita) (*EstadisticasVisita, error)
	GetEstadisticasPorEscuelaProfesional(convocatoriaID *uint64) ([]*EstadisticasPorEscuelaProfesional, error)
	GetEstadisticasPorLugarProcedencia(convocatoriaID *uint64) ([]*EstadisticasPorLugarProcedencia, error)
	GetAlumnosPendientesPorDepartamento(convocatoriaID uint64, departamento string) ([]*AlumnoPendienteVisitaPorDepartamento, error)
	GetTodosAlumnosPorConvocatoria(convocatoriaID uint64) ([]*AlumnoCompleto, error)
	ExistsByAlumnoID(alumnoID string) (bool, error)
}

type service struct {
	repository ServiceVisitaResidenteRepository
	user       *models.User
	txID       string
}

func NewVisitaResidenteService(repository ServiceVisitaResidenteRepository, user *models.User, TxID string) PortsServerVisitaResidente {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateVisitaResidente(alumnoID, estado, comentario, imagenURL string) (*VisitaDomiciliaria, int, error) {
	alumnoIDUint64, err := strconv.ParseUint(alumnoID, 10, 64)
	if err != nil {
		logger.Error.Println(s.txID, " - invalid alumno ID:", err)
		return nil, 15, fmt.Errorf("invalid alumno ID")
	}

	exists, err := s.repository.existsByAlumnoID(alumnoIDUint64)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't check if alumno has visit:", err)
		return nil, 3, err
	}

	if exists {
		logger.Error.Println(s.txID, " - alumno already has a visit registered")
		return nil, 15, fmt.Errorf("el alumno ya tiene una visita domiciliaria registrada")
	}

	var comentarioPtr *string
	if comentario != "" {
		comentarioPtr = &comentario
	}
	
	var imagenURLPtr *string
	if imagenURL != "" {
		imagenURLPtr = &imagenURL
	}

	var userIDPtr *uint64
	if s.user != nil && s.user.ID > 0 {
		userID := uint64(s.user.ID)
		userIDPtr = &userID
	}

	m := NewVisitaDomiciliaria(alumnoIDUint64, estado, comentarioPtr, imagenURLPtr, userIDPtr)
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
		logger.Error.Println(s.txID, " - couldn't create VisitaResidente :", err)
		return nil, 3, err
	}
	return m, 29, nil
}

func (s *service) GetAllVisitaResidente() ([]*VisitaDomiciliaria, error) {
	m, err := s.repository.getAll()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get all VisitaResidente :", err)
		return nil, err
	}
	return m, nil
}

func (s *service) GetVisitaResidenteByID(id string) (*VisitaDomiciliaria, int, error) {
	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error.Println(s.txID, " - invalid ID:", err)
		return nil, 15, fmt.Errorf("invalid ID")
	}

	m, err := s.repository.getByID(idUint64)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't get VisitaResidente :", err)
		return nil, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateVisitaResidente(id, alumnoID, estado, comentario, imagenURL string) (*VisitaDomiciliaria, int, error) {
	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error.Println(s.txID, " - invalid ID:", err)
		return nil, 15, fmt.Errorf("invalid ID")
	}

	alumnoIDUint64, err := strconv.ParseUint(alumnoID, 10, 64)
	if err != nil {
		logger.Error.Println(s.txID, " - invalid alumno ID:", err)
		return nil, 15, fmt.Errorf("invalid alumno ID")
	}

	var comentarioPtr *string
	if comentario != "" {
		comentarioPtr = &comentario
	}
	
	var imagenURLPtr *string
	if imagenURL != "" {
		imagenURLPtr = &imagenURL
	}

	var userIDPtr *uint64
	if s.user != nil && s.user.ID > 0 {
		userID := uint64(s.user.ID)
		userIDPtr = &userID
	}

	m := NewVisitaDomiciliaria(alumnoIDUint64, estado, comentarioPtr, imagenURLPtr, userIDPtr)
	m.ID = idUint64 
	
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
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't update VisitaResidente :", err)
		return nil, 3, err
	}
	return m, 29, nil
}

func (s *service) DeleteVisitaResidente(id string) (int, error) {
	idUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error.Println(s.txID, " - invalid ID:", err)
		return 15, fmt.Errorf("invalid ID")
	}

	if err := s.repository.delete(idUint64); err != nil {
		if err.Error() == "rows affected error" {
			return 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't delete VisitaResidente :", err)
		return 3, err
	}
	return 29, nil
}

func (s *service) GetAlumnosPendientesVisita(convocatoriaID *string) ([]*AlumnoPendienteVisita, error) {
	if convocatoriaID != nil && *convocatoriaID != "" {
		convocatoriaIDUint64, err := strconv.ParseUint(*convocatoriaID, 10, 64)
		if err != nil {
			logger.Error.Println(s.txID, " - invalid convocatoria ID:", err)
			return nil, fmt.Errorf("invalid convocatoria ID")
		}
		
		alumnos, err := s.repository.getAlumnosPendientesVisitaPorConvocatoria(convocatoriaIDUint64)
		if err != nil {
			logger.Error.Println(s.txID, " - couldn't get alumnos pendientes por convocatoria:", err)
			return nil, err
		}
		return alumnos, nil
	}

	alumnos, err := s.repository.getTodosAlumnosPendientesVisita()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get todos los alumnos pendientes:", err)
		return nil, err
	}
	return alumnos, nil
}

func (s *service) GetEstadisticasVisitas(filtros *FiltrosVisita) (*EstadisticasVisita, error) {
	var convocatoriaID *uint64
	if filtros != nil && filtros.ConvocatoriaID != nil {
		convID := uint64(*filtros.ConvocatoriaID)
		convocatoriaID = &convID
	}
	
	stats, err := s.repository.getEstadisticas(convocatoriaID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get estadisticas:", err)
		return nil, err
	}
	return stats, nil
}

func (s *service) GetEstadisticasPorEscuelaProfesional(convocatoriaID *uint64) ([]*EstadisticasPorEscuelaProfesional, error) {
	estadisticas, err := s.repository.getEstadisticasPorEscuelaProfesional(convocatoriaID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get estadisticas por facultad:", err)
		return nil, err
	}
	return estadisticas, nil
}

func (s *service) GetEstadisticasPorLugarProcedencia(convocatoriaID *uint64) ([]*EstadisticasPorLugarProcedencia, error) {
	estadisticas, err := s.repository.getEstadisticasPorLugarProcedencia(convocatoriaID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get estadisticas por lugar de procedencia:", err)
		return nil, err
	}
	return estadisticas, nil
}

func (s *service) GetAlumnosPendientesPorDepartamento(convocatoriaID uint64, departamento string) ([]*AlumnoPendienteVisitaPorDepartamento, error) {
	alumnos, err := s.repository.getAlumnosPendientesPorDepartamento(convocatoriaID, departamento)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get alumnos pendientes por departamento:", err)
		return nil, err
	}
	return alumnos, nil
}

func (s *service) GetTodosAlumnosPorConvocatoria(convocatoriaID uint64) ([]*AlumnoCompleto, error) {
	alumnos, err := s.repository.getTodosAlumnosPorConvocatoria(convocatoriaID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get todos los alumnos por convocatoria:", err)
		return nil, err
	}
	return alumnos, nil
}

func (s *service) ExistsByAlumnoID(alumnoID string) (bool, error) {
	alumnoIDUint64, err := strconv.ParseUint(alumnoID, 10, 64)
	if err != nil {
		logger.Error.Println(s.txID, " - invalid alumno ID:", err)
		return false, fmt.Errorf("invalid alumno ID")
	}

	exists, err := s.repository.existsByAlumnoID(alumnoIDUint64)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't check if alumno has visit:", err)
		return false, err
	}
	return exists, nil
}