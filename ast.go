package anococo

import "github.com/khicago/gococo"

func MockCocoList() []gococo.Coco{
	return []gococo.Coco{
		gococo.NewCoco([]string{ "MockCmd", "MockArg", "_" }),
		gococo.NewCoco([]string{ "MockCmd2", "MockArg", "_" }),
	}
}