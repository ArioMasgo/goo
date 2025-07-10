package low_code_residences

import (
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"dbu-api/pkg/residence"
	"dbu-api/pkg/residence/rooms"
	"dbu-api/pkg/submission"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"strconv"
	"time"
)

type ResidenceService struct {
	db   *sqlx.DB
	usr  *models.User
	txID string
}

type PortsServerResidence interface {
	CreateResidenceLowCode(m *models.Residence) (int, error)
	GetResidenceLowCode() ([]*models.Residence, int, error)
	GetStudentsByResidenceLowCode(residenceID string, submissionID int64, page, limit int, filter string) ([]*models.Student, int, int, error)
	GetRoomsByResidenceLowCode(residenceID string, submissionID int64, page, limit int) ([]*models.RoomStudentsSimple, int, int, error)
	AssignmentRoom(roomID string, submissionID int64, studentID int64) (*models.AssignmentRoom, int, error)
}

func NewResidence(db *sqlx.DB, usr *models.User, txID string) PortsServerResidence {
	return &ResidenceService{db: db, usr: usr, txID: txID}
}

func (s *ResidenceService) CreateResidenceLowCode(m *models.Residence) (int, error) {
	// Validación de pisos
	if m.Floors == nil {
		logger.Error.Printf("%s - floors must not be nil", s.txID)
		return 2, errors.New("floors must not be nil")
	}

	// Validación del modelo
	valid, err := m.ValidResidence()
	if err != nil {
		logger.Error.Printf("%s - validation error: %v", s.txID, err)
		return 2, err
	}

	if !valid {
		logger.Error.Printf("%s - invalid residence data", s.txID)
		return 2, errors.New("residence data is invalid")
	}

	// Crear residencia
	srvService := residence.NewServerResidence(s.db, s.usr, s.txID)
	dataResidence, _, err := srvService.SrvResidence.CreateResidence(
		m.ID,
		m.Name,
		m.Description,
		m.Gender,
		m.Address,
		m.Status,
	)
	if err != nil {
		logger.Error.Printf("%s - couldn't create residence: %v", s.txID, err)
		return 12, err
	}

	if dataResidence == nil {
		logger.Error.Printf("%s - insufficient permissions", s.txID)
		return 10, errors.New("insufficient permissions")
	}

	// Crear configuración de residencia
	residenceConfiguration := models.Configuration{
		ID:                      uuid.New().String(),
		PercentageFcea:          70,
		PercentageEngineering:   30,
		MinimumGradeFcea:        16,
		MinimumGradeEngineering: 12,
	}

	dataResidenceConfiguration, _, err := srvService.SrvResidenceConfiguration.CreateResidenceConfiguration(
		residenceConfiguration.ID,
		residenceConfiguration.PercentageFcea,
		residenceConfiguration.PercentageEngineering,
		residenceConfiguration.MinimumGradeFcea,
		residenceConfiguration.MinimumGradeEngineering,
		m.ID,
		false,
	)
	if err != nil {
		logger.Error.Printf("%s - couldn't create residence configuration: %v", s.txID, err)
		return 12, err
	}

	if dataResidenceConfiguration == nil {
		logger.Error.Printf("%s - couldn't get residence configuration", s.txID)
		return 4, fmt.Errorf("residence configuration not found")
	}

	// Crear habitaciones
	dataRoom, err := srvService.SrvRoom.MultiCreate(m.ID, m.Floors)
	if err != nil {
		logger.Error.Printf("%s - couldn't create rooms: %v", s.txID, err)
		return 12, err
	}

	if dataRoom == nil {
		logger.Error.Printf("%s - couldn't get room data", s.txID)
		return 4, fmt.Errorf("room data not found")
	}

	return 211, nil
}

