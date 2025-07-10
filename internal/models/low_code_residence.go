package models

import (
	"github.com/asaskevich/govalidator"
	"time"
)

type Residence struct {
	ID          string         `json:"id" valid:"uuid,required"`
	Name        string         `json:"name" valid:"required"`
	Gender      string         `json:"gender" valid:"in(femenino|masculino),required"`
	Description string         `json:"description" valid:"required"`
	Address     string         `json:"address" valid:"required"`
	Status      string         `json:"status" valid:"in(mantenimiento|deshabilitado|habilitado),required"`
	Floors      []*Floor       `json:"floors,omitempty" valid:"-"`
	Config      *Configuration `json:"configuration,omitempty" valid:"-"`
	Submissions []*Submission  `json:"submissions,omitempty" valid:"-"`
}

type Submission struct {
	ID    int64  `json:"id" valid:"uuid,required"`
	Name  string `json:"name" valid:"required"`
	Start string `json:"start" valid:"required"`
	State bool   `json:"state" valid:"required"`
}

type Floor struct {
	Floor int    `json:"floor" valid:"numeric,required"`
	Rooms []Room `json:"rooms" valid:"required"`
}

type Room struct {
	ID       string `json:"id" valid:"uuid,required"`
	Number   int    `json:"number" valid:"numeric,required"`
	Capacity int    `json:"capacity" valid:"numeric,required"`
	Status   string `json:"status" valid:"in(mantenimiento|deshabilitado|habilitado),required"`
}

type RoomStudentsSimple struct {
	ID       string                   `json:"id" valid:"uuid,required"`
	Number   int                      `json:"number" valid:"numeric,required"`
	Capacity int                      `json:"capacity" valid:"numeric,required"`
	Status   string                   `json:"status" valid:"in(mantenimiento|deshabilitado|habilitado),required"`
	Floor    int                      `json:"floor" valid:"numeric,required"`
	Students []StudentInfoWithoutRoom `json:"students,omitempty" valid:"-"`
}

type Configuration struct {
	ID                      string  `json:"id" valid:"uuid"`
	PercentageFcea          float64 `json:"percentage_fcea" valid:"required"`
	PercentageEngineering   float64 `json:"percentage_engineering" valid:"required"`
	MinimumGradeFcea        float64 `json:"minimum_grade_fcea" valid:"required"`
	MinimumGradeEngineering float64 `json:"minimum_grade_engineering" valid:"required"`
	IsNewbie                bool    `json:"is_newbie"`
}

type Student struct {
	StudentInfo   StudentInfo    `json:"student"`
	AssignedGoods []AssignedGood `json:"assigned_goods"`
	RoomMates     []RoomMate     `json:"room_mates"`
	Sanctions     []Sanction     `json:"sanctions"`
}

type StudentInfo struct {
	ID                   int64  `json:"id" db:"id" valid:"uuid,required"`
	NumberIdentification string `json:"number_identification" db:"number_identification"`
	FullName             string `json:"full_name" db:"full_name"`
	Department           string `json:"department,omitempty" db:"department"`
	Sex                  string `json:"sex,omitempty" db:"sex"`
	Code                 string `json:"code" db:"code"`
	ProfessionalSchool   string `json:"professional_school" db:"professional_school"`
	Faculty              string `json:"faculty" db:"faculty"`
	Room                 *Room  `json:"room" db:"room"`
	Residence            string `json:"residence,omitempty" db:"residence"`
	AdmissionDate        string `json:"admission_date" db:"admission_date"`
}

type StudentInfoWithoutRoom struct {
	ID                 int64  `json:"id" db:"id" valid:"uuid,required"`
	FullName           string `json:"full_name" db:"full_name"`
	Code               string `json:"code" db:"code"`
	ProfessionalSchool string `json:"professional_school" db:"professional_school"`
	Faculty            string `json:"faculty" db:"faculty"`
	AdmissionDate      string `json:"admission_date" db:"admission_date"`
}

type AssignedGood struct {
	Code string `json:"code"`
}

type RoomMate struct {
	FullName string `json:"full_name"`
	Code     string `json:"code"`
}

type Sanction struct {
	Description string `json:"description"`
	Date        string `json:"date"`
}

type DataWithoutAssignation struct {
	Residence ResidenceForAssignation `json:"residence"`
	Students  []StudentForAssignation `json:"students"`
}

type ResidenceForAssignation struct {
	ID          string        `json:"id" valid:"uuid,required"`
	Name        string        `json:"name" valid:"matches(^[\\p{L} ]+$),required"`
	Gender      string        `json:"gender" valid:"in(femenino|masculino),required"`
	Description string        `json:"description" valid:"matches(^[\\p{L} ]+$),required"`
	Status      string        `json:"status" valid:"in(mantenimiento|deshabilitado|habilitado),required"`
	Rooms       []Room        `json:"rooms" valid:"-"`
	Config      Configuration `json:"configuration" valid:"-"`
}

type StudentForAssignation struct {
	ID                    int64  `json:"id"`
	StudenCode            string `json:"codigo_estudiante"`
	DNI                   string `json:"DNI"`
	Sex                   string `json:"sexo"`
	Facultad              string `json:"facultad"`
	ProfessionalSchool    string `json:"escuela_profesional"`
	NumSemestersCompleted string `json:"num_semestres_cursados"`
	Age                   int32  `json:"edad"`
	PPS                   string `json:"pps"`
	PPA                   string `json:"ppa"`
	TCA                   string `json:"tca"`
}

type ResidenceGPT struct {
	ResidenceId string    `json:"residence_id"`
	Rooms       []RoomGPT `json:"rooms"`
}

type RoomGPT struct {
	RoomId   string       `json:"room_id"`
	Capacity int32        `json:"capacity"`
	Students []StudentGPT `json:"students"`
}

type StudentGPT struct {
	StudentId  int64  `json:"student_id"`
	CodStudent string `json:"cod_student"`
}

type StudentExcel struct {
	ID                    int64  `json:"id" db:"id" valid:"uuid,required"`
	Code                  string `json:"code" db:"code"`
	FullName              string `json:"full_name" db:"full_name"`
	Sex                   string `json:"sex" db:"sex"`
	Home                  string `json:"home" db:"home"`
	Address               string `json:"address" db:"address"`
	NumSemestersCompleted string `json:"number_semesters_completed" db:"number_semesters_completed"`
	PPS                   string `json:"pps" db:"pps"`
	ProfessionalSchool    string `json:"professional_school" db:"professional_school"`
	Faculty               string `json:"faculty" db:"faculty"`
	Room                  string `json:"room" db:"room"`
	Residence             string `json:"residence,omitempty" db:"residence"`
	AdmissionDate         string `json:"admission_date" db:"admission_date"`
	Status                string `json:"status" db:"status"`
	UpdateDate            string `json:"update_date" db:"update_date"`
}

type AssignmentRoom struct {
	ID             string    `json:"id" valid:"uuid,required"`
	AssignmentDate time.Time `json:"fecha_asignacion" valid:"-"`
	Status         string    `json:"estado" valid:"required"`
}

func (m *Residence) ValidResidence() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *Residence) ValidStudent() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (m *Configuration) ValidConfig() (bool, error) {
	result, err := govalidator.ValidateStruct(m)
	if err != nil {
		return result, err
	}
	return result, nil
}
