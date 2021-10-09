package main

import "github.com/TheNeoCarvalho/go-api/server"

func main() {
	server := server.NewServer()

	server.Run()
}
