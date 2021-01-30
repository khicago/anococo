package anococo

import (
	"go/ast"
	"go/token"
)

// ForEachImportSpec performs the operation fn on all sepc of
// a genDecl of type `token.IMPORT`.
// If the type of genDecl is not `token.IMPORT`, nothing will happen.
// For each fn execution, the Comment of genDecl will also be
// passed in.
// @see https://golang.org/pkg/go/ast/#GenDecl
func ForEachImportSpec(genDecl *ast.GenDecl, fn func(spec *ast.ImportSpec, declComment *ast.CommentGroup)) {
	if genDecl.Tok != token.IMPORT {
		return // skip
	}
	for _, spec := range genDecl.Specs {
		s := spec.(*ast.ImportSpec)
		fn(s, genDecl.Doc)
	}
}

// ForEachConstSpec performs the operation fn on all sepc of
// a genDecl of type `token.CONST`.
// If the type of genDecl is not `token.CONST`, nothing will happen.
// For each fn execution, the Comment of genDecl will also be
// passed in.
// @see https://golang.org/pkg/go/ast/#GenDecl
func ForEachConstSpec(genDecl *ast.GenDecl, fn func(spec *ast.ValueSpec, declComment *ast.CommentGroup)) {
	if genDecl.Tok != token.CONST {
		return // skip
	}
	for _, spec := range genDecl.Specs {
		s := spec.(*ast.ValueSpec)
		fn(s, genDecl.Doc)
	}
}

// ForEachVarSpec performs the operation fn on all sepc of
// a genDecl of type `token.VAR`.
// If the type of genDecl is not `token.VAR`, nothing will happen.
// For each fn execution, the Comment of genDecl will also be
// passed in.
// @see https://golang.org/pkg/go/ast/#GenDecl
func ForEachVarSpec(genDecl *ast.GenDecl, fn func(spec *ast.ValueSpec, declComment *ast.CommentGroup)) {
	if genDecl.Tok != token.VAR {
		return // skip
	}
	for _, spec := range genDecl.Specs {
		s := spec.(*ast.ValueSpec)
		fn(s, genDecl.Doc)
	}
}

// ForEachTypeSpec performs the operation fn on all sepc of
// a genDecl of type `token.TYPE`.
// If the type of genDecl is not `token.TYPE`, nothing will happen.
// For each fn execution, the Comment of genDecl will also be
// passed in.
// @see https://golang.org/pkg/go/ast/#GenDecl
func ForEachTypeSpec(genDecl *ast.GenDecl, fn func(spec *ast.TypeSpec, declComment *ast.CommentGroup)) {
	if genDecl.Tok != token.TYPE {
		return // skip
	}
	for _, spec := range genDecl.Specs {
		s := spec.(*ast.TypeSpec)
		fn(s, genDecl.Doc)
	}
}
