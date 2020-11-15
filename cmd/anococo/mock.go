package main

import (
	"fmt"

	lib "github.com/khicago/anococo"
	"github.com/urfave/cli/v2"
)

var cmdMock = cli.Command{
	Name:     "mock",
	Aliases:  []string{"m"},
	Usage:    "return mocked coco cmds",
	Category: "Func",
	Action:   actionMock,
}

func actionMock(ctx *cli.Context) error {
	fmt.Println(lib.MockCocoList())
	return nil
}
