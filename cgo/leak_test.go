package main

/*
#cgo CFLAGS: -std=c99
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
	long long int a;
	char b[1024*1024];
} X;
void P(X* px) {
	if(px->a % 1000 == 0) {
		printf("%lld\n", px->a);
	}
	return;
}

X* get_in_C(long long int i) {
	X* out = (X*)malloc(sizeof(X));
	out->a = i;
	memset(out->b, 0, sizeof(char) * 1024*1024);
	out->b[0] = 'a';
	return out;
}

void LoopP() {
	for(long long int i=0;1;i++) {
		X* px = get_in_C(i);
		P(px);
	}
}
*/
import "C"

//import "fmt"
//import "unsafe"

func get(i int64) *C.X {
	out := &C.X{a: C.longlong(i)}
	return out
}

func main() {

	// C.LoopP() // leak
	var x *C.X
	for i := int64(0); true; i++ {
		x = get(i) // won't leak
		//x = C.get_in_C(C.longlong(i)) // leak
		//defer C.free(unsafe.Pointer(x))
		//defer C.free(x)
		C.P(x)
	}
	return
}
