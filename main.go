package main

import (
	"log"
	"os"

	"github.com/kzbandai/mirin/commands"
	"github.com/urfave/cli"
)

var (
	author = "kzbandai"
)

type Mirin struct {
	*cli.App
}

func main() {
	app := newMirin()

	app.setInfo()
	app.setCommands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func newMirin() *Mirin {
	return &Mirin{cli.NewApp()}
}

func (m *Mirin) setInfo() {
	m.Author = author
}

func (m *Mirin) setCommands() {
	m.Commands = []cli.Command{
		commands.GetHomebrew(),
	}
}
