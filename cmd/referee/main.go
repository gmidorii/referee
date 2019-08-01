package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "referee"
	app.Commands = []cli.Command{
		{
			Name:    "judge",
			Aliases: []string{"j"},
			Action: func(c *cli.Context) error {
				fmt.Println("Hello referee")
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
