package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestParse1(t *testing.T) {
	fset := token.NewFileSet()
	ast0, err := parser.ParseFile(fset, "parser_input.go", nil,  parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}

	decl0 := ast0.Decls[0]
	if decl0 == nil {
		t.Fatal("decl0 is nil")
	}

	gen0, ok := decl0.(*ast.GenDecl)
	if gen0 == nil || !ok {
		t.Fatal("gen0, cast failed")
	}

	spec0 := gen0.Specs[0]
	imp0, ok := spec0.(*ast.ImportSpec)
	if imp0 == nil || !ok {
		t.Fatal("imp0, cast failed")
	}

	bl0 := imp0.Path
	if bl0 == nil {
		t.Fatal("bl0 is nil")
	}
	fmt.Printf("import: %s\n", bl0.Value)

	//if gen0.Tok == token.IMPORT {
	//	if gen0.Specs
	//}

	fmt.Printf("comments, len = %d\n", len(ast0.Comments))
	fmt.Println("done")
}
