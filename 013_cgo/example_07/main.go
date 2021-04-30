package main

/*
 #include <stdlib.h>
 int* create_array(int n) {
	int* result;
	result = (int*)malloc(sizeof(int)*n);
	for( int i = 0; i < n; i++){
		result[i] = i;
	}
	return result;
 }
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	array := C.create_array(5)
	defer C.free(unsafe.Pointer(array))
	fmt.Printf("C returns a pointer: %d\n", array)
	aux := (*[5]int32)(unsafe.Pointer(array))
	fmt.Printf("We extract the content: %d\n", aux)
	for _, i := range aux {
		fmt.Printf("->%d\n", i)
	}
}