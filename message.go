package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/nlopes/slack"
)

type Message struct {
	client *slack.Slack
}

func (m *Message) Help() string {
	return "oshirase message -c <channelID> -u <username> <message>"
}

func (m *Message) Run(args []string) int {
	var (
		channelID string
		user      string
	)
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	f.StringVar(&channelID, "c", "", "channelID")
	f.StringVar(&user, "u", appName(), "userName")
	f.Parse(args)

	msg := f.Arg(0)
	if msg == "" {
		buf, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return 1
		}
		if len(buf) == 0 {
			fmt.Fprintln(os.Stderr, "error: please input message")
			return 1
		}
		msg = string(buf)
	}

	if channelID == "" {
		fmt.Fprintln(os.Stderr, "error: channelID is required")
		return 1
	}
	api := client()
	param := slack.PostMessageParameters{}
	param.Username = user
	_, _, err := api.PostMessage(channelID, msg, param)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return 1
	}

	return 0
}

func (m *Message) Synopsis() string {
	return "post message to selected channel"
}
