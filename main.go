package main

import (
	"github.com/1pedrohfreitas/pcams_back_go/database"
	"github.com/1pedrohfreitas/pcams_back_go/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
