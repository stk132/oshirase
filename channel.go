package main

import (
	"fmt"
	"os"
)

type Channel struct{}

func (c *Channel) Help() string {
	return "slacker channel"
}

func (c *Channel) Run(args []string) int {
	api := client()
	channels, err := api.GetChannels(false)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}

	if len(channels) != 0 {
		fmt.Println("channelName,ID")
		for _, channel := range channels {
			fmt.Printf("%s,%s\n", channel.Name, channel.Id)
		}
	}
	return 0
}

func (c *Channel) Synopsis() string {
	return "show channelName channelID"
}
