package low_code_medical_area

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"dbu-api/pkg/medical_area"
	"dbu-api/pkg/medical_area/patient_background"
	"dbu-api/pkg/medical_area/patients"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type PatientService struct {
	db   *sqlx.DB
	usr  *models.User
	txID string
}

type PortsServerPatient interface {
	CreatePatientLowCode(m *models.Patients) (int, error)
	UpdatePatientLowCode(m *models.Patients) (int, error)
	DeletePatientLowCode(id string) (int, error)
	GetPatientByIdLowCode(id string) (*models.Patients, int, error)
	GetAllPatientLowCode() ([]*models.Patients, error)
	GetPatients(dni, names, surnames string, limit, offset int64) ([]*models.Patients, error)
	GetMetadata(dni, names, surnames string, limit, offset int64) (*models.Metadata, error)
	GetPatientByDNILowCode(dni string) (*models.Patients, int, error)
}

func NewPatient(db *sqlx.DB, usr *models.User, txID string) PortsServerPatient {
	return &PatientService{db: db, usr: usr, txID: txID}
}

func (s *PatientService) CreatePatientLowCode(m *models.Patients) (int, error) {
	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	dataPatient, code, err := srv.SrvPatient.CreatePatients(m.ID, m.CodigoSGA, m.DNI, m.Nombres, m.Apellidos, m.Sexo, m.Edad, m.EstadoCivil, m.GrupoSanguineo, m.FechaNacimiento, m.LugarNacimiento, m.Procedencia, m.EscuelaProfesional, m.Ocupacion, m.CorreoElectronico, m.NumeroCelular, m.Direccion, m.TipoPersona, m.FactorRH, m.Alergias, m.RAM)
	if err != nil {
		logger.Error.Printf("couldn't create patient, error: %v", dataPatient, err)
		return code, err
	}

	if m.Antecedentes != nil {
		for _, antecedente := range m.Antecedentes {
			code, err := srv.SrvPatientBackground.CreatePatientBackground(antecedente.ID, m.ID, antecedente.Nombre, antecedente.Estado)
			if err != nil {
				logger.Error.Printf("couldn't create background, error: %v", antecedente, err)
				return code, err
			}
		}
	}

	return 29, nil
}

