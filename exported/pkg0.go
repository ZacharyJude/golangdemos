package unexported

type StructA struct {
	F0 string
	F1 int
}

type structB struct {
	F0 int
	F1 string
}

func (sa *StructA) GetStructB() *structB {
	return &structB{F0: 10, F1: "hey"}
}
