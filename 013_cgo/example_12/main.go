package main

/*
#include <stdio.h>
extern void goFunction();

static inline void cFunction() {
	printf("cFunction will call goFunction\n");
	goFunction();
}
*/
import "C"
import "fmt"

// call function no parameters
func main() {
	C.cFunction()
}

//export goFunction
func goFunction() {
	fmt.Println("goFunction was called")
}
