package main

import (
	"golang-mongo-auth/cmd"
	server "golang-mongo-auth/internal/api"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/joho/godotenv"
)

func main() {
	if loadEnvErr := godotenv.Load(); loadEnvErr != nil {
		log.Println("Warning: No .env file found")
	}

	app := &cli.App{
		Name:     "cmd",
		Usage:    "Seed DB or run server",
		Commands: cmd.SeedCmd,
		Action: func(c *cli.Context) error {
			log.Println("Starting Gin server...")
			server.Start()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
