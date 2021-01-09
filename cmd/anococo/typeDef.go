package main

import (
	"fmt"

	"github.com/bagaking/easycmd"

	lib "github.com/khicago/anococo"
	"github.com/urfave/cli/v2"
)

var cmdTypeDef = easycmd.New("typedef").Set.
	Alias("t").Usage("scan coco cmds from type annotations").End.Action(actionTypeDef)

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
