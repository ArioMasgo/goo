package participantes

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/jmoiron/sqlx"
)

type ParticipanteRepository struct {
	DB *sqlx.DB
}

func NewParticipanteRepository(db *sqlx.DB) *ParticipanteRepository {
	return &ParticipanteRepository{DB: db}
}

func (r *ParticipanteRepository) GetAll(page, pageSize int) ([]Participante, int, error) {
	var participantes []Participante
	offset := (page - 1) * pageSize
	var total int
	err := r.DB.Get(&total, "SELECT COUNT(*) FROM participantes")
	if err != nil {
		return nil, 0, err
	}
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	query := "SELECT * FROM participantes LIMIT ? OFFSET ?"
	err = r.DB.Select(&participantes, query, pageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	return participantes, totalPages, nil
}

func (r *ParticipanteRepository) GetByID(id int) (*Participante, error) {
	var participante Participante
	err := r.DB.Get(&participante, "SELECT * FROM participantes WHERE id_participante = ?", id)
	return &participante, err
}

func (r *ParticipanteRepository) GetByDNI(dni string) (*Participante, error) {
	var participante Participante
	err := r.DB.Get(&participante, "SELECT * FROM participantes WHERE dni = ?", dni)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &participante, nil
}

func (r *ParticipanteRepository) Create(p *Participante) (int, error) {
	query := `
        INSERT INTO participantes (
            tipo_participante, nombre, apellido, dni, estado, created_at, updated_at, 
            colegio_procedencia, anio_ingreso, escuela, codigo_estudiante, 
            fecha_nacimiento, edad, lugar_nacimiento, modalidad_ingreso, 
            numero_atencion, sexo, num_telefono, estado_evaluacion, 
            diagnostico, con_quienes_vive_actualmente, semestre_cursa, direccion,
			profesion, estado_civil, labora_en_unas, grado_instruccion
        ) VALUES (
            :tipo_participante, :nombre, :apellido, :dni, :estado, :created_at, :updated_at, 
            :colegio_procedencia, :anio_ingreso, :escuela, :codigo_estudiante, 
            :fecha_nacimiento, :edad, :lugar_nacimiento, :modalidad_ingreso, 
            :numero_atencion, :sexo, :num_telefono, :estado_evaluacion, 
            :diagnostico, :con_quienes_vive_actualmente, :semestre_cursa, :direccion, 
			:profesion, :estado_civil, :labora_en_unas, :grado_instruccion
        )`

	result, err := r.DB.NamedExec(query, p)
	if err != nil {
		return 0, fmt.Errorf("error al insertar participante: %w", err)
	}

	idParticipante, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("error al obtener el ID del participante: %w", err)
	}

	return int(idParticipante), nil
}

