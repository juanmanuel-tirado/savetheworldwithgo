package main

/*
#include <stdio.h>
extern int goFunction(int number);

static inline void cFunction(int number) {
	printf("cFunction will call goFunction with %d\n",number);
	int x = goFunction(number);
	printf("goFunction returned %d\n", x);
}
*/
import "C"
import "fmt"

// call function with parameters
func main() {
	C.cFunction(42)
}

//export goFunction
func goFunction(number int32) int32 {
	fmt.Println("goFunction was called with", number)
	return number * 2
}
