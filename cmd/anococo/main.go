package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"os"
)

var App *cli.App

func makeApp(name string) *cli.App {


	cli.HelpPrinter = func(out io.Writer, tpl string, data interface{}) {
		cli.HelpPrinterCustom(out, tpl, data, nil)
		os.Exit(0)
	}

	return &cli.App{
		Name:            name,
		Commands:        makeCommands(),
		HideHelpCommand: true,
		HideVersion:     true,
	}

}

func makeCommands() []*cli.Command {
	return []*cli.Command {
		&cmdMock,
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
