package consultation_integral_attention

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"fmt"
	"github.com/google/uuid"
)

type PortsServerConsultationIntegralAttention interface {
	CreateConsultationIntegralAttention(id, consulta_id, fecha, hora, edad, motivo_consulta, tiempo_enfermedad, apetito, sed, suenio, estado_animo, orina, deposiciones, temperatura, presion_arterial, frecuencia_cardiaca, frecuencia_respiratoria, peso, talla, indice_masa_corporal, diagnostico, tratamiento, examenes_axuliares, referencia, observacion, numero_recibo, costo, fecha_pago string) (*ConsultationIntegralAttention, int, error)
	UpdateConsultationIntegralAttention(id, consulta_id, fecha, hora, edad, motivo_consulta, tiempo_enfermedad, apetito, sed, suenio, estado_animo, orina, deposiciones, temperatura, presion_arterial, frecuencia_cardiaca, frecuencia_respiratoria, peso, talla, indice_masa_corporal, diagnostico, tratamiento, examenes_axuliares, referencia, observacion, numero_recibo, costo, fecha_pago string) (*ConsultationIntegralAttention, int, error)
	DeleteConsultationIntegralAttention(id string) (int, error)
	DeleteConsultationIntegralAttentionByIDConsultation(id string) (int, error)
	GetConsultationIntegralAttentionByID(id string) (*ConsultationIntegralAttention, int, error)
	GetConsultationIntegralAttentionByIDConsultation(id string) (*ConsultationIntegralAttention, int, error)
	GetAllConsultationIntegralAttention() ([]*ConsultationIntegralAttention, error)
	GetConsultationIntegralAttentionExcel(fecha_inicio, fecha_fin string) ([]*models.ConsultationIntegralAttentionExcel, error)
}

type service struct {
	repository ServicesConsultationIntegralAttentionRepository
	user       *models.User
	txID       string
}

func NewConsultationIntegralAttentionService(repository ServicesConsultationIntegralAttentionRepository, user *models.User, TxID string) PortsServerConsultationIntegralAttention {
	return &service{repository: repository, user: user, txID: TxID}
}

func (s *service) CreateConsultationIntegralAttention(id, consulta_id, fecha, hora, edad, motivo_consulta, tiempo_enfermedad, apetito, sed, suenio, estado_animo, orina, deposiciones, temperatura, presion_arterial, frecuencia_cardiaca, frecuencia_respiratoria, peso, talla, indice_masa_corporal, diagnostico, tratamiento, examenes_axuliares, referencia, observacion, numero_recibo, costo, fecha_pago string) (*ConsultationIntegralAttention, int, error) {

	m := NewConsultationIntegralAttention(id, consulta_id, fecha, hora, edad, motivo_consulta, tiempo_enfermedad, apetito, sed, suenio, estado_animo, orina, deposiciones, temperatura, presion_arterial, frecuencia_cardiaca, frecuencia_respiratoria, peso, talla, indice_masa_corporal, diagnostico, tratamiento, examenes_axuliares, referencia, observacion, numero_recibo, costo, fecha_pago)
	if valid, err := m.valid(); !valid {
		logger.Error.Println(s.txID, " - don't meet validations:", err)
		return m, 15, err
	}
	if err := s.repository.create(m); err != nil {
		if err.Error() == "rows affected error" {
			return m, 108, nil
		}
		logger.Error.Println(s.txID, " - couldn't create consultation integral attention :", err)
		return m, 3, err
	}
	return m, 29, nil
}

func (s *service) UpdateConsultationIntegralAttention(id, consulta_id, fecha, hora, edad, motivo_consulta, tiempo_enfermedad, apetito, sed, suenio, estado_animo, orina, deposiciones, temperatura, presion_arterial, frecuencia_cardiaca, frecuencia_respiratoria, peso, talla, indice_masa_corporal, diagnostico, tratamiento, examenes_axuliares, referencia, observacion, numero_recibo, costo, fecha_pago string) (*ConsultationIntegralAttention, int, error) {
	m := NewConsultationIntegralAttention(id, consulta_id, fecha, hora, edad, motivo_consulta, tiempo_enfermedad, apetito, sed, suenio, estado_animo, orina, deposiciones, temperatura, presion_arterial, frecuencia_cardiaca, frecuencia_respiratoria, peso, talla, indice_masa_corporal, diagnostico, tratamiento, examenes_axuliares, referencia, observacion, numero_recibo, costo, fecha_pago)
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
		logger.Error.Println(s.txID, " - couldn't update consultation integral attention :", err)
		return nil, 18, err
	}
	return m, 29, nil
}

func (s *service) DeleteConsultationIntegralAttention(id string) (int, error) {
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

func (s *service) DeleteConsultationIntegralAttentionByIDConsultation(id string) (int, error) {
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

func (s *service) GetConsultationIntegralAttentionByID(id string) (*ConsultationIntegralAttention, int, error) {
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

func (s *service) GetConsultationIntegralAttentionByIDConsultation(id string) (*ConsultationIntegralAttention, int, error) {
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

func (s *service) GetAllConsultationIntegralAttention() ([]*ConsultationIntegralAttention, error) {
	return s.repository.getAll()
}

func (s *service) GetConsultationIntegralAttentionExcel(fecha_inicio, fecha_fin string) ([]*models.ConsultationIntegralAttentionExcel, error) {
	m, err := s.repository.getAllByDateExcel(fecha_inicio, fecha_fin)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get consultation integran attention excel:", err)
		return nil, err
	}
	return m, nil
}
