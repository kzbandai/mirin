package main

import (
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

var (
	author      = "kzbandai"
	copyright   = "MIT"
	description = ""
	hideHelp    = true
	name        = "mirin"
	usage       = "Updates manager by Go"
	version     = "0.0.1"
)

var s []cli.Command

var commands = []Definition{
	{
		"anyenv",
		[]string{"update"},
	},
	{
		"brew",
		[]string{"upgrade"},
	},
	{
		"composer",
		[]string{"self-update"},
	},
	{
		"gcloud",
		[]string{"container", "update"},
	},
	{
		"npm",
		[]string{"install", "-g", "npm@latest"},
	},
	{
		"yarn",
		[]string{"self-update"},
	},
}

type Definition struct {
	Name string
	Args []string
}

type Mirin struct {
	*cli.App
}

func main() {
	app := newMirin()

	app.setInfo()
	app.setCommands()

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}

func newMirin() *Mirin {
	return &Mirin{cli.NewApp()}
}

func (m *Mirin) setInfo() {
	m.Author = author
	m.Copyright = copyright
	m.Description = description
	m.HideHelp = hideHelp
	m.Name = name
	m.Usage = usage
	m.Version = version
}

func (m *Mirin) setCommands() {
	for _, d := range commands {
		path, err := exec.LookPath(d.Name)
		if err != nil {

		}

		s = append(s, cli.Command{
			Name: d.Name,
			Action: func(c *cli.Context) error {
				cmd := exec.Command(path, d.Args...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Start()

				return cmd.Wait()
			},
		})
	}

	m.Commands = s
}
