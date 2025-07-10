package dentistry_consultation_buccal_procedure

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerBuccalProcedure interface {
	CreateBuccalProcedure(id, consulta_odontologia_id, tipo_procedimiento, recibo, costo, fecha_pago, pieza_dental, comentarios string) (*BuccalProcedure, int, error)
	UpdateBuccalProcedure(id, consulta_odontologia_id, tipo_procedimiento, recibo, costo, fecha_pago, pieza_dental, comentarios string) (*BuccalProcedure, int, error)
	DeleteBuccalProcedure(id string) (int, error)
	DeleteBuccalProcedureByIDConsultation(id string) (int, error)
	GetBuccalProcedureByID(id string) (*BuccalProcedure, int, error)
	GetBuccalProcedureByIDConsultation(id string) (*BuccalProcedure, int, error)
	GetAllBuccalProcedure() ([]*BuccalProcedure, error)
	GetBuccalProceduresExcel(fecha_inicio, fecha_fin string) ([]*models.PerformedProceduresExcel, error)
}

type service struct {
	repository ServicesBuccalProcedureRepository
	user       *models.User
	txID       string
}

func NewBuccalProcedureService(repository ServicesBuccalProcedureRepository, user *models.User, TxID string) PortsServerBuccalProcedure {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateBuccalProcedure(id, consulta_odontologia_id, tipo_procedimiento, recibo, costo, fecha_pago, pieza_dental, comentarios string) (*BuccalProcedure, int, error) {

	//TODO: RECIBO
	/*if recibo != "" {
		exists, err := s.repository.existsByReceipt(recibo)
		if err != nil {
			logger.Error.Println(s.txID, " - error checking if buccal procedure exists by receipt:", err)
			return nil, 3, err
		}
		if exists {
			logger.Error.Println(s.txID, " - buccal procedure with receipt already exists")
			return nil, 16, fmt.Errorf("buccal procedure with receipt already exists")
		}
	}*/
	m := NewBuccalProcedure(id, consulta_odontologia_id, tipo_procedimiento, recibo, costo, fecha_pago, pieza_dental, comentarios)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create buccal procedure :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateBuccalProcedure(id, consulta_odontologia_id, tipo_procedimiento, recibo, costo, fecha_pago, pieza_dental, comentarios string) (*BuccalProcedure, int, error) {
	m := NewBuccalProcedure(id, consulta_odontologia_id, tipo_procedimiento, recibo, costo, fecha_pago, pieza_dental, comentarios)
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
		logger.Error.Println(s.txID, " - couldn't update buccal procedure :", err)
		return nil, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteBuccalProcedure(id string) (int, error) {
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

func (s *service) DeleteBuccalProcedureByIDConsultation(id string) (int, error) {
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

func (s *service) GetBuccalProcedureByID(id string) (*BuccalProcedure, int, error) {
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

func (s *service) GetBuccalProcedureByIDConsultation(id string) (*BuccalProcedure, int, error) {
	if err := uuid.Validate(id); err != nil {
		logger.Error.Println(s.txID, " - don't meet validations:", fmt.Errorf("id is required"))
		return nil, 15, fmt.Errorf("id is required")
	}
	m, err := s.repository.getByIDConsultation(id)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, 22, err
	}
	return m, 29, nil
}

func (s *service) GetAllBuccalProcedure() ([]*BuccalProcedure, error) {
	return s.repository.getAll()
}

func (s *service) GetBuccalProceduresExcel(fecha_inicio, fecha_fin string) ([]*models.PerformedProceduresExcel, error) {
	m, err := s.repository.getBuccalProceduresExcel(fecha_inicio, fecha_fin)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn`t getByID row:", err)
		return nil, err
	}
	return m, nil
}
