package main

import (
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Usage = "count up and down"
	app.Commands = []cli.Command{
		{
			Name: "up", ShortName: "u",
			Usage: "count up",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "stop, s",
					Usage: "value to count up to 10",
					Value: 20,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("stop")
				if start <= 0 {
					fmt.Println("stop can not be negative")
				}
				fmt.Println(start)
				for i := 1; i <= start; i++ {
					fmt.Println(i)
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
