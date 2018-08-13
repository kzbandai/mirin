package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"fmt"

	"github.com/urfave/cli"
)

var (
	author = "kzbandai"

	homebrew        = "brew"
	homebrewDoctor  = "doctor"
	homebrewUpgrade = "upgrade"
)

func main() {
	app := cli.NewApp()
	setInfo(app)
	setCommands(app)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func setInfo(a *cli.App) {
	a.Author = author
}

func setCommands(a *cli.App) {
	a.Commands = []cli.Command{
		{
			Name: homebrew,
			Action: func(c *cli.Context) error {
				path, err := exec.LookPath(homebrew)
				if err != nil {
					log.Fatal(err)
				}

				out, err := exec.Command(path, homebrewUpgrade).Output()
				if err != nil {
					fmt.Println("error!")
				}

				fmt.Printf("%s", out)

				return err
			},
			Subcommands: []cli.Command{
				{
					Name: "doctor",
					Action: func(c *cli.Context) error {
						path, err := exec.LookPath(homebrew)
						if err != nil {
							log.Fatal(err)
						}

						cmd := exec.Command(path, homebrewDoctor)
						stderr, err := cmd.StderrPipe()
						if err != nil {
							log.Fatal(err)
						}

						if err = cmd.Start(); err != nil {
							log.Fatal(err)
						}

						slurp, _ := ioutil.ReadAll(stderr)
						if slurp != nil {
							fmt.Printf("%s", slurp)
						}

						if err := cmd.Wait(); err != nil {
							log.Fatal(err)
						}

						return nil
					},
				},
			},
		},
	}
}