func (s *ResidenceService) GetResidenceLowCode() ([]*models.Residence, int, error) {
	srvService := residence.NewServerResidence(s.db, s.usr, s.txID)
	residences, err := srvService.SrvResidence.GetAllResidence()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get user by username", err)
		return nil, 15, err
	}

	if residences == nil {
		return []*models.Residence{}, 214, nil
	}

	var data []*models.Residence

	roomsByResidence, err := srvService.SrvRoom.GetAllRooms()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get user by username", err)
		return nil, 15, errors.New("permissions not found")
	}

	if roomsByResidence == nil {
		logger.Error.Println(s.txID, " - couldn't get user by username")
		return nil, 4, errors.New("permissions not found")
	}

	roomsMap := make(map[string]map[int][]models.Room)
	for _, room := range roomsByResidence {
		if _, exists := roomsMap[room.ResidenceID]; !exists {
			roomsMap[room.ResidenceID] = make(map[int][]models.Room)
		}
		roomsMap[room.ResidenceID][room.Floor] = append(roomsMap[room.ResidenceID][room.Floor], models.Room{
			ID:       room.ID,
			Number:   room.Number,
			Capacity: room.Capacity,
			Status:   room.Status,
		})
	}

	var listSubmissions []*models.Submission

	months := map[string]string{
		"January":   "Enero",
		"February":  "Febrero",
		"March":     "Marzo",
		"April":     "Abril",
		"May":       "Mayo",
		"June":      "Junio",
		"July":      "Julio",
		"August":    "Agosto",
		"September": "Septiembre",
		"October":   "Octubre",
		"November":  "Noviembre",
		"December":  "Diciembre",
	}

	//id 2 is for residence service
	srvServiceSubmission := submission.NewServerSubmission(s.db, s.usr, s.txID)
	submissions, err := srvServiceSubmission.SrvConvocatorias.GetAllSubmissionsByService(2)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get user by username", err)
		return nil, 15, err
	}

	if submissions == nil {
		listSubmissions = []*models.Submission{}
	}

	for _, submissionInfo := range submissions {
		now := time.Now()
		dataSubmission := &models.Submission{
			ID:   submissionInfo.ID,
			Name: submissionInfo.Nombre,
		}
		if submissionInfo.FechaInicio == nil || submissionInfo.FechaFin == nil {
			continue
		}

		month := submissionInfo.FechaInicio.Format("January")
		year := submissionInfo.FechaInicio.Format("2006")
		monthSpanish := months[month]
		dataSubmission.Start = fmt.Sprintf("%s %s", monthSpanish, year)

		if !now.Before(*submissionInfo.FechaInicio) && !now.After(*submissionInfo.FechaFin) {
			dataSubmission.State = true
		}

		listSubmissions = append(listSubmissions, dataSubmission)
	}

	for _, dataResidence := range residences {
		configuration, _, errConfig := srvService.SrvResidenceConfiguration.GetResidenceConfigurationByResidenceID(dataResidence.ID)
		if errConfig != nil {
			logger.Error.Println(s.txID, " - couldn't get residence configuration", err)
			return nil, 15, err
		}

		if configuration == nil {
			logger.Error.Println(s.txID, " - No found residence configuration", err)
			return nil, 4, errors.New("residence configuration not found")
		}

		data = append(data, &models.Residence{
			ID:          dataResidence.ID,
			Name:        dataResidence.Name,
			Gender:      dataResidence.Gender,
			Description: dataResidence.Description,
			Address:     dataResidence.Address,
			Status:      dataResidence.Status,
			Config: &models.Configuration{
				ID:                      configuration.ID,
				PercentageFcea:          configuration.PercentageFcea,
				PercentageEngineering:   configuration.PercentageEngineering,
				MinimumGradeFcea:        configuration.MinimumGradeFcea,
				MinimumGradeEngineering: configuration.MinimumGradeEngineering,
				IsNewbie:                configuration.IsNewbie,
			},
			Submissions: listSubmissions,
		})
	}

	return data, 214, nil
}

func (s *ResidenceService) GetStudentsByResidenceLowCode(residenceID string, submissionID int64, page, limit int, filter string) ([]*models.Student, int, int, error) {
	srvService := residence.NewServerResidence(s.db, s.usr, s.txID)
	residenceData, _, err := srvService.SrvResidence.GetResidenceByID(residenceID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get user by username", err)
		return nil, 15, 0, err
	}

	if residenceData == nil {
		return make([]*models.Student, 0), 4, 0, nil
	}

	roomsResidence, err := srvService.SrvRoom.GetAllRoomsByResidenceID(residenceData.ID)
	if err != nil {
		logger.Error.Printf("%s - error getting students: %v", s.txID, err)
		return nil, 15, 0, err
	}

	if roomsResidence == nil {
		roomsResidence = make([]*rooms.Room, 0)
	}

	roomInfo := make(map[string]*models.Room)

	for _, room := range roomsResidence {
		roomKey := strconv.Itoa(room.Number)
		if roomKey != "" {
			roomInfo[roomKey] = &models.Room{
				ID:       room.ID,
				Number:   room.Number,
				Capacity: room.Capacity,
				Status:   room.Status,
			}
		}
	}

	srvSubmission := submission.NewServerSubmission(s.db, s.usr, s.txID)

	totalStudents, err := srvSubmission.SrvAlumnos.GetTotalStudentsByResidenceANDBySubmission(residenceData.ID, submissionID, filter)
	if err != nil {
		logger.Error.Printf("%s - error getting total students by residence: %v", s.txID, err)
		return nil, 15, 0, err
	}

	studentsResidence, err := srvSubmission.SrvAlumnos.GetStudentsByResidenceANDBySubmission(residenceData.ID,
		submissionID, page, limit, filter)
	if err != nil {
		logger.Error.Printf("%s - error getting students: %v", s.txID, err)
		return nil, 15, 0, err
	}

	if studentsResidence == nil {
		return make([]*models.Student, 0), 214, totalStudents, nil
	}

	roomStudents := make(map[string][]models.RoomMate)
	var students []*models.Student

	for _, studentInfo := range studentsResidence {
		roomStudents[studentInfo.Room] = append(roomStudents[studentInfo.Room], models.RoomMate{
			FullName: studentInfo.FullName,
			Code:     studentInfo.Code,
		})

		student := &models.Student{
			StudentInfo: models.StudentInfo{
				ID:                   studentInfo.ID,
				NumberIdentification: studentInfo.DNI,
				FullName:             studentInfo.FullName,
				Code:                 studentInfo.Code,
				ProfessionalSchool:   studentInfo.ProfessionalSchool,
				Faculty:              studentInfo.Faculty,
				Room:                 roomInfo[studentInfo.Room],
				Residence:            studentInfo.Residence,
				AdmissionDate:        studentInfo.AdmissionDate,
			},
		}
		students = append(students, student)
	}

	for i, student := range students {
		students[i].RoomMates = []models.RoomMate{}

		students[i].AssignedGoods = []models.AssignedGood{}
		students[i].Sanctions = []models.Sanction{}

		if student.StudentInfo.Room.Number == 0 {
			continue
		}

		for _, rm := range roomStudents[strconv.Itoa(student.StudentInfo.Room.Number)] {
			if rm.Code != student.StudentInfo.Code {
				students[i].RoomMates = append(students[i].RoomMates, rm)
			}
		}
	}

	return students, 214, totalStudents, nil
}