func (s *PatientService) UpdatePatientLowCode(input *models.Patients) (int, error) {

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	patient := &patients.Patients{
		ID:                 input.ID,
		TipoPersona:        input.TipoPersona,
		CodigoSGA:          input.CodigoSGA,
		DNI:                input.DNI,
		Nombres:            input.Nombres,
		Apellidos:          input.Apellidos,
		Sexo:               input.Sexo,
		Edad:               input.Edad,
		EstadoCivil:        input.EstadoCivil,
		GrupoSanguineo:     input.GrupoSanguineo,
		FechaNacimiento:    input.FechaNacimiento,
		LugarNacimiento:    input.LugarNacimiento,
		Procedencia:        input.Procedencia,
		FactorRH:           input.FactorRH,
		RAM:                input.RAM,
		Alergias:           input.Alergias,
		EscuelaProfesional: input.EscuelaProfesional,
		Ocupacion:          input.Ocupacion,
		CorreoElectronico:  input.CorreoElectronico,
		NumeroCelular:      input.NumeroCelular,
		Direccion:          input.Direccion,
	}

	code, err := srv.SrvPatient.UpdatePatients(patient)
	if err != nil {
		logger.Error.Printf("couldn't update patient, error: %v", err)
		return code, err
	}

	antecedentesExistentes, code, err := srv.SrvPatientBackground.GetPatientBackgroundByIDPatient(input.ID)
	if err != nil {
		logger.Error.Printf("couldn't get antecedentes, error: %v", err)
		return code, err
	}

	// Caso 1: input.Antecedentes es nil
	if input.Antecedentes == nil {
		if len(antecedentesExistentes) > 0 {
			for _, antecedente := range antecedentesExistentes {
				code, err = srv.SrvPatientBackground.DeletePatientBackground(antecedente.ID)
				if err != nil {
					logger.Error.Printf("couldn't delete antecedentes, error: %v", antecedente.ID, err)
					return code, err
				}
			}
		}
		return 29, nil // Finaliza aquí si no hay antecedentes en input
	}

	// Caso 2: input.Antecedentes no es nil
	if input.Antecedentes != nil {
		existingMap := make(map[string]*patient_background.PatientBackground) // Mapa de antecedentes existentes en la base
		for _, antecedentesExistente := range antecedentesExistentes {
			existingMap[antecedentesExistente.Nombre] = antecedentesExistente
		}

		inputMap := make(map[string]*models.RequestAntecedents) // Mapa de nuevos antecedentes en el input
		for _, antecedente := range input.Antecedentes {
			inputMap[antecedente.Nombre] = antecedente
		}

		var nuevosAntecedentes []*models.RequestAntecedents
		var antecedentesParaActualizar []*models.RequestAntecedents
		var antecedentesParaEliminar []*patient_background.PatientBackground

		for _, antecedenteInput := range input.Antecedentes {
			if _, existe := existingMap[antecedenteInput.Nombre]; existe {
				antecedentesParaActualizar = append(antecedentesParaActualizar, antecedenteInput)
			} else {
				nuevosAntecedentes = append(nuevosAntecedentes, antecedenteInput)
			}
		}

		for _, antecedenteExistente := range antecedentesExistentes {
			if _, existe := inputMap[antecedenteExistente.Nombre]; !existe {
				antecedentesParaEliminar = append(antecedentesParaEliminar, antecedenteExistente)
			}
		}

		if len(nuevosAntecedentes) > 0 {
			for _, nuevoAntecedente := range nuevosAntecedentes {
				code, err := srv.SrvPatientBackground.CreatePatientBackground(nuevoAntecedente.ID, input.ID, nuevoAntecedente.Nombre, nuevoAntecedente.Estado)
				if err != nil {
					logger.Error.Printf("couldn't create antecedente, error: %v", nuevoAntecedente.Nombre, err)
					return code, err
				}
			}
		}

		if len(antecedentesParaActualizar) > 0 {
			for _, antecedente := range antecedentesParaActualizar {
				code, err := srv.SrvPatientBackground.UpdatePatientBackground(antecedente.ID, input.ID, antecedente.Nombre, antecedente.Estado)
				if err != nil {
					logger.Error.Printf("couldn't update antecedente, error: %v", antecedente.Nombre, err)
					return code, err
				}
			}
		}

		if len(antecedentesParaEliminar) > 0 {
			for _, antecedente := range antecedentesParaEliminar {
				code, err := srv.SrvPatientBackground.DeletePatientBackground(antecedente.ID)
				if err != nil {
					logger.Error.Printf("couldn't delete antecedente, error: %v", antecedente.Nombre, err)
					return code, err
				}
			}
		}
	}

	return StatusSuccess, nil

}

func (s *PatientService) DeletePatientLowCode(id string) (int, error) {

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	codeBackground, err := srv.SrvPatientBackground.DeletePatientBackgroundByIDPatient(id)
	if err != nil {
		logger.Error.Printf("couldn't delete background, error: %v", err)
		return codeBackground, err
	}

	codePatient, err := srv.SrvPatient.DeletePatients(id)
	if err != nil {
		logger.Error.Printf("couldn't delete patient, error: %v", err)
		return codePatient, err
	}

	return 29, nil
}

func (s *PatientService) GetPatientByIdLowCode(id string) (*models.Patients, int, error) {
	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	data, err := fetchPatient(srv, id, s.txID)
	if err != nil {
		logger.Error.Println(s.txID, " - failed to get Patient:", err)
		return nil, 0, err
	}

	return data, 29, nil
}

func (s *PatientService) GetPatientByDNILowCode(dni string) (*models.Patients, int, error) {
	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	data, err := fetchPatientByDNI(srv, dni, s.txID)
	if err != nil {
		logger.Error.Println(s.txID, " - failed to get Patient:", err)
		return nil, 0, err
	}

	return data, 29, nil
}

