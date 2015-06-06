package main

import (
	//"io/ioutil"
	"os"
	"text/template"
)

type MemberSliceGetterFuncTmpl struct {
	Prefix    string
	FieldName string
	FieldType string
	OwnerType string
}

func main() {
	t, _ := template.ParseFiles(
		"T0.tmpl",
		"T1.tmpl",
	)
	_ = t.ExecuteTemplate(
		os.Stdout,
		"FuncOfMemberSliceGetter",
		MemberSliceGetterFuncTmpl{
			Prefix:    "hidata",
			FieldName: "PoiUid",
			FieldType: "string",
			OwnerType: "House",
		})
}