func (r *ParticipanteRepository) Update(id int, p *Participante) error {
	var updates []string
	params := make(map[string]interface{})

	if p.Tipo != "" {
		updates = append(updates, "tipo_participante=:tipo_participante")
		params["tipo_participante"] = p.Tipo
	}
	if p.Nombre != "" {
		updates = append(updates, "nombre=:nombre")
		params["nombre"] = p.Nombre
	}
	if p.Apellido != "" {
		updates = append(updates, "apellido=:apellido")
		params["apellido"] = p.Apellido
	}
	if p.DNI != "" {
		updates = append(updates, "dni=:dni")
		params["dni"] = p.DNI
	}
	if p.Estado != "" {
		updates = append(updates, "estado=:estado")
		params["estado"] = p.Estado
	}
	if p.ColegioProcedencia != "" {
		updates = append(updates, "colegio_procedencia=:colegio_procedencia")
		params["colegio_procedencia"] = p.ColegioProcedencia
	}
	if p.AnioIngreso != 0 {
		updates = append(updates, "anio_ingreso=:anio_ingreso")
		params["anio_ingreso"] = p.AnioIngreso
	}
	if p.Escuela != "" {
		updates = append(updates, "escuela=:escuela")
		params["escuela"] = p.Escuela
	}
	if p.CodigoEstudiante != "" {
		updates = append(updates, "codigo_estudiante=:codigo_estudiante")
		params["codigo_estudiante"] = p.CodigoEstudiante
	}
	if p.FechaNacimiento != "" {
		updates = append(updates, "fecha_nacimiento=:fecha_nacimiento")
		params["fecha_nacimiento"] = p.FechaNacimiento
	}
	if p.Edad != 0 {
		updates = append(updates, "edad=:edad")
		params["edad"] = p.Edad
	}
	if p.LugarNacimiento != "" {
		updates = append(updates, "lugar_nacimiento=:lugar_nacimiento")
		params["lugar_nacimiento"] = p.LugarNacimiento
	}
	if p.ModalidadIngreso != "" {
		updates = append(updates, "modalidad_ingreso=:modalidad_ingreso")
		params["modalidad_ingreso"] = p.ModalidadIngreso
	}
	if p.NumeroAtencion != 0 {
		updates = append(updates, "numero_atencion=:numero_atencion")
		params["numero_atencion"] = p.NumeroAtencion
	}
	if p.Sexo != "" {
		updates = append(updates, "sexo=:sexo")
		params["sexo"] = p.Sexo
	}
	if p.NumTelefono != "" {
		updates = append(updates, "num_telefono=:num_telefono")
		params["num_telefono"] = p.NumTelefono
	}
	if p.EstadoEvaluacion != "" {
		updates = append(updates, "estado_evaluacion=:estado_evaluacion")
		params["estado_evaluacion"] = p.EstadoEvaluacion
	}
	if p.Diagnostico != "" {
		updates = append(updates, "diagnostico=:diagnostico")
		params["diagnostico"] = p.Diagnostico
	}
	if p.ConQuienesVive != "" {
		updates = append(updates, "con_quienes_vive_actualmente=:con_quienes_vive_actualmente")
		params["con_quienes_vive_actualmente"] = p.ConQuienesVive
	}
	if p.SemestreCursa != "" {
		updates = append(updates, "semestre_cursa=:semestre_cursa")
		params["semestre_cursa"] = p.SemestreCursa
	}

	if len(updates) == 0 {
		return fmt.Errorf("no fields to update")
	}

	query := "UPDATE participantes SET " + strings.Join(updates, ", ") + " WHERE id_participante=:id"
	params["id"] = id

	_, err := r.DB.NamedExec(query, params)
	if err != nil {
		return fmt.Errorf("error updating participante: %w", err)
	}
	return nil
}

func (r *ParticipanteRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM participantes WHERE id_participante = ?", id)
	return err
}

func (s *ParticipanteRepository) SearchParticipants(criteria map[string]interface{}) ([]Participante, error) {
	const sqlBase = `
        SELECT 
            p.id_participante, p.tipo_participante, p.nombre, p.apellido, 
            p.dni, p.estado, p.created_at, p.updated_at, p.num_telefono, 
            p.diagnostico, p.estado_evaluacion, p.colegio_procedencia, p.anio_ingreso, 
            p.escuela, p.codigo_estudiante, p.fecha_nacimiento, p.edad, 
            p.lugar_nacimiento, p.modalidad_ingreso, p.numero_atencion, p.sexo,
            p.con_quienes_vive_actualmente, p.semestre_cursa
        FROM participantes p
        WHERE 1=1
    `
	query := sqlBase
	args := []interface{}{}

	for key, value := range criteria {
		if value != "" && value != nil {
			query += fmt.Sprintf(" AND p.%s LIKE ?", key)
			args = append(args, "%"+strings.TrimSpace(fmt.Sprintf("%v", value))+"%")
		}
	}

	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error preparing query: %w", err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}
	defer rows.Close()

	var participantes []Participante
	for rows.Next() {
		var p Participante
		err := rows.Scan(
			&p.ID, &p.Tipo, &p.Nombre, &p.Apellido, &p.DNI,
			&p.Estado, &p.CreatedAt, &p.UpdatedAt, &p.NumTelefono, &p.Diagnostico,
			&p.EstadoEvaluacion, &p.ColegioProcedencia, &p.AnioIngreso, &p.Escuela,
			&p.CodigoEstudiante, &p.FechaNacimiento, &p.Edad, &p.LugarNacimiento,
			&p.ModalidadIngreso, &p.NumeroAtencion, &p.Sexo, &p.ConQuienesVive, &p.SemestreCursa,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		participantes = append(participantes, p)
	}

	return participantes, nil
}
