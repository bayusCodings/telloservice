package main

import (
	"fmt"
	"log"

	"github.com/bayuscodings/telloservice"
	"github.com/bayuscodings/telloservice/seeding/seeders"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Some error occured. Err: %s", err)
	}
	fmt.Println("Start an independent seeder....")
	app := &telloservice.ApplicationHandler{}

	app.PrepareConfig()
	app.InitDB()

	seeders.Init(app)

}