func (s *PatientService) GetAllPatientLowCode() ([]*models.Patients, error) {
	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	patients, err := srv.SrvPatient.GetAllPatients()
	if err != nil {
		logger.Error.Println(s.txID, " - failed to get Patient:", err)
		return nil, fmt.Errorf("failed to get Patient: %w", err)
	}

	var data []*models.Patients
	for _, patient := range patients {

		antecedentes, err := fetchBackground(srv, patient.ID, s.txID)
		if err != nil {
			logger.Error.Println(s.txID, " - failed to get Background:", err)
			return nil, fmt.Errorf("failed to get Background: %w", err)
		}

		data = append(data, &models.Patients{
			ID:                 patient.ID,
			CodigoSGA:          patient.CodigoSGA,
			DNI:                patient.DNI,
			Nombres:            patient.Nombres,
			Apellidos:          patient.Apellidos,
			Sexo:               patient.Sexo,
			Edad:               patient.Edad,
			EstadoCivil:        patient.EstadoCivil,
			GrupoSanguineo:     patient.GrupoSanguineo,
			FechaNacimiento:    patient.FechaNacimiento,
			LugarNacimiento:    patient.LugarNacimiento,
			Procedencia:        patient.Procedencia,
			EscuelaProfesional: patient.EscuelaProfesional,
			Ocupacion:          patient.Ocupacion,
			CorreoElectronico:  patient.CorreoElectronico,
			NumeroCelular:      patient.NumeroCelular,
			Direccion:          patient.Direccion,
			TipoPersona:        patient.TipoPersona,
			FactorRH:           patient.FactorRH,
			Alergias:           patient.Alergias,
			RAM:                patient.RAM,
			Antecedentes:       antecedentes,
		})
	}

	return data, nil
}

func (s *PatientService) GetPatients(dni, names, surnames string, limit, offset int64) ([]*models.Patients, error) {
	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	patients, err := srv.SrvPatient.SearchPaginationPatients(dni, names, surnames, limit, offset)
	if err != nil {
		logger.Error.Println(s.txID, " - failed to get Patient:", err)
		return nil, fmt.Errorf("failed to get Patient: %w", err)
	}

	var data []*models.Patients
	for _, patient := range patients {

		antecedentes, err := fetchBackground(srv, patient.ID, s.txID)
		if err != nil {
			logger.Error.Println(s.txID, " - failed to get Background:", err)
			return nil, fmt.Errorf("failed to get Background: %w", err)
		}

		data = append(data, &models.Patients{
			ID:                 patient.ID,
			CodigoSGA:          patient.CodigoSGA,
			DNI:                patient.DNI,
			Nombres:            patient.Nombres,
			Apellidos:          patient.Apellidos,
			Sexo:               patient.Sexo,
			Edad:               patient.Edad,
			EstadoCivil:        patient.EstadoCivil,
			GrupoSanguineo:     patient.GrupoSanguineo,
			FechaNacimiento:    patient.FechaNacimiento,
			LugarNacimiento:    patient.LugarNacimiento,
			Procedencia:        patient.Procedencia,
			EscuelaProfesional: patient.EscuelaProfesional,
			Ocupacion:          patient.Ocupacion,
			CorreoElectronico:  patient.CorreoElectronico,
			NumeroCelular:      patient.NumeroCelular,
			Direccion:          patient.Direccion,
			TipoPersona:        patient.TipoPersona,
			FactorRH:           patient.FactorRH,
			Alergias:           patient.Alergias,
			RAM:                patient.RAM,
			Antecedentes:       antecedentes,
		})
	}

	return data, nil
}

func (s *PatientService) GetMetadata(dni, names, surnames string, limit, offset int64) (*models.Metadata, error) {
	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	totalPatient, err := srv.SrvPatient.CountPatients(dni, names, surnames)
	if err != nil {
		logger.Error.Println(s.txID, " - failed to get Patient:", err)
		return nil, fmt.Errorf("failed to get Patient: %w", err)
	}

	metadata := &models.Metadata{
		Links:      models.Links{},
		Total:      totalPatient,
		Offset:     offset,
		Limit:      limit,
		Page:       offset / limit,
		TotalPages: totalPatient / limit,
	}

	if metadata.Page < metadata.TotalPages {
		metadata.Links.Next = fmt.Sprintf("?limit=%d&offset=%d", limit, offset+limit)
	} else {
		metadata.Links.Next = ""
	}

	if metadata.Page > 0 {
		metadata.Links.Prev = fmt.Sprintf("?limit=%d&offset=%d", limit, offset-limit)
	} else {
		metadata.Links.Prev = ""
	}

	return metadata, nil
}

