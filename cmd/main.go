package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/bayuscodings/telloservice"
	"github.com/rs/zerolog/log"

	"github.com/bayuscodings/telloservice/app/routes"
	"github.com/bayuscodings/telloservice/seeding/seeders"
	"github.com/joho/godotenv"
)

// @title           TelloService API
// @version         1.0
// @description     API documentation for TelloService.
// @termsOfService  http://swagger.io/terms/
// @contact.name    Ogunbayo Abayomi
// @contact.url     https://github.com/bayuscodings
// @contact.email   ogunbayo.abayo@gmail.com
// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html
// @host            localhost:50055
// @BasePath        /v1
// @securityDefinitions.apikey  BearerAuth
// @in              header
// @name            Authorization

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Some error occured. Err: %s", err)
	}

	StartServer()
}

func StartServer() {
	app := &telloservice.ApplicationHandler{}

	// prepare config from environment variables
	app.PrepareConfig()

	// initialize the DB connection and seeding
	app.InitDB()
	if app.Config.Database.RunSeeding {
		seeders.Init(app)
	}

	// initialize JWT token maker
	app.InitTokenMaker()

	// initialize ZeroLog
	app.InitLogger()

	// prepare Rest API endpoints
	router := routes.BuildRoute(app)

	// awareness
	host, _ := os.Hostname()
	log.Info().Msg(fmt.Sprintf("Starting the server at http://%s:%v\n", host, app.Config.AppPort))

	// start the webserver and keep the app alive
	err := http.ListenAndServe(":"+app.Config.AppPort, router)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Failed to start server: %v", err))
	}
}
