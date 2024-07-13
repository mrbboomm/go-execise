package main

import (
	"github.com/mrbboomm/go-execise/app/loader"
	"github.com/mrbboomm/go-execise/app/runner"
)


func main() {
	// app := fiber.New()
	// runnerAPI := &handler.Runner{}
	// app.Post("/execute", runnerAPI.Execute)
	// app.Listen(":8080")
	loader := &loader.MockLoader{}
	_runner := runner.NewRunner(loader)
	_runner.Run()
	_runner.LogResult()
}

