package low_code_medical_area

import (
	"bytes"
	"dbu-api/internal/file"
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"dbu-api/pkg/medical_area"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
	"time"
)

type ReportsMedicalAreaService struct {
	db   *sqlx.DB
	usr  *models.User
	txID string
}

type PortsServerReportsMedicalArea interface {
	GetReportMedicalConsultationByMedicalAreaLowCode(areaMedica, fechaInicio, fechaFin string) (string, int, error)
	GetReportNursingFrameLowCode(numeroCuadro, month, year, startYear, endYear string) (string, int, error)
	GetReportDentistryFrameLowCode(numeroCuadro, month, year string) (string, int, error)
	GetReportMedicalFrameLowCode(numeroCuadro, month, year string) (string, int, error)
}

func NewReportsMedicalArea(db *sqlx.DB, usr *models.User, txID string) PortsServerReportsMedicalArea {
	return &ReportsMedicalAreaService{db: db, usr: usr, txID: txID}
}

var rangeDateTrimester = map[string][2]string{
	"1,2,3":    {"01-01", "03-31"},
	"4,5,6":    {"04-01", "06-30"},
	"7,8,9":    {"07-01", "09-30"},
	"10,11,12": {"10-01", "12-31"},
}

var months = map[string]string{
	"1":  "ENERO",
	"2":  "FEBRERO",
	"3":  "MARZO",
	"4":  "ABRIL",
	"5":  "MAYO",
	"6":  "JUNIO",
	"7":  "JULIO",
	"8":  "AGOSTO",
	"9":  "SETIEMBRE",
	"10": "OCTUBRE",
	"11": "NOVIEMBRE",
	"12": "DICIEMBRE",
}

var schools = []string{
	"AGRONOMIA",
	"ZOOTECNIA",
	"INGENIERIA EN INDUSTRIAS ALIMENTARIAS",
	"INGENIERIA FORESTAL",
	"INGENIERIA EN CONSERVACION DE SUELOS Y AGUA",
	"INGENIERIA AMBIENTAL",
	"INGENIERIA EN RECURSOS NATURALES RENOVABLES",
	"ECONOMIA",
	"ADMINISTRACION",
	"CONTABILIDAD",
	"INGENIERIA EN INFORMATICA Y SISTEMAS",
	"INGENIERIA MECANICA ELECTRICA",
	"INGENIERIA EN CIBERSEGURIDAD",
	"INGENIERIA CIVIL",
	"TURISMO Y HOTELERIA",
}

var nursingProceduresHeader = []string{
	"INYECT.",
	"C.F.V",
	"LAVADO DE OÍDO",
	"CIRUG. MENOR",
	"VENOCLISIS",
	"OTROS",
}

var nursingProceduresComparation = []string{
	"INYECTABLES",
	"CONTROL DE FUNCIONES VITALES",
	"LAVADO DE OIDO",
	"CIRUGIA MENOR",
	"VENOCLISIS",
	"OTROS",
}

var nursingProcedures = []string{
	"INYECTABLES",
	"CONTROL DE FUNCIONES VITALES",
	"LAVADO DE OIDO",
	"BAÑO RAYOS INFRARROJOS",
	"EXTRACCION DE UÑERO",
	"CIRUGIA MENOR",
	"VENOCLISIS",
	"OTROS",
}

var nursingStaffs = []string{
	"ADMINISTRATIVOS",
	"DOCENTES",
	"OTROS",
}

var nursingStaffsComparation = []string{
	"Administrativo",
	"Docente",
	"Otros",
}

var dentistryProcedures = []string{
	"PERIODONCIA PROFILAXIS",
	"OPERATORIA RESINA",
	"CIRUGÍA EXODONCIA",
	"DIAGNÓSTICO EXAMEN BUCAL",
	"CONSULTA DENTAL",
	"PREVENCIÓN FLUORIZACIÓN",
	"PREVENCIÓN: COLUTORIO CH2+CPC 0.12%",
	"PREVENCIÓN: IHO - TCA CEPILLADO",
	"RADIOGRAFÍA DENTAL",
}

var dentistryStaffs = []string{
	"ADMINISTRATIVOS",
	"DOCENTES",
}

var dentistryStaffsComparation = []string{
	"Administrativo",
	"Docente",
}

func (s *ReportsMedicalAreaService) GetReportMedicalConsultationByMedicalAreaLowCode(areaMedica, fechaInicio, fechaFin string) (string, int, error) {
	srvNursingConsultation := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)

	if fechaInicio == "" {
		fechaInicio = "2000-01-01"
	}
	if fechaFin == "" {
		fechaFin = time.Now().Format("2006-01-02")
	}
	fechaInicioTime, _ := time.Parse("2006-01-02", fechaInicio)
	fechaFinTime, _ := time.Parse("2006-01-02", fechaFin)
	fechaFinTime = fechaFinTime.AddDate(0, 0, 1)

	fechaInicio = fechaInicioTime.Format("2006-01-02")
	fechaFin = fechaFinTime.Format("2006-01-02")

	if areaMedica == "" {
		return "", 1, errors.New("missing areaMedica")
	}

	now := time.Now()
	excel := models.ExcelFile{
		Name: fmt.Sprintf("reporte-consultas-%s-%s.xlsx", areaMedica, now.Format("20060102_150405")),
		Path: "./reports/medical_area",
		Page: []models.ExcelPage{},
	}

	var rows []models.ExcelPageRow

	// ENFERMERÍA
	if areaMedica == "enfermería" {
		consultations, err := srvNursingConsultation.SrvConsultationMedicalArea.GetAllConsultationMedicalAreaNursingByDateExcel(areaMedica, fechaInicio, fechaFin)
		if err != nil {
			logger.Error.Println(s.txID, " - couldn't get nursing consultations", err)
			return "", 15, err
		}

		rows = append(rows, createNursingExcelRow(nil, 1)) // encabezado
		for i, c := range consultations {
			dateBirth, err := time.Parse("2006-01-02", c.FechaNacimiento)
			if err != nil {
				logger.Error.Println(s.txID, " - couldn't parse date:", err)
				return "", 15, err
			}
			c.FechaNacimiento = dateBirth.Format("02/01/2006")
			rows = append(rows, createNursingExcelRow(c, i+2))
		}

	} else if areaMedica == "odontología" {
		consultations, err := srvNursingConsultation.SrvConsultationMedicalArea.GetAllConsultationMedicalAreaDentistryByDateExcel(areaMedica, fechaInicio, fechaFin)
		if err != nil {
			logger.Error.Println(s.txID, " - couldn't get dentistry consultations", err)
			return "", 15, err
		}

		rows = append(rows, createDentistryExcelRow(nil, 1)) // encabezado
		for i, c := range consultations {
			dateBirth, err := time.Parse("2006-01-02", c.FechaNacimiento)
			if err != nil {
				logger.Error.Println(s.txID, " - couldn't parse date:", err)
				return "", 15, err
			}
			c.FechaNacimiento = dateBirth.Format("02/01/2006")
			rows = append(rows, createDentistryExcelRow(c, i+2))
		}

	} else {
		consultations, err := srvNursingConsultation.SrvConsultationMedicalArea.GetAllConsultationMedicalAreaMedicalByDateExcel(areaMedica, fechaInicio, fechaFin)
		if err != nil {
			logger.Error.Println(s.txID, " - couldn't get medical consultations", err)
			return "", 15, err
		}

		rows = append(rows, createMedicalExcelRow(nil, 1))
		for i, c := range consultations {
			dateBirth, err := time.Parse("2006-01-02", c.FechaNacimiento)
			if err != nil {
				logger.Error.Println(s.txID, " - couldn't parse date:", err)
				return "", 15, err
			}
			c.FechaNacimiento = dateBirth.Format("02/01/2006")
			rows = append(rows, createMedicalExcelRow(c, i+2))
		}
	}

	excel.Page = append(excel.Page, models.ExcelPage{
		Name: "General",
		Rows: rows,
	})

	base64Str, codErr := file.CreateExcelFile(&excel)
	if codErr != 223 {
		return "", codErr, errors.New("error to generate excel file")
	}

	return base64Str, 0, nil
}

func (s *ReportsMedicalAreaService) GetReportNursingFrameLowCode(numeroCuadro, month, year, startYear, endYear string) (string, int, error) {

	if numeroCuadro == "" {
		return "", 1, errors.New("missing numeroCuadro")
	}

	if month == "" {
		month = "1,2,3"
	}

	if year == "" {
		year = time.Now().Format("2006")
	}

	if startYear == "" {
		startYear = time.Now().Format("2006")
	}
	if endYear == "" {
		endYear = time.Now().Format("2006")
	}

	if numeroCuadro == "1" {
		base64Str, code, err := s.GetReportNursingFrame1LowCode(month, year)
		if err != nil {
			logger.Error.Println("Error generating report nursing frame 1:", err)
			return "", code, err
		}
		return base64Str, code, err
	} else if numeroCuadro == "2" {
		base64Str, code, err := s.GetReportNursingFrame2LowCode(month, year)
		if err != nil {
			logger.Error.Println("Error generating report nursing frame 2:", err)
			return "", code, err
		}
		return base64Str, code, err
	} else if numeroCuadro == "3" {
		base64Str, code, err := s.GetReportNursingFrame3LowCode(month, year)
		if err != nil {
			logger.Error.Println("Error generating report nursing frame 3:", err)
			return "", code, err
		}
		return base64Str, code, err
	} else if numeroCuadro == "4" {
		base64Str, code, err := s.GetReportNursingFrame4LowCode(month, year)
		if err != nil {
			logger.Error.Println("Error generating report nursing frame 4:", err)
			return "", code, err
		}
		return base64Str, code, err
	} else if numeroCuadro == "5" {
		base64Str, code, err := s.GetReportNursingFrame5LowCode(month, year)
		if err != nil {
			logger.Error.Println("Error generating report nursing frame 5:", err)
			return "", code, err
		}
		return base64Str, code, err
	} else if numeroCuadro == "6" {
		base64Str, code, err := s.GetReportNursingFrame6LowCode(month, year)
		if err != nil {
			logger.Error.Println("Error generating report nursing frame 6:", err)
			return "", code, err
		}
		return base64Str, code, err
	} else if numeroCuadro == "7" {
		base64Str, code, err := s.GetReportNursingFrame7LowCode(month, year)
		if err != nil {
			logger.Error.Println("Error generating report nursing frame 7:", err)
			return "", code, err
		}
		return base64Str, code, err
	} else if numeroCuadro == "8" {
		base64Str, code, err := s.GetReportNursingFrame8LowCode(month, year)
		if err != nil {
			logger.Error.Println("Error generating report nursing frame 8:", err)
			return "", code, err
		}
		return base64Str, code, err
	} else if numeroCuadro == "9" {
		base64Str, code, err := s.GetReportNursingFrame9LowCode(startYear, endYear)
		if err != nil {
			logger.Error.Println("Error generating report nursing frame 9:", err)
			return "", code, err
		}
		return base64Str, code, err
	}

	return "", 1, errors.New("invalid numeroCuadro")

}

func (s *ReportsMedicalAreaService) GetReportDentistryFrameLowCode(numeroCuadro, month, year string) (string, int, error) {

	if numeroCuadro == "" {
		return "", 1, errors.New("missing numeroCuadro")
	}
	if month == "" {
		month = "1,2,3"
	}

	if year == "" {
		year = time.Now().Format("2006")
	}

	if numeroCuadro == "1" {
		base64Str, code, err := s.GetReportDentistryFrame1LowCode(month, year)
		if err != nil {
			logger.Error.Println("Error generating report dentistry frame 1:", err)
			return "", code, err
		}
		return base64Str, code, err
	} else if numeroCuadro == "2" {
		base64Str, code, err := s.GetReportDentistryFrame2LowCode(month, year)
		if err != nil {
			logger.Error.Println("Error generating report dentistry frame 2:", err)
			return "", code, err
		}
		return base64Str, code, err
	}

	return "", 1, errors.New("invalid numeroCuadro")

}

func (s *ReportsMedicalAreaService) GetReportMedicalFrameLowCode(numeroCuadro, month, year string) (string, int, error) {

	if numeroCuadro == "" {
		return "", 1, errors.New("missing numeroCuadro")
	}
	if month == "" {
		month = "1,2,3"
	}

	if year == "" {
		year = time.Now().Format("2006")
	}

	if numeroCuadro == "1" {
		base64Str, code, err := s.GetReportMedicalFrame1LowCode(month, year)
		if err != nil {
			logger.Error.Println("Error generating report medical frame 1:", err)
			return "", code, err
		}
		return base64Str, code, err
	} else if numeroCuadro == "3" {
		base64Str, code, err := s.GetReportMedicalFrame3LowCode(month, year)
		if err != nil {
			logger.Error.Println("Error generating report medical frame 3:", err)
			return "", code, err
		}
		return base64Str, code, err
	} else if numeroCuadro == "8" {
		base64Str, code, err := s.GetReportMedicalFrame8LowCode(month, year)
		if err != nil {
			logger.Error.Println("Error generating report medical frame 8:", err)
			return "", code, err
		}
		return base64Str, code, err
	} else if numeroCuadro == "12" {
		base64Str, code, err := s.GetReportMedicalFrame12LowCode(month, year)
		if err != nil {
			logger.Error.Println("Error generating report medical frame 12:", err)
			return "", code, err
		}
		return base64Str, code, err
	}

	return "", 1, errors.New("invalid numeroCuadro")

}

