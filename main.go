package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
	"github.com/nlopes/slack"
)

func client() *slack.Slack {
	return slack.New(os.Getenv("SLACK_API_TOKEN"))
}

func main() {
	c := cli.NewCLI("slacker", "0.0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"message": func() (cli.Command, error) {
			return &Message{}, nil
		},
		"channel": func() (cli.Command, error) {
			return &Channel{}, nil
		},
		"group": func() (cli.Command, error) {
			return &Group{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}

	os.Exit(exitStatus)

}
