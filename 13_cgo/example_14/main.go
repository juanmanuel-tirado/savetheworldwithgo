package main

import (
	"fmt"
	"unsafe"
)

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L . -lclibrary

#include "clibrary.h"

void goFunction_cgo(int number);
*/
import "C"

func main() {
	fmt.Println("Go main() calls a_callback_func")
	C.a_callback_func((C.callback_f)(unsafe.Pointer(C.goFunction_cgo)))
}

//export goFunction
func goFunction(number int) {
	fmt.Println("Go goFunction was finally called with", number)
}