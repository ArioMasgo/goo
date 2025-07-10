package low_code_residences

import (
	"dbu-api/internal/env"
	"dbu-api/internal/logger"
	"dbu-api/internal/models"
	"dbu-api/internal/ws"
	"dbu-api/pkg/residence"
	"dbu-api/pkg/submission"
	"dbu-api/pkg/submission/alumnos"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type RobotResidenceService struct {
	db   *sqlx.DB
	usr  *models.User
	txID string
}

type PortsServerRobotResidence interface {
	AssignationResidenceLowCode(residenceID string) (string, int, error)
}

func NewRobotResidence(db *sqlx.DB, usr *models.User, txID string) PortsServerRobotResidence {
	return &RobotResidenceService{db: db, usr: usr, txID: txID}
}

func (s *RobotResidenceService) AssignationResidenceLowCode(residenceID string) (string, int, error) {
	srv := submission.NewServerSubmission(s.db, s.usr, s.txID)
	activeSubmission, _, err := srv.SrvConvocatorias.GetActiveSubmissions()
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get active submission", err)
		return "", 201, err
	}

	if activeSubmission == nil {
		return "", 301, errors.New("no active submission")
	}

	srvService := residence.NewServerResidence(s.db, s.usr, s.txID)
	residenceData, codErrResidence, err := srvService.SrvResidence.GetResidenceByID(residenceID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get user by username", err)
		return "", codErrResidence, err
	}

	if residenceData == nil {
		return "", 301, nil
	}

	roomsByResidence, err := srvService.SrvRoom.GetAllRoomsBySubmissionIDResidenceID(activeSubmission.ID, residenceData.ID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get user by username", err)
		return "", 301, errors.New("permissions not found")
	}

	if roomsByResidence == nil {
		logger.Error.Println(s.txID, " - couldn't get user by username")
		return "", 301, errors.New("permissions not found")
	}

	var roomForAssignment []models.Room
	for _, room := range roomsByResidence {
		roomForAssignment = append(roomForAssignment, models.Room{
			ID:       room.ID,
			Number:   room.Number,
			Capacity: room.Capacity,
			Status:   room.Status,
		})
	}

	configuration, _, errConfig := srvService.SrvResidenceConfiguration.GetResidenceConfigurationByResidenceID(residenceData.ID)
	if errConfig != nil {
		logger.Error.Println(s.txID, " - couldn't get residence configuration", err)
		return "", 310, err
	}

	if configuration == nil {
		logger.Error.Println(s.txID, " - No found residence configuration", err)
		return "", 310, errors.New("residence configuration not found")
	}

	residenceForAssignation := models.ResidenceForAssignation{
		ID:          residenceData.ID,
		Name:        residenceData.Name,
		Gender:      residenceData.Gender,
		Description: residenceData.Description,
		Status:      residenceData.Status,
		Rooms:       roomForAssignment,
		Config: models.Configuration{
			ID:                      configuration.ID,
			PercentageFcea:          configuration.PercentageFcea,
			PercentageEngineering:   configuration.PercentageEngineering,
			MinimumGradeFcea:        configuration.MinimumGradeFcea,
			MinimumGradeEngineering: configuration.MinimumGradeEngineering,
			IsNewbie:                configuration.IsNewbie,
		},
	}

	var students []*alumnos.Alumno

	if configuration.IsNewbie {
		students, err = srv.SrvAlumnos.GetStudentsAcceptedBySubmissionNewbie(activeSubmission.ID)
		if err != nil {
			logger.Error.Printf("%s - couldn't get students: %v", s.txID, err)
			return "", 301, err
		}
	} else {
		students, err = srv.SrvAlumnos.GetStudentsAcceptedBySubmission(activeSubmission.ID)
		if err != nil {
			logger.Error.Printf("%s - couldn't get students: %v", s.txID, err)
			return "", 301, err
		}
	}

	// Construir estudiantes para asignación
	studentsForAssignation := make([]models.StudentForAssignation, 0, len(students))
	for _, student := range students {
		studentForAssignation := models.StudentForAssignation{
			ID:                    student.ID,
			StudenCode:            student.CodigoEstudiante,
			DNI:                   student.DNI,
			Sex:                   student.Sexo,
			Facultad:              student.Facultad,
			ProfessionalSchool:    student.EscuelaProfesional,
			NumSemestersCompleted: student.NumSemestresCursados,
			Age:                   student.Edad,
			PPS:                   student.PPS,
			PPA:                   student.PPA,
			TCA:                   student.TCA,
		}
		studentsForAssignation = append(studentsForAssignation, studentForAssignation)
	}

	// Construir la respuesta final
	response := models.DataWithoutAssignation{
		Residence: residenceForAssignation,
		Students:  studentsForAssignation,
	}

	if len(response.Students) == 0 {
		return "", 200, nil
	}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return "", 23, nil
	}

	//systemPrompt := "You are a JSON generator. Always respond with valid JSON that matches the specified structure. Do not include any explanatory text outside the JSON."
	systemPrompt := "You are a JSON generator. Your entire response must be a single valid JSON object. Do not include markdown formatting, code blocks, backticks, or any other text. Start your response with { and end with }. Any other format will cause parsing errors."
	userPrompt := fmt.Sprintf(`
Actúa como un sistema automatizado de asignación de residencias universitarias. Tu tarea es asignar estudiantes a habitaciones siguiendo un conjunto específico de reglas y criterios de priorización.

### Datos de Entrada
%s

### Reglas de Asignación

#### 1. Validación Inicial de Estudiantes
- **Género**: Los estudiantes deben coincidir con el género de la residencia
- **Estado**: Solo procesar estudiantes con status "habilitado"

#### 2. Criterios para Ingresantes (is_newbie = true)
- Solo aceptar estudiantes con num_semestres_cursados ≤ 1
- Para estudiantes de FCEA:
  * PPS ≥ minimum_grade_fcea
  * Si PPS = 0, considerar como válido solo para ingresantes del primer semestre
- Para estudiantes de Ingeniería:
  * PPS ≥ minimum_grade_engineering
  * Si PPS = 0, considerar como válido solo para ingresantes del primer semestre

#### 3. Distribución por Facultades
- FCEA (Facultad de Ciencias Económicas y Administrativas):
  * Porcentaje de cupos = percentage_fcea
  * Incluye: Administración, Contabilidad, Economía
- Ingeniería:
  * Porcentaje de cupos = percentage_engineering
  * Incluye: Todas las carreras que contengan "INGENIERIA" en facultad o escuela_profesional

#### 4. Priorización de Asignación
1. Para estudiantes con PPS > 0:
   - Ordenar de mayor a menor PPS dentro de cada facultad
2. Para estudiantes con PPS = 0 (nuevos ingresantes):
   - Asignar por orden de código de estudiante (más reciente primero)

#### 5. Asignación de Habitaciones
- Respetar la capacidad máxima de cada habitación
- Distribuir equitativamente entre habitaciones disponibles
- No dejar habitaciones parcialmente vacías si hay estudiantes pendientes
- Mantener agrupados estudiantes de la misma facultad cuando sea posible

### Formato de Salida
{
    "residence_id": "string",
    "rooms": [
        {
            "room_id": "string",
            "capacity": number,
            "students": [
                {
                    "student_id": number,
                    "cod_student": "string"
                }
            ]
        }
    ]
}

### Validaciones Adicionales
1. Verificar que la suma de porcentajes (percentage_fcea + percentage_engineering) = 100
2. Validar que los datos de estudiantes estén completos (no nulos)

### Manejo de Casos Especiales
1. Si no se llena la cuota de una facultad, los espacios pueden ser asignados a la otra
2. En caso de empate en PPS, usar el código de estudiante como desempate
3. Si una habitación queda con un espacio libre y no hay más estudiantes de la misma facultad, se puede asignar a un estudiante de otra facultad`, string(jsonBytes))

	e := env.NewConfiguration()

	request := models.GPTRequest{
		Model: e.IA.Model,
		Messages: []models.GPTMessages{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userPrompt},
		},
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return "", 111, fmt.Errorf("error marshaling request: %w", err)
	}

	requestBody := ws.WebServiceSchemeRequest{
		Url:      e.IA.UrlApi,
		Method:   "POST",
		Duration: time.Duration(3) * time.Minute,
		Headers: map[string]string{
			"Content-Type":  "application/json",
			"Authorization": fmt.Sprintf("Bearer %v", e.IA.Token),
		},
		Payload: string(jsonData),
	}

	responseGPT, _, err := ws.CallWebService(requestBody)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't consume gpt", err)
		return "", 400, err
	}

	if len(responseGPT) == 0 {
		return "", 203, fmt.Errorf("datos vacíos")
	}

	var structResponseGPT models.GPTResponse

	if err := json.Unmarshal(responseGPT, &structResponseGPT); err != nil {
		return "", 201, fmt.Errorf("error deserializar JSON with response GPT: %w", err)
	}

	residenceRobot, _, err := srvService.SrvResidenceRobot.CreateResidenceRobot(residenceData.ID, structResponseGPT.Usage.PromptTokens,
		structResponseGPT.Usage.CompletionTokens, structResponseGPT.Usage.TotalTokens)
	if err != nil {
		return "", 201, fmt.Errorf("error register response GPT: %w", err)
	}

	if len(structResponseGPT.Choices) == 0 {
		return "", 231, fmt.Errorf("no found choices")
	}

	var structResponseResidenceGPT models.ResidenceGPT

	if err := json.Unmarshal([]byte(structResponseGPT.Choices[0].Message.Content), &structResponseResidenceGPT); err != nil {
		return "", 201, fmt.Errorf("error al deserializar JSON with response GPT: %w", err)
	}

	codErrAssignment, err := srvService.SrvAssignmentRoom.MultiAssignRoom(activeSubmission.ID, structResponseResidenceGPT.Rooms)

	if err != nil {
		return "", codErrAssignment, fmt.Errorf("error register assignment room: %w", err)
	}

	return residenceRobot.ID, 200, nil
}
