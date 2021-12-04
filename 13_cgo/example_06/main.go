package main

/*
 #include<stdio.h>
 int sum_array(int a[], int length)  {
	 int i,result = 0;
	 for (i = 0; i < length; i++) {
		result += a[i];
	 }
	 return result;
 }
*/
import "C"
import "fmt"

func main() {
	a := []int32{0, 1, 2, 3, 4}
	first := (*C.int)(&a[0])
	res := C.sum_array(first, C.int(len(a)))
	fmt.Printf("The result is: %d\n", res)
}