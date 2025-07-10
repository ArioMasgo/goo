package dbx

import (
	"dbu-api/internal/env"
	"dbu-api/internal/logger"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	//_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbx      *sqlx.DB
	once     sync.Once
	DBEngine string
)

func init() {
	once.Do(func() {
		setConnection()
	})
}

func setConnection() {
	var err error
	c := env.NewConfiguration()
	DBEngine = c.DB.Engine

	// Check the connection
	dbx, err = sqlx.Open(DBEngine, connectionString("data"))
	if err != nil {
		logger.Error.Printf("no se puede conectar a la base de datos: %v", err)
		panic(err)
	}
	err = dbx.Ping()
	if err != nil {
		logger.Error.Printf("couldn't connect to database: %v", err)
		panic(err)
	}
	dbx.SetMaxIdleConns(5)
	dbx.SetConnMaxLifetime(2 * time.Minute)
	dbx.SetMaxOpenConns(95)
}

func GetConnection() *sqlx.DB {
	if dbx == nil {
		setConnection()
	}
	return dbx
}

func connectionString(t string) string {
	c := env.NewConfiguration()

	var host, database, username, password, instance string
	var port int
	switch t {
	case "data":
		host = c.DB.Server
		database = c.DB.Name
		username = c.DB.User
		password = c.DB.Password
		instance = c.DB.Instance
		port = c.DB.Port
	default:
		logger.Error.Print("El tipo de conexión no correspondea data/logs")
		return ""
	}
	switch strings.ToLower(DBEngine) {
	case "postgres":
		return fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=disable", database, username, password, host, port)
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	case "sqlserver":
		return fmt.Sprintf(
			"server=%s\\%s;User id=%s;database=%s;password=%s;port=%d", host, instance, username, database, password, port)
	}
	logger.Error.Print("el motor de bases de datos solicitado no está configurado aún")

	return ""
}
