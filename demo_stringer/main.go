package main

import (
	"fmt"
)

const _enumA_name = "v0v1v2v3"

var _enumA_index = [...]uint8{0, 2, 4, 6, 8}

func (i enumA) String() string {
	if i < 0 || i+1 >= enumA(len(_enumA_index)) {
		return fmt.Sprintf("enumA(%d)", i)
	}
	return _enumA_name[_enumA_index[i]:_enumA_index[i+1]]
}

type enumA int

//go:generate stringer -type=enumA
const (
	v0 enumA = iota
	v1
	v2
	v3
)

var stringToEnumA map[string]enumA

func init() {
	stringToEnumA = make(map[string]enumA, 0)

	for i := 0; i < 4; i++ {
		stringToEnumA[enumA(i).String()] = enumA(i)
	}
}

func main() {

	x := stringToEnumA["v3"]

	if x == v3 {
		fmt.Printf("gocha!\n")
	}
}
