package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mrbboomm/go-execise/app/runner"
)

type Runner struct {

}

func (e *Runner) Execute(c *fiber.Ctx) error {
	_runner := runner.New()
	_runner.Run()
	return c.JSON(_runner.Result())
}