func (s *ReportsMedicalAreaService) GetReportNursingFrame1LowCode(month, year string) (string, int, error) {

	f := excelize.NewFile()
	sheet := f.GetSheetName(0)

	styleHeaderID, styleBorderID := styleExcel(f)

	trimester := getTrimester(month)

	headerMedicalArea := models.HeaderMedicalAreaExcel{
		Frame: "CUADRO N° 01",
		Title: fmt.Sprintf("ATENCIONES DE CONSULTAS DE ENFERMERIA A ESTUDIANTES, SEGÚN ESCUELA PROFESIONAL - %s AÑO %s", trimester, year),
		Area:  "ESCUELA PROFESIONAL",
	}

	f = HeaderNursingConsultingTrimesterExcel(f, sheet, headerMedicalArea, styleHeaderID, month, trimester, year)

	for i, school := range schools {
		cell := fmt.Sprintf("A%d", 6+i)
		cellEnd := fmt.Sprintf("M%d", 6+i)
		_ = f.SetCellValue(sheet, cell, school)
		_ = f.SetCellStyle(sheet, cell, cellEnd, styleBorderID)
	}

	_ = f.SetColWidth(sheet, "A", "A", 50)

	dateStart, dateEnd := getDateRange(month, year)

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)
	integralAttentions, err := srv.SrvConsultationIntegralAttention.GetConsultationIntegralAttentionExcel(dateStart, dateEnd)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get performed procedures:", err)
		return "", 15, err
	}

	startRow := 6
	totalRow := 5

	type DateRange struct {
		Start time.Time
		End   time.Time
	}

	startTime, _ := time.Parse("2006-01-02", dateStart)
	endTime, _ := time.Parse("2006-01-02", dateEnd)
	nextStart := startTime.AddDate(0, 1, 0)
	prevEnd := endTime.AddDate(0, -1, 0)

	dateRanges := []DateRange{
		{Start: startTime, End: nextStart},
		{Start: nextStart, End: prevEnd},
		{Start: prevEnd, End: endTime},
	}

	columnTotalTrimester := []string{"B", "C", "D"}
	columnsTotalMonthStart := []string{"E", "H", "K"}

	shouldCount := func(attn *models.ConsultationIntegralAttentionExcel, school string, sexo string) bool {
		if attn.TipoPersona != "Estudiante" {
			return false
		}
		if attn.EscuelaProfesional != school {
			return false
		}
		if sexo == "" {
			return true
		}
		return attn.Sexo == sexo
	}

	for i, sexo := range []string{"", "M", "F"} {
		totalCount := 0

		for j, school := range schools {
			row := startRow + j
			count := 0

			for _, attn := range integralAttentions {
				if shouldCount(attn, school, sexo) {
					count++
					totalCount++
				}
			}
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[i], row), count)
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[i], totalRow), totalCount)
	}

	for monthIdx, dateRange := range dateRanges {
		for i, sexo := range []string{"", "M", "F"} {
			colNumStart, _ := excelize.ColumnNameToNumber(columnsTotalMonthStart[monthIdx])
			col, _ := excelize.ColumnNumberToName(colNumStart + i)
			totalCount := 0

			for j, school := range schools {
				row := startRow + j
				count := 0

				for _, attn := range integralAttentions {
					date, err := time.Parse("2006-01-02", attn.FechaConsulta[:10])
					if err != nil {
						logger.Error.Println(s.txID, " - couldn't parse date:", err)
						return "", 15, err
					}

					if (date.After(dateRange.Start)) && date.Before(dateRange.End) {
						if shouldCount(attn, school, sexo) {
							count++
							totalCount++
						}
					}
				}
				_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, row), count)
			}
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, totalRow), totalCount)
		}
	}

	base64Str, code, err := SaveExcelAndReturnBase64(f, "enfermeria_cuadro_01")

	return base64Str, code, err
}

func (s *ReportsMedicalAreaService) GetReportNursingFrame2LowCode(month, year string) (string, int, error) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)
	styleHeaderID, styleBorderID := styleExcel(f)

	trimester := getTrimester(month)

	headerMedicalArea := models.HeaderMedicalAreaExcel{
		Frame: "CUADRO N° 02",
		Title: fmt.Sprintf("PROCEDIMIENTOS DE ENFERMERIA REALIZADOS A ESTUDIANTES, SEGÚN ESCUELA PROFESIONAL Y MESES - %s AÑO %s", trimester, year),
		Area:  "ESCUELA PROFESIONAL",
	}

	f = HeaderNursingProceduresTrimesterExcel(f, sheet, headerMedicalArea, styleHeaderID, month, trimester, year)

	for i, sName := range schools {
		row := 7 + i
		start, end := fmt.Sprintf("A%d", row), fmt.Sprintf("AO%d", row)
		_ = f.SetCellValue(sheet, start, sName)
		_ = f.SetCellStyle(sheet, start, end, styleBorderID)
	}

	_ = f.SetColWidth(sheet, "A", "A", 50)

	dateStart, dateEnd := getDateRange(month, year)

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)
	performedProcedures, err := srv.SrvNursingConsultationPerformedProcedures.GetPerformedProceduresExcel(dateStart, dateEnd)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get performed procedures:", err)
		return "", 15, err
	}

	totalRow := 6
	startRow := 7

	type ProcedureInfo struct {
		Escuela string
		Fecha   time.Time
		Tipo    string
		Sexo    string
	}

	type DateRange struct {
		Start time.Time
		End   time.Time
	}

	sNameIdx := make(map[string]int)
	for idx, s := range schools {
		sNameIdx[s] = idx
	}

	startTime, _ := time.Parse("2006-01-02", dateStart)
	endTime, _ := time.Parse("2006-01-02", dateEnd)
	nextStart := startTime.AddDate(0, 1, 0)
	prevEnd := endTime.AddDate(0, -1, 0)

	dateRanges := []DateRange{
		{Start: startTime, End: nextStart},
		{Start: nextStart, End: prevEnd},
		{Start: prevEnd, End: endTime},
	}

	columnsTotalMonth := []string{"C", "P", "AC"}
	columnsStartProcedures := []string{"D", "Q", "AD"}
	columnTotalTrimester := "B"

	totalTrimester := make([]int, len(schools))
	monthlyTotals := make([][]int, 3)
	for i := 0; i < 3; i++ {
		monthlyTotals[i] = make([]int, len(schools))
	}

	knownProcedures := make(map[string]struct{})
	for _, proc := range nursingProceduresComparation[:len(nursingProceduresComparation)-1] {
		knownProcedures[proc] = struct{}{}
	}

	totalProceduresM := make([][][]int, 3) // trimestre -> procedimiento -> escuela -> M
	totalProceduresF := make([][][]int, 3) // trimestre -> procedimiento -> escuela -> F

	for i := range totalProceduresM {
		totalProceduresM[i] = make([][]int, len(nursingProceduresComparation))
		totalProceduresF[i] = make([][]int, len(nursingProceduresComparation))
		for j := range nursingProceduresComparation {
			totalProceduresM[i][j] = make([]int, len(schools))
			totalProceduresF[i][j] = make([]int, len(schools))
		}
	}

	for _, p := range performedProcedures {
		if p.TipoPersona != "Estudiante" {
			continue
		}
		date, err := time.Parse("2006-01-02", p.FechaConsulta[:10])
		if err != nil {
			logger.Error.Println(s.txID, " - couldn't parse date:", err)
			return "", 15, err
		}

		idxSchool, ok := sNameIdx[p.EscuelaProfesional]
		if !ok {
			continue
		}

		totalTrimester[idxSchool]++

		for idxRange, dr := range dateRanges {
			if date.After(dr.Start) && date.Before(dr.End) {
				monthlyTotals[idxRange][idxSchool]++

				for idxProc, proc := range nursingProceduresComparation {
					if proc == "OTROS" {
						if _, known := knownProcedures[p.TipoProcedimiento]; known {
							continue
						}
					} else {
						if proc != p.TipoProcedimiento {
							continue
						}
					}

					if p.Sexo == "M" {
						totalProceduresM[idxRange][idxProc][idxSchool]++
					} else {
						totalProceduresF[idxRange][idxProc][idxSchool]++
					}
				}
				break
			}
		}
	}

	// Escribir procedimientos totales de trimestre
	total := 0
	for idx, cnt := range totalTrimester {
		row := startRow + idx
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester, row), cnt)
		total += cnt
	}
	_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester, totalRow), total)

	// Escribir procedimientos totales por mes
	for i, col := range columnsTotalMonth {
		total = 0
		for j, cnt := range monthlyTotals[i] {
			row := startRow + j
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, row), cnt)
			total += cnt
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, totalRow), total)
	}

	// Escribir procedimientos
	for idxTrim, colStart := range columnsStartProcedures {
		startColNum, _ := excelize.ColumnNameToNumber(colStart)
		for idxProc := range nursingProceduresComparation {
			colM, _ := excelize.ColumnNumberToName(startColNum + idxProc*2)
			colF, _ := excelize.ColumnNumberToName(startColNum + idxProc*2 + 1)

			totalM, totalF := 0, 0
			for idxSchool := range schools {
				row := startRow + idxSchool
				m := totalProceduresM[idxTrim][idxProc][idxSchool]
				fem := totalProceduresF[idxTrim][idxProc][idxSchool]

				_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colM, row), m)
				_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colF, row), fem)

				totalM += m
				totalF += fem
			}
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colM, totalRow), totalM)
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colF, totalRow), totalF)
		}
	}

	base64Str, code, err := SaveExcelAndReturnBase64(f, "enfermeria_cuadro_02")

	return base64Str, code, err
}

func (s *ReportsMedicalAreaService) GetReportNursingFrame3LowCode(month, year string) (string, int, error) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)
	styleHeaderID, styleBorderID := styleExcel(f)

	trimester := getTrimester(month)

	headerMedicalArea := models.HeaderMedicalAreaExcel{
		Frame: "CUADRO N° 03",
		Title: fmt.Sprintf("PROCEDIMIENTOS DE ENFERMERIA REALIZADOS A ESTUDIANTES, SEGÚN ESCUELA PROFESIONAL - %s AÑO %s", trimester, year),
		Area:  "ESCUELA PROFESIONAL",
	}

	f = HeaderNursingProceduresExcel(f, sheet, headerMedicalArea, styleHeaderID, trimester)

	for i, sName := range schools {
		row := 7 + i
		start, end := fmt.Sprintf("A%d", row), fmt.Sprintf("N%d", row)
		_ = f.SetCellValue(sheet, start, sName)
		_ = f.SetCellStyle(sheet, start, end, styleBorderID)
	}

	_ = f.SetColWidth(sheet, "A", "A", 50)

	dateStart, dateEnd := getDateRange(month, year)

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)
	performedProcedures, err := srv.SrvNursingConsultationPerformedProcedures.GetPerformedProceduresExcel(dateStart, dateEnd)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get performed procedures:", err)
		return "", 15, err
	}

	totalRow := 6
	startRow := 7

	startTime, _ := time.Parse("2006-01-02", dateStart)
	endTime, _ := time.Parse("2006-01-02", dateEnd)

	columnTotalTrimester := "B"
	columnsStartProcedures := "C"

	knownProcedures := make(map[string]struct{})
	for _, proc := range nursingProceduresComparation[:len(nursingProceduresComparation)-1] {
		knownProcedures[proc] = struct{}{}
	}

	totalTrimester := make([]int, len(schools))

	totalProceduresM := make([][]int, len(nursingProceduresComparation))
	totalProceduresF := make([][]int, len(nursingProceduresComparation))

	for i := range nursingProceduresComparation {
		totalProceduresM[i] = make([]int, len(schools))
		totalProceduresF[i] = make([]int, len(schools))
	}

	sNameIdx := make(map[string]int)
	for idx, s := range schools {
		sNameIdx[s] = idx
	}

	for _, p := range performedProcedures {
		if p.TipoPersona != "Estudiante" {
			continue
		}
		date, err := time.Parse("2006-01-02", p.FechaConsulta[:10])
		if err != nil {
			logger.Error.Println(s.txID, " - couldn't parse date:", err)
			return "", 15, err
		}

		idxSchool, ok := sNameIdx[p.EscuelaProfesional]
		if !ok {
			continue
		}

		totalTrimester[idxSchool]++

		if date.After(startTime) && date.Before(endTime) {
			for idxProc, proc := range nursingProceduresComparation {
				if proc == "OTROS" {
					if _, known := knownProcedures[p.TipoProcedimiento]; known {
						continue
					}
				} else {
					if proc != p.TipoProcedimiento {
						continue
					}
				}

				if p.Sexo == "M" {
					totalProceduresM[idxProc][idxSchool]++
				} else {
					totalProceduresF[idxProc][idxSchool]++
				}
			}
		}

	}

	// Escribir procedimientos totales de trimestre
	countTotal := 0
	for i, cnt := range totalTrimester {
		row := startRow + i
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester, row), cnt)
		countTotal += cnt
	}
	_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester, totalRow), countTotal)

	// Escribir procedimientos
	startColNum, _ := excelize.ColumnNameToNumber(columnsStartProcedures)
	for idxProc := range nursingProceduresComparation {
		colM, _ := excelize.ColumnNumberToName(startColNum + idxProc*2)
		colF, _ := excelize.ColumnNumberToName(startColNum + idxProc*2 + 1)

		totalM, totalF := 0, 0
		for idxSchool := range schools {
			row := startRow + idxSchool
			m := totalProceduresM[idxProc][idxSchool]
			fem := totalProceduresF[idxProc][idxSchool]

			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colM, row), m)
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colF, row), fem)

			totalM += m
			totalF += fem
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colM, totalRow), totalM)
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colF, totalRow), totalF)
	}

	base64Str, code, err := SaveExcelAndReturnBase64(f, "enfermeria_cuadro_03")

	return base64Str, code, err
}

