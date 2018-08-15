package commands

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/urfave/cli"
)

var (
	homebrew        = "brew"
	homebrewDoctor  = "doctor"
	homebrewUpgrade = "upgrade"
)

func GetHomebrew() cli.Command {
	path, err := exec.LookPath(homebrew)
	if err != nil {
		log.Fatal(err)
	}

	return cli.Command{
		Name: homebrew,
		Action: func(c *cli.Context) error {
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
	}
}
