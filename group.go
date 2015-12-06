package main

import (
	"fmt"
	"os"
)

type Group struct{}

func (g *Group) Help() string {
	return "hogehoge"
}

func (g *Group) Run(args []string) int {
	api := Client()
	groups, err := api.GetGroups(false)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return 1
	}

	for _, group := range groups {
		fmt.Printf("groupName: %s, id: %s \n", group.Name, group.Id)
	}
	return 0
}

func (g *Group) Synopsis() string {
	return "hogehoge"
}
