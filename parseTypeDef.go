package anococo

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/khicago/gococo"
)

type ParseResult map[string]TypeDefParseResult

type TypeDefParseResult struct {
	TypeDoc string
	Doc     string
	Cocos   []gococo.Coco
}

func (result TypeDefParseResult) String() string {
	return fmt.Sprintf(`-- TypeDoc ---
%s
-- Doc -------
%s
-- Cocos -----
%v
--------------`, result.TypeDoc, result.Doc, result.Cocos)
}

func (result ParseResult) tryRecordTypeSpec(typeSpec *ast.TypeSpec, typeDoc string) {
	doc := typeSpec.Doc.Text()
	cocos := make([]gococo.Coco, 0)
	if cmd, ok := gococo.Parse(doc); ok {
		cocos = append(cocos, cmd...)
	}
	if cmd, ok := gococo.Parse(typeDoc); ok {
		cocos = append(cocos, cmd...)
	}
	if len(cocos) <= 0 {
		return
	}
	typeName := typeSpec.Name.String()
	result[typeName] = TypeDefParseResult{
		TypeDoc: typeDoc,
		Doc:     doc,
		Cocos:   cocos,
	}
}

func handlerParseTypeDef(result ParseResult) InspectHandler {
	return func(pth string, node ast.Node) bool {
		var genDecl *ast.GenDecl
		var ok bool
		if genDecl, ok = node.(*ast.GenDecl); !ok {
			return true
		}
		if genDecl.Tok != token.TYPE {
			return true
		}

		for _, spec := range genDecl.Specs {
			switch s := spec.(type) {
			case *ast.TypeSpec:
				result.tryRecordTypeSpec(s, genDecl.Doc.Text())
			}
		}

		return true
	}
}

func ParseTypeDef(pth string) (ParseResult, error) {
	result := make(ParseResult)
	if err := inspectAst(pth, handlerParseTypeDef(result)); err != nil {
		return nil, err
	}
	return result, nil
}
