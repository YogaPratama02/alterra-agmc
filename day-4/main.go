package main

import (
	"alterra-agmc/day-4/routes"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf("Error read env file with err: %s", errEnv)
	}
	e := routes.Init()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))))
}
