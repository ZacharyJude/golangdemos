package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func e(a int, b string) {
	fmt.Printf("a %v\tb %v\n", a, b)
}

func main() {
	methodType := reflect.ValueOf(e).Type()
	fmt.Printf("name of method:%s\n", runtime.FuncForPC(reflect.ValueOf(e).Pointer()).Name())
	for i := 0; i < methodType.NumIn(); i++ {
		paramType := methodType.In(i)
		fmt.Printf("%v\n", paramType)
	}
}
