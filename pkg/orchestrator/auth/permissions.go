package auth_orchestrator

import (
	"dbu-api/internal/logger"
	"dbu-api/pkg/core"
	"errors"
	"github.com/jmoiron/sqlx"
	"regexp"
	"strconv"
	"strings"
)

type PermissionsService struct {
	db   *sqlx.DB
	txID string
}

type PortsServerPermissions interface {
	Permissions(levelUserID int64, method, path string) error
}

func NewPermissions(db *sqlx.DB, txID string) PortsServerPermissions {
	return &PermissionsService{db: db, txID: txID}
}

func (s *PermissionsService) Permissions(levelUserID int64, method, path string) error {
	srvCoreService := core.NewServerCore(s.db, nil, s.txID)
	levelUPermissions, err := srvCoreService.SrvLevelUserPermissions.GetAllByLevelUser(levelUserID)
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get user by username", err)
		return err
	}

	if levelUPermissions == nil {
		logger.Error.Println(s.txID, " - no found permissions")
		return errors.New("permissions not found")
	}

	var ids []string

	for _, levelUPermission := range levelUPermissions {
		ids = append(ids, strconv.FormatInt(levelUPermission.PermissionID, 10))
	}

	if ids == nil {
		return errors.New("Permissions not found")
	}

	if len(ids) == 0 {
		return errors.New("Permissions not found")
	}

	permissions, err := srvCoreService.SrvPermissions.GetAllByIDs(strings.Join(ids, ","))
	if err != nil {
		logger.Error.Println(s.txID, " - couldn't get user by username", err)
		return err
	}

	if permissions == nil {
		logger.Error.Println(s.txID, " - no found permissions")
		return errors.New("permissions not found")
	}

	for _, permission := range permissions {
		pattern := "^" + regexp.QuoteMeta(permission.Path) + "$"
		pattern = regexp.MustCompile(`\\\{[a-zA-Z0-9_]+\\\}`).ReplaceAllString(pattern, `[^/]+`)

		// Validar el path usando la expresión regular
		matched, err := regexp.MatchString(pattern, path)

		if err != nil {
			continue
		}
		if matched && permission.Method == method {
			return nil
		}
	}

	return errors.New("permissions not found")
}