func (s *ReportsMedicalAreaService) GetReportNursingFrame4LowCode(month, year string) (string, int, error) {

	f := excelize.NewFile()
	sheet := f.GetSheetName(0)

	styleHeaderID, styleBorderID := styleExcel(f)

	trimester := getTrimester(month)

	headerMedicalArea := models.HeaderMedicalAreaExcel{
		Frame: "CUADRO N° 04",
		Title: fmt.Sprintf("ATENCIONES DE CONSULTAS DE ENFERMERIA POR MES Y SEXO, SEGÚN TIPO DE PERSONAL - %s AÑO %s", trimester, year),
		Area:  "PERSONAL",
	}

	f = HeaderNursingConsultingTrimesterExcel(f, sheet, headerMedicalArea, styleHeaderID, month, trimester, year)

	for i, sName := range nursingStaffs {
		row := 6 + i
		start, end := fmt.Sprintf("A%d", row), fmt.Sprintf("M%d", row)
		_ = f.SetCellValue(sheet, start, sName)
		_ = f.SetCellStyle(sheet, start, end, styleBorderID)
	}

	_ = f.SetColWidth(sheet, "A", "A", 20)

	dateStart, dateEnd := getDateRange(month, year)

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)
	integralAttentions, err := srv.SrvConsultationIntegralAttention.GetConsultationIntegralAttentionExcel(dateStart, dateEnd)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get performed procedures:", err)
		return "", 15, err
	}

	startRow := 6
	totalRow := 5

	type DateRange struct {
		Start time.Time
		End   time.Time
	}

	startTime, _ := time.Parse("2006-01-02", dateStart)
	endTime, _ := time.Parse("2006-01-02", dateEnd)
	nextStart := startTime.AddDate(0, 1, 0)
	prevEnd := endTime.AddDate(0, -1, 0)

	dateRanges := []DateRange{
		{Start: startTime, End: nextStart},
		{Start: nextStart, End: prevEnd},
		{Start: prevEnd, End: endTime},
	}

	columnTotalTrimester := []string{"B", "C", "D"}
	columnsTotalMonthStart := []string{"E", "H", "K"}

	shouldCount := func(attn *models.ConsultationIntegralAttentionExcel, staff string, sexo string) bool {
		if attn.TipoPersona == "Estudiante" {
			return false
		}
		if attn.TipoPersona != staff {
			return false
		}
		if sexo == "" {
			return true
		}
		return attn.Sexo == sexo
	}

	for i, sexo := range []string{"", "M", "F"} {
		totalCount := 0

		for j, staff := range nursingStaffsComparation {
			row := startRow + j
			count := 0

			for _, attn := range integralAttentions {
				if shouldCount(attn, staff, sexo) {
					count++
					totalCount++
				}
			}
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[i], row), count)
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[i], totalRow), totalCount)
	}

	for monthIdx, dateRange := range dateRanges {
		for i, sexo := range []string{"", "M", "F"} {
			colNumStart, _ := excelize.ColumnNameToNumber(columnsTotalMonthStart[monthIdx])
			col, _ := excelize.ColumnNumberToName(colNumStart + i)
			totalCount := 0

			for j, staff := range nursingStaffsComparation {
				row := startRow + j
				count := 0

				for _, attn := range integralAttentions {
					date, err := time.Parse("2006-01-02", attn.FechaConsulta[:10])
					if err != nil {
						logger.Error.Println(s.txID, " - couldn't parse date:", err)
						return "", 15, err
					}

					if (date.Equal(dateRange.Start) || date.After(dateRange.Start)) && date.Before(dateRange.End) {
						if shouldCount(attn, staff, sexo) {
							count++
							totalCount++
						}
					}
				}
				_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, row), count)
			}
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, totalRow), totalCount)
		}
	}

	base64Str, code, err := SaveExcelAndReturnBase64(f, "enfermeria_cuadro_04")

	return base64Str, code, err
}

func (s *ReportsMedicalAreaService) GetReportNursingFrame5LowCode(month, year string) (string, int, error) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)
	styleHeaderID, styleBorderID := styleExcel(f)

	trimester := getTrimester(month)

	headerMedicalArea := models.HeaderMedicalAreaExcel{
		Frame: "CUADRO N° 05",
		Title: fmt.Sprintf("ATENCIONES DE ENFERMERIA POR PROCEDIMIENTOS REALIZADOS SEGÚN PERSONAL ATENDIDO Y MES - %s AÑO %s", trimester, year),
		Area:  "PERSONAL",
	}

	f = HeaderNursingProceduresTrimesterExcel(f, sheet, headerMedicalArea, styleHeaderID, month, trimester, year)

	for i, sName := range nursingStaffs {
		row := 7 + i
		start, end := fmt.Sprintf("A%d", row), fmt.Sprintf("AO%d", row)
		_ = f.SetCellValue(sheet, start, sName)
		_ = f.SetCellStyle(sheet, start, end, styleBorderID)
	}

	_ = f.SetColWidth(sheet, "A", "A", 20)

	dateStart, dateEnd := getDateRange(month, year)

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)
	performedProcedures, err := srv.SrvNursingConsultationPerformedProcedures.GetPerformedProceduresExcel(dateStart, dateEnd)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get performed procedures:", err)
		return "", 15, err
	}

	totalRow := 6
	startRow := 7

	type ProcedureInfo struct {
		Escuela string
		Fecha   time.Time
		Tipo    string
		Sexo    string
	}

	type DateRange struct {
		Start time.Time
		End   time.Time
	}

	sNameIdx := make(map[string]int)
	for idx, s := range nursingStaffsComparation {
		sNameIdx[s] = idx
	}

	startTime, _ := time.Parse("2006-01-02", dateStart)
	endTime, _ := time.Parse("2006-01-02", dateEnd)
	nextStart := startTime.AddDate(0, 1, 0)
	prevEnd := endTime.AddDate(0, -1, 0)

	dateRanges := []DateRange{
		{Start: startTime, End: nextStart},
		{Start: nextStart, End: prevEnd},
		{Start: prevEnd, End: endTime},
	}

	columnsTotalMonth := []string{"C", "P", "AC"}
	columnsStartProcedures := []string{"D", "Q", "AD"}
	columnTotalTrimester := "B"

	totalTrimester := make([]int, len(nursingStaffsComparation))
	monthlyTotals := make([][]int, 3)
	for i := 0; i < 3; i++ {
		monthlyTotals[i] = make([]int, len(nursingStaffsComparation))
	}

	knownProcedures := make(map[string]struct{})
	for _, proc := range nursingProceduresComparation[:len(nursingProceduresComparation)-1] {
		knownProcedures[proc] = struct{}{}
	}

	totalProceduresM := make([][][]int, 3) // trimestre -> procedimiento -> escuela -> M
	totalProceduresF := make([][][]int, 3) // trimestre -> procedimiento -> escuela -> F

	for i := range totalProceduresM {
		totalProceduresM[i] = make([][]int, len(nursingProceduresComparation))
		totalProceduresF[i] = make([][]int, len(nursingProceduresComparation))
		for j := range nursingProceduresComparation {
			totalProceduresM[i][j] = make([]int, len(nursingStaffsComparation))
			totalProceduresF[i][j] = make([]int, len(nursingStaffsComparation))
		}
	}

	for _, p := range performedProcedures {
		if p.TipoPersona == "Estudiante" {
			continue
		}

		date, err := time.Parse("2006-01-02", p.FechaConsulta[:10])
		if err != nil {
			logger.Error.Println(s.txID, " - couldn't parse date:", err)
			return "", 15, err
		}

		idxNursingStaff, ok := sNameIdx[p.TipoPersona]
		if !ok {
			continue
		}

		totalTrimester[idxNursingStaff]++

		for idxRange, dr := range dateRanges {
			if date.After(dr.Start) && date.Before(dr.End) {
				monthlyTotals[idxRange][idxNursingStaff]++

				for idxProc, proc := range nursingProceduresComparation {
					if proc == "OTROS" {
						if _, known := knownProcedures[p.TipoProcedimiento]; known {
							continue
						}
					} else {
						if proc != p.TipoProcedimiento {
							continue
						}
					}

					if p.Sexo == "M" {
						totalProceduresM[idxRange][idxProc][idxNursingStaff]++
					} else {
						totalProceduresF[idxRange][idxProc][idxNursingStaff]++
					}
				}
				break
			}
		}
	}

	// Escribir procedimientos totales de trimestre
	total := 0
	for idx, cnt := range totalTrimester {
		row := startRow + idx
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester, row), cnt)
		total += cnt
	}
	_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester, totalRow), total)

	// Escribir procedimientos totales por mes
	for i, col := range columnsTotalMonth {
		total = 0
		for j, cnt := range monthlyTotals[i] {
			row := startRow + j
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, row), cnt)
			total += cnt
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, totalRow), total)
	}

	// Escribir procedimientos
	for idxTrim, colStart := range columnsStartProcedures {
		startColNum, _ := excelize.ColumnNameToNumber(colStart)
		for idxProc := range nursingProceduresComparation {
			colM, _ := excelize.ColumnNumberToName(startColNum + idxProc*2)
			colF, _ := excelize.ColumnNumberToName(startColNum + idxProc*2 + 1)

			totalM, totalF := 0, 0
			for idxNursingStaff := range nursingStaffsComparation {
				row := startRow + idxNursingStaff
				m := totalProceduresM[idxTrim][idxProc][idxNursingStaff]
				fem := totalProceduresF[idxTrim][idxProc][idxNursingStaff]

				_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colM, row), m)
				_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colF, row), fem)

				totalM += m
				totalF += fem
			}
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colM, totalRow), totalM)
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colF, totalRow), totalF)
		}
	}

	base64Str, code, err := SaveExcelAndReturnBase64(f, "enfermeria_cuadro_05")

	return base64Str, code, err
}

func (s *ReportsMedicalAreaService) GetReportNursingFrame6LowCode(month, year string) (string, int, error) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)
	styleHeaderID, styleBorderID := styleExcel(f)

	headerMedicalArea := models.HeaderMedicalAreaExcel{
		Frame: "CUADRO N° 06",
		Title: fmt.Sprintf("ATENCIONES DE ENFERMERÍA A ESTUDIANTES POR TRIMESTRE Y SEXO, SEGÚN ESCUELA PROFESIONAL - AÑO %s", year),
		Area:  "ESCUELA PROFESIONAL",
	}

	f = HeaderNursingConsultingYearlyExcel(f, sheet, headerMedicalArea, styleHeaderID, year)

	for i, sName := range schools {
		row := 6 + i
		start, end := fmt.Sprintf("A%d", row), fmt.Sprintf("P%d", row)
		_ = f.SetCellValue(sheet, start, sName)
		_ = f.SetCellStyle(sheet, start, end, styleBorderID)
	}

	_ = f.SetColWidth(sheet, "A", "A", 50)

	nextYear, _ := strconv.Atoi(year)
	nextYear++

	dateStart := fmt.Sprintf("%s-01-01", year)
	dateEnd := fmt.Sprintf("%d-01-01", nextYear)

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)
	integralAttentions, err := srv.SrvConsultationIntegralAttention.GetConsultationIntegralAttentionExcel(dateStart, dateEnd)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get performed procedures:", err)
		return "", 15, err
	}

	totalRow := 5
	startRow := 6

	type DateRange struct {
		Start time.Time
		End   time.Time
	}

	parseDate := func(dateStr string) time.Time {
		t, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			logger.Error.Println(s.txID, " - couldn't parse date:", err)
		}
		return t
	}

	getDateRangeParsed := func(months string, year string) DateRange {
		start, end := getDateRange(months, year)
		return DateRange{
			Start: parseDate(start),
			End:   parseDate(end),
		}
	}

	dateRanges := []DateRange{
		getDateRangeParsed("1,2,3", year),
		getDateRangeParsed("4,5,6", year),
		getDateRangeParsed("7,8,9", year),
		getDateRangeParsed("10,11,12", year),
	}

	startTime := parseDate(dateStart)
	endTime := parseDate(dateEnd)

	columnTotalTrimester := []string{"B", "C", "D"}
	columnsTrimesterStart := []string{"E", "H", "K", "N"}

	type PreprocessedAttention struct {
		*models.ConsultationIntegralAttentionExcel
		ParsedDate time.Time
	}

	var processedAttentions []PreprocessedAttention
	for _, attn := range integralAttentions {
		parsed := parseDate(attn.FechaConsulta[:10])
		processedAttentions = append(processedAttentions, PreprocessedAttention{attn, parsed})
	}

	shouldCount := func(attn PreprocessedAttention, school string, sexo string, dateStart, dateEnd time.Time) bool {
		if attn.TipoPersona != "Estudiante" || attn.EscuelaProfesional != school {
			return false
		}
		if !(attn.ParsedDate.After(dateStart) && attn.ParsedDate.Before(dateEnd)) {
			return false
		}
		if sexo == "" {
			return true
		}
		return attn.Sexo == sexo
	}

	// Primera parte: total general
	for i, sexo := range []string{"", "M", "F"} {
		totalCount := 0
		for j, school := range schools {
			row := startRow + j
			count := 0
			for _, attn := range processedAttentions {
				if shouldCount(attn, school, sexo, startTime, endTime) {
					count++
					totalCount++
				}
			}
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[i], row), count)
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[i], totalRow), totalCount)
	}

	// Segunda parte: por trimestre
	for trimesterIdx, dateRange := range dateRanges {
		for i, sexo := range []string{"", "M", "F"} {
			colNumStart, _ := excelize.ColumnNameToNumber(columnsTrimesterStart[trimesterIdx])
			col, _ := excelize.ColumnNumberToName(colNumStart + i)
			totalCount := 0

			for j, school := range schools {
				row := startRow + j
				count := 0
				for _, attn := range processedAttentions {
					if shouldCount(attn, school, sexo, dateRange.Start, dateRange.End) {
						count++
						totalCount++
					}
				}
				_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, row), count)
			}
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, totalRow), totalCount)
		}
	}

	base64Str, code, err := SaveExcelAndReturnBase64(f, "enfermeria_cuadro_06")

	return base64Str, code, err
}

