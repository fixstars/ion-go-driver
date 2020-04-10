package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	ion "gitlab.fixstars.com/ion/ion-go"
)

func main() {

	app := &cli.App{
		Name:  "ion-go-driver",
		Usage: "The driver interface for ion-go",
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:  "with-bb-module",
				Usage: "Specify path to building block list.",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "compile",
				Usage: "Compile ion graph into static library and header.",
				Action: func(c *cli.Context) error {
					graph := c.Path("graph")
					modules := c.StringSlice("with-bb-module")
					target := c.String("target")
					funcName := c.String("func-name")

					b, err := ion.NewBuilder()
					if err != nil {
						return err
					}

					if err = b.SetTarget(target); err != nil {
						return err
					}

					var f *os.File
					if f, err = os.Open(graph); err != nil {
						return err
					}
					defer f.Close()

					if err = b.LoadFromReader(f); err != nil {
						return err
					}

					for _, m := range modules {
						if err = b.WithBBModule(m); err != nil {
							return err
						}
					}

					if err = b.Compile(funcName); err != nil {
						return err
					}
					return nil
				},
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:     "graph",
						Usage:    "Specify graph path to be built.",
						Required: true,
					},
					&cli.StringFlag{
						Name:        "target",
						Usage:       "Specify target string for builder.",
						DefaultText: "host",
					},
					&cli.StringFlag{
						Name:     "func-name",
						Usage:    "Specify func name to be built.",
						Required: true,
					},
				},
			}, {
				Name:  "metadata",
				Usage: "Retrieve metadata from building block.",
				Action: func(c *cli.Context) error {
					modules := c.StringSlice("with-bb-module")

					b, err := ion.NewBuilder()
					if err != nil {
						return err
					}

					for _, m := range modules {
						if err = b.WithBBModule(m); err != nil {
							return err
						}
					}

					var md string
					if md, err = b.BBMetadata(); err != nil {
						return err
					}

					fmt.Printf("%v", md)

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
