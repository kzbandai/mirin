package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/BurntSushi/toml"
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

type Definitions struct {
	Definitions []Definition `toml:"definition"`
}

type Definition struct {
	Name string   `toml:"name"`
	Args []string `toml:"args"`
}

type Mirin struct {
	*cli.App
	Definitions []Definition
}

func main() {
	app := newMirin()

	app.setInfo()
	app.loadDefinitions()
	app.setDefinitions()

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}

func newMirin() *Mirin {
	return &Mirin{
		App: cli.NewApp(),
	}
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

func (m *Mirin) loadDefinitions() {
	var d Definitions
	toml.DecodeFile("definitions.toml", &d)

	m.Definitions = d.Definitions
}

func (m *Mirin) setDefinitions() {
	for _, d := range m.Definitions {
		path, _ := exec.LookPath(d.Name)
		m.Commands = append(m.Commands, cli.Command{
			Name: d.Name,
			Action: func(c *cli.Context) error {
				cmd := exec.Command(path, getArgs(c)...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Start()

				return cmd.Wait()
			},
			Usage: getUsage(d.Args),
		})
	}
}

func getArgs(c *cli.Context) []string {
	return strings.Split(c.Command.Usage, " ")
}

func getUsage(ss []string) string {
	return strings.Join(ss, " ")
}
