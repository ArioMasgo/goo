package preguntas

import (
	"fmt"
)

type PortsServerPreguntas interface {
	GetAll(page, pageSize int) ([]Pregunta, error)
	GetByID(id int) (*Pregunta, error)
	Create(p *Pregunta) error
	Update(id int, p *Pregunta) error
	Delete(id int) error
}

type service struct {
	repository ServicesPreguntaRepository
	txID       string
}

func NewPreguntaService(repository ServicesPreguntaRepository, txID string) PortsServerPreguntas {
	return &service{repository: repository, txID: txID}
}

func (s *service) GetAll(page, pageSize int) ([]Pregunta, error) {
	return s.repository.GetAll(page, pageSize)
}

func (s *service) GetByID(id int) (*Pregunta, error) {
	return s.repository.GetByID(id)
}

func (s *service) Create(p *Pregunta) error {
	if p.TextoPregunta == "" || p.Type == "" {
		return fmt.Errorf("texto_pregunta y type son obligatorios")
	}
	return s.repository.Create(p)
}

func (s *service) Update(id int, p *Pregunta) error {
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
