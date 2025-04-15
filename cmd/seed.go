package cmd

import (
	"golang-mongo-auth/pkg/common/database"
	"golang-mongo-auth/pkg/config"
	"golang-mongo-auth/pkg/seeder"

	"github.com/urfave/cli/v2"
)

var SeedCmd = []*cli.Command{
	{
		Name:  "seed",
		Usage: "Seed database",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "dry-run"},
			&cli.BoolFlag{Name: "reset", Value: false},
		},
		ArgsUsage: "[all|rbac]",
		Action: func(c *cli.Context) error {
			if c.Args().Len() < 1 {
				return cli.Exit("Specify what to seed: all, rbac", 1)
			}

			db := database.Init(config.GetMongoURI(), config.GetDbName())

			isDryRun := c.Bool("dry-run")
			shouldReset := c.Bool("reset")

			switch c.Args().Get(0) {
			case "rbac":
				seeder.SeedRbac(db, isDryRun, shouldReset)
			default:
				return cli.Exit("Invalid option", 1)
			}
			return nil
		},
	},
}
