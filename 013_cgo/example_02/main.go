package main

/*
 #include <stdio.h>
 // This is a comment
 void func_in_c() {
	printf("printed with C code\n");
 }
*/
import "C"
import "fmt"

func main() {
	C.func_in_c()
	fmt.Println("printed with Go code")
}
