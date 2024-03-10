package main

import (
	"context"

	"github.com/dalmazox/todo-api/internal/server"
)

// @title TODO Api
// @version 1.0
// @description This is a simple TODO API
func main() {
	server.NewServer(server.ServerOptions{
		Context: context.Background(),
	}).Start()
}
