package main

import "fmt"
import "reflect"

import "github.com/zacharyjude/golangdemos/extend_struct/pkg0"
import "github.com/zacharyjude/golangdemos/extend_struct/pkg1"

func main() {

	s0 := &pkg0.S0{}
	s0.F0 = "kakakakaka"

	ss0PtrTyp := reflect.TypeOf((*pkg1.SS0)(nil))

	i1 := reflect.ValueOf(s0).Convert(ss0PtrTyp).Interface().(pkg1.I1)

	fmt.Printf("%s\n", (*pkg1.SS0)(s0).Bar())
	fmt.Printf("%s\n", i1.Bar())
}
