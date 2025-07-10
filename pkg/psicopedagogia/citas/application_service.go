package citas

import (
	"fmt"
)

type PortsServerCita interface {
	GetByID(id int) (*Cita, error)
	Create(c *Cita) (int64, error)
	Update(c *Cita) error
	Delete(id int) error
	GetAll() ([]*Cita, error)
}

type service struct {
	repository ServicesCitaRepository
	txID       string
}

func NewCitaService(repository ServicesCitaRepository, txID string) PortsServerCita {
	return &service{repository: repository, txID: txID}
}

func (s *service) GetByID(id int) (*Cita, error) {
	if id == 0 {
		return nil, fmt.Errorf("id is required")
	}
	c, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *service) GetAll() ([]*Cita, error) {
	c, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *service) Create(c *Cita) (int64, error) {
	dateAlreadyRegistered, err := s.repository.ExisteCitaEnFecha(c.FechaInicio, c.DNI)
	if err != nil {
		return 0, err
	}
	if dateAlreadyRegistered {
		return 0, fmt.Errorf("ya existe una cita en el rango de fecha proporcionado")
	}

	idCita, err := s.repository.Create(c)
	if err != nil {
		return 0, err
	}
	return idCita, nil
}

func (s *service) Update(c *Cita) error {
	err := s.repository.Update(c)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Delete(id int) error {
	if id == 0 {
		return fmt.Errorf("id is required")
	}
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