func (s *ReportsMedicalAreaService) GetReportNursingFrame7LowCode(month, year string) (string, int, error) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)
	styleHeaderID, styleBorderID := styleExcel(f)

	headerMedicalArea := models.HeaderMedicalAreaExcel{
		Frame: "CUADRO N° 07",
		Title: fmt.Sprintf("ATENCIONES DE ENFERMERÍA POR TRIMESTRE Y SEXO, SEGÚN TIPO DE PERSONAL - AÑO %s", year),
		Area:  "PERSONAL",
	}

	f = HeaderNursingConsultingYearlyExcel(f, sheet, headerMedicalArea, styleHeaderID, year)

	for i, sName := range nursingStaffs {
		row := 6 + i
		start, end := fmt.Sprintf("A%d", row), fmt.Sprintf("P%d", row)
		_ = f.SetCellValue(sheet, start, sName)
		_ = f.SetCellStyle(sheet, start, end, styleBorderID)
	}

	_ = f.SetColWidth(sheet, "A", "A", 20)

	nextYear, _ := strconv.Atoi(year)
	nextYear++

	dateStart := fmt.Sprintf("%s-01-01", year)
	dateEnd := fmt.Sprintf("%d-01-01", nextYear)

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)
	integralAttentions, err := srv.SrvConsultationIntegralAttention.GetConsultationIntegralAttentionExcel(dateStart, dateEnd)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get performed procedures:", err)
		return "", 15, err
	}

	totalRow := 5
	startRow := 6

	type DateRange struct {
		Start time.Time
		End   time.Time
	}

	parseDate := func(dateStr string) time.Time {
		t, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			logger.Error.Println(s.txID, " - couldn't parse date:", err)
		}
		return t
	}

	getDateRangeParsed := func(months string, year string) DateRange {
		start, end := getDateRange(months, year)
		return DateRange{
			Start: parseDate(start),
			End:   parseDate(end),
		}
	}

	dateRanges := []DateRange{
		getDateRangeParsed("1,2,3", year),
		getDateRangeParsed("4,5,6", year),
		getDateRangeParsed("7,8,9", year),
		getDateRangeParsed("10,11,12", year),
	}

	startTime := parseDate(dateStart)
	endTime := parseDate(dateEnd)

	columnTotalTrimester := []string{"B", "C", "D"}
	columnsTrimesterStart := []string{"E", "H", "K", "N"}

	type PreprocessedAttention struct {
		*models.ConsultationIntegralAttentionExcel
		ParsedDate time.Time
	}

	var processedAttentions []PreprocessedAttention
	for _, attn := range integralAttentions {
		parsed := parseDate(attn.FechaConsulta[:10])
		processedAttentions = append(processedAttentions, PreprocessedAttention{attn, parsed})
	}

	shouldCount := func(attn PreprocessedAttention, staff string, sexo string, dateStart, dateEnd time.Time) bool {
		if attn.TipoPersona == "Estudiante" {
			return false
		}
		if attn.TipoPersona != staff {
			return false
		}
		if !(attn.ParsedDate.After(dateStart) && attn.ParsedDate.Before(dateEnd)) {
			return false
		}
		if sexo == "" {
			return true
		}
		return attn.Sexo == sexo
	}

	// Primera parte: total general
	for i, sexo := range []string{"", "M", "F"} {
		totalCount := 0
		for j, staff := range nursingStaffsComparation {
			row := startRow + j
			count := 0
			for _, attn := range processedAttentions {
				if shouldCount(attn, staff, sexo, startTime, endTime) {
					count++
					totalCount++
				}
			}
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[i], row), count)
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[i], totalRow), totalCount)
	}

	// Segunda parte: por trimestre
	for trimesterIdx, dateRange := range dateRanges {
		for i, sexo := range []string{"", "M", "F"} {
			colNumStart, _ := excelize.ColumnNameToNumber(columnsTrimesterStart[trimesterIdx])
			col, _ := excelize.ColumnNumberToName(colNumStart + i)
			totalCount := 0

			for j, staff := range nursingStaffsComparation {
				row := startRow + j
				count := 0
				for _, attn := range processedAttentions {
					if shouldCount(attn, staff, sexo, dateRange.Start, dateRange.End) {
						count++
						totalCount++
					}
				}
				_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, row), count)
			}
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, totalRow), totalCount)
		}
	}

	base64Str, code, err := SaveExcelAndReturnBase64(f, "enfermeria_cuadro_07")

	return base64Str, code, err
}

func (s *ReportsMedicalAreaService) GetReportNursingFrame8LowCode(month, year string) (string, int, error) {

	f := excelize.NewFile()
	sheet := f.GetSheetName(0)
	styleHeaderID, styleBorderID := styleExcel(f)

	headerMedicalArea := models.HeaderMedicalAreaExcel{
		Frame: "CUADRO N° 08",
		Title: fmt.Sprintf("ATENCIONES DE ENFERMERIA A ESTUDIANTES POR MES, SEGÚN PROCEDIMIENTOS REALIZADOS - AÑO %s", year),
		Area:  "PROCEDIMIENTOS",
	}

	f = HeaderNursingProceduresYearlyExcel(f, sheet, headerMedicalArea, styleHeaderID)

	for i, sName := range nursingProcedures {
		row := 6 + i
		start, end := fmt.Sprintf("A%d", row), fmt.Sprintf("N%d", row)
		_ = f.SetCellValue(sheet, start, sName)
		_ = f.SetCellStyle(sheet, start, end, styleBorderID)
	}

	_ = f.SetColWidth(sheet, "A", "A", 30)

	nextYear, _ := strconv.Atoi(year)
	nextYear++

	dateStart := fmt.Sprintf("%s-01-01", year)
	dateEnd := fmt.Sprintf("%d-01-01", nextYear)

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)
	performedProcedures, err := srv.SrvNursingConsultationPerformedProcedures.GetPerformedProceduresExcel(dateStart, dateEnd)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get performed procedures:", err)
		return "", 15, err
	}

	totalRow := 5
	startRow := 6

	type DateRange struct {
		Start time.Time
		End   time.Time
	}

	parseDate := func(dateStr string) (time.Time, error) {
		t, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			logger.Error.Println(s.txID, " - couldn't parse date:", err)
		}
		return t, err
	}

	var dateRanges []DateRange
	dateRanges = make([]DateRange, 12)
	for i := 0; i < 12; i++ {
		tStart := fmt.Sprintf("%s-%02d-01", year, i+1)
		dStart, err := parseDate(tStart)
		if err != nil {
			return "", 15, err
		}
		dEnd := dStart.AddDate(0, 1, 0)
		dateRanges[i] = DateRange{Start: dStart, End: dEnd}
	}

	columnTotal := "B"
	columnsMonths := []string{"C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N"}

	countProcedures := func(sName string, dateRange *DateRange) int {
		count := 0
		for _, p := range performedProcedures {

			if p.TipoProcedimiento != sName || p.TipoPersona != "Estudiante" {
				continue
			}

			if dateRange == nil {
				count++
				continue
			}

			date, err := time.Parse("2006-01-02", p.FechaConsulta[:10])
			if err != nil {
				logger.Error.Println(s.txID, " - couldn't parse date:", err)
				continue
			}

			if date.After(dateRange.Start) && date.Before(dateRange.End) {
				count++
			}
		}
		return count
	}

	// Escribir procedimientos totales
	countTotal := 0
	for i, sName := range nursingProcedures {
		row := startRow + i
		count := countProcedures(sName, nil)
		countTotal += count
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotal, row), count)
	}
	_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotal, totalRow), countTotal)

	// Escribir procedimientos de cada mes
	for idx, col := range columnsMonths {
		countTotal = 0
		for i, sName := range nursingProcedures {
			row := startRow + i
			count := countProcedures(sName, &dateRanges[idx])
			countTotal += count
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, row), count)
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, totalRow), countTotal)
	}

	base64Str, code, err := SaveExcelAndReturnBase64(f, "enfermeria_cuadro_08")

	return base64Str, code, err
}

func (s *ReportsMedicalAreaService) GetReportNursingFrame9LowCode(startYear, endYear string) (string, int, error) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)
	styleHeaderID, styleBorderID := styleExcel(f)

	headerMedicalArea := models.HeaderMedicalAreaExcel{
		Frame: "CUADRO N° 09",
		Title: fmt.Sprintf("EVOLUCIÓN DE ATENCIONES DE SERVICIO DE TÓPICO POR PROCEDIMIENTO REALIZADO, SEGÚN ESCUELA PROFESIONAL - AÑOS: %s - %s", startYear, endYear),
		Area:  "ESCUELA PROFESIONAL",
	}

	columnEnd := ""
	f, columnEnd = HeaderNursingProceduresYearsExcel(f, sheet, headerMedicalArea, styleHeaderID, startYear, endYear)

	for i, sName := range schools {
		row := 6 + i
		start, end := fmt.Sprintf("A%d", row), fmt.Sprintf("%s%d", columnEnd, row)
		_ = f.SetCellValue(sheet, start, sName)
		_ = f.SetCellStyle(sheet, start, end, styleBorderID)
	}

	_ = f.SetColWidth(sheet, "A", "A", 50)

	iStartYear, _ := strconv.Atoi(startYear)
	iEndYear, _ := strconv.Atoi(endYear)

	dateStart := fmt.Sprintf("%s-01-01", startYear)
	dateEnd := fmt.Sprintf("%s-01-01", endYear)

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)
	performedProcedures, err := srv.SrvNursingConsultationPerformedProcedures.GetPerformedProceduresExcel(dateStart, dateEnd)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get performed procedures:", err)
		return "", 15, err
	}

	totalRow := 5
	startRow := 6

	type DateRange struct {
		Start time.Time
		End   time.Time
	}

	parseDate := func(dateStr string) (time.Time, error) {
		t, err := time.Parse("2006-01-02", dateStr)
		if err != nil {
			logger.Error.Println(s.txID, " - couldn't parse date:", err)
		}
		return t, err
	}

	var dateRanges []DateRange
	dateRanges = make([]DateRange, (iEndYear-iStartYear)+1)
	idx := 0
	for i := iStartYear; i <= iEndYear; i++ {
		tStart := fmt.Sprintf("%d-01-01", i)
		dStart, err := parseDate(tStart)
		if err != nil {
			return "", 15, err
		}
		dEnd := dStart.AddDate(1, 0, 0)
		dateRanges[idx] = DateRange{Start: dStart, End: dEnd}
		idx++
	}

	columnStart := "B"
	col := ""

	for i, date := range dateRanges {
		countTotal := 0

		for j, sName := range schools {
			count := 0
			row := startRow + j
			colNum, _ := excelize.ColumnNameToNumber(columnStart)
			col, _ = excelize.ColumnNumberToName(colNum + i)
			for _, p := range performedProcedures {
				if p.TipoPersona != "Estudiante" || p.EscuelaProfesional != sName {
					continue
				}

				start, err := parseDate(p.FechaConsulta[:10])
				if err != nil {
					logger.Error.Println(s.txID, " - couldn't parse date:", err)
					return "", 15, err
				}

				if start.After(date.Start) && start.Before(date.End) {
					count++
				}

			}
			countTotal += count
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, row), count)
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, totalRow), countTotal)
	}

	base64Str, code, err := SaveExcelAndReturnBase64(f, "enfermeria_cuadro_09")

	return base64Str, code, err
}

