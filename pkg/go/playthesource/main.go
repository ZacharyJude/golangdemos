// comment associated to package definition
package main

// comment associated to these import statements
import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
)

type SampleStruct0 struct {
	Field0 string `json:"field_0" sql:"col_field_0"`
	Field1 int    `json:"field_1"`
	field2 []byte
}

// comment associated to the main function
// some more comment in the next line
func main() {

	out := "should I print something out?"
	fmt.Println(out)

	fset := token.NewFileSet() // I don't know what is this for now
	nodeForThisFile, err := parser.ParseFile(
		fset,
		"main.go",
		nil,
		parser.ParseComments,
	)

	if nil != err {
		fmt.Println(err)
	}

	fmt.Printf("list comments here start --------------------\n")
	for _, commentGroup := range nodeForThisFile.Comments {
		fmt.Printf(
			"pos:%d\tend:%d\t%s\n",
			commentGroup.Pos(),
			commentGroup.End(),
			commentGroup.Text(),
		)
		/*
			for comment := range commentGroup.List {
				fmt.Println(reflect.TypeOf(comment))
			}
		*/
	}
	fmt.Printf("\n\n\nlist decls here start -----------------\n")

	for _, decl := range nodeForThisFile.Decls {
		fmt.Printf(
			"pos:%d\tend:%d  %s\n",
			decl.Pos(),
			decl.End(),
			reflect.TypeOf(decl),
		)
		visitDecl(decl, "")
		fmt.Printf("\n")
	}

	fmt.Printf("\n\n\nwalk start ----------------------------\n")

	ast.Walk(TryVisitor{}, nodeForThisFile)
}

func visitImportSpec(specImpl *ast.ImportSpec, tabPrefix string) {
	fmt.Printf("%s%s\n", tabPrefix, specImpl.Path.Value)
}

func visitGenDecl(genDecl *ast.GenDecl, tabPrefix string) {
	fmt.Printf("%s%s\n", tabPrefix, genDecl.Tok)
	for _, spec := range genDecl.Specs {
		visitSpec(spec, tabPrefix+"  ")
	}
}

func visitTag(in *ast.BasicLit, tabPrefix string) {
	if nil != in {
		fmt.Printf("%s%s\n", tabPrefix, in.Value)
	} else {
		fmt.Printf("%sno tag\n", tabPrefix)
	}
}

func visitField(in *ast.Field, tabPrefix string) {
	fmt.Printf("%s%s\n", tabPrefix, in.Names[0].Name)
	visitTag(in.Tag, tabPrefix+"  ")
	visitExpr(in.Type, tabPrefix+"  ")
}

func visitStructType(in *ast.StructType, tabPrefix string) {
	for _, field := range in.Fields.List {
		visitField(field, tabPrefix)
	}
}

func visitTypeSpec(specImpl *ast.TypeSpec, tabPrefix string) {
	fmt.Printf("%s%s\n", tabPrefix, specImpl.Name.Name)
	visitExpr(specImpl.Type, tabPrefix+"  ")
}

func visitArrayType(in *ast.ArrayType, tabPrefix string) {
	fmt.Printf("%s%s\n", tabPrefix, "array element")
	visitExpr(in.Elt, tabPrefix+"  ")
}

func visitIdent(in *ast.Ident, tabPrefix string) {
	if nil != in.Obj {
		fmt.Printf("%s%s,%s\n", tabPrefix, in.Obj.Name, in.Name)
	} else {
		fmt.Printf("%s%s\n", tabPrefix, in.Name)
	}
}

func visitExpr(expr ast.Expr, tabPrefix string) {
	switch exprImpl := expr.(type) {
	case *ast.Ident:
		visitIdent(exprImpl, tabPrefix)
	case *ast.StructType:
		visitStructType(exprImpl, tabPrefix)
	case *ast.ArrayType:
		visitArrayType(exprImpl, tabPrefix)
	default:
		fmt.Printf("%sunknown:%s\n", tabPrefix, reflect.TypeOf(exprImpl))
	}
}

func visitSpec(spec ast.Spec, tabPrefix string) {
	switch specImpl := spec.(type) {
	case *ast.ImportSpec:
		visitImportSpec(specImpl, tabPrefix)
	case *ast.TypeSpec:
		visitTypeSpec(specImpl, tabPrefix)
	default:
		fmt.Printf("%sunknown:%s\n", tabPrefix, reflect.TypeOf(specImpl))
	}
}

func visitDecl(decl ast.Decl, tabPrefix string) {
	switch declImpl := decl.(type) {
	case *ast.GenDecl:
		visitGenDecl(declImpl, tabPrefix)
	default:
		fmt.Printf("%sunknown decl:%s", tabPrefix, reflect.TypeOf(declImpl))
	}
}

type TryVisitor struct {
}

func (v TryVisitor) Visit(node ast.Node) (w ast.Visitor) {
	if nil != node {
		fmt.Printf(
			"pos:%d,end:%d,type:%s\n",
			node.Pos(),
			node.End(),
			reflect.TypeOf(node),
		)
	}
	return v
}
