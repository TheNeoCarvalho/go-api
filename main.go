package main

import (
	"github.com/TheNeoCarvalho/go-api/database"
	"github.com/TheNeoCarvalho/go-api/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
