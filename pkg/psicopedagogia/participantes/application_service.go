package participantes

import (
	"fmt"
)

type PortsServerParticipantes interface {
	GetAll(page, pageSize int) ([]Participante, int, error)
	GetByID(id int) (*Participante, error)
	Create(p *Participante) error
	Update(id int, p *Participante) error
	Delete(id int) error
	SearchParticipants(criteria map[string]interface{}) ([]Participante, int, error)
}

type service struct {
	repository ServicesParticipanteRepository
	txID       string
}

func NewParticipanteService(repository ServicesParticipanteRepository, txID string) PortsServerParticipantes {
	return &service{repository: repository, txID: txID}
}

func (s *service) GetAll(page, pageSize int) ([]Participante, int, error) {
	return s.repository.GetAll(page, pageSize)
}

func (s *service) GetByID(id int) (*Participante, error) {
	return s.repository.GetByID(id)
}

func (s *service) Create(p *Participante) error {
	if p.Nombre == "" || p.Apellido == "" {
		return fmt.Errorf("nombre, apellido son obligatorios")
	}
	partID, err := s.repository.Create(p)
	fmt.Print(partID)
	return err
}

func (s *service) Update(id int, p *Participante) error {
	if id == 0 {
		return fmt.Errorf("id es requerido")
	}
	return s.repository.Update(id, p)
}

func (s *service) Delete(id int) error {
	if id == 0 {
		return fmt.Errorf("id es requerido")
	}
	return s.repository.Delete(id)
}

func (s *service) SearchParticipants(criteria map[string]interface{}) ([]Participante, int, error) {
	result, err := s.repository.SearchParticipants(criteria)
	if err != nil {
		return nil, 500, err
	}
	return result, 200, nil
}
