package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "hello_cli"
	app.Usage = "Print Hello World"
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "name, n"},
		cli.StringFlag{Name: "surname,s"},
		cli.IntFlag{Name: "age, a", Value: 100},
	}
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		surname := c.GlobalString("surname")
		age := c.GlobalInt("age")
		fmt.Printf("Hello %s-%s-%d\n", name, surname, age)
		return nil
	}
	app.Run(os.Args)
}
