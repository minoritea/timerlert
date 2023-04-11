package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}
	app.Usage = "A simple timer command with alerts for MacOS"
	app.HideHelpCommand = true
	app.ArgsUsage = "message"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "title",
			Usage:   "set notification title",
			Aliases: []string{"t"},
		},
		&cli.StringFlag{
			Name:    "wait",
			Usage:   "waiting time until the alert is displayed",
			Aliases: []string{"w"},
		},
	}

	app.Action = func(cc *cli.Context) error {
		wait := cc.String("wait")
		var waitDuration time.Duration
		if wait != "" {
			var err error
			waitDuration, err = time.ParseDuration(wait)
			if err != nil {
				return err
			}
		}

		message := cc.Args().First()
		if message == "" {
			return fmt.Errorf("message should not be blank")
		}
		script := fmt.Sprintf("display notification %q", message)

		title := cc.String("title")
		if title != "" {
			script += fmt.Sprintf(" with title %q", title)
		}

		time.Sleep(waitDuration)

		return exec.Command("osascript", "-e", script).Run()
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
