package main

import (
	"fmt"
	"os"

	"github.com/bagaking/easycmd"
	"github.com/urfave/cli/v2"
)

var App *cli.App

func makeApp(name string) *cli.App {
	easycmd.SetCustomOptions(easycmd.CustomOption{
		ExitAfterPrintHelpMsg: true,
	})

	return &cli.App{
		Name: name,
		Commands: []*cli.Command{
			cmdMock.Root(), cmdTypeDef.Root(),
		},
		HideVersion: true,
	}
}

func init() {
	App = makeApp("anococo")
}

func main() {
	fmt.Println("execute the anococo program")

	if err := App.Run(os.Args); err != nil {
		panic(err)
	}
}
