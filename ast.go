package anococo

import (
	"github.com/khicago/gococo"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

func MockCocoList() []gococo.Coco {
	return []gococo.Coco{
		gococo.NewCoco([]string{"MockCmd", "MockArg", "_"}),
		gococo.NewCoco([]string{"MockCmd2", "MockArg", "_"}),
	}
}

type (
	// InspectHandler
	InspectHandler func(pth string, node ast.Node) bool

	astInspector struct {
		fSet *token.FileSet
		root string
	}
)

func newSearcher(pth string) *astInspector {
	return &astInspector{
		fSet: token.NewFileSet(),
		root: pth,
	}
}

func (s *astInspector) walk(handler InspectHandler) error {
	return filepath.Walk(
		s.root,
		func(pth string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				return nil
			}
			pkgLst, err := parser.ParseDir(s.fSet, pth, nil, parser.ParseComments)
			if err != nil {
				return err
			}
			for _, v := range pkgLst {
				ast.Inspect(v, func(node ast.Node) bool { return handler(pth, node) })
			}
			return nil
		})
}

func inspectAst(pth string, handler InspectHandler) error {
	s := newSearcher(pth)
	return s.walk(handler)
}
