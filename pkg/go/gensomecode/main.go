package main

//go:generate higen arg...

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"regexp"
	//	"reflect"
)

// taghere
// hihi
// what?
type SampleStruct0 struct {
	Field0 string `json:"field_0" sql:"col_field_0"`
	Field1 int    `json:"field_1"`
	field2 []byte
}

// taghere
type SampleStruct1 struct {
	Field3 string `json:"field_3" sql:"col_field_3"`
	Field4 int    `json:"field_4"`
	field5 []byte
}

type x struct{}

func genDeclWithTagInDoc(tag string, node ast.Node) []*ast.GenDecl {
	out := make([]*ast.GenDecl, 0)
	ast.Inspect(node, func(n ast.Node) bool {

		if nil == n {
			return false
		}

		switch nodeImpl := n.(type) {
		case *ast.GenDecl:

			if token.TYPE != nodeImpl.Tok {
				return false
			}

			if 1 != len(nodeImpl.Specs) {
				return false
			}

			pattern := fmt.Sprintf("^\\s*%s\\s*", tag)

			matched, _ := regexp.MatchString(
				pattern,
				nodeImpl.Doc.Text(),
			)

			if !matched {
				return false
			}

			out = append(out, nodeImpl)
		}
		return true
	})

	return out

}

func Gen0() {

	gen0 := `
package main

func SaySomething() {
	fmt.Printf("%s\n", "hehehehe!!")
}
`

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "gen0.go", gen0, 0)
	if nil != err {
		panic(err)
	}

	ast.Print(fset, f)
	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Printf("%s", buf.Bytes())
}

func Gen1() {

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", nil, parser.ParseComments)
	if nil != err {
		panic(err)
	}

	ast.Print(fset, f)
}

func Gen2() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", nil, parser.ParseComments)
	if nil != err {
		panic(err)
	}

	genDecls := genDeclWithTagInDoc("taghere", f)
	for _, x := range genDecls {
		y, _ := x.Specs[0].(*ast.TypeSpec)
		fmt.Printf("%s\n", y.Name.Name)
	}
}

func main() {
	Gen1()
}
