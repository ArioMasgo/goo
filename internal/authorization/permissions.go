package authorization

import (
	"dbu-api/internal/middleware"
	"dbu-api/internal/models"
	auth_orchestrator "dbu-api/pkg/orchestrator/auth"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"net/http"
	"strings"
)

func Permissions(db *sqlx.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {
		method := c.Method()
		path := strings.ToLower(c.Path())
		params := c.AllParams()

		user, err := middleware.GetUser("", db)
		if err != nil {
			return c.Status(http.StatusUnauthorized).SendString("Unauthorized: " + err.Error())
		}

		if user != nil {
			return c.Status(http.StatusUnauthorized).SendString("Unauthorized: invalid user")
		}

		fmt.Println(method, " ", path, " ", params)
		return nil
		//return c.Status(http.StatusUnauthorized).SendString("You don't have a permissions")
	}
}

func ValidPermissions(user *models.User, db *sqlx.DB, c *fiber.Ctx) error {
	if user == nil {
		return errors.New("user is nil")
	}

	if user.StatusID != 3 {
		return errors.New("user doesn't active")
	}

	srv := auth_orchestrator.NewPermissions(db, uuid.New().String())
	err := srv.Permissions(user.IDLevelUser, c.Method(), c.Path())
	if err != nil {
		return err
	}

	return nil
}
