package main

// #include <stdio.h>
// void func_in_c() {
//	printf("printed with C code\n");
// }
import "C"
import "fmt"

func main() {
	C.func_in_c()
	fmt.Println("printed with Go code")
}
