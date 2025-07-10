package low_code_submissions

import (
	"dbu-api/internal/env"
	"dbu-api/internal/file"
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"dbu-api/pkg/residence"
	"dbu-api/pkg/residence/rooms"
	"dbu-api/pkg/submission"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type SubmissionService struct {
	db   *sqlx.DB
	usr  *models.User
	txID string
}

type PortsServerSubmission interface {
	GetStudentsBySubmissionsLowCode(submissionID int, page, limit int, gender string) ([]*models.Student, int, int, error)
	GetReportBySubmissionsLowCode(submissionID int) (string, int, error)
}

func NewSubmission(db *sqlx.DB, usr *models.User, txID string) PortsServerSubmission {
	return &SubmissionService{db: db, usr: usr, txID: txID}
}

func (s *SubmissionService) GetStudentsBySubmissionsLowCode(submissionID int, page, limit int, gender string) ([]*models.Student, int, int, error) {
	cfg := env.NewConfiguration()
	srvSubmission := submission.NewServerSubmission(s.db, s.usr, s.txID)

	totalStudents, err := srvSubmission.SrvAlumnos.GetTotalStudentsBySubmission(submissionID, gender, "aprobado", cfg.Department.RequirementId)
	if err != nil {
		logger.Error.Printf("%s - error getting total students: %v", s.txID, err)
		return nil, 15, 0, err
	}

	studentsResidence, err := srvSubmission.SrvAlumnos.GetStudentsBySubmission(submissionID, page, limit, gender, "aprobado", cfg.Department.RequirementId)
	if err != nil {
		logger.Error.Printf("%s - error getting students: %v", s.txID, err)
		return nil, 15, 0, err
	}

	if studentsResidence == nil {
		return []*models.Student{}, 214, totalStudents, nil
	}

	srvResidence := residence.NewServerResidence(s.db, s.usr, s.txID)
	roomsResidence, err := srvResidence.SrvRoom.GetAllRooms()
	if err != nil {
		logger.Error.Printf("%s - error getting students: %v", s.txID, err)
		return nil, 15, 0, err
	}

	if roomsResidence == nil {
		roomsResidence = make([]*rooms.Room, 0)
	}

	roomInfo := make(map[string]*models.Room)

	for _, room := range roomsResidence {
		roomKey := fmt.Sprintf("%s_%d", room.ResidenceID, room.Number)
		if roomKey != "_" {
			roomInfo[roomKey] = &models.Room{
				ID:       room.ID,
				Number:   room.Number,
				Capacity: room.Capacity,
				Status:   room.Status,
			}
		}

	}

	roomStudents := make(map[string][]models.RoomMate)
	var students []*models.Student

	for _, student := range studentsResidence {
		roomKey := fmt.Sprintf("%s_%s", student.Residence, student.Room)
		if roomKey != "_" {
			roomStudents[roomKey] = append(roomStudents[roomKey], models.RoomMate{
				FullName: student.FullName,
				Code:     student.Code,
			})
		}

		data := &models.Student{
			StudentInfo: models.StudentInfo{
				ID:                   student.ID,
				NumberIdentification: student.DNI,
				FullName:             student.FullName,
				Sex:                  student.Sex,
				Department:           student.Department,
				Code:                 student.Code,
				ProfessionalSchool:   student.ProfessionalSchool,
				Faculty:              student.Faculty,
				Room:                 roomInfo[roomKey],
				Residence:            student.Residence,
				AdmissionDate:        student.AdmissionDate,
			},
		}
		students = append(students, data)
	}

	for i, student := range students {
		students[i].RoomMates = []models.RoomMate{}

		//TODO: Search assignment objects and sanctions
		students[i].AssignedGoods = []models.AssignedGood{}
		students[i].Sanctions = []models.Sanction{}

		roomKey := fmt.Sprintf("%s_%s", student.StudentInfo.Residence, student.StudentInfo.Room)

		if roomKey == "_" {
			continue
		}

		for _, rm := range roomStudents[roomKey] {
			if rm.Code != student.StudentInfo.Code {
				students[i].RoomMates = append(students[i].RoomMates, rm)
			}
		}
	}

	return students, 214, totalStudents, nil
}

func (s *SubmissionService) GetReportBySubmissionsLowCode(submissionID int) (string, int, error) {
	srvResidence := residence.NewServerResidence(s.db, s.usr, s.txID)
	residences, err := srvResidence.SrvResidence.GetAllResidence()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get user by username", err)
		return "", 15, err
	}

	if residences == nil {
		return "", 214, errors.New("residence not found")
	}

	now := time.Now()
	fmt.Println(now.Format("20060102_150405"))

	excel := models.ExcelFile{
		Name: fmt.Sprintf("reporte-residencia-%s.xlsx", now.Format("20060102_150405")),
		Path: "./reports/residence",
		Page: make([]models.ExcelPage, 0, len(residences)+1),
	}

	srvSubmission := submission.NewServerSubmission(s.db, s.usr, s.txID)
	studentsResidence, err := srvSubmission.SrvAlumnos.GetStudentsBySubmissionExcel(submissionID)
	if err != nil {
		logger.Error.Printf("%s - error getting students: %v", s.txID, err)
		return "", 15, err
	}

	if studentsResidence == nil {
		//return []*models.Student{}, 214, nil
		studentsResidence = []*models.StudentExcel{}
	}

	// Crear mapa de estudiantes por residencia
	studentsByResidence := make(map[string][]*models.StudentExcel)
	for _, student := range studentsResidence {
		studentsByResidence[student.Residence] = append(studentsByResidence[student.Residence], student)
	}

	pagesResidences := make([]models.ExcelPage, 0, len(residences))
	residenceMap := make(map[string]string)

	for _, res := range residences {
		residenceMap[res.ID] = res.Name

		page := models.ExcelPage{
			Name: res.Name,
			Rows: make([]models.ExcelPageRow, 0, len(studentsByResidence[res.ID])),
		}

		page.Rows = append(page.Rows, createResidentStudentExcelData(nil, 1))

		// Agregar estudiantes de esta residencia
		indexActives := 1
		for _, student := range studentsByResidence[res.ID] {
			if student.Status == "activo" {
				indexActives++
				page.Rows = append(page.Rows, createResidentStudentExcelData(student, indexActives))
			}
		}

		pagesResidences = append(pagesResidences, page)
	}

	rowFirstPage := make([]models.ExcelPageRow, 0, len(studentsResidence))
	rowFirstPage = append(rowFirstPage, createGeneralStudentExcelData(nil, 1))

	var rowLeaveStudentsPage []models.ExcelPageRow
	rowLeaveStudentsPage = append(rowLeaveStudentsPage, createLeaveStudentExcelData(nil, 1))

	leaveIndex := 1
	for i, stu := range studentsResidence {
		stu.Residence = residenceMap[stu.Residence]
		rowPage := i + 2
		rowFirstPage = append(rowFirstPage, createGeneralStudentExcelData(stu, rowPage))

		if stu.Status != "activo" {
			leaveIndex++
			rowLeaveStudentsPage = append(rowLeaveStudentsPage, createLeaveStudentExcelData(stu, leaveIndex))
		}
	}

	excel.Page = append(excel.Page, models.ExcelPage{
		Name: "General",
		Rows: rowFirstPage,
	})

	excel.Page = append(excel.Page, models.ExcelPage{
		Name: "Alumnos Retirados-Suspendidos",
		Rows: rowLeaveStudentsPage,
	})

	excel.Page = append(excel.Page, pagesResidences...)

	base64, codErr := file.CreateExcelFile(&excel)
	if codErr != 223 {
		return "", codErr, errors.New("error to generate excel file")
	}

	return base64, 214, nil
}

