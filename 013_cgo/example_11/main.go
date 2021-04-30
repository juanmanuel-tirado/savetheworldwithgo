package main

/*
 #cgo LDFLAGS: -lm
 #include <math.h>
*/
import "C"
import "fmt"

func main() {
	fmt.Printf("Log10 10: %f\n", C.log10(C.double(10)))
	fmt.Printf("Sine Pi: %f\n", C.sin(C.M_PI/2))
	fmt.Printf("Cosine Pi: %f\n", C.cos(C.M_PI))
}
