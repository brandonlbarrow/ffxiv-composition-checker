package main

import (
	"fmt"
	"github.com/brandonlbarrow/ffxiv-composition-checker/pkg/checker"
)

func main() {
	players := []checker.Player{
		{
			Name: "fred",
			Roles: []string{
				"PLD",
			},
		}, {
			Name: "joe",
			Roles: []string{
				"DRG",
			},
		}, {
			Name: "bob",
			Roles: []string{
				"AST",
			},
		}, {
			Name: "jen",
			Roles: []string{
				"DNC",
			},
		}, {
			Name: "kit",
			Roles: []string{
				"SCH",
			},
		}, {
			Name: "tom",
			Roles: []string{
				"DRK",
			},
		}, {
			Name: "jers",
			Roles: []string{
				"BLM",
			},
		}, {
			Name: "ben",
			Roles: []string{
				"MNK",
			},
		},
	}
	fp, err := checker.Allocate(players, checker.CompArgs{Format: checker.Any})
	if err != nil {
		panic(err)
	}
	fmt.Println(fp)
}
