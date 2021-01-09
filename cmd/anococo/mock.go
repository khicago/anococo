package main

import (
	"fmt"

	"github.com/bagaking/easycmd"

	lib "github.com/khicago/anococo"
	"github.com/urfave/cli/v2"
)

var cmdMock = easycmd.New("mock").Set.
	Alias("m").Usage("return mocked coco cmds").End.Action(func(*cli.Context) error {
	fmt.Println(lib.MockCocoList())
	return nil
})
