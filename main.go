package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/pflag"
)

func main() {
	var (
		title, wait string
		help        bool
	)
	pflag.StringVarP(&title, "title", "t", "", "set notification title")
	pflag.StringVarP(&wait, "wait", "w", "", "waiting time until the alert is displayed")
	pflag.BoolVarP(&help, "help", "h", false, "show help")
	pflag.Parse()
	message := pflag.Arg(0)

	if help {
		fmt.Fprintln(os.Stderr, `NAME:`)
		fmt.Fprintln(os.Stderr, `   timerlert - A simple timer command with alerts for MacOS`)
		fmt.Fprintln(os.Stderr, ``)
		fmt.Fprintln(os.Stderr, `USAGE:`)
		fmt.Fprintln(os.Stderr, `   timerlert [global options] message`)
		fmt.Fprintln(os.Stderr, ``)
		fmt.Fprintln(os.Stderr, `OPTIONS:`)
		pflag.PrintDefaults()
		return
	}

	err := run(message, title, wait)
	if err != nil {
		log.Fatal(err)
	}
}

func run(message, title, wait string) error {
	var waitDuration time.Duration
	if wait != "" {
		var err error
		waitDuration, err = time.ParseDuration(wait)
		if err != nil {
			return err
		}
	}

	if message == "" {
		return fmt.Errorf("message should not be blank")
	}
	script := fmt.Sprintf("display notification %q", message)

	if title != "" {
		script += fmt.Sprintf(" with title %q", title)
	}

	time.Sleep(waitDuration)

	return exec.Command("osascript", "-e", script).Run()
}