func fetchPatient(srvService *medical_area.ServerMedicalArea, id, txID string) (*models.Patients, error) {
	patient, _, err := srvService.SrvPatient.GetPatientsByID(id)
	if err != nil {
		logger.Error.Println(txID, " - failed to get Patient:", err)
		return nil, fmt.Errorf("failed to get Patient: %w", err)
	}
	if patient == nil {
		return nil, fmt.Errorf("Patient not found")
	}

	antecedentes, err := fetchBackground(srvService, id, txID)
	if err != nil {
		logger.Error.Println(txID, " - failed to get Background:", err)
		return nil, fmt.Errorf("failed to get Background: %w", err)
	}

	return &models.Patients{
		ID:                 patient.ID,
		CodigoSGA:          patient.CodigoSGA,
		DNI:                patient.DNI,
		Nombres:            patient.Nombres,
		Apellidos:          patient.Apellidos,
		Sexo:               patient.Sexo,
		Edad:               patient.Edad,
		EstadoCivil:        patient.EstadoCivil,
		GrupoSanguineo:     patient.GrupoSanguineo,
		FechaNacimiento:    patient.FechaNacimiento,
		LugarNacimiento:    patient.LugarNacimiento,
		Procedencia:        patient.Procedencia,
		EscuelaProfesional: patient.EscuelaProfesional,
		Ocupacion:          patient.Ocupacion,
		CorreoElectronico:  patient.CorreoElectronico,
		NumeroCelular:      patient.NumeroCelular,
		Direccion:          patient.Direccion,
		TipoPersona:        patient.TipoPersona,
		FactorRH:           patient.FactorRH,
		Alergias:           patient.Alergias,
		RAM:                patient.RAM,
		Antecedentes:       antecedentes,
	}, nil
}

func fetchPatientByDNI(srvService *medical_area.ServerMedicalArea, dni, txID string) (*models.Patients, error) {
	patient, _, err := srvService.SrvPatient.GetPatientsByDNI(dni)
	if err != nil {
		logger.Error.Println(txID, " - failed to get Patient:", err)
		return nil, fmt.Errorf("failed to get Patient: %w", err)
	}
	if patient == nil {
		return nil, nil
	}

	antecedentes, err := fetchBackground(srvService, patient.ID, txID)
	if err != nil {
		logger.Error.Println(txID, " - failed to get Background:", err)
		return nil, fmt.Errorf("failed to get Background: %w", err)
	}

	return &models.Patients{
		ID:                 patient.ID,
		CodigoSGA:          patient.CodigoSGA,
		DNI:                patient.DNI,
		Nombres:            patient.Nombres,
		Apellidos:          patient.Apellidos,
		Sexo:               patient.Sexo,
		Edad:               patient.Edad,
		EstadoCivil:        patient.EstadoCivil,
		GrupoSanguineo:     patient.GrupoSanguineo,
		FechaNacimiento:    patient.FechaNacimiento,
		LugarNacimiento:    patient.LugarNacimiento,
		Procedencia:        patient.Procedencia,
		EscuelaProfesional: patient.EscuelaProfesional,
		Ocupacion:          patient.Ocupacion,
		CorreoElectronico:  patient.CorreoElectronico,
		NumeroCelular:      patient.NumeroCelular,
		Direccion:          patient.Direccion,
		TipoPersona:        patient.TipoPersona,
		FactorRH:           patient.FactorRH,
		Alergias:           patient.Alergias,
		RAM:                patient.RAM,
		Antecedentes:       antecedentes,
	}, nil
}

func fetchResposePatientInfo(srvService *medical_area.ServerMedicalArea, id, txID string) (*models.ResponsePatientInfo, error) {
	patient, _, err := srvService.SrvPatient.GetPatientsByDNI(id)
	if err != nil {
		logger.Error.Println(txID, " - failed to get Patient:", err)
		return nil, fmt.Errorf("failed to get Patient: %w", err)
	}
	if patient == nil {
		return nil, fmt.Errorf("Patient not found")
	}

	return &models.ResponsePatientInfo{
		CodigoSGA:   patient.CodigoSGA,
		DNI:         patient.DNI,
		Nombres:     patient.Nombres,
		Apellidos:   patient.Apellidos,
		TipoPersona: patient.TipoPersona,
	}, nil
}

func fetchBackground(srvService *medical_area.ServerMedicalArea, id, txID string) ([]*models.RequestAntecedents, error) {
	backgrounds, _, err := srvService.SrvPatientBackground.GetPatientBackgroundByIDPatient(id)
	if err != nil {
		logger.Error.Println(txID, " - failed to get Background:", err)
		return nil, fmt.Errorf("failed to get Background: %w", err)
	}

	var data []*models.RequestAntecedents
	for _, background := range backgrounds {
		data = append(data, &models.RequestAntecedents{
			ID:     background.ID,
			Nombre: background.Nombre,
			Estado: background.Estado,
		})
	}

	return data, nil
}
