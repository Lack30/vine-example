package main

import (
	"fmt"

	"github.com/lack-io/cli"
	vine "github.com/lack-io/vine/service"
)

func main() {
	service := vine.NewService(
		vine.Flags(&cli.StringFlag{
			Name:  "environment",
			Usage: "The environment",
		}),
	)

	service.Init(
		vine.Action(func(c *cli.Context) error {
			env := c.String("environment")
			if len(env) > 0 {
				fmt.Println("Environment set to", env)
			}

			return nil
		}),
	)
}
