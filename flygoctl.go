package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func main() {
	var language string
	app := cli.NewApp()
	app.Name = "CtlTest"
	app.Usage = "hello world"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "port, p",
			Value: 8000,
			Usage: "listening port",
		},
		cli.StringFlag{
			Name:        "lang, l",
			Value:       "english",
			Usage:       "read from `FILE`",
			Destination: &language,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:     "add",
			Aliases:  []string{"a"},
			Usage:    "calc 1+1",
			Category: "arithmetic",
			Action: func(c *cli.Context) error {
				fmt.Println("1 + 1 = ", 1+1)
				return nil
			},
		},
		{
			Name:     "sub",
			Aliases:  []string{"s"},
			Usage:    "calc 5-3",
			Category: "arithmetic",
			Action: func(c *cli.Context) error {
				fmt.Println("5 - 3 = ", 5-3)
				return nil
			},
		},
		{
			Name:     "db",
			Usage:    "database operations",
			Category: "database",
			Subcommands: []cli.Command{
				{
					Name:  "insert",
					Usage: "insert data",
					Action: func(c *cli.Context) error {
						fmt.Println("insert subcommand")
						return nil
					},
				},
				{
					Name:  "delete",
					Usage: "delete data",
					Action: func(c *cli.Context) error {
						fmt.Println("delete subcommand")
						return nil
					},
				},
			},
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Println("BOOM!")
		return nil
	}
	app.Before = func(c *cli.Context) error {
		fmt.Println("app Before")
		return nil
	}
	app.After = func(c *cli.Context) error {
		fmt.Println("app After")
		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "Help!Help!",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "print-version, v",
		Usage: "print version",
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

}
