package main

import "echo-service/internal/server"

func main() {
	srv := server.New(server.Config{})

	srv.Start()
}
