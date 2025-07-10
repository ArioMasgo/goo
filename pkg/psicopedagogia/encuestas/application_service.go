package encuestas

import (
	"fmt"
)

type PortsServerEncuestas interface {
	GetAll(page, pageSize int) ([]Encuesta, error)
	GetByID(id int) (*Encuesta, error)
	Create(e *Encuesta) (int64, error)
	Update(id int, e *Encuesta) error
	Delete(id int) error
	HasActiveSRQ() (int, string, bool, error)
}

type service struct {
	repository ServicesEncuestaRepository
	txID       string
}

func NewEncuestaService(repository ServicesEncuestaRepository, txID string) PortsServerEncuestas {
	return &service{repository: repository, txID: txID}
}

func (s *service) GetAll(page, pageSize int) ([]Encuesta, error) {
	return s.repository.GetAll(page, pageSize)
}

func (s *service) GetByID(id int) (*Encuesta, error) {
	return s.repository.GetByID(id)
}

func (s *service) Create(e *Encuesta) (int64, error) {
	if e.Nombre == "" {
		return 0, fmt.Errorf("nombre_encuesta es obligatorio")
	}
	return s.repository.Create(e)
}

func (s *service) Update(id int, e *Encuesta) error {
	if id == 0 {
		return fmt.Errorf("id es requerido")
	}
	return s.repository.Update(id, e)
}

func (s *service) Delete(id int) error {
	if id == 0 {
		return fmt.Errorf("id es requerido")
	}
	return s.repository.Delete(id)
}

func (s *service) HasActiveSRQ() (int, string, bool, error) {
	return s.repository.GetActiveSRQ()
}