func (s *ReportsMedicalAreaService) GetReportDentistryFrame1LowCode(month, year string) (string, int, error) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)
	styleHeaderID, styleBorderID := styleExcel(f)

	trimester := getTrimester(month)

	headerMedicalArea := models.HeaderMedicalAreaExcel{
		Frame: "CUADRO N° 01",
		Title: fmt.Sprintf("ATENCIONES DE SERVICIO DE ODONTOLOGÍA POR PROCEDIMIENTOS REALIZADO, SEGÚN TIPO DE PERSONAL - %s", trimester),
		Area:  "PERSONAL",
	}

	f = HeaderDentistryProceduresTrimesterExcel(f, sheet, headerMedicalArea, styleHeaderID)

	for i, sName := range dentistryStaffs {
		row := 7 + i
		start, end := fmt.Sprintf("A%d", row), fmt.Sprintf("AE%d", row)
		_ = f.SetCellValue(sheet, start, sName)
		_ = f.SetCellStyle(sheet, start, end, styleBorderID)
	}

	cols := getExcelColumns("B", "AE")
	for _, col := range cols {
		_ = f.SetColWidth(sheet, col, col, 4)
	}

	_ = f.SetColWidth(sheet, "A", "A", 28)

	dateStart, dateEnd := getDateRange(month, year)

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)
	performedProcedures, err := srv.SrvDentistryConsultationBuccalProcedure.GetBuccalProceduresExcel(dateStart, dateEnd)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get buccal procedures:", err)
		return "", 15, err
	}

	totalRow := 6
	startRow := 7

	startTime, _ := time.Parse("2006-01-02", dateStart)
	endTime, _ := time.Parse("2006-01-02", dateEnd)

	columnTotalTrimester := []string{"B", "C", "D"}
	columnsStartProcedures := "E"

	knownProcedures := make(map[string]struct{})
	for _, proc := range dentistryProcedures[:len(dentistryProcedures)-1] {
		knownProcedures[proc] = struct{}{}
	}

	totalTrimester := make([]int, len(dentistryStaffsComparation))
	totalTrimesterM := make([]int, len(dentistryStaffsComparation))
	totalTrimesterF := make([]int, len(dentistryStaffsComparation))

	totalProcedures := make([][]int, len(dentistryProcedures))
	totalProceduresM := make([][]int, len(dentistryProcedures))
	totalProceduresF := make([][]int, len(dentistryProcedures))

	for i := range dentistryProcedures {
		totalProcedures[i] = make([]int, len(dentistryStaffsComparation))
		totalProceduresM[i] = make([]int, len(dentistryStaffsComparation))
		totalProceduresF[i] = make([]int, len(dentistryStaffsComparation))
	}

	sNameIdx := make(map[string]int)
	for idx, s := range dentistryStaffsComparation {
		sNameIdx[s] = idx
	}

	for _, p := range performedProcedures {
		if p.TipoPersona == "Estudiante" {
			continue
		}
		date, err := time.Parse("2006-01-02", p.FechaConsulta[:10])
		if err != nil {
			logger.Error.Println(s.txID, " - couldn't parse date:", err)
			return "", 15, err
		}

		idxSchool, ok := sNameIdx[p.TipoPersona]
		if !ok {
			continue
		}

		totalTrimester[idxSchool]++

		if p.Sexo == "M" {
			totalTrimesterM[idxSchool]++
		} else {
			totalTrimesterF[idxSchool]++
		}

		if date.After(startTime) && date.Before(endTime) {
			for idxProc, proc := range dentistryProcedures {
				if proc != p.TipoProcedimiento {
					continue
				}

				totalProcedures[idxProc][idxSchool]++

				if p.Sexo == "M" {
					totalProceduresM[idxProc][idxSchool]++
				} else {
					totalProceduresF[idxProc][idxSchool]++
				}
			}
		}

	}

	// Escribir procedimientos totales de trimestre
	countTotal := 0
	countTotalM := 0
	countTotalF := 0
	for i, cnt := range totalTrimester {
		row := startRow + i
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[0], row), cnt)
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[1], row), totalTrimesterM[i])
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[2], row), totalTrimesterF[i])
		countTotal += cnt
		countTotalM += totalTrimesterM[i]
		countTotalF += totalTrimesterF[i]
	}
	_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[0], totalRow), countTotal)
	_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[1], totalRow), countTotalM)
	_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[2], totalRow), countTotalF)

	// Escribir procedimientos
	startColNum, _ := excelize.ColumnNameToNumber(columnsStartProcedures)
	for idxProc := range dentistryProcedures {
		colT, _ := excelize.ColumnNumberToName(startColNum + idxProc*3)
		colM, _ := excelize.ColumnNumberToName(startColNum + idxProc*3 + 1)
		colF, _ := excelize.ColumnNumberToName(startColNum + idxProc*3 + 2)

		total, totalM, totalF := 0, 0, 0
		for idxSchool := range dentistryStaffsComparation {
			row := startRow + idxSchool
			t := totalProcedures[idxProc][idxSchool]
			m := totalProceduresM[idxProc][idxSchool]
			fem := totalProceduresF[idxProc][idxSchool]

			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colT, row), t)
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colM, row), m)
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colF, row), fem)

			total += t
			totalM += m
			totalF += fem
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colT, totalRow), total)
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colM, totalRow), totalM)
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colF, totalRow), totalF)
	}

	base64Str, code, err := SaveExcelAndReturnBase64(f, "odontologia_cuadro_01")

	return base64Str, code, err
}

func (s *ReportsMedicalAreaService) GetReportDentistryFrame2LowCode(month, year string) (string, int, error) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)
	styleHeaderID, styleBorderID := styleExcel(f)

	trimester := getTrimester(month)

	headerMedicalArea := models.HeaderMedicalAreaExcel{
		Frame: "CUADRO N° 02",
		Title: fmt.Sprintf("ATENCIONES DE SERVICIO DE ODONTOLOGÍA POR PROCEDIMIENTOS REALIZADO, SEGÚN ESCUELA PROFESIONAL - %s", trimester),
		Area:  "ESCUELA PROFESIONAL",
	}

	f = HeaderDentistryProceduresTrimesterExcel(f, sheet, headerMedicalArea, styleHeaderID)

	for i, sName := range schools {
		row := 7 + i
		start, end := fmt.Sprintf("A%d", row), fmt.Sprintf("AE%d", row)
		_ = f.SetCellValue(sheet, start, sName)
		_ = f.SetCellStyle(sheet, start, end, styleBorderID)
	}

	_ = f.SetColWidth(sheet, "A", "A", 28)

	cols := getExcelColumns("B", "AE")
	for _, col := range cols {
		_ = f.SetColWidth(sheet, col, col, 4)
	}

	dateStart, dateEnd := getDateRange(month, year)

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)
	performedProcedures, err := srv.SrvDentistryConsultationBuccalProcedure.GetBuccalProceduresExcel(dateStart, dateEnd)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get buccal procedures:", err)
		return "", 15, err
	}

	totalRow := 6
	startRow := 7

	startTime, _ := time.Parse("2006-01-02", dateStart)
	endTime, _ := time.Parse("2006-01-02", dateEnd)

	columnTotalTrimester := []string{"B", "C", "D"}
	columnsStartProcedures := "E"

	knownProcedures := make(map[string]struct{})
	for _, proc := range dentistryProcedures[:len(dentistryProcedures)-1] {
		knownProcedures[proc] = struct{}{}
	}

	totalTrimester := make([]int, len(schools))
	totalTrimesterM := make([]int, len(schools))
	totalTrimesterF := make([]int, len(schools))

	totalProcedures := make([][]int, len(dentistryProcedures))
	totalProceduresM := make([][]int, len(dentistryProcedures))
	totalProceduresF := make([][]int, len(dentistryProcedures))

	for i := range dentistryProcedures {
		totalProcedures[i] = make([]int, len(schools))
		totalProceduresM[i] = make([]int, len(schools))
		totalProceduresF[i] = make([]int, len(schools))
	}

	sNameIdx := make(map[string]int)
	for idx, s := range schools {
		sNameIdx[s] = idx
	}

	for _, p := range performedProcedures {
		if p.TipoPersona != "Estudiante" {
			continue
		}
		date, err := time.Parse("2006-01-02", p.FechaConsulta[:10])
		if err != nil {
			logger.Error.Println(s.txID, " - couldn't parse date:", err)
			return "", 15, err
		}

		idxSchool, ok := sNameIdx[p.EscuelaProfesional]
		if !ok {
			continue
		}

		totalTrimester[idxSchool]++

		if p.Sexo == "M" {
			totalTrimesterM[idxSchool]++
		} else {
			totalTrimesterF[idxSchool]++
		}

		if date.After(startTime) && date.Before(endTime) {
			for idxProc, proc := range dentistryProcedures {
				if proc != p.TipoProcedimiento {
					continue
				}

				totalProcedures[idxProc][idxSchool]++

				if p.Sexo == "M" {
					totalProceduresM[idxProc][idxSchool]++
				} else {
					totalProceduresF[idxProc][idxSchool]++
				}
			}
		}

	}

	// Escribir procedimientos totales de trimestre
	countTotal := 0
	countTotalM := 0
	countTotalF := 0
	for i, cnt := range totalTrimester {
		row := startRow + i
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[0], row), cnt)
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[1], row), totalTrimesterM[i])
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[2], row), totalTrimesterF[i])
		countTotal += cnt
		countTotalM += totalTrimesterM[i]
		countTotalF += totalTrimesterF[i]
	}
	_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[0], totalRow), countTotal)
	_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[1], totalRow), countTotalM)
	_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[2], totalRow), countTotalF)

	// Escribir procedimientos
	startColNum, _ := excelize.ColumnNameToNumber(columnsStartProcedures)
	for idxProc := range dentistryProcedures {
		colT, _ := excelize.ColumnNumberToName(startColNum + idxProc*3)
		colM, _ := excelize.ColumnNumberToName(startColNum + idxProc*3 + 1)
		colF, _ := excelize.ColumnNumberToName(startColNum + idxProc*3 + 2)

		total, totalM, totalF := 0, 0, 0
		for idxSchool := range schools {
			row := startRow + idxSchool
			t := totalProcedures[idxProc][idxSchool]
			m := totalProceduresM[idxProc][idxSchool]
			fem := totalProceduresF[idxProc][idxSchool]

			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colT, row), t)
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colM, row), m)
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colF, row), fem)

			total += t
			totalM += m
			totalF += fem
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colT, totalRow), total)
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colM, totalRow), totalM)
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", colF, totalRow), totalF)
	}

	base64Str, code, err := SaveExcelAndReturnBase64(f, "odontologia_cuadro_02")

	return base64Str, code, err
}

func (s *ReportsMedicalAreaService) GetReportMedicalFrame1LowCode(month, year string) (string, int, error) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)
	styleHeaderID, styleBorderID := styleExcel(f)

	trimester := getTrimester(month)

	headerMedicalArea := models.HeaderMedicalAreaExcel{
		Frame: "CUADRO N° 01",
		Title: "CONSULTAS A ESTUDIANTES POR MES Y SEXO SEGÚN ESCUELA PROFESIONAL",
		Area:  "ESCUELA PROFESIONAL",
	}

	f = HeaderMedicalConsultingTrimesterExcel(f, sheet, headerMedicalArea, styleHeaderID, month, trimester, year)

	startRow := 6
	var mergeColumnsBody []models.MergeRange

	countSchools := len(schools)

	for i, sName := range schools {
		rowStart := startRow + (i * 2)
		rowEnd := rowStart + 1

		_ = f.SetCellValue(sheet, fmt.Sprintf("A%d", rowStart), sName)
		_ = f.SetCellValue(sheet, fmt.Sprintf("B%d", rowStart), "M")
		_ = f.SetCellValue(sheet, fmt.Sprintf("B%d", rowEnd), "F")

		mergeColumnsBody = append(mergeColumnsBody, models.MergeRange{
			Column1: fmt.Sprintf("A%d", rowStart),
			Column2: fmt.Sprintf("A%d", rowEnd),
		})
		mergeColumnsBody = append(mergeColumnsBody, models.MergeRange{
			Column1: fmt.Sprintf("G%d", rowStart),
			Column2: fmt.Sprintf("G%d", rowEnd),
		})
		_ = f.SetCellStyle(sheet, fmt.Sprintf("B%d", rowStart), fmt.Sprintf("B%d", rowEnd), styleHeaderID)
		_ = f.SetCellStyle(sheet, fmt.Sprintf("C%d", rowStart), fmt.Sprintf("F%d", rowEnd), styleBorderID)
	}

	finalRow := startRow + (countSchools * 2)
	_ = f.SetCellValue(sheet, fmt.Sprintf("A%d", finalRow), "TOTAL GENERAL")
	_ = f.SetCellStyle(sheet, fmt.Sprintf("A%d", finalRow), fmt.Sprintf("A%d", finalRow), styleHeaderID)
	_ = f.SetCellStyle(sheet, fmt.Sprintf("B%d", finalRow), fmt.Sprintf("G%d", finalRow), styleBorderID)

	for _, merge := range mergeColumnsBody {
		_ = f.MergeCell(sheet, merge.Column1, merge.Column2)
		_ = f.SetCellStyle(sheet, merge.Column1, merge.Column2, styleBorderID)
	}

	_ = f.SetColWidth(sheet, "A", "A", 50)

	dateStart, dateEnd := getDateRange(month, year)

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)
	generalMedicineConsultation, err := srv.SrvMedicalGeneralMedicineConsultation.GetGeneralMedicineConsultationExcel(dateStart, dateEnd)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get performed procedures:", err)
		return "", 15, err
	}

	totalRow := (len(schools) * 2) + startRow

	type DateRange struct {
		Start time.Time
		End   time.Time
	}

	startTime, _ := time.Parse("2006-01-02", dateStart)
	endTime, _ := time.Parse("2006-01-02", dateEnd)
	nextStart := startTime.AddDate(0, 1, 0)
	prevEnd := endTime.AddDate(0, -1, 0)

	dateRanges := []DateRange{
		{Start: startTime, End: nextStart},
		{Start: nextStart, End: prevEnd},
		{Start: prevEnd, End: endTime},
	}

	columnsTotalMonthStart := []string{"C", "D", "E"}

	countTotal := 0

	for idx, column := range columnsTotalMonthStart {
		countTotal = 0
		for i, school := range schools {
			rowM := startRow + (i * 2)
			rowF := startRow + (i * 2) + 1
			countM, countF := 0, 0
			for _, attn := range generalMedicineConsultation {
				if attn.TipoPersona != "Estudiante" {
					continue
				}
				if attn.EscuelaProfesional != school {
					continue
				}
				date, err := time.Parse("2006-01-02", attn.FechaConsulta[:10])
				if err != nil {
					logger.Error.Println(s.txID, " - couldn't parse date:", err)
					return "", 15, err
				}

				if (date.Equal(dateRanges[idx].Start) || date.After(dateRanges[idx].Start)) && date.Before(dateRanges[idx].End) {
					if attn.Sexo == "M" {
						countM++
					} else {
						countF++
					}
					countTotal++
				}
			}
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", column, rowM), countM)
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", column, rowF), countF)
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", column, totalRow), countTotal)
	}

	columnSubtotal := "F"
	countTotal = 0
	for i, school := range schools {
		rowM := startRow + (i * 2)
		rowF := startRow + (i * 2) + 1
		countM, countF := 0, 0
		for _, attn := range generalMedicineConsultation {
			if attn.TipoPersona != "Estudiante" {
				continue
			}
			if attn.EscuelaProfesional != school {
				continue
			}
			if attn.Sexo == "M" {
				countM++
			} else {
				countF++
			}
			countTotal++
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnSubtotal, rowM), countM)
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnSubtotal, rowF), countF)
	}
	_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnSubtotal, totalRow), countTotal)

	columnTotal := "G"
	countTotal = 0
	for i, school := range schools {
		row := startRow + (i * 2)
		count := 0
		for _, attn := range generalMedicineConsultation {
			if attn.TipoPersona != "Estudiante" {
				continue
			}
			if attn.EscuelaProfesional != school {
				continue
			}
			count++
			countTotal++
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotal, row), count)
	}
	_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotal, totalRow), countTotal)

	base64Str, code, err := SaveExcelAndReturnBase64(f, "medicina_cuadro_01")

	return base64Str, code, err
}

