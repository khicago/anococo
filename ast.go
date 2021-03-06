package anococo

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"

	"github.com/khicago/gococo"
)

func MockCocoList() []gococo.Coco {
	return []gococo.Coco{
		gococo.NewCoco([]string{"MockCmd", "MockArg", "_"}),
		gococo.NewCoco([]string{"MockCmd2", "MockArg", "_"}),
	}
}

type (
	// InspectHandler
	InspectHandler func(pth string, p *ast.Package, node ast.Node) bool

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
			for _, pkg := range pkgLst {
				ast.Inspect(pkg, func(node ast.Node) bool { return handler(pth, pkg, node) })
			}
			return nil
		})
}

func inspectAst(pth string, handler InspectHandler) error {
	s := newSearcher(pth)
	return s.walk(handler)
}
