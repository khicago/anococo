package utils

import (
	"io"
	"os"
	"strings"
	"text/tabwriter"
	"text/template"

	"github.com/urfave/cli/v2"
)

func CreateCli(name, usage string, flags []cli.Flag, action cli.ActionFunc) *cli.App {

	// exit when help flag
	// @see github.com/urfave/cli/v2/help.go:260
	cli.HelpPrinter = func (out io.Writer, tpl string, data interface{}) {
		funcMap := template.FuncMap{
			"join": strings.Join,
		}
		w := tabwriter.NewWriter(out, 0, 8, 1, '\t', 0)
		t := template.Must(template.New("help").Funcs(funcMap).Parse(tpl))
		err := t.Execute(w, data)
		if err != nil {
			panic(err)
		}
		err = w.Flush()
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}

	return &cli.App{
		Name:            name,
		Usage:           usage,
		Flags:           flags,
		Action:          action,
		HideHelpCommand: true,
		HideVersion:     true,
	}
}