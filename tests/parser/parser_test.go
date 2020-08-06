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

		fmt.Printf("name = %v\n", st.Name)
	}

	fmt.Println("done")
}

func getStructDecl(decl ast.Decl) (result GenStruct, ok bool) {
	gd, ok := decl.(*ast.GenDecl)
	if !ok {
		return GenStruct{}, false
	}

	if len(gd.Specs) == 0 {
		return GenStruct{}, false
	}
	sp := gd.Specs[0]

	ts, ok := sp.(*ast.TypeSpec)
	if !ok {
		return GenStruct{}, false
	}

	st, ok := ts.Type.(*ast.StructType)

	result.StructType = st
	result.Decl = decl
	result.TypeSpec = ts
	if ts.Name != nil {
		result.Name = ts.Name.Name
	}
	return
}

type GenStruct struct {
	// ast fields
	Decl 		ast.Decl
	StructType  *ast.StructType
	TypeSpec 	*ast.TypeSpec

	// name of struct
	Name 		string
	// skip serialization
	Skip 		bool
}

type GenField struct {

}

func ParseFile(goFile string) ([]GenStruct, error) {
	ast0, err := parser.ParseFile(token.NewFileSet(), goFile, nil,  parser.ParseComments)
	if err != nil {
		return nil, err
	}

	result := make([]GenStruct, 0, 4)
	for _, decl := range ast0.Decls {
		gs, ok := getStructDecl(decl)
		if !ok {
			continue
		}

		result = append(result, gs)
	}

	return result, nil
}

func TestParseFile1(t *testing.T) {
	result, err := ParseFile("parser_input.go")
	if err != nil {
		t.Fatal(err)
	}
	if len(result) == 0 {
		t.Fatal("len is 0")
	}
}