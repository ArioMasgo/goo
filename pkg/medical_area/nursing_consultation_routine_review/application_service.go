package nursing_consultation_routine_review

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerRoutineReview interface {
	CreateRoutineReview(id, consulta_enfermeria_id, fiebre_ultimo_quince_dias, tos_mas_quince_dias, secrecion_lesion_genitales, fecha_ultima_regla, comentarios string) (*RoutineReview, int, error)
	UpdateRoutineReview(id, consulta_enfermeria_id, fiebre_ultimo_quince_dias, tos_mas_quince_dias, secrecion_lesion_genitales, fecha_ultima_regla, comentarios string) (*RoutineReview, int, error)
	DeleteRoutineReview(id string) (int, error)
	DeleteRoutineReviewByIDConsultation(id string) (int, error)
	GetRoutineReviewByID(id string) (*RoutineReview, int, error)
	GetAllRoutineReview() ([]*RoutineReview, error)
}

type service struct {
	repository ServicesRoutineReviewRepository
	user       *models.User
	txID       string
}

func NewRoutineReviewService(repository ServicesRoutineReviewRepository, user *models.User, TxID string) PortsServerRoutineReview {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateRoutineReview(id, consulta_enfermeria_id, fiebre_ultimo_quince_dias, tos_mas_quince_dias, secrecion_lesion_genitales, fecha_ultima_regla, comentarios string) (*RoutineReview, int, error) {

	m := NewRoutineReview(id, consulta_enfermeria_id, fiebre_ultimo_quince_dias, tos_mas_quince_dias, secrecion_lesion_genitales, fecha_ultima_regla, comentarios)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create routine review :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateRoutineReview(id, consulta_enfermeria_id, fiebre_ultimo_quince_dias, tos_mas_quince_dias, secrecion_lesion_genitales, fecha_ultima_regla, comentarios string) (*RoutineReview, int, error) {
	m := NewRoutineReview(id, consulta_enfermeria_id, fiebre_ultimo_quince_dias, tos_mas_quince_dias, secrecion_lesion_genitales, fecha_ultima_regla, comentarios)
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

func (s *service) DeleteRoutineReview(id string) (int, error) {
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

func (s *service) DeleteRoutineReviewByIDConsultation(id string) (int, error) {
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

func (s *service) GetRoutineReviewByID(id string) (*RoutineReview, int, error) {
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

func (s *service) GetAllRoutineReview() ([]*RoutineReview, error) {
	return s.repository.getAll()
}
