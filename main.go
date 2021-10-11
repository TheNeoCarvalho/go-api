package main

import (
	"github.com/TheNeoCarvalho/go-api/database"
	"github.com/TheNeoCarvalho/go-api/server"
)

func main() {
	database.StartDb()
	server := server.NewServer()

	server.Run()
}
