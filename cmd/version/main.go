package main

import (
	"fmt"
	"log"
	"os"

	cli "github.com/urfave/cli/v2"

	npm "github.com/khulnasoft-lab/go-npm-version/pkg"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "compare",
				Aliases: []string{"c"},
				Usage:   "compare two versions",
				Action: func(c *cli.Context) error {
					s1 := c.Args().Get(0)
					v1, err := npm.NewVersion(s1)
					if err != nil {
						log.Fatalf("failed to parse npm version (%s): %s", s1, err)
					}

					s2 := c.Args().Get(1)
					v2, err := npm.NewVersion(s2)
					if err != nil {
						log.Fatalf("failed to parse npm version (%s): %s", s2, err)
					}

					fmt.Println(v1.Compare(v2))
					return nil
				},
			},
			{
				Name:    "satisfy",
				Aliases: []string{"s"},
				Usage:   "check if the version satisfies the constraint",
				Action: func(c *cli.Context) error {
					s1 := c.Args().Get(0)
					v, err := npm.NewVersion(s1)
					if err != nil {
						log.Fatalf("failed to parse npm version (%s): %s", s1, err)
					}

					s2 := c.Args().Get(1)
					constraint, err := npm.NewConstraints(s2)
					if err != nil {
						log.Fatalf("failed to parse npm constraint (%s): %s", s2, err)
					}

					fmt.Println(constraint.Check(v))
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