func createGeneralStudentExcelData(student *models.StudentExcel, row int) models.ExcelPageRow {
	if row == 1 {
		return models.ExcelPageRow{
			Row: 1,
			Columns: []models.ExcelPageColumn{
				{Column: "A", Value: "Código"},
				{Column: "B", Value: "Nombre Completo"},
				{Column: "C", Value: "Género"},
				{Column: "D", Value: "Lugar de Procedencia"},
				{Column: "E", Value: "Dirección"},
				{Column: "F", Value: "Ponderado Semestral"},
				{Column: "G", Value: "Semestres Cursados"},
				{Column: "H", Value: "Escuela Profesional"},
				{Column: "I", Value: "Residencia"},
				{Column: "J", Value: "Fecha de Admisión"},
				{Column: "K", Value: "Estado Asignación"},
			},
		}
	}

	return models.ExcelPageRow{
		Row: row,
		Columns: []models.ExcelPageColumn{
			{Column: "A", Value: student.Code},
			{Column: "B", Value: student.FullName},
			{Column: "C", Value: student.Sex},
			{Column: "D", Value: student.Home},
			{Column: "E", Value: student.Address},
			{Column: "F", Value: student.PPS},
			{Column: "G", Value: student.NumSemestersCompleted},
			{Column: "H", Value: student.ProfessionalSchool},
			{Column: "I", Value: student.Residence},
			{Column: "J", Value: student.AdmissionDate},
			{Column: "K", Value: student.Status},
		},
	}
}

