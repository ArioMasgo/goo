package main

import (
	"dbu-api/api"
	_ "dbu-api/docs"
	"dbu-api/internal/env"
)

// @title DBU API
// @version 1.0
// @description DBU API
// @termsOfService https://dbu.com/terms/
// @contact.name API Support
// @contact.email juanm.campos@unas.edu.pe
// @license.name Comercial
// @license.url http://bjungle.net/licenses
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @scheme bearer
// @bearerFormat JWT
func main() {
	e := env.NewConfiguration()
	api.Start(e.App.Port, e.App.ServiceName, e.Router.LoggerHttp, e.Router.AllowedDomains)
}
