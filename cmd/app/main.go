package main

import (
	"fmt"
	"log"


	"github.com/ASeegull/edriver-space/api/server"
	"github.com/ASeegull/edriver-space/config"
	"github.com/ASeegull/edriver-space/logger"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello Lv-644.Go!")

	// Path to env
	if err := godotenv.Load("env/docker/postgres.env", "env/docker/app.env"); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	//Initializing Logger
	logger.LogInit()

	// Loading config values
	conf, err := config.LoadConfig("")

	if err != nil {
		logger.LogErr(err)
	}

	//Creating and starting server
	s := server.NewServer()
	logger.LogFatal(s.Start(":" + conf.ServerPort))

	// Verify if connection is ok
	conn := MustGetConnection()
	conErr := conn.Ping()
	if conErr != nil {
		log.Fatal(conErr)
	}
	fmt.Println("Successfully connected ✓")
	defer func() {
		if conErr := conn.Close(); conErr != nil {
			fmt.Println("db connection closed.")
		}
	}()
}
