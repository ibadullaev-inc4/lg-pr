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
	}
	app.Action = func(c *cli.Context) error {
		name := c.GlobalString("name")
		fmt.Printf("Hello %s\n", name)
		return nil
	}
	app.Run(os.Args)
}
