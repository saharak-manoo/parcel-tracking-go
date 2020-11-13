package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/saharak/parcel-tracking-go/app/controllers"
	"github.com/saharak/parcel-tracking-go/app/db"
)

var server = controllers.Server{}

func init() {
	// Loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print(".env file found")
	}
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("API_PORT")
		log.Print("No Port In Heroku " + port)
	}

	return ":" + port
}

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		log.Print("We are getting values")
	}

	server.Initialize(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	// DB Migration
	migration := db.Migration{}
	migration.Load(server.DB)

	// DB Seed
	seeder := db.Seeder{}
	seeder.Load(server.DB)

	apiPort := getPort()
	log.Print("Running on ENV: " + os.Getenv("APP_ENV"))
	log.Print("Listening to port: " + apiPort)

	server.Run(apiPort)
}