func (s *ReportsMedicalAreaService) GetReportMedicalFrame3LowCode(month, year string) (string, int, error) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)
	styleHeaderID, styleBorderID := styleExcel(f)

	trimester := getTrimester(month)

	headerMedicalArea := models.HeaderMedicalAreaExcel{
		Frame: "CUADRO N° 03",
		Title: "CONSULTAS POR MES Y SEXO SEGÚN TIPO DE PERSONAL",
		Area:  "PERSONAL",
	}

	f = HeaderMedicalConsultingTrimesterExcel(f, sheet, headerMedicalArea, styleHeaderID, month, trimester, year)

	startRow := 6
	var mergeColumnsBody []models.MergeRange
	countStaffs := len(nursingStaffs)

	for i, sName := range nursingStaffs {
		rowStart := startRow + (i * 2)
		rowEnd := rowStart + 1

		_ = f.SetCellValue(sheet, fmt.Sprintf("A%d", rowStart), sName)
		_ = f.SetCellValue(sheet, fmt.Sprintf("B%d", rowStart), "M")
		_ = f.SetCellValue(sheet, fmt.Sprintf("B%d", rowEnd), "F")

		mergeColumnsBody = append(mergeColumnsBody, models.MergeRange{
			Column1: fmt.Sprintf("A%d", rowStart),
			Column2: fmt.Sprintf("A%d", rowEnd),
		})
		mergeColumnsBody = append(mergeColumnsBody, models.MergeRange{
			Column1: fmt.Sprintf("G%d", rowStart),
			Column2: fmt.Sprintf("G%d", rowEnd),
		})
		_ = f.SetCellStyle(sheet, fmt.Sprintf("B%d", rowStart), fmt.Sprintf("B%d", rowEnd), styleHeaderID)
		_ = f.SetCellStyle(sheet, fmt.Sprintf("C%d", rowStart), fmt.Sprintf("F%d", rowEnd), styleBorderID)
	}

	finalRow := startRow + (countStaffs * 2)
	_ = f.SetCellValue(sheet, fmt.Sprintf("A%d", finalRow), "TOTAL GENERAL")
	_ = f.SetCellStyle(sheet, fmt.Sprintf("A%d", finalRow), fmt.Sprintf("A%d", finalRow), styleHeaderID)
	_ = f.SetCellStyle(sheet, fmt.Sprintf("B%d", finalRow), fmt.Sprintf("G%d", finalRow), styleBorderID)

	for _, merge := range mergeColumnsBody {
		_ = f.MergeCell(sheet, merge.Column1, merge.Column2)
		_ = f.SetCellStyle(sheet, merge.Column1, merge.Column2, styleBorderID)
	}

	_ = f.SetColWidth(sheet, "A", "A", 20)

	dateStart, dateEnd := getDateRange(month, year)

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)
	generalMedicineConsultation, err := srv.SrvMedicalGeneralMedicineConsultation.GetGeneralMedicineConsultationExcel(dateStart, dateEnd)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get performed procedures:", err)
		return "", 15, err
	}

	totalRow := (len(nursingStaffsComparation) * 2) + startRow

	type DateRange struct {
		Start time.Time
		End   time.Time
	}

	startTime, _ := time.Parse("2006-01-02", dateStart)
	endTime, _ := time.Parse("2006-01-02", dateEnd)
	nextStart := startTime.AddDate(0, 1, 0)
	prevEnd := endTime.AddDate(0, -1, 0)

	dateRanges := []DateRange{
		{Start: startTime, End: nextStart},
		{Start: nextStart, End: prevEnd},
		{Start: prevEnd, End: endTime},
	}

	columnsTotalMonthStart := []string{"C", "D", "E"}

	countTotal := 0

	for idx, column := range columnsTotalMonthStart {
		countTotal = 0
		for i, staff := range nursingStaffsComparation {
			rowM := startRow + (i * 2)
			rowF := startRow + (i * 2) + 1
			countM, countF := 0, 0
			for _, attn := range generalMedicineConsultation {
				if attn.TipoPersona == "Estudiante" {
					continue
				}
				if attn.TipoPersona != staff {
					continue
				}
				date, err := time.Parse("2006-01-02", attn.FechaConsulta[:10])
				if err != nil {
					logger.Error.Println(s.txID, " - couldn't parse date:", err)
					return "", 15, err
				}

				if (date.Equal(dateRanges[idx].Start) || date.After(dateRanges[idx].Start)) && date.Before(dateRanges[idx].End) {
					if attn.Sexo == "M" {
						countM++
					} else {
						countF++
					}
					countTotal++
				}
			}
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", column, rowM), countM)
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", column, rowF), countF)
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", column, totalRow), countTotal)
	}

	columnSubtotal := "F"
	countTotal = 0
	for i, staff := range nursingStaffsComparation {
		rowM := startRow + (i * 2)
		rowF := startRow + (i * 2) + 1
		countM, countF := 0, 0
		for _, attn := range generalMedicineConsultation {
			if attn.TipoPersona == "Estudiante" {
				continue
			}
			if attn.TipoPersona != staff {
				continue
			}
			if attn.Sexo == "M" {
				countM++
			} else {
				countF++
			}
			countTotal++
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnSubtotal, rowM), countM)
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnSubtotal, rowF), countF)
	}
	_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnSubtotal, totalRow), countTotal)

	columnTotal := "G"
	countTotal = 0
	for i, staff := range nursingStaffsComparation {
		row := startRow + (i * 2)
		count := 0
		for _, attn := range generalMedicineConsultation {
			if attn.TipoPersona == "Estudiante" {
				continue
			}
			if attn.TipoPersona != staff {
				continue
			}
			count++
			countTotal++
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotal, row), count)
	}
	_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotal, totalRow), countTotal)

	base64Str, code, err := SaveExcelAndReturnBase64(f, "medicina_cuadro_03")

	return base64Str, code, err
}

func (s *ReportsMedicalAreaService) GetReportMedicalFrame8LowCode(month, year string) (string, int, error) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)
	styleHeaderID, styleBorderID := styleExcel(f)

	trimester := getTrimester(month)

	headerMedicalArea := models.HeaderMedicalAreaExcel{
		Frame: "CUADRO N° 08",
		Title: fmt.Sprintf("ATENCIONES DE CONSULTAS MÉDICAS A ESTUDIANTES DEL %s, SEGÚN ESCUELA PROFESIONAL - AÑO %s", trimester, year),
		Area:  "ESCUELA PROFESIONAL",
	}

	f = HeaderNursingConsultingTrimesterExcel(f, sheet, headerMedicalArea, styleHeaderID, month, trimester, year)

	for i, sName := range nursingStaffs {
		row := 6 + i
		start, end := fmt.Sprintf("A%d", row), fmt.Sprintf("M%d", row)
		_ = f.SetCellValue(sheet, start, sName)
		_ = f.SetCellStyle(sheet, start, end, styleBorderID)
	}

	_ = f.SetColWidth(sheet, "A", "A", 20)

	dateStart, dateEnd := getDateRange(month, year)

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)
	generalMedicineConsultation, err := srv.SrvMedicalGeneralMedicineConsultation.GetGeneralMedicineConsultationExcel(dateStart, dateEnd)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get performed procedures:", err)
		return "", 15, err
	}

	startRow := 6
	totalRow := 5

	type DateRange struct {
		Start time.Time
		End   time.Time
	}

	startTime, _ := time.Parse("2006-01-02", dateStart)
	endTime, _ := time.Parse("2006-01-02", dateEnd)
	nextStart := startTime.AddDate(0, 1, 0)
	prevEnd := endTime.AddDate(0, -1, 0)

	dateRanges := []DateRange{
		{Start: startTime, End: nextStart},
		{Start: nextStart, End: prevEnd},
		{Start: prevEnd, End: endTime},
	}

	columnTotalTrimester := []string{"B", "C", "D"}
	columnsTotalMonthStart := []string{"E", "H", "K"}

	shouldCount := func(attn *models.ConsultationIntegralAttentionExcel, staff string, sexo string) bool {
		if attn.TipoPersona == "Estudiante" {
			return false
		}
		if attn.TipoPersona != staff {
			return false
		}
		if sexo == "" {
			return true
		}
		return attn.Sexo == sexo
	}

	for i, sexo := range []string{"", "M", "F"} {
		totalCount := 0

		for j, staff := range nursingStaffsComparation {
			row := startRow + j
			count := 0

			for _, attn := range generalMedicineConsultation {
				if shouldCount(attn, staff, sexo) {
					count++
					totalCount++
				}
			}
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[i], row), count)
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[i], totalRow), totalCount)
	}

	for monthIdx, dateRange := range dateRanges {
		for i, sexo := range []string{"", "M", "F"} {
			colNumStart, _ := excelize.ColumnNameToNumber(columnsTotalMonthStart[monthIdx])
			col, _ := excelize.ColumnNumberToName(colNumStart + i)
			totalCount := 0

			for j, staff := range nursingStaffsComparation {
				row := startRow + j
				count := 0

				for _, attn := range generalMedicineConsultation {
					date, err := time.Parse("2006-01-02", attn.FechaConsulta[:10])
					if err != nil {
						logger.Error.Println(s.txID, " - couldn't parse date:", err)
						return "", 15, err
					}

					if (date.Equal(dateRange.Start) || date.After(dateRange.Start)) && date.Before(dateRange.End) {
						if shouldCount(attn, staff, sexo) {
							count++
							totalCount++
						}
					}
				}
				_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, row), count)
			}
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, totalRow), totalCount)
		}
	}

	base64Str, code, err := SaveExcelAndReturnBase64(f, "medicina_cuadro_08")

	return base64Str, code, err
}

func (s *ReportsMedicalAreaService) GetReportMedicalFrame12LowCode(month, year string) (string, int, error) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)
	styleHeaderID, styleBorderID := styleExcel(f)

	trimester := getTrimester(month)

	headerMedicalArea := models.HeaderMedicalAreaExcel{
		Frame: "CUADRO N° 12",
		Title: fmt.Sprintf("ATENCIONES DE CONSULTAS MÉDICAS A ESTUDIANTES DEL %s, SEGÚN TIPO DE PERSONAL - AÑO %s", trimester, year),
		Area:  "PERSONAL",
	}

	f = HeaderNursingConsultingTrimesterExcel(f, sheet, headerMedicalArea, styleHeaderID, month, trimester, year)

	for i, school := range schools {
		cell := fmt.Sprintf("A%d", 6+i)
		cellEnd := fmt.Sprintf("M%d", 6+i)
		_ = f.SetCellValue(sheet, cell, school)
		_ = f.SetCellStyle(sheet, cell, cellEnd, styleBorderID)
	}

	_ = f.SetColWidth(sheet, "A", "A", 50)

	dateStart, dateEnd := getDateRange(month, year)

	srv := medical_area.NewServerMedicalArea(s.db, s.usr, s.txID)
	generalMedicineConsultation, err := srv.SrvMedicalGeneralMedicineConsultation.GetGeneralMedicineConsultationExcel(dateStart, dateEnd)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get performed procedures:", err)
		return "", 15, err
	}

	startRow := 6
	totalRow := 5

	type DateRange struct {
		Start time.Time
		End   time.Time
	}

	startTime, _ := time.Parse("2006-01-02", dateStart)
	endTime, _ := time.Parse("2006-01-02", dateEnd)
	nextStart := startTime.AddDate(0, 1, 0)
	prevEnd := endTime.AddDate(0, -1, 0)

	dateRanges := []DateRange{
		{Start: startTime, End: nextStart},
		{Start: nextStart, End: prevEnd},
		{Start: prevEnd, End: endTime},
	}

	columnTotalTrimester := []string{"B", "C", "D"}
	columnsTotalMonthStart := []string{"E", "H", "K"}

	shouldCount := func(attn *models.ConsultationIntegralAttentionExcel, school string, sexo string) bool {
		if attn.TipoPersona != "Estudiante" {
			return false
		}
		if attn.EscuelaProfesional != school {
			return false
		}
		if sexo == "" {
			return true
		}
		return attn.Sexo == sexo
	}

	for i, sexo := range []string{"", "M", "F"} {
		totalCount := 0

		for j, school := range schools {
			row := startRow + j
			count := 0

			for _, attn := range generalMedicineConsultation {
				if shouldCount(attn, school, sexo) {
					count++
					totalCount++
				}
			}
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[i], row), count)
		}
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", columnTotalTrimester[i], totalRow), totalCount)
	}

	for monthIdx, dateRange := range dateRanges {
		for i, sexo := range []string{"", "M", "F"} {
			colNumStart, _ := excelize.ColumnNameToNumber(columnsTotalMonthStart[monthIdx])
			col, _ := excelize.ColumnNumberToName(colNumStart + i)
			totalCount := 0

			for j, school := range schools {
				row := startRow + j
				count := 0

				for _, attn := range generalMedicineConsultation {
					date, err := time.Parse("2006-01-02", attn.FechaConsulta[:10])
					if err != nil {
						logger.Error.Println(s.txID, " - couldn't parse date:", err)
						return "", 15, err
					}

					if (date.Equal(dateRange.Start) || date.After(dateRange.Start)) && date.Before(dateRange.End) {
						if shouldCount(attn, school, sexo) {
							count++
							totalCount++
						}
					}
				}
				_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, row), count)
			}
			_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", col, totalRow), totalCount)
		}
	}

	base64Str, code, err := SaveExcelAndReturnBase64(f, "medicina_cuadro_12")

	return base64Str, code, err
}

