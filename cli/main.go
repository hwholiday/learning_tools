package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "howie"
	app.Version = "1.0.0"
	app.Email = "520@520.com"
	app.Commands = []cli.Command{
		{
			Name:    "prot",
			Aliases: []string{"p"},
			Usage:   "用户端口",
			Before: func(c *cli.Context) error {
				fmt.Println("Before")
				return nil
			},
			Action: func(c *cli.Context) {
				fmt.Println("Action")
				fmt.Println(c.Args().First())
			},
			After: func(c *cli.Context) error {
				fmt.Println("After")
				return nil
			},
		},
		{
			Name:    "isdebug",
			Aliases: []string{"p"},
			Usage:   "is debug",
			Before: func(c *cli.Context) error {
				fmt.Println("Before")
				return nil
			},
			Action: func(c *cli.Context) {
				fmt.Println("Action")
				fmt.Println(c.Args().First())
			},
			After: func(c *cli.Context) error {
				fmt.Println("After")
				return nil
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
