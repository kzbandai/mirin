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

var definitions = []Definition{
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
		[]string{"components", "update"},
	},
	{
		"npm",
		[]string{"install", "-g", "npm@latest"},
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
	for _, d := range definitions {
		path, _ := exec.LookPath(d.Name)
		m.Commands = append(m.Commands, cli.Command{
			Name: d.Name,
			Action: func(c *cli.Context) error {
				cmd := exec.Command(path, getArgs(c.Command.Name)...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Start()

				return cmd.Wait()
			},
			Usage: getUsage(getArgs(d.Name)),
		})
	}
}

func getArgs(s string) []string {
	for _, d := range definitions {
		if d.Name == s {
			return d.Args
		}
	}

	return nil
}

func getUsage(ss []string) string {
	var r string
	for _, s := range ss {
		r += s + " "
	}
	return r
}
