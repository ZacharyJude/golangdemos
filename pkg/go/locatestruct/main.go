// comment associated to package definition
package main

// comment associated to these import statements
import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"regexp"
)

// location for SampleStruct0
type SampleStruct0 struct {
	Field0 string `json:"field_0" sql:"col_field_0"`
	Field1 int    `json:"field_1"`
	field2 []byte
}

// comment associated to the main function
// some more comment in the next line
func main() {

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

	ast.Walk(TryVisitor{}, nodeForThisFile)
}

type TryVisitor struct {
	allowPrint bool
}

func (v TryVisitor) Visit(node ast.Node) (w ast.Visitor) {
	if nil == node {
		return v
	}

	fmt.Printf("%s\n", &v)

	if commentNode, ok := node.(*ast.Comment); ok {
		if matched, _ := regexp.MatchString("//\\s*location\\s", commentNode.Text); matched {
			fmt.Printf(
				"I got the comment:%s\n",
				commentNode.Text,
			)

			v.allowPrint = true
		}
	}
	if v.allowPrint {
		fmt.Printf(
			"pos:%d,end:%d,type:%s\n",
			node.Pos(),
			node.End(),
			reflect.TypeOf(node),
		)
	}
	return &v
}
