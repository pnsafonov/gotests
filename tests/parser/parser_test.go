package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

const (
	all = parser.PackageClauseOnly | parser.ImportsOnly | parser.ParseComments | parser.Trace |
		parser.DeclarationErrors | parser.SpuriousErrors | parser.AllErrors
)

func parseAst0(t *testing.T) *ast.File {
	fset := token.NewFileSet()
	ast0, err := parser.ParseFile(fset, "parser_input.go", nil,  parser.ParseComments)
	//ast0, err := parser.ParseFile(fset, "parser_input.go", nil,  all)
	if err != nil {
		t.Fatal(err)
	}
	return ast0
}

func TestParse1(t *testing.T) {
	ast0 := parseAst0(t)

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

func TestParse2(t *testing.T) {
	ast0 := parseAst0(t)

	for _, decl := range ast0.Decls {
		st, ok := getStructDecl(decl)
		if !ok {
			continue
		}

		fmt.Printf("incomplete = %v\n", st.Incomplete)
	}

	fmt.Println("done")
}

func getStructDecl(decl ast.Decl) (st *ast.StructType, ok bool) {
	gd, ok := decl.(*ast.GenDecl)
	if !ok {
		return nil, false
	}

	if len(gd.Specs) == 0 {
		return nil, false
	}
	sp := gd.Specs[0]

	ts, ok := sp.(*ast.TypeSpec)
	if !ok {
		return nil, false
	}

	st, ok = ts.Type.(*ast.StructType)
	return
}


