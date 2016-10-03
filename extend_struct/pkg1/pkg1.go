package pkg1

import "github.com/zacharyjude/golangdemos/extend_struct/pkg0"

type I1 interface {
	Bar() string
}

type SS0 pkg0.S0

func (ss0 *SS0) Bar() string {
	return (*pkg0.S0)(ss0).Foo() + " " + "bar in ss0"
}
