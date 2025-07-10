package residence

import (
	"dbu-api/internal/models"
	"dbu-api/pkg/residence/asignaciones_cuartos"
	"dbu-api/pkg/residence/residence_configuration"
	"dbu-api/pkg/residence/residence_robot"
	"dbu-api/pkg/residence/residences"
	"dbu-api/pkg/residence/rooms"
	"github.com/jmoiron/sqlx"
)

type ServerResidence struct {
	SrvResidence              residences.PortsServerResidence
	SrvRoom                   rooms.PortsServerRoom
	SrvResidenceConfiguration residence_configuration.PortsServerResidenceConfiguration
	SrvAssignmentRoom         asignaciones_cuartos.RoomAssignmentServer
	SrvResidenceRobot         residence_robot.PortsServerResidenceRobot
}

func NewServerResidence(db *sqlx.DB, usr *models.User, txID string) *ServerResidence {
	return &ServerResidence{
		SrvResidence:              residences.NewResidenceService(residences.FactoryStorage(db, txID), usr, txID),
		SrvRoom:                   rooms.NewRoomService(rooms.FactoryStorage(db, usr, txID), usr, txID),
		SrvResidenceConfiguration: residence_configuration.NewResidenceConfigurationService(residence_configuration.FactoryStorage(db, txID), usr, txID),
		SrvAssignmentRoom:         asignaciones_cuartos.NewRoomAssignmentService(asignaciones_cuartos.FactoryStorage(db, usr, txID), usr, txID),
		SrvResidenceRobot:         residence_robot.NewResidenceRobotService(residence_robot.FactoryStorage(db, txID), usr, txID),
	}
}
