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
	api := client()
	groups, err := api.GetGroups(false)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return 1
	}

	if len(groups) == 0 {
		return 0
	}

	fmt.Println("groupName,ID")

	for _, group := range groups {
		fmt.Printf("%s,%s\n", group.Name, group.Id)
	}
	return 0
}

func (g *Group) Synopsis() string {
	return "show group list"
}
