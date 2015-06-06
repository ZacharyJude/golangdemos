package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type A struct {
	F0 string
	F1 int
}

func main() {
	var buf bytes.Buffer
	tmpl, _ := template.ParseFiles("./t0.tmpl")
	tmpl.ExecuteTemplate(&buf, "t02", map[string]string{
		"abc": "shit!",
	})

	fmt.Printf("%s\n", buf.String())
}
