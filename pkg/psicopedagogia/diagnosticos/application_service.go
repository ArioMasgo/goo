package diagnosticos

import (
	"fmt"
)

type PortsServerDiagnostico interface {
	GetByID(id int) (*Diagnostico, error)
	Create(d *Diagnostico) (int64, error)
	Update(d *Diagnostico) error
	Delete(id int) error
	GetAll() ([]*Diagnostico, error)
}

type service struct {
	repository ServicesDiagnosticoRepository
	txID       string
}

func NewDiagnosticoService(repository ServicesDiagnosticoRepository, txID string) PortsServerDiagnostico {
	return &service{repository: repository, txID: txID}
}

func (s *service) GetByID(id int) (*Diagnostico, error) {
	if id == 0 {
		return nil, fmt.Errorf("id is required")
	}
	d, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (s *service) GetAll() ([]*Diagnostico, error) {
	d, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (s *service) Create(d *Diagnostico) (int64, error) {
	idDiagnostico, err := s.repository.Create(d)
	if err != nil {
		return 0, err
	}
	return idDiagnostico, nil
}

func (s *service) Update(d *Diagnostico) error {
	err := s.repository.Update(d)
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
