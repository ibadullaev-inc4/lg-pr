package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "test cli app"
	app.Description = "New cli application"
	app.Usage = "New cli application"
	app.Author = "Nariman"
	app.Email = "nariman@inc4.net"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "name,n",
			Value: "No name",
		},
		cli.StringFlag{
			Name:  "surname,s",
			Value: "No Surname",
		},
	}
	app.Action = func(c *cli.Context) error {
		name := c.String("name")
		surname := c.String("surname")
		fmt.Println(name, surname)
		return nil
	}

	app.Run(os.Args)
}