func getTrimester(month string) string {
	trimester := ""
	if month == "1,2,3" {
		trimester = "I TRIMESTRE"
	}
	if month == "4,5,6" {
		trimester = "II TRIMESTRE"
	}
	if month == "7,8,9" {
		trimester = "III TRIMESTRE"
	}
	if month == "10,11,12" {
		trimester = "IV TRIMESTRE"
	}
	return trimester
}

func getDateRange(month, year string) (string, string) {

	rangeDates := rangeDateTrimester[month]

	dateStart := fmt.Sprintf("%s-%s", year, rangeDates[0])
	dateEnd := fmt.Sprintf("%s-%s", year, rangeDates[1])

	tEnd, _ := time.Parse("2006-01-02", dateEnd)
	tEnd = tEnd.AddDate(0, 0, 1)
	dateEnd = tEnd.Format("2006-01-02")

	return dateStart, dateEnd
}

func HeaderNursingConsultingYearlyExcel(f *excelize.File, sheet string, headerMedicalArea models.HeaderMedicalAreaExcel, styleHeaderID int, year string) *excelize.File {

	var mergeColumns = []models.MergeRange{
		{Column1: "A1", Column2: "P1"},
		{Column1: "A2", Column2: "P2"},
		{Column1: "A3", Column2: "A4"},
		{Column1: "B3", Column2: "D3"},
		{Column1: "E3", Column2: "G3"},
		{Column1: "H3", Column2: "J3"},
		{Column1: "K3", Column2: "M3"},
		{Column1: "N3", Column2: "P3"},
	}

	for _, merge := range mergeColumns {
		_ = f.MergeCell(sheet, merge.Column1, merge.Column2)
		_ = f.SetCellStyle(sheet, merge.Column1, merge.Column2, styleHeaderID)
	}

	cellValues := map[string]string{
		"A1": headerMedicalArea.Frame,
		"A2": headerMedicalArea.Title,
		"A3": headerMedicalArea.Area,
		"A5": "TOTAL",
		"B3": fmt.Sprintf("TOTAL ANUAL %s", year),
	}

	for col := 'B'; col <= 'P'; col++ {
		cell := fmt.Sprintf("%c5", col)
		cellValues[cell] = ""
	}

	cols := []string{"E3", "H3", "K3", "N3"}
	trimester := []string{"I TRIMESTRE(Ene-Mar)", "II TRIMESTRE(Abr-Jun)", "III TRIMESTRE(Jul-Set)", "IV TRIMESTRE(Oct-Dic)"}
	for i := range trimester {
		if i < len(cols) {
			cellValues[cols[i]] = trimester[i]
		}
	}

	for cell, value := range cellValues {
		_ = f.SetCellValue(sheet, cell, value)
		_ = f.SetCellStyle(sheet, cell, cell, styleHeaderID)
	}

	columns := []string{
		"B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
	}

	values := []string{"T", "M", "F"}

	for i, col := range columns {
		value := values[i%len(values)]
		cell := col + "4"
		_ = f.SetCellStyle(sheet, cell, cell, styleHeaderID)
		_ = f.SetCellValue(sheet, cell, value)
	}

	return f
}

func HeaderNursingConsultingTrimesterExcel(f *excelize.File, sheet string, headerMedicalArea models.HeaderMedicalAreaExcel, styleHeaderID int, month, trimester, year string) *excelize.File {

	var mergeColumns = []models.MergeRange{
		{Column1: "A1", Column2: "M1"},
		{Column1: "A2", Column2: "M2"},
		{Column1: "A3", Column2: "A4"},
		{Column1: "B3", Column2: "D3"},
		{Column1: "E3", Column2: "G3"},
		{Column1: "H3", Column2: "J3"},
		{Column1: "K3", Column2: "M3"},
	}

	for _, merge := range mergeColumns {
		_ = f.MergeCell(sheet, merge.Column1, merge.Column2)
		_ = f.SetCellStyle(sheet, merge.Column1, merge.Column2, styleHeaderID)
	}

	cellValues := map[string]string{
		"A1": headerMedicalArea.Frame,
		"A2": headerMedicalArea.Title,
		"A3": headerMedicalArea.Area,
		"A5": "TOTAL",
		"B3": fmt.Sprintf("TOTAL %s %s", trimester, year),
	}

	for col := 'B'; col <= 'M'; col++ {
		cell := fmt.Sprintf("%c5", col)
		cellValues[cell] = ""
	}

	cols := []string{"E3", "H3", "K3"}
	monthCodes := strings.Split(month, ",")
	for i, code := range monthCodes {
		if i < len(cols) {
			cellValues[cols[i]] = months[code]
		}
	}

	for cell, value := range cellValues {
		_ = f.SetCellValue(sheet, cell, value)
		_ = f.SetCellStyle(sheet, cell, cell, styleHeaderID)
	}

	columns := []string{
		"B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
	}

	values := []string{"T", "M", "F"}
	for i, col := range columns {
		value := values[i%len(values)]
		cell := col + "4"
		_ = f.SetCellStyle(sheet, cell, cell, styleHeaderID)
		_ = f.SetCellValue(sheet, cell, value)
		_ = f.SetColWidth(sheet, col, col, 6)
	}

	_ = f.SetRowHeight(sheet, 3, 38)

	return f
}

func HeaderNursingProceduresYearsExcel(f *excelize.File, sheet string, headerMedicalArea models.HeaderMedicalAreaExcel, styleHeaderID int, startYear, endYear string) (*excelize.File, string) {

	iStartYear, _ := strconv.Atoi(startYear)
	iEndYear, _ := strconv.Atoi(endYear)

	cellHeader := 4
	columnStart := "B"
	columnEnd := ""

	idx := 0
	for i := iStartYear; i <= iEndYear; i++ {
		columnNumber, _ := excelize.ColumnNameToNumber(columnStart)
		column, _ := excelize.ColumnNumberToName(columnNumber + idx)
		_ = f.SetCellValue(sheet, fmt.Sprintf("%s%d", column, cellHeader), i)
		_ = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", column, cellHeader), fmt.Sprintf("%s%d", column, cellHeader), styleHeaderID)
		_ = f.SetCellStyle(sheet, fmt.Sprintf("%s%d", column, cellHeader+1), fmt.Sprintf("%s%d", column, cellHeader+1), styleHeaderID)
		columnEnd = column
		idx++
	}

	var mergeColumns = []models.MergeRange{
		{Column1: "A1", Column2: fmt.Sprintf("%s1", columnEnd)},
		{Column1: "A2", Column2: fmt.Sprintf("%s2", columnEnd)},
		{Column1: "A3", Column2: "A4"},
		{Column1: "B3", Column2: fmt.Sprintf("%s3", columnEnd)},
	}

	for _, merge := range mergeColumns {
		_ = f.MergeCell(sheet, merge.Column1, merge.Column2)
		_ = f.SetCellStyle(sheet, merge.Column1, merge.Column2, styleHeaderID)
	}

	cellValues := map[string]string{
		"A1": headerMedicalArea.Frame,
		"A2": headerMedicalArea.Title,
		"A3": headerMedicalArea.Area,
		"A5": "TOTAL",
		"B3": "AÑOS",
	}

	for cell, value := range cellValues {
		_ = f.SetCellValue(sheet, cell, value)
		_ = f.SetCellStyle(sheet, cell, cell, styleHeaderID)
	}

	return f, columnEnd
}

func HeaderNursingProceduresYearlyExcel(f *excelize.File, sheet string, headerMedicalArea models.HeaderMedicalAreaExcel, styleHeaderID int) *excelize.File {

	var mergeColumns = []models.MergeRange{
		{Column1: "A1", Column2: "N1"},
		{Column1: "A2", Column2: "N2"},
		{Column1: "A3", Column2: "A4"},
		{Column1: "B3", Column2: "B4"},
		{Column1: "C3", Column2: "N3"},
	}

	for _, merge := range mergeColumns {
		_ = f.MergeCell(sheet, merge.Column1, merge.Column2)
		_ = f.SetCellStyle(sheet, merge.Column1, merge.Column2, styleHeaderID)
	}

	cellValues := map[string]string{
		"A1": headerMedicalArea.Frame,
		"A2": headerMedicalArea.Title,
		"A3": headerMedicalArea.Area,
		"A5": "TOTAL",
		"B3": "TOTAL",
		"C3": "ESTUDIANTES BENEFICIADOS",
		"C4": "Ene.",
		"D4": "Feb.",
		"E4": "Mar.",
		"F4": "Abr.",
		"G4": "May.",
		"H4": "Jun.",
		"I4": "Jul.",
		"J4": "Ago.",
		"K4": "Set.",
		"L4": "Oct.",
		"M4": "Nov.",
		"N4": "Dic.",
	}

	for col := 'B'; col <= 'N'; col++ {
		cell := fmt.Sprintf("%c5", col)
		cellValues[cell] = ""
	}

	for cell, value := range cellValues {
		_ = f.SetCellValue(sheet, cell, value)
		_ = f.SetCellStyle(sheet, cell, cell, styleHeaderID)
	}

	return f
}

func HeaderNursingProceduresTrimesterExcel(f *excelize.File, sheet string, headerMedicalArea models.HeaderMedicalAreaExcel, styleHeaderID int, month, trimester, year string) *excelize.File {
	mergeCells := []models.MergeRange{
		{"A1", "AO1"},
		{"A2", "AO2"},
		{"A3", "A5"},
		{"B3", "B5"},
		{"C3", "C5"},
		{"D3", "O3"},
		{"D4", "E4"},
		{"F4", "G4"},
		{"H4", "I4"},
		{"J4", "K4"},
		{"L4", "M4"},
		{"N4", "O4"},
		{"P3", "P5"},
		{"Q3", "AB3"},
		{"Q4", "R4"},
		{"S4", "T4"},
		{"U4", "V4"},
		{"W4", "X4"},
		{"Y4", "Z4"},
		{"AA4", "AB4"},
		{"AC3", "AC5"},
		{"AD3", "AO3"},
		{"AD4", "AE4"},
		{"AF4", "AG4"},
		{"AH4", "AI4"},
		{"AJ4", "AK4"},
		{"AL4", "AM4"},
		{"AN4", "AO4"},
	}
	for _, m := range mergeCells {
		_ = f.MergeCell(sheet, m.Column1, m.Column2)
		_ = f.SetCellStyle(sheet, m.Column1, m.Column2, styleHeaderID)
	}

	setCells := func(cells map[string]string, styleID int) {
		for cell, value := range cells {
			_ = f.SetCellValue(sheet, cell, value)
			_ = f.SetCellStyle(sheet, cell, cell, styleID)
		}
	}
	cellValues := map[string]string{
		"A1": headerMedicalArea.Frame,
		"A2": headerMedicalArea.Title,
		"A3": headerMedicalArea.Area,
		"A6": "TOTAL",
		"B3": fmt.Sprintf("TOTAL %s %s", trimester, year),
		"C3": "TOTAL MES", "P3": "TOTAL MES", "AC3": "TOTAL MES",
	}

	_ = f.SetColWidth(sheet, "C", "C", 5)
	_ = f.SetColWidth(sheet, "P", "P", 5)
	_ = f.SetColWidth(sheet, "AC", "AC", 5)

	cols := []string{"D3", "Q3", "AD3"}
	monthCodes := strings.Split(month, ",")
	for i, code := range monthCodes {
		if i < len(cols) {
			cellValues[cols[i]] = months[code]
		}
	}
	monthStartCols := []string{"D", "Q", "AD"}

	for _, mCol := range monthStartCols {
		colNum, _ := excelize.ColumnNameToNumber(mCol)
		for j, proc := range nursingProceduresHeader {
			procCol, _ := excelize.ColumnNumberToName(colNum + j*2)
			cellValues[procCol+"4"] = proc
		}
	}

	setCells(cellValues, styleHeaderID)

	for _, startCol := range monthStartCols {
		colNum, _ := excelize.ColumnNameToNumber(startCol)
		for i := 0; i < 6; i++ {
			mCol, _ := excelize.ColumnNumberToName(colNum + i*2)
			fCol, _ := excelize.ColumnNumberToName(colNum + i*2 + 1)

			_ = f.SetCellValue(sheet, mCol+"5", "M")
			_ = f.SetCellValue(sheet, fCol+"5", "F")
			_ = f.SetCellStyle(sheet, mCol+"5", mCol+"5", styleHeaderID)
			_ = f.SetCellStyle(sheet, fCol+"5", fCol+"5", styleHeaderID)
			_ = f.SetColWidth(sheet, mCol, fCol, 4.2)
		}
	}
	_ = f.SetRowHeight(sheet, 4, 25)
	for i := 1; i <= 41; i++ {
		col, _ := excelize.ColumnNumberToName(i)
		cell := col + "6"
		if i != 1 {
			_ = f.SetCellValue(sheet, cell, "")
		}
		_ = f.SetCellStyle(sheet, cell, cell, styleHeaderID)
	}

	return f
}

