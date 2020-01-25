package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"github.com/wozz/oryx/internal/command/use"
)

func main() {
	app := NewApp()
	app.Run(os.Args)
}

// App embeds an urfave/cli.App
type App struct {
	*cli.App
}

// NewApp creates a new Application.
func NewApp() *App {

	log.SetPrefix("[oryx] ")

	// Create Cli inner app
	app := cli.NewApp()
	app.Usage = "useful utilities for working with bazel"
	app.Version = "dev"

	// Add commands
	app.Commands = []cli.Command{
		*use.Command,
	}

	instance := &App{
		App: app,
	}

	return instance
}
