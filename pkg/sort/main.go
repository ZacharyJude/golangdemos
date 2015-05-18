package main

import "fmt"
import "sort"

type SampleStruct struct {
	f0 string
	f1 int
}

type SampleStructCol []SampleStruct

func (ssc SampleStructCol) Len() int {
	return len(ssc)
}

func (ssc SampleStructCol) Less(i, j int) bool {
	return ssc[i].f0 < ssc[j].f0
}

func (ssc SampleStructCol) Swap(i, j int) {
	ssc[i], ssc[j] = ssc[j], ssc[i]
}

func main() {
	arr := [3]SampleStruct{
		{
			f0: "100003",
			f1: 100003,
		},
		{
			f0: "100003",
			f1: 100002,
		},
		{
			f0: "100001",
			f1: 100001,
		},
	}

	sli := SampleStructCol(arr[0:3])
	fmt.Printf("sli:%s\n", sli)
	fmt.Printf("arr:%s\n", arr)
	sort.Sort(sli)
	fmt.Printf("sli:%s\n", sli)
	fmt.Printf("arr:%s\n", arr)
}
