package anococo

import (
	"fmt"
	"go/ast"

	"github.com/khicago/gococo"
)

type ParseResult struct {
	PackageName string
	Import      map[string]GenDeclParseResult
	Type        map[string]GenDeclParseResult
	Const       map[string]GenDeclParseResult
	Value       map[string]GenDeclParseResult
}

type GenDeclParseResult struct {
	DeclComment string
	Doc         string
	Cocos       []gococo.Coco
}

func (result *ParseResult) String() string {
	return fmt.Sprintf(`%s{
  gd.Import: %v
  gd.Type:   %v
  gd.Const:  %v
  gd.Value:  %v
}`, result.PackageName, result.Import, result.Type, result.Const, result.Value)
}

func (result GenDeclParseResult) String() string {
	return fmt.Sprintf("cocos%v", result.Cocos)
}

func parseCocoInAllDocs(docs ...*ast.CommentGroup) []gococo.Coco {
	cocos := make([]gococo.Coco, 0)
	for _, doc := range docs {
		if doc == nil {
			continue
		}
		for _, l := range doc.List {
			if cmd, ok := gococo.Parse(l.Text); ok {
				cocos = append(cocos, cmd...)
			}
		}
	}
	return cocos
}

func (result *ParseResult) tryRecordValueSpec(valueSpec *ast.ValueSpec, declComment *ast.CommentGroup) {
	cocos := parseCocoInAllDocs(valueSpec.Doc, declComment)
	if len(cocos) <= 0 {
		return
	}
	vd := GenDeclParseResult{
		DeclComment: declComment.Text(),
		Doc:         valueSpec.Doc.Text(),
		Cocos:       cocos,
	}
	for _, valueName := range valueSpec.Names {
		result.Value[valueName.Name] = vd
	}
}

func (result *ParseResult) tryRecordConstSpec(valueSpec *ast.ValueSpec, declComment *ast.CommentGroup) {
	cocos := parseCocoInAllDocs(valueSpec.Doc, declComment)
	if len(cocos) <= 0 {
		return
	}
	vd := GenDeclParseResult{
		DeclComment: declComment.Text(),
		Doc:         valueSpec.Doc.Text(),
		Cocos:       cocos,
	}
	for _, valueName := range valueSpec.Names {
		result.Const[valueName.Name] = vd
	}
}

func (result *ParseResult) tryRecordTypeSpec(typeSpec *ast.TypeSpec, declComment *ast.CommentGroup) {
	if typeSpec.Name == nil {
		return
	}
	cocos := parseCocoInAllDocs(typeSpec.Doc, declComment)
	if len(cocos) <= 0 {
		return
	}
	vd := GenDeclParseResult{
		DeclComment: declComment.Text(),
		Doc:         typeSpec.Doc.Text(),
		Cocos:       cocos,
	}
	result.Type[typeSpec.Name.Name] = vd
}

func (result *ParseResult) tryRecordImportSpec(importSpec *ast.ImportSpec, declComment *ast.CommentGroup) {
	if importSpec.Name == nil {
		return
	}
	cocos := parseCocoInAllDocs(importSpec.Doc, declComment)
	if len(cocos) <= 0 {
		return
	}
	vd := GenDeclParseResult{
		DeclComment: declComment.Text(),
		Doc:         importSpec.Doc.Text(),
		Cocos:       cocos,
	}
	result.Import[importSpec.Name.Name] = vd
}

func handlerParseDef(resultMap map[string]*ParseResult) InspectHandler {
	return func(pth string, p *ast.Package, node ast.Node) bool {
		if _, ok := resultMap[pth]; !ok {
			resultMap[pth] = NewResult()
		}
		result := resultMap[pth]

		var genDecl *ast.GenDecl
		var ok bool
		if genDecl, ok = node.(*ast.GenDecl); !ok {
			return true
		}
		result.PackageName = p.Name
		ForEachImportSpec(genDecl, result.tryRecordImportSpec)
		ForEachTypeSpec(genDecl, result.tryRecordTypeSpec)
		ForEachVarSpec(genDecl, result.tryRecordValueSpec)
		ForEachConstSpec(genDecl, result.tryRecordConstSpec)

		return true
	}
}

func NewResult() *ParseResult {
	return &ParseResult{
		Import: make(map[string]GenDeclParseResult),
		Type:   make(map[string]GenDeclParseResult),
		Const:  make(map[string]GenDeclParseResult),
		Value:  make(map[string]GenDeclParseResult),
	}
}

func ParseTypeDef(pth string) (map[string]*ParseResult, error) {
	resultMap := make(map[string]*ParseResult)
	if err := inspectAst(pth, handlerParseDef(resultMap)); err != nil {
		return nil, err
	}
	return resultMap, nil
}
