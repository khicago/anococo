package main

import (
	"fmt"

	lib "github.com/khicago/anococo"
	"github.com/urfave/cli/v2"
)

var cmdTypeDef = cli.Command{
	Name:     "typedef",
	Aliases:  []string{"t"},
	Usage:    "scan coco cmds from type annotations",
	Category: "Type",
	Action:   actionTypeDef,
}

func actionTypeDef(ctx *cli.Context) error {
	pth := "."
	if ctx.NArg() > 0 {
		pth = ctx.Args().Get(0)
	}
	result, _ := lib.ParseTypeDef(pth)
	for k, v := range result {
		fmt.Printf("[%s] => \n%s\n\n", k, v)
	}
	return nil
}
