package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
	"github.com/nlopes/slack"
)

func Client() *slack.Slack {
	return slack.New(os.Getenv("SLACK_API_TOKEN"))
}

func main() {
	c := cli.NewCLI("slacker", "0.0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"message": func() (cli.Command, error) {
			return &Message{}, nil
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
	// api := slack.New(os.Getenv("SLACK_API_TOKEN"))
	// param := slack.PostMessageParameters{}
	// param.Username = "shimabukuro"
	// api.PostMessage("hogehoge", "hogehoge", param)
	// groups, err := api.GetGroups(false)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, err.Error())
	// 	os.Exit(1)
	// }
	//
	// for _, group := range groups {
	// 	fmt.Printf("groupName: %s, id: %s \n", group.Name, group.Id)
	// }

}