func createResidentStudentExcelData(student *models.StudentExcel, row int) models.ExcelPageRow {
	if row == 1 {
		return models.ExcelPageRow{
			Row: 1,
			Columns: []models.ExcelPageColumn{
				{Column: "A", Value: "Código"},
				{Column: "B", Value: "Nombre Completo"},
				{Column: "C", Value: "Género"},
				{Column: "D", Value: "Ponderado Semestral"},
				{Column: "E", Value: "Semestres Cursados"},
				{Column: "F", Value: "Escuela Profesional"},
				{Column: "G", Value: "Cuarto"},
			},
		}
	}

	return models.ExcelPageRow{
		Row: row,
		Columns: []models.ExcelPageColumn{
			{Column: "A", Value: student.Code},
			{Column: "B", Value: student.FullName},
			{Column: "C", Value: student.Sex},
			{Column: "D", Value: student.PPS},
			{Column: "E", Value: student.NumSemestersCompleted},
			{Column: "F", Value: student.ProfessionalSchool},
			{Column: "G", Value: student.Room},
		},
	}
}

func createLeaveStudentExcelData(student *models.StudentExcel, row int) models.ExcelPageRow {
	if row == 1 {
		return models.ExcelPageRow{
			Row: 1,
			Columns: []models.ExcelPageColumn{
				{Column: "A", Value: "Código"},
				{Column: "B", Value: "Nombre Completo"},
				{Column: "C", Value: "Género"},
				{Column: "D", Value: "Lugar de Procedencia"},
				{Column: "E", Value: "Dirección"},
				{Column: "F", Value: "Ponderado Semestral"},
				{Column: "G", Value: "Semestres Cursados"},
				{Column: "H", Value: "Escuela Profesional"},
				{Column: "I", Value: "Residencia"},
				{Column: "J", Value: "Fecha actualización"},
			},
		}
	}

	return models.ExcelPageRow{
		Row: row,
		Columns: []models.ExcelPageColumn{
			{Column: "A", Value: student.Code},
			{Column: "B", Value: student.FullName},
			{Column: "C", Value: student.Sex},
			{Column: "D", Value: student.Home},
			{Column: "E", Value: student.Address},
			{Column: "F", Value: student.PPS},
			{Column: "G", Value: student.NumSemestersCompleted},
			{Column: "H", Value: student.ProfessionalSchool},
			{Column: "I", Value: student.Residence},
			{Column: "J", Value: student.UpdateDate},
		},
	}
}
