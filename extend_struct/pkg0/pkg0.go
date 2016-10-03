package pkg0

type I0 interface {
	Foo() string
}

type S0 struct {
	F0 string
}

func (s0 *S0) Foo() string {
	return "foo in s0" + " " + s0.F0
}

type S1 struct {
	F0 string
}

func (s1 *S1) Foo() string {
	return "foo in s1" + " " + s1.F0
}
