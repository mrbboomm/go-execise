package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrbboomm/go-execise/app/loader"
	"github.com/mrbboomm/go-execise/app/runner"
)

type Runner struct {

}

func (e *Runner) Execute(c *fiber.Ctx) error {
	_runner := runner.NewRunner(&loader.MockLoader{})
	_runner.Run()
	return c.JSON(_runner.Result())
}