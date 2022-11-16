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
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				stop := c.Int("stop")
				if stop <= 0 {
					fmt.Println("stop can not be negative")
				}
				for i := 1; i <= stop; i++ {
					fmt.Println(i)
				}
				return nil
			},
		},

		{
			Name: "down", ShortName: "d",
			Usage: "count down",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "start, s",
					Usage: "value to count down from 10",
					Value: 10,
				},
			},
			Action: func(c *cli.Context) error {
				start := c.Int("start")
				if start <= 0 {
					fmt.Println("stop can not be negative")
				}
				for i := start; i >= 0; i-- {
					fmt.Println(i)
				}
				return nil
			},
		},
	}
	app.Run(os.Args)
}