func (s *ResidenceService) GetRoomsByResidenceLowCode(residenceID string, submissionID int64, page, limit int) ([]*models.RoomStudentsSimple, int, int, error) {
	srvService := residence.NewServerResidence(s.db, s.usr, s.txID)
	residenceData, _, err := srvService.SrvResidence.GetResidenceByID(residenceID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get user by username", err)
		return nil, 15, 0, err
	}

	if residenceData == nil {
		return nil, 4, 0, nil
	}

	roomsData, err := srvService.SrvRoom.GetRoomsByResidence(residenceData.ID, page, limit)
	if err != nil {
		logger.Error.Printf("%s - error getting students: %v", s.txID, err)
		return nil, 15, 0, err
	}

	if roomsData == nil {
		return make([]*models.RoomStudentsSimple, 0), 214, 0, nil
	}

	roomStudents := make([]string, 0, len(roomsData))
	for _, roomInfo := range roomsData {
		roomStudents = append(roomStudents, roomInfo.ID)
	}

	srvSubmission := submission.NewServerSubmission(s.db, s.usr, s.txID)
	studentsRoom, err := srvSubmission.SrvAlumnos.GetStudentsByRooms(roomStudents, submissionID)
	if err != nil {
		logger.Error.Printf("%s - error getting students: %v", s.txID, err)
		return nil, 15, 0, err
	}

	if studentsRoom == nil {
		roomsEmptyStudents := make([]*models.RoomStudentsSimple, 0, len(roomsData))
		for _, roomInfo := range roomsData {
			roomsEmptyStudents = append(roomsEmptyStudents, &models.RoomStudentsSimple{
				ID:       roomInfo.ID,
				Number:   roomInfo.Number,
				Capacity: roomInfo.Capacity,
				Status:   roomInfo.Status,
				Floor:    roomInfo.Floor,
			})
		}
		return roomsEmptyStudents, 214, 0, nil
	}

	mapRoomStudents := make(map[string][]models.StudentInfoWithoutRoom)
	for _, studentInfo := range studentsRoom {
		mapRoomStudents[studentInfo.Room] = append(mapRoomStudents[studentInfo.Room], models.StudentInfoWithoutRoom{
			ID:                 studentInfo.ID,
			FullName:           studentInfo.FullName,
			Code:               studentInfo.Code,
			ProfessionalSchool: studentInfo.ProfessionalSchool,
			Faculty:            studentInfo.Faculty,
			AdmissionDate:      studentInfo.AdmissionDate,
		})
	}

	var rooms []*models.RoomStudentsSimple
	for _, roomInfo := range roomsData {
		rooms = append(rooms, &models.RoomStudentsSimple{
			ID:       roomInfo.ID,
			Number:   roomInfo.Number,
			Capacity: roomInfo.Capacity,
			Status:   roomInfo.Status,
			Floor:    roomInfo.Floor,
			Students: mapRoomStudents[strconv.Itoa(roomInfo.Number)],
		})
	}

	if rooms == nil {
		rooms = make([]*models.RoomStudentsSimple, 0)
	}

	return rooms, 214, 0, nil
}