func HeaderNursingProceduresExcel(f *excelize.File, sheet string, headerMedicalArea models.HeaderMedicalAreaExcel, styleHeaderID int, trimester string) *excelize.File {
	mergeCells := []models.MergeRange{
		{"A1", "N1"},
		{"A2", "N2"},
		{"A3", "A5"},
		{"B3", "B5"},
		{"C3", "N3"},
		{"C4", "D4"},
		{"E4", "F4"},
		{"G4", "H4"},
		{"I4", "J4"},
		{"K4", "L4"},
		{"M4", "N4"},
	}
	for _, m := range mergeCells {
		_ = f.MergeCell(sheet, m.Column1, m.Column2)
		_ = f.SetCellStyle(sheet, m.Column1, m.Column2, styleHeaderID)
	}

	setCells := func(cells map[string]string, styleID int) {
		for cell, value := range cells {
			_ = f.SetCellValue(sheet, cell, value)
			_ = f.SetCellStyle(sheet, cell, cell, styleID)
		}
	}
	cellValues := map[string]string{
		"A1": headerMedicalArea.Frame,
		"A2": headerMedicalArea.Title,
		"A3": headerMedicalArea.Area,
		"A6": "TOTAL",
		"B3": fmt.Sprintf("TOTAL %s 2025", trimester),
		"C3": "PROCEDIMIENTOS REALIZADOS",
	}

	for i, proc := range nursingProceduresHeader {
		procCol, _ := excelize.ColumnNumberToName(1 + (i+1)*2)
		cellValues[procCol+"4"] = proc
	}

	setCells(cellValues, styleHeaderID)

	columns := []string{
		"C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N",
	}

	values := []string{"M", "F"}

	for i, col := range columns {
		value := values[i%len(values)]
		cell := col + "5"
		_ = f.SetCellStyle(sheet, cell, cell, styleHeaderID)
		_ = f.SetCellValue(sheet, cell, value)
		_ = f.SetColWidth(sheet, col, col, 6)
	}

	for i := 1; i <= 14; i++ {
		col, _ := excelize.ColumnNumberToName(i)
		cell := col + "6"
		if i != 1 {
			_ = f.SetCellValue(sheet, cell, "")
		}
		_ = f.SetCellStyle(sheet, cell, cell, styleHeaderID)
	}

	return f
}

func HeaderDentistryProceduresTrimesterExcel(f *excelize.File, sheet string, headerMedicalArea models.HeaderMedicalAreaExcel, styleHeaderID int) *excelize.File {
	mergeCells := []models.MergeRange{
		{"A1", "AE1"},
		{"A2", "AE2"},
		{"A3", "A5"},
		{"B3", "D4"},
		{"E3", "AE3"},
		{"E4", "G4"},
		{"H4", "J4"},
		{"K4", "M4"},
		{"N4", "P4"},
		{"Q4", "S4"},
		{"T4", "V4"},
		{"W4", "Y4"},
		{"Z4", "AB4"},
		{"AC4", "AE4"},
	}
	for _, m := range mergeCells {
		_ = f.MergeCell(sheet, m.Column1, m.Column2)
		_ = f.SetCellStyle(sheet, m.Column1, m.Column2, styleHeaderID)
	}

	setCells := func(cells map[string]string, styleID int) {
		for cell, value := range cells {
			_ = f.SetCellValue(sheet, cell, value)
			_ = f.SetCellStyle(sheet, cell, cell, styleID)
		}
	}
	cellValues := map[string]string{
		"A1": headerMedicalArea.Frame,
		"A2": headerMedicalArea.Title,
		"A3": headerMedicalArea.Area,
		"A6": "TOTAL",
		"B3": "TOTAL",
		"E3": "PROCEDIMIENTOS REALIZADOS",
	}

	for i, proc := range dentistryProcedures {
		procCol, _ := excelize.ColumnNumberToName(2 + (i+1)*3)
		cellValues[procCol+"4"] = proc
	}

	setCells(cellValues, styleHeaderID)

	columns := []string{
		"B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "AA", "AB", "AC", "AD", "AE",
	}

	values := []string{"T", "M", "F"}

	for i, col := range columns {
		value := values[i%len(values)]
		cell := col + "5"
		_ = f.SetCellStyle(sheet, cell, cell, styleHeaderID)
		_ = f.SetCellStyle(sheet, col+"6", col+"6", styleHeaderID)
		_ = f.SetCellValue(sheet, cell, value)
		_ = f.SetColWidth(sheet, col, col, 6)
	}

	_ = f.SetRowHeight(sheet, 4, 38)

	return f
}

func HeaderMedicalConsultingTrimesterExcel(f *excelize.File, sheet string, headerMedicalArea models.HeaderMedicalAreaExcel, styleHeaderID int, month, trimester, year string) *excelize.File {

	var mergeColumns = []models.MergeRange{
		{Column1: "A1", Column2: "G1"},
		{Column1: "A2", Column2: "G2"},
		{Column1: "A3", Column2: "G3"},
		{Column1: "A4", Column2: "A5"},
		{Column1: "B4", Column2: "B5"},
		{Column1: "C4", Column2: "E4"},
		{Column1: "F4", Column2: "F5"},
		{Column1: "G4", Column2: "G5"},
	}

	for _, merge := range mergeColumns {
		_ = f.MergeCell(sheet, merge.Column1, merge.Column2)
		_ = f.SetCellStyle(sheet, merge.Column1, merge.Column2, styleHeaderID)
	}

	setCells := func(cells map[string]string, styleID int) {
		for cell, value := range cells {
			_ = f.SetCellValue(sheet, cell, value)
			_ = f.SetCellStyle(sheet, cell, cell, styleID)
		}
	}

	cellValues := map[string]string{
		"A1": headerMedicalArea.Frame,
		"A2": headerMedicalArea.Title,
		"A3": fmt.Sprintf("%s %s", trimester, year),
		"A4": headerMedicalArea.Area,
		"B4": "SEXO",
		"C4": "MESES",
		"F4": "SUB TOTAL",
		"G4": "TOTAL",
	}

	cols := []string{"C5", "D5", "E5"}
	monthCodes := strings.Split(month, ",")
	for i, code := range monthCodes {
		if i < len(cols) {
			cellValues[cols[i]] = months[code]
		}
	}

	setCells(cellValues, styleHeaderID)

	for i, proc := range dentistryProcedures {
		procCol, _ := excelize.ColumnNumberToName(2 + (i+1)*3)
		cellValues[procCol+"4"] = proc
	}

	return f
}

func styleExcel(f *excelize.File) (int, int) {
	styleHeaderID, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			WrapText:   true,
			Horizontal: "center",
			Vertical:   "center",
		},
		Font: &excelize.Font{
			Bold: true,
			Size: 8,
		},
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#AFDF9F"},
			Pattern: 1,
		},
	})
	if err != nil {
		logger.Error.Println("Error creating style:", err)
	}

	styleBorderID, err := f.NewStyle(&excelize.Style{
		Border: []excelize.Border{
			{Type: "left", Color: "000000", Style: 1},
			{Type: "top", Color: "000000", Style: 1},
			{Type: "right", Color: "000000", Style: 1},
			{Type: "bottom", Color: "000000", Style: 1},
		},
	})
	if err != nil {
		logger.Error.Println("Error creating style:", err)
	}

	return styleHeaderID, styleBorderID
}

func SaveExcelAndReturnBase64(f *excelize.File, fileName string) (string, int, error) {
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("%s_%s.xlsx", fileName, timestamp)
	outputPath := fmt.Sprintf("./reports/medical_area/%s", filename)

	if err := f.SaveAs(outputPath); err != nil {
		return "", 0, fmt.Errorf("error saving Excel file: %w", err)
	}

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return "", 0, fmt.Errorf("error writing Excel to buffer: %w", err)
	}
	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())

	return base64Str, 0, nil
}

func getExcelColumns(from, to string) []string {
	var cols []string
	for col := from; ; {
		cols = append(cols, col)
		if col == to {
			break
		}
		col = nextExcelCol(col)
	}
	return cols
}

func nextExcelCol(col string) string {
	n := len(col)
	for i := n - 1; i >= 0; i-- {
		if col[i] < 'Z' {
			return col[:i] + string(col[i]+1) + col[i+1:]
		}
		col = col[:i] + "A" + col[i+1:]
	}
	return "A" + col
}

func createNursingExcelRow(consultation *models.ConsultationPatientsMedicalAreaExcel, row int) models.ExcelPageRow {
	if row == 1 {
		return models.ExcelPageRow{
			Row: 1,
			Columns: []models.ExcelPageColumn{
				{Column: "A", Value: "Fecha consulta"},
				{Column: "B", Value: "Tipo de persona"},
				{Column: "C", Value: "Código SGA"},
				{Column: "D", Value: "DNI"},
				{Column: "E", Value: "Nombre Completo"},
				{Column: "F", Value: "Sexo"},
				{Column: "G", Value: "Fecha de nacimiento"},
				{Column: "H", Value: "Número de teléfono"},
				{Column: "I", Value: "Servicios"},
				{Column: "J", Value: "Lugar de procedencia"},
				{Column: "K", Value: "Dirección residencia"},
				{Column: "L", Value: "Procedimiento"},
				{Column: "M", Value: "Número de recibo"},
				{Column: "N", Value: "Monto"},
			},
		}
	}

	return models.ExcelPageRow{
		Row: row,
		Columns: []models.ExcelPageColumn{
			{Column: "A", Value: consultation.FechaConsulta},
			{Column: "B", Value: consultation.TipoPersona},
			{Column: "C", Value: consultation.CodigoSGA},
			{Column: "D", Value: consultation.DNI},
			{Column: "E", Value: consultation.NombreCompleto},
			{Column: "F", Value: consultation.Sexo},
			{Column: "G", Value: consultation.FechaNacimiento},
			{Column: "H", Value: consultation.NumeroCelular},
			{Column: "I", Value: consultation.Servicios},
			{Column: "J", Value: consultation.Procedencia},
			{Column: "K", Value: consultation.DireccionResidencia},
			{Column: "L", Value: consultation.Procedimiento},
			{Column: "M", Value: consultation.NumeroRecibo},
			{Column: "N", Value: consultation.Monto},
		},
	}
}
func createDentistryExcelRow(consultation *models.ConsultationPatientsMedicalAreaExcel, row int) models.ExcelPageRow {
	if row == 1 {
		return models.ExcelPageRow{
			Row: 1,
			Columns: []models.ExcelPageColumn{
				{Column: "A", Value: "Fecha consulta"},
				{Column: "B", Value: "Tipo de persona"},
				{Column: "C", Value: "DNI"},
				{Column: "D", Value: "Nombre Completo"},
				{Column: "E", Value: "Sexo"},
				{Column: "F", Value: "Fecha de nacimiento"},
				{Column: "G", Value: "Ocupación"},
				{Column: "H", Value: "Número de teléfono"},
				{Column: "I", Value: "Servicios"},
				{Column: "J", Value: "Tipo de procedimiento"},
				{Column: "K", Value: "Recibo"},
				{Column: "L", Value: "Costo"},
				{Column: "M", Value: "Fecha de pago"},
				{Column: "N", Value: "Pieza dental"},
			},
		}
	}

	return models.ExcelPageRow{
		Row: row,
		Columns: []models.ExcelPageColumn{
			{Column: "A", Value: consultation.FechaConsulta},
			{Column: "B", Value: consultation.TipoPersona},
			{Column: "C", Value: consultation.DNI},
			{Column: "D", Value: consultation.NombreCompleto},
			{Column: "E", Value: consultation.Sexo},
			{Column: "F", Value: consultation.FechaNacimiento},
			{Column: "G", Value: consultation.Ocupacion},
			{Column: "H", Value: consultation.NumeroCelular},
			{Column: "I", Value: consultation.Servicios},
			{Column: "J", Value: consultation.TipoProcedimiento},
			{Column: "K", Value: consultation.Recibo},
			{Column: "L", Value: consultation.Costo},
			{Column: "M", Value: consultation.FechaPago},
			{Column: "N", Value: consultation.PiezaDental},
		},
	}
}
func createMedicalExcelRow(consultation *models.ConsultationPatientsMedicalAreaExcel, row int) models.ExcelPageRow {
	if row == 1 {
		return models.ExcelPageRow{
			Row: 1,
			Columns: []models.ExcelPageColumn{
				{Column: "A", Value: "Fecha consulta"},
				{Column: "B", Value: "Tipo de persona"},
				{Column: "C", Value: "DNI"},
				{Column: "D", Value: "Nombre Completo"},
				{Column: "E", Value: "Sexo"},
				{Column: "F", Value: "Fecha de nacimiento"},
				{Column: "G", Value: "Número de teléfono"},
				{Column: "H", Value: "Servicios"},
			},
		}
	}

	return models.ExcelPageRow{
		Row: row,
		Columns: []models.ExcelPageColumn{
			{Column: "A", Value: consultation.FechaConsulta},
			{Column: "B", Value: consultation.TipoPersona},
			{Column: "C", Value: consultation.DNI},
			{Column: "D", Value: consultation.NombreCompleto},
			{Column: "E", Value: consultation.Sexo},
			{Column: "F", Value: consultation.FechaNacimiento},
			{Column: "G", Value: consultation.NumeroCelular},
			{Column: "H", Value: consultation.Servicios},
		},
	}
}
