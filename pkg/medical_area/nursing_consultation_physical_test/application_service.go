package nursing_consultation_physical_test

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerPhysicalTest interface {
	CreatePhysicalTest(id, consulta_enfermeria_id, talla_peso, perimetro_cintura, indice_masa_corporal_img, presion_arterial, comentarios string) (*PhysicalTest, int, error)
	UpdatePhysicalTest(id, consulta_enfermeria_id, talla_peso, perimetro_cintura, indice_masa_corporal_img, presion_arterial, comentarios string) (*PhysicalTest, int, error)
	DeletePhysicalTest(id string) (int, error)
	DeletePhysicalTestByIDConsultation(id string) (int, error)
	GetPhysicalTestByID(id string) (*PhysicalTest, int, error)
	GetAllPhysicalTest() ([]*PhysicalTest, error)
}

type service struct {
	repository ServicesPhysicalTestRepository
	user       *models.User
	txID       string
}

func NewPhysicalTestService(repository ServicesPhysicalTestRepository, user *models.User, TxID string) PortsServerPhysicalTest {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreatePhysicalTest(id, consulta_enfermeria_id, talla_peso, perimetro_cintura, indice_masa_corporal_img, presion_arterial, comentarios string) (*PhysicalTest, int, error) {

	m := NewPhysicalTest(id, consulta_enfermeria_id, talla_peso, perimetro_cintura, indice_masa_corporal_img, presion_arterial, comentarios)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create physical test :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdatePhysicalTest(id, consulta_enfermeria_id, talla_peso, perimetro_cintura, indice_masa_corporal_img, presion_arterial, comentarios string) (*PhysicalTest, int, error) {
	m := NewPhysicalTest(id, consulta_enfermeria_id, talla_peso, perimetro_cintura, indice_masa_corporal_img, presion_arterial, comentarios)
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
		logger.Error.Println(s.txID, " - couldn't update nursing consultation :", err)
		return nil, 18, err
	}
	return m, 29, nil
}

func (s *service) DeletePhysicalTest(id string) (int, error) {
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

func (s *service) DeletePhysicalTestByIDConsultation(id string) (int, error) {
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

func (s *service) GetPhysicalTestByID(id string) (*PhysicalTest, int, error) {
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

func (s *service) GetAllPhysicalTest() ([]*PhysicalTest, error) {
	return s.repository.getAll()
}
