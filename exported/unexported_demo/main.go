package main

import (
	"fmt"

	"github.com/golangdemos/exported"
)

func main() {
	sa := &unexported.StructA{}

	sb := sa.GetStructB()

	fmt.Printf("%v\n", sb)
}
