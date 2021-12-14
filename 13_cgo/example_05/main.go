package main

/*
 #include <stdlib.h>
 #include <stdio.h>
 void func_c_print(char* input) {
	 printf("C prints: %s\n",input);
 }
 char* func_c_str() {
	 char* str = "this is a C string";
	 return str;
 }
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	str := "this is a Go string"
	goToC := C.CString(str)
	C.func_c_print(goToC)
	defer C.free(unsafe.Pointer(goToC))

	cStr := C.func_c_str()
	cToGo := C.GoString(cStr)
	fmt.Printf("Go prints: %s\n", cToGo)

	nCToGo := C.GoStringN(cStr, 10)
	fmt.Printf("Go prints: %s\n", nCToGo)
}