func (s *ResidenceService) AssignmentRoom(roomID string, submissionID int64, studentID int64) (*models.AssignmentRoom, int, error) {
	srv := residence.NewServerResidence(s.db, s.usr, s.txID)
	room, _, err := srv.SrvRoom.GetRoomByID(roomID)
	if err != nil {
		logger.Error.Printf("couldn't get Rooms, error: %v", err)
		return nil, 15, fmt.Errorf("Error al obtener el registro: %v", err)
	}

	if room == nil {
		logger.Error.Printf("no found room")
		return nil, 4, fmt.Errorf("Recurso no encontrado")
	}

	srvSubmission := submission.NewServerSubmission(s.db, s.usr, s.txID)
	studentAccepted, _, err := srvSubmission.SrvAlumnos.GetStudentAcceptedBySubmission(submissionID, studentID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get student", err)
		return nil, 15, fmt.Errorf("Error al obtener el registro: %v", err)
	}

	if studentAccepted == nil {
		logger.Error.Println(s.txID, " - no found student")
		return nil, 96, fmt.Errorf("El estudiante no está aceptado en esta convocatoria")
	}

	activeSubmission, _, err := srvSubmission.SrvConvocatorias.GetSubmissionsByID(submissionID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get active submission", err)
		return nil, 15, fmt.Errorf("Error al obtener el registro: %v", err)
	}

	if activeSubmission == nil {
		logger.Error.Println(s.txID, " - no found active submission")
		return nil, 4, fmt.Errorf("Recurso no encontrado")
	}

	allAssignmentsRoom, err := srv.SrvAssignmentRoom.GetRoomAssignmentByRoomIDSubmissionID(roomID, activeSubmission.ID)
	if err != nil {
		logger.Error.Printf("couldn't get assignment room, error: %v", err)
		return nil, 15, fmt.Errorf("Error al obtener el registro: %v", err)
	}

	if allAssignmentsRoom == nil {
		logger.Error.Printf("no found assignment room")
		return nil, 4, fmt.Errorf("Recurso no encontrado")
	}

	countAssignmentsRoom := 0
	for _, assignment := range allAssignmentsRoom {
		if assignment.Status == "activo" {
			countAssignmentsRoom++
		}
		if assignment.StudentID == studentID && assignment.CallID == submissionID && assignment.Status == "activo" {
			logger.Error.Println("Student is already assigned to this room for the given submission")
			return nil, 94, fmt.Errorf("El estudiante ya está asignado a esta habitación")
		}
	}

	if room.Capacity <= countAssignmentsRoom {
		logger.Error.Printf("couldn't assignment room")
		return nil, 93, fmt.Errorf("Habitación sin capacidad disponible")
	}

	assignmentsRoomByStudentAndSubmission, err := srv.SrvAssignmentRoom.GetAllRoomAssignmentsByStudentIDANDSubmissionID(studentID, activeSubmission.ID)
	if err != nil {
		logger.Error.Printf("couldn't get assignment room, error: %v", err)
		return nil, 15, fmt.Errorf("Error al obtener el registro: %v", err)
	}

	if len(assignmentsRoomByStudentAndSubmission) == 0 {
		assignmentRoom, _, err := srv.SrvAssignmentRoom.CreateRoomAssignment(studentID, room.ID, activeSubmission.ID, time.Now(), "activo")
		if err != nil {
			logger.Error.Printf("couldn't create assignment room, error: %v", err)
			return nil, 12, fmt.Errorf("Error al crear el registro: %v", err)
		}

		if assignmentRoom == nil {
			logger.Error.Printf("no created assignment room")
			return nil, 12, fmt.Errorf("Error al crear el registro")
		}

		responseAssignment := &models.AssignmentRoom{
			ID:             assignmentRoom.ID,
			Status:         assignmentRoom.Status,
			AssignmentDate: assignmentRoom.AssignmentDate,
		}

		return responseAssignment, 211, nil
	}

	assignmentRoom, _, err := srv.SrvAssignmentRoom.UpdateRoomAssignment(assignmentsRoomByStudentAndSubmission[0].ID,
		assignmentsRoomByStudentAndSubmission[0].StudentID, assignmentsRoomByStudentAndSubmission[0].RoomID,
		assignmentsRoomByStudentAndSubmission[0].CallID, time.Now(), assignmentsRoomByStudentAndSubmission[0].Status)

	if err != nil {
		logger.Error.Printf("couldn't update assignment room, error: %v", err)
		return nil, 13, fmt.Errorf("Error al actualizar el registro: %v", err)
	}

	if assignmentRoom == nil {
		logger.Error.Printf("no update assignment room")
		return nil, 13, fmt.Errorf("Error al actualizar el registro")
	}

	responseAssignment := &models.AssignmentRoom{
		ID:             assignmentRoom.ID,
		Status:         assignmentRoom.Status,
		AssignmentDate: assignmentRoom.AssignmentDate,
	}

	return responseAssignment, 212, nil
